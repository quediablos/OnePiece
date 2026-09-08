package connection

import (
	"OnePiece/core"
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

// ParsedRequest holds the parsed fields from an incoming HTTP request.
type ParsedRequest struct {
	Method  string
	Path    string
	Query   string
	Headers http.Header
	Body    []byte
}

func ReadHTTPRequest(conn net.Conn) (*ParsedRequest, error) {
	reader := bufio.NewReader(conn)

	req, err := http.ReadRequest(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTTP request: %w", err)
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	return &ParsedRequest{
		Method:  req.Method,
		Path:    req.URL.Path,
		Query:   req.URL.RawQuery,
		Headers: req.Header,
		Body:    body,
	}, nil
}

func (r *ParsedRequest) ParseURL() (core.Operation, string, error) {
	// Trim leading slash and split: ["lock", "42"]
	segments := strings.Split(strings.TrimPrefix(r.Path, "/"), "/")
	if len(segments) != 2 || segments[0] == "" || segments[1] == "" {
		return "", "", fmt.Errorf("invalid path %q: expected /<operation>/<resource_id>", r.Path)
	}

	op := core.Operation(segments[0])
	switch op {
	case core.Lock:
	case core.Unlock:
		// valid
	default:
		return "", "", fmt.Errorf("unknown operation %q", segments[0])
	}

	return op, segments[1], nil
}

// ReleaseClientHttp Writes a response to the client and releases its hold.
func ReleaseClientHttp(conn net.Conn, response string) {

	conn.Write([]byte(response))
	conn.Close()
}

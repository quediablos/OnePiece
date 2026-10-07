package network

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

func (r *ParsedRequest) ParseURL() (core.Operation, string, []string, error) {

	segments := strings.Split(strings.TrimPrefix(r.Path, "/"), "/")
	if segments[0] == "" || segments[1] == "" {
		return "", "", nil, fmt.Errorf("invalid path %q: expected /<operation>/<resource_id>", r.Path)
	}

	var extraParams []string

	op := core.Operation(segments[0])

	switch op {
	case core.OpLock:
	case core.OpUnlock:
	case core.OpReserveStock:
	case core.OpReleaseStock:
		break
	case core.OpCreateStock:
		extraParams = append(extraParams, segments[2])
		break

	case core.OpSetupRateLimiter:
		extraParams = append(extraParams, segments[2]) //userId
		extraParams = append(extraParams, segments[3]) //timeFrame
		extraParams = append(extraParams, segments[4]) //rate
		break
	case core.OpWaitForRateLimiter:
		extraParams = append(extraParams, segments[2]) //userId
		break

	default:
		return "", "", nil, fmt.Errorf("unknown operation %q", segments[0])
	}

	return op, segments[1], extraParams, nil
}

// ReleaseClient Writes a response to the client and releases its hold.
func ReleaseClient(conn net.Conn, response string) {

	conn.Write([]byte(response))
	conn.Close()
}

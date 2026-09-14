package core

import (
	"net"
)

// Operation represents the type of operation to be performed.
type Operation string

const (
	OpLock         Operation = "lock"
	OpUnlock       Operation = "unlock"
	OpReserveStock Operation = "reserve_stock"
	OpReleaseStock Operation = "release_stock"
	OpCreateStock  Operation = "create_stock"
)

type OperationData struct {
	Operation   Operation
	ResourceId  string
	ExtraParams []string
	Conn        net.Conn
}

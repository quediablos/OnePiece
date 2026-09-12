package core

// Operation represents the type of operation to be performed.
type Operation string

const (
	OpLock         Operation = "lock"
	OpUnlock       Operation = "unlock"
	OpReserveStock Operation = "reserve_stock"
	OpReleaseStock Operation = "release_stock"
)

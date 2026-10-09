package core

type PrioritizedClient struct {
	OperationData OperationData
	Priority      float64 // lower value = higher priority (min-heap)
}

type ClientQueue []PrioritizedClient

func (q ClientQueue) Len() int           { return len(q) }
func (q ClientQueue) Less(i, j int) bool { return q[i].Priority < q[j].Priority } // min-heap
func (q ClientQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *ClientQueue) Push(x any) {
	*q = append(*q, x.(PrioritizedClient))
}

func (q *ClientQueue) Pop() any {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[:n-1]
	return item
}

package prior

import (
	"container/heap"
	"sync"
)

type Node struct {
	Key      interface{}
	Value    interface{}
	Priority int
	Index    int
	mutex    sync.RWMutex
}

func NewNode(key interface{}, v interface{}, priority int) *Node {
	return &Node{
		Key:      key,
		Value:    v,
		Priority: priority,
		Index:    -1,
	}
}

func (n *Node) GetKey() interface{} {
	defer n.mutex.RUnlock()
	n.mutex.RLock()
	return n.Key
}

func (n *Node) GetValue() interface{} {
	defer n.mutex.RUnlock()
	n.mutex.RLock()
	return n.Value
}

func (n *Node) GetIndex() int {
	defer n.mutex.RUnlock()
	n.mutex.RLock()
	return n.Index
}

func (n *Node) UpdatePriority(newPrio int) {
	defer n.mutex.Unlock()
	n.mutex.Lock()
	n.Priority = newPrio
}

type Nodes []*Node

func (nodes Nodes) Len() int {
	return len(nodes)
}

func (nodes Nodes) Less(i, j int) bool { return nodes[i].Priority < nodes[j].Priority }

func (nodes Nodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
	nodes[i].Index = i
	nodes[j].Index = j
}

func (nodes *Nodes) Push(v interface{}) {
	node := v.(*Node)
	node.Index = len(*nodes)
	*nodes = append(*nodes, node)
}

func (nodes *Nodes) Pop() interface{} {
	old := *nodes
	size := len(old)
	node := old[size-1]
	// for safety
	node.Index = -1
	*nodes = old[0 : size-1]
	return node
}

type PriorityQueue struct {
	nodes Nodes
	mutex sync.RWMutex
}

func (pq *PriorityQueue) Push(n *Node) {
	defer pq.mutex.Unlock()
	pq.mutex.Lock()
	heap.Push(&(pq.nodes), n)
}

func (pq *PriorityQueue) Pop() *Node {
	defer pq.mutex.RUnlock()
	pq.mutex.RLock()
	if len(pq.nodes) <= 0 {
		return nil
	}
	n := heap.Pop(&(pq.nodes))
	return n.(*Node)
}

func (pq *PriorityQueue) Remove(index int) {
	pq.mutex.RLock()
	if index < 0 || index >= len(pq.nodes) {
		return
	}
	pq.mutex.RUnlock()
	pq.mutex.Lock()
	heap.Remove(&(pq.nodes), index)
	pq.mutex.Unlock()
}

func (pq *PriorityQueue) Length() int {
	defer pq.mutex.RUnlock()
	pq.mutex.RLock()
	return len(pq.nodes)
}

func NewPriorityQueue() *PriorityQueue {
	pq := &PriorityQueue{nodes: make(Nodes, 0, 1024)}
	heap.Init(&(pq.nodes))
	return pq
}

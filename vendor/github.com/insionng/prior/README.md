# About prior

prior is a priority queue based on golang container/heap.

# Explation

**Node**: node is the unit insert to queue. Node has attributes:  
    Key:      key associated with node, can be nil  
    value:    value of key, can be nil  
    Priority:priority of node  
    Index:    index in queue  

# API

**Push(Node)**   : push a node into queue, O(logN) where N is queue length  
**Pop(Node)**    : fetch node with max priority, O(1)  
**Remove(index)**: remove a specified node of index, O(logN) where N is queue length  
**Length()**     : queue length  

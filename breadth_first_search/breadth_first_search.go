package breadth_first_search

import "github.com/oleksiivelychko/go-code-helpers/array"

/*
BreadthFirstSearch (BFS)
An algorithm for searching a tree data structure for a node that satisfies a given property.

O(V+E), V is number of vertices, E is number of edges.
Only for unweighted graphs.

1. Put first node as node-source into dequeue.
2. Pop the node from dequeue and if is a target then finish.
3. If not - put into dequeue all its non-processed successors.
4. Repeat steps 2,3 until dequeue is non-empty.
*/
func BreadthFirstSearch(graph map[string][]string, target, source string) string {
	var dequeue []string
	dequeue = append(dequeue, graph[source]...)

	var searched []string

	for len(dequeue) > 0 {
		item := dequeue[0]
		dequeue = array.Pop(dequeue, 0)

		if array.In(searched, item) {
			continue
		}

		if found := item == target; found {
			return item
		}

		dequeue = append(dequeue, graph[item]...)
		searched = append(searched, item)
	}

	return ""
}

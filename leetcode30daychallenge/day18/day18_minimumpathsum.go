package day18

import "math"

func minPathSum(grid [][]int) int {
	cost := make([][]int, len(grid))
	for row := range grid {
		cost[row] = make([]int, len(grid[row]))

		for col := range grid[row] {
			cost[row][col] = math.MaxInt32
		}
	}

	cost[0][0] = grid[0][0]

	for row := range grid {
		for col := range grid[row] {
			if row > 0 {
				if cost[row-1][col]+grid[row][col] < cost[row][col] {
					cost[row][col] = cost[row-1][col] + grid[row][col]
				}
			}

			if col > 0 {
				if cost[row][col-1]+grid[row][col] < cost[row][col] {
					cost[row][col] = cost[row][col-1] + grid[row][col]
				}
			}
		}
	}

	return cost[len(grid)-1][len(grid[0])-1]
}

// Solution using Dijsktra's algorithm implemented with heap. It ran too long, but
// I'll leave it here for posterity.

// import (
// 	"container/heap"
// 	"math"
// )

// type node struct {
// 	paths  []*path
// 	weight int
// 	parent *node
// 	index  int
// }

// type path struct {
// 	distance int
// 	node     *node
// }

// type priorityQueue []*node

// // heap implementation for our priority queue
// func (pq priorityQueue) Len() int {
// 	return len(pq)
// }

// func (pq priorityQueue) Less(i, j int) bool {
// 	return pq[i].weight < pq[j].weight
// }

// func (pq priorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// 	pq[i].index = i
// 	pq[j].index = j
// }

// func (pq *priorityQueue) Push(n interface{}) {
// 	index := len(*pq)
// 	item := n.(*node)
// 	item.index = index
// 	*pq = append(*pq, item)
// }

// func (pq *priorityQueue) Pop() interface{} {
// 	old := *pq
// 	index := len(old) - 1
// 	node := old[index]
// 	old[index] = nil
// 	node.index = -1
// 	*pq = old[:index]

// 	return node
// }

// func minPathSum(grid [][]int) int {
// 	first, last, allNodes := makeGrid(grid)

// 	dijkstra(first, allNodes)

// 	return last.weight
// }

// func dijkstra(start *node, nodes []*node) {
// 	start.weight = 0
// 	pq := priorityQueue{}

// 	for i := range nodes {
// 		pq.Push(nodes[i])
// 	}

// 	heap.Init(&pq)

// 	for pq.Len() > 0 {
// 		node := heap.Pop(&pq).(*node)

// 		for i := range node.paths {
// 			if node.paths[i].distance+node.weight < node.paths[i].node.weight {
// 				node.paths[i].node.weight = node.paths[i].distance + node.weight
// 				node.paths[i].node.parent = node
// 			}
// 		}

// 		heap.Init(&pq)
// 	}
// }

// func makeGrid(grid [][]int) (*node, *node, []*node) {
// 	allNodes := []*node{}
// 	nodeGrid := make([][]*node, len(grid))

// 	// make nodes
// 	for row := range grid {
// 		nodeGrid[row] = make([]*node, len(grid[row]))

// 		for col := range grid[row] {
// 			nodeGrid[row][col] = newNode()
// 			allNodes = append(allNodes, nodeGrid[row][col])
// 		}
// 	}

// 	// add paths
// 	for row := range grid {
// 		for col := range grid[row] {
// 			if row > 0 {
// 				nodeGrid[row-1][col].paths = append(nodeGrid[row-1][col].paths, &path{
// 					node:     nodeGrid[row][col],
// 					distance: grid[row-1][col],
// 				})
// 			}

// 			if col > 0 {
// 				nodeGrid[row][col-1].paths = append(nodeGrid[row][col-1].paths, &path{
// 					node:     nodeGrid[row][col],
// 					distance: grid[row][col-1],
// 				})
// 			}
// 		}
// 	}

// 	// add end node
// 	end := newNode()
// 	nodeGrid[len(grid)-1][len(grid[0])-1].paths = append(nodeGrid[len(grid)-1][len(grid[0])-1].paths, &path{
// 		node:     end,
// 		distance: grid[len(grid)-1][len(grid[0])-1],
// 	})

// 	allNodes = append(allNodes, end)

// 	return nodeGrid[0][0], end, allNodes
// }

// func newNode() *node {
// 	return &node{
// 		paths:  []*path{},
// 		weight: math.MaxInt32,
// 	}
// }

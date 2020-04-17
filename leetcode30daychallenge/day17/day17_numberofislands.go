package day17

func numIslands(grid [][]byte) int {
	count := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				markNeighbours(&grid, i, j)
			}
		}
	}

	return count
}

func markNeighbours(g *[][]byte, i, j int) {
	// check out of bounds
	if i < 0 || j < 0 || i > len(*g)-1 || j > len((*g)[0])-1 {
		return
	}

	if (*g)[i][j] == '1' {
		// part of the island, mark it so we don't count it again
		(*g)[i][j] = '2'

		// check neighbours to see if the island continues
		markNeighbours(g, i+1, j)
		markNeighbours(g, i-1, j)
		markNeighbours(g, i, j+1)
		markNeighbours(g, i, j-1)
	}
}

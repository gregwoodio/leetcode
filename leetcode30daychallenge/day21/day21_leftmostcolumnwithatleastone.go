package day21

/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get(int, int) int
 *     Dimensions() []int
 * }
 */

type BinaryMatrix struct {
	nums     [][]int
	apiCalls int
}

func (b BinaryMatrix) Get(x, y int) int {
	b.apiCalls++
	return b.nums[x][y]
}

func (b BinaryMatrix) Dimensions() []int {
	b.apiCalls++
	return []int{
		len(b.nums),
		len(b.nums[0]),
	}
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dimensions := binaryMatrix.Dimensions()

	if dimensions[0] == 0 || dimensions[1] == 0 {
		return -1
	}

	// pointer to top right corner
	pX, pY := 0, dimensions[0]-1

	// track whether the given column has a one in it
	hasOnes := make([]bool, dimensions[1])

	for {
		curr := binaryMatrix.Get(pX, pY)
		if curr == 0 {
			pX++
		} else if curr == 1 {
			hasOnes[pY] = true
			pY--
		}

		if pX == dimensions[0] || pY < 0 {
			break
		}
	}

	// return the first column with a one
	for col, hasOne := range hasOnes {
		if hasOne {
			return col
		}
	}

	return -1
}

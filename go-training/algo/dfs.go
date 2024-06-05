package algo

// https://leetcode.com/problems/number-of-islands/

type Position struct {
	row int
	col int
}

func numIslands(grid [][]byte) int {
	vs := map[Position]bool{}
	var count int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !vs[Position{i, j}] && grid[i][j] == '1' {
				count++
				dfs(grid, i, j, vs)
			}
		}
	}

	return count
}

func dfs(grid [][]byte, row int, col int, vs map[Position]bool) {
	rowLen := len(grid)
	colLen := len(grid[0])

	if row < 0 || row >= rowLen || col < 0 || col >= colLen {
		return
	}

	if grid[row][col] == '0' {
		return
	}

	vs[Position{row, col}] = true

	dx := []int{-1, 0, 1, 0}
	dy := []int{0, 1, 0, -1}

	for i := 0; i < len(dx); i++ {
		newX := row + dx[i]
		newY := col + dy[i]
		if !vs[Position{newX, newY}] {
			dfs(grid, newX, newY, vs)
		}
	}
}

// More cleaner
func numIslandsC2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	var count int

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || c < 0 || r >= rows || c >= cols || grid[r][c] == '0' {
			return
		}
		grid[r][c] = '0' // Mark as visited
		dfs(r-1, c)      // Up
		dfs(r+1, c)      // Down
		dfs(r, c-1)      // Left
		dfs(r, c+1)      // Right
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				dfs(r, c)
			}
		}
	}

	return count
}

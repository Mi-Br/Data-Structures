package main

func numIslands(grid [][]byte) int {

	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	visited := makeRadar(grid) // grid to track visited
	island_count := 0

	//go through each cell on grid starting 0,0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if visited[i][j] == false {
				if grid[i][j] == '1' {
					island_count++
					dfs(grid, i, j, visited)
				} else {
					visited[i][j] = true
				}
			}

		}
	}
	return island_count
}

func dfs(grid [][]byte, i, j int, visited [][]bool) {
	m := len(grid)
	n := len(grid[0])
	if visited[i][j] {
		return
	} else {
		visited[i][j] = true
		if grid[i][j] == '1' {
			//north
			if i > 0 {
				dfs(grid, i-1, j, visited)
			}
			//south
			if i+1 < m {
				dfs(grid, i+1, j, visited)
			}

			//west
			if j+1 < n {
				dfs(grid, i, j+1, visited)
			}
			//east
			if j > 0 {
				dfs(grid, i, j-1, visited)
			}
		}
	}

}

func makeRadar(grid [][]byte) [][]bool {
	var outp [][]bool = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		outp[i] = make([]bool, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			outp[i][j] = false
		}
	}
	return outp
}

func main() {

}

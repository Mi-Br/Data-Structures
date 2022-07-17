// https://leetcode.com/problems/flood-fill/

// An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.

// You are also given three integers sr, sc, and color. You should perform a flood fill on the image starting from the pixel image[sr][sc].

// To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with color.

// Return the modified image after performing the flood fill.

package main

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] == color {
		return image
	} else {
		target := image[sr][sc]
		dfs(image, sr, sc, color, target)
	}
	return image
}

func dfs(image [][]int, r, c, color, target int) {
	//basecase
	if image[r][c] != target {
		return
	} else {
		image[r][c] = color
	}

	if c >= 1 {
		dfs(image, r, c-1, color, target)
	}
	if c+1 < len(image[r]) {
		dfs(image, r, c+1, color, target)
	}
	if r >= 1 {
		dfs(image, r-1, c, color, target)
	}
	if r+1 < len(image) {
		dfs(image, r+1, c, color, target)
	}
}

func main() {}

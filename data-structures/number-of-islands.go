package datastructures

import "fmt"

// 蔓延函数，蔓延它的上下左右
func spread(i, j int, grid [][]byte) {
	// 周围都是水了
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	// 蔓延它的上下
	spread(i+1, j, grid)
	spread(i-1, j, grid)
	spread(i, j+1, grid)
	spread(i, j-1, grid)
}

// problem-link : https://leetcode.cn/problems/number-of-islands/
// 可以把每块陆地看成是一块待浇水的农田，旁边的水会蔓延进整个农田
// 当水无处可蔓延时，也就说明这块区域没有农田了
// 遍历 grid 遇到土地就要去灌溉，然后等灌溉完整块地结束，然后土地 + 1
func numIslands(grid [][]byte) int {
	numIslands := 0
	for i, it := range grid {
		for j, _ := range it {
			if grid[i][j] == '1' {
				spread(i, j, grid)
				numIslands++
			}
		}
	}
	return numIslands
}

func RunNumIslands() {
	var grid = [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	res := numIslands(grid)
	fmt.Print(res)
}

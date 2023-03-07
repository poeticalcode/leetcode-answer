package datastructures

import "fmt"

func dfs(i, j int, board [][]byte) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'Z'
	// 上下左右搜索
	dfs(i+1, j, board)
	dfs(i-1, j, board)
	dfs(i, j+1, board)
	dfs(i, j-1, board)
	return
}

// https://leetcode.cn/problems/surrounded-regions/submissions/
// 反过来考虑，找到与边界 O 相连的区域，然后将其他区域输出为 X
func solve(board [][]byte) {

	rowNum := len(board)

	colNum := len(board[0])

	for i := 0; i < rowNum; i++ {
		dfs(i, 0, board)
		dfs(i, colNum-1, board)
	}

	for i := 0; i < colNum; i++ {
		dfs(0, i, board)
		dfs(rowNum-1, i, board)
	}

	for i, boardI := range board {
		for j, item := range boardI {
			if item == 'Z' {
				board[i][j] = 'O'
			}
			if item == 'O' {
				board[i][j] = 'X'
			}
		}
	}

}

func RunSolve() {
	var board = [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'X', 'X'},
	}
	solve(board)

	for i := range board {
		for _, item := range board[i] {
			fmt.Printf("%c ", item)
		}
		fmt.Println()
	}
}

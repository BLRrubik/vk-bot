package utils

func evaluate(board [][]string) int {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			if board[i][0] == "O" {
				return 10
			} else if board[i][0] == "X" {
				return -10
			}
		}
		if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			if board[0][i] == "O" {
				return 10
			} else if board[0][i] == "X" {
				return -10
			}
		}
	}
	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		if board[0][0] == "O" {
			return 10
		} else if board[0][0] == "X" {
			return -10
		}
	}
	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		if board[0][2] == "O" {
			return 10
		} else if board[0][2] == "X" {
			return -10
		}
	}
	return 0
}

func isMovesLeft(board [][]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "▢" {
				return true
			}
		}
	}
	return false
}

func minimax(board [][]string, depth int, isMaximizingPlayer bool) int {
	// Получаем оценку доски, если игра закончена
	score := evaluate(board)
	if score == 10 {
		return score - depth
	}
	if score == -10 {
		return score + depth
	}
	if isMovesLeft(board) == false {
		return 0
	}
	if isMaximizingPlayer {
		bestScore := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == "▢" {
					board[i][j] = "O"
					score := minimax(board, depth+1, false)
					board[i][j] = "▢"
					if score > bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == "▢" {
					board[i][j] = "X"
					score := minimax(board, depth+1, true)
					board[i][j] = "▢"
					if score < bestScore {
						bestScore = score
					}
				}
			}
		}
		return bestScore
	}
}

func FindBestMove(board [][]string) (int, int) {
	bestScore := -1000
	row := -1
	col := -1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "▢" {
				board[i][j] = "O"
				score := minimax(board, 0, false)
				board[i][j] = "▢"
				if score > bestScore {
					bestScore = score
					row = i
					col = j
				}
			}
		}
	}
	return row, col
}

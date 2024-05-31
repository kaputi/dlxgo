package ssudoku

func getNextEmpty(board *boardT) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j].Value == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func isValid(board *boardT, row, col, value int) bool {
	if value == 0 {
		return true
	}

	for i := 0; i < 9; i++ {
		if board[row][i].Value == value && i != col {
			return false
		}
		if board[i][col].Value == value && i != row {
			return false
		}
	}

	firstRowInBox := row - row%3
	firstColInBox := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			checkRow := firstRowInBox + i
			checkCol := firstColInBox + j
			if checkRow == row && checkCol == col {
				continue
			}
			if board[checkRow][checkCol].Value == value {
				return false
			}
		}
	}

	return true
}

func boardIsFull(board *boardT) bool {
  for i := 0; i < 9; i++ {
    for j := 0; j < 9; j++ {
      if board[i][j].Value == 0 {
        return false
      }
    }
  }
  return true
}

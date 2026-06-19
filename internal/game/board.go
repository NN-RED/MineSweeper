package game

import (
	"math/rand"
)

const (
	Playing = iota // Состояния игры
	Win            //----------------
	Lose
)

type Cell struct { // Cтруктура для клетки со всеми нужными параметрами
	IsMine        bool
	Found         bool
	IsFlagged     bool
	NeighborMines int
}

type Board struct { // Структура для поля
	Width  int
	Height int
	Grid   [][]Cell
	Status int
}

func CreatingBoard(width, height, mines int) *Board {
	board := &Board{
		Width:  width,
		Height: height,
		Grid:   make([][]Cell, height),
		Status: Playing,
	}

	for i := range board.Grid { // Заполняем наше поле своими клетками
		board.Grid[i] = make([]Cell, width)
	}

	count := 0
	for count < mines { // Рандомно заполняем поле минами
		x := rand.Intn(board.Width)
		y := rand.Intn(board.Height)

		if !board.Grid[y][x].IsMine {
			board.Grid[y][x].IsMine = true
			count++
		}
	}

	for y := 0; y != board.Height; y++ { //Подсчет мин вблизи клетки
		for x := 0; x != board.Width; x++ {
			if !board.Grid[y][x].IsMine {
				board.Grid[y][x].NeighborMines = board.countNeighbors(x, y)
			}
		}
	}
	return board
}

func (board *Board) countNeighbors(x, y int) int { //Метод для нашей структуры - выполняет подсчет мин
	count := 0

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < board.Width && ny >= 0 && ny < board.Height {
				if board.Grid[ny][nx].IsMine {
					count++
				}
			}
		}
	}

	// (0,0) (0,1) (0,2) ...
	// (1,0) (1,1 - мина) (1,2)
	// (2,0) (2,1) (2,2)
	// - 1 итерация === Берем клетку (1,1) и прибавляем -1 => (0,0)
	// - 2 итерация ==== (1, 1) + (-1,0) ==> (0,1)
	// и тд

	return count
}

func (board *Board) ToggleFlag(x, y int) { // Логика выставления флага
	if board.Status != Playing || board.Grid[y][x].Found { // Если игра закончена или клетка уже найдена, то ничего не делаем
		return
	}

	board.Grid[y][x].IsFlagged = !board.Grid[y][x].IsFlagged
}

func (board *Board) OpenCell(x, y int) { //Основа

	if y < 0 || y >= board.Height || x < 0 || x >= board.Width {
		return
	}

	if board.Status != Playing || board.Grid[y][x].Found || board.Grid[y][x].IsFlagged {
		return // Делаем пустой return если игра закончена или клетка найдена или клетка помечена флагом
	}

	board.Grid[y][x].Found = true

	if board.Grid[y][x].IsMine { // Заканчиваем игру если нарвались на мину
		board.Status = Lose
		return
	}

	if board.Grid[y][x].NeighborMines == 0 && !board.Grid[y][x].IsMine { //Рядом нет мин -- ищем клетки с числами
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx != 0 || dy != 0 {
					board.OpenCell(x+dx, y+dy) // Рекурсивный вызов пока не найдем клетку с числом
				}
			}
		}
	}

	if board.checkWin() { // Смотрим выиграл или нет
		board.Status = Win
	}

}

func (board *Board) checkWin() bool { //Проверка на выигрыш
	for y := 0; y != board.Height; y++ {
		for x := 0; x != board.Width; x++ {
			if !board.Grid[y][x].IsMine && !board.Grid[y][x].Found { //Если клетка не мина, и она не найдена ==> Игрок не выиграл
				return false
			}
		}
	}

	return true
}

func (board *Board) CountFlags() int {
	count_flags := 0
	for y := 0; y != board.Height; y++ {
		for x := 0; x != board.Width; x++ {
			if board.Grid[y][x].IsFlagged {
				count_flags++
			}
		}
	}

	return count_flags
}

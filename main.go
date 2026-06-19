package main

import (
	"fmt"
	"log"
	"minesweeper/internal/game"
	"os"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	fmt.Println("=== САПЕР НА GO ===")
	fmt.Println("1. Новичок (9x9) | 2. Любитель (16x16) | 3. Профи (30x16)")
	fmt.Print("Выбери 1-3: ")

	var choice int
	fmt.Scanln(&choice)

	var width, height, mines int
	switch choice {
	case 2:
		width, height, mines = 16, 16, 40
	case 3:
		width, height, mines = 30, 16, 99
	default:
		width, height, mines = 9, 9, 10
	}

	board := game.CreatingBoard(width, height, mines)

	cursorX := width / 2 //Начальная позиция нашего курсора
	cursorY := height / 2

	var startTime time.Time
	gameStarted := false
	timeElapsed := 0

	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	// Очищаем экран перед стартом
	fmt.Print("\033[2J\033[H")

	//Запускаем го рутину для отрисовки и счетчика времени
	go func() {
		for {
			if gameStarted && board.Status == game.Playing {
				timeElapsed = int(time.Since(startTime).Seconds())
			}

			// Возвращаем курсор в начало (0,0)
			fmt.Print("\033[H")

			minesLeft := mines - board.CountFlags()
			if minesLeft < 0 {
				minesLeft = 0
			}

			minutes := timeElapsed / 60
			seconds := timeElapsed % 60
			fmt.Printf(" 🚩 Мин осталось: %02d  |  ⏱️  Время: %02d:%02d", minesLeft, minutes, seconds)
			fmt.Println()
			fmt.Println()

			// Вывод игрового поля
			for y := 0; y < board.Height; y++ {
				fmt.Print(" ")
				for x := 0; x < board.Width; x++ {

					if x == cursorX && y == cursorY && board.Status == game.Playing {
						fmt.Print("⭐ ")
						continue
					}

					cell := board.Grid[y][x]
					cellStr := ""

					// Каждый элемент строго 2 символа в ширину
					if board.Status == game.Lose && cell.IsMine {
						cellStr = "💣"
					} else if cell.IsFlagged {
						cellStr = "🚩"
					} else if !cell.Found {
						cellStr = "⬜"
					} else if cell.NeighborMines > 0 {
						// Цифры с пробелом справа для выравнивания
						switch cell.NeighborMines {
						case 1:
							cellStr = fmt.Sprintf("\033[34m%d \033[0m", cell.NeighborMines) // Синяя
						case 2:
							cellStr = fmt.Sprintf("\033[32m%d \033[0m", cell.NeighborMines) // Зеленая
						case 3:
							cellStr = fmt.Sprintf("\033[31m%d \033[0m", cell.NeighborMines) // Красная
						default:
							cellStr = fmt.Sprintf("\033[35m%d \033[0m", cell.NeighborMines) // Пурпурная
						}
					} else {
						cellStr = "  "
					}

					fmt.Printf("%s ", cellStr)
				}
				fmt.Println()
			}

			switch board.Status {
			case game.Lose:
				fmt.Println("\n 💀 БУМ! Вы подорвались на мине. (Esc для выхода)")

			case game.Win:
				minutes := timeElapsed / 60
				seconds := timeElapsed % 60

				if minutes > 0 {
					fmt.Printf("\n 🎉 ПОЗДРАВЛЯЕМ! Вы победили за %d мин %d сек! (Esc для выхода)\n", minutes, seconds)
				} else {
					fmt.Printf("\n 🎉 ПОЗДРАВЛЯЕМ! Вы победили за %d сек! (Esc для выхода)\n", seconds)
				}

			}

			os.Stdout.Sync()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		if key == keyboard.KeyEsc {
			break
		}

		if board.Status != game.Playing {
			continue
		}

		switch {
		case key == keyboard.KeyArrowUp || char == 'w' || char == 'W':
			if cursorY > 0 {
				cursorY--
			}
		case key == keyboard.KeyArrowDown || char == 's' || char == 'S':
			if cursorY < height-1 {
				cursorY++
			}
		case key == keyboard.KeyArrowLeft || char == 'a' || char == 'A':
			if cursorX > 0 {
				cursorX--
			}
		case key == keyboard.KeyArrowRight || char == 'd' || char == 'D':
			if cursorX < width-1 {
				cursorX++
			}
		case key == keyboard.KeySpace || key == keyboard.KeyEnter:
			if !gameStarted {
				gameStarted = true
				startTime = time.Now()
			}
			board.OpenCell(cursorX, cursorY)
		case char == 'f' || char == 'F':
			board.ToggleFlag(cursorX, cursorY)
		}
	}
}

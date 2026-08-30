# Minesweeper-TUI

Minimalistic realization game Minesweeper in terminal, based on Go. Project was released with pure architecture and high rendering perfomance without heavy GUI frameworks.

## Project features

* **Pure architecture:** All math and logic of the game are located in `internal/game` (`board.go`). Rendering part (`main.go`) knows nothing about mines generation rules, its just show current state of the game.
* **Smooth ANSI-rendering:** Interface rendering based in separate **goroutine**. Using Esc-queue `\033[H` can update display without terminal shimmer.
* **Cross-platform capability:** Game compiles in one independent binary file for Windows, Linux and macOs.

## Project structure
```text
├── go.mod                # Go-module and dependencies
├── go.sum                # Control summs (security of libraries)
├── main.go               # Entry point: initialization, rendering goroutine and input processing
└── internal/
    └── game/
        └── board.go      # The game's mathematical core 
```

## Game Controls
 * Arrows / WASD — Moving the cursor (⭐) on game map.
 * Space / Enter — Open selected cell.
 * F / f — Install/remove flag (🚩) on selected cell.
 * Esc — Instant exit from the game.

## Startup instructions

###     Requirements
 * Go version 1.18 or higher installed.

###     Quick start:
```Bash
go run main.go
```

* Go will automatically download the minimal dependency for keyboard handling
github.com/eiannone/keyboard in first startup.

---

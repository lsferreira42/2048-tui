package main

import (
	"embed"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

//go:embed move_sound.mp3 ending.mp3
var soundFiles embed.FS

const size = 4

var cellColors = map[int]string{
	2:    "#FFEBCD",
	4:    "#FFE4B5",
	8:    "#FFDAB9",
	16:   "#FFA07A",
	32:   "#FF8C00",
	64:   "#FF7F50",
	128:  "#FF6347",
	256:  "#FF4500",
	512:  "#FF3030",
	1024: "#FF0000",
	2048: "#8B0000",
}

type model struct {
	board      [size][size]int
	score      int
	gameOver   bool
	termWidth  int
	termHeight int
}

func initialModel() (model, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width, height = 80, 24
	}

	m := model{
		board:      [size][size]int{},
		score:      0,
		gameOver:   false,
		termWidth:  width,
		termHeight: height,
	}

	initAudio()

	m.addTile()
	m.addTile()
	return m, nil
}

func (m *model) addTile() {
	emptyCells := [][2]int{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if m.board[i][j] == 0 {
				emptyCells = append(emptyCells, [2]int{i, j})
			}
		}
	}
	if len(emptyCells) > 0 {
		cell := emptyCells[rand.Intn(len(emptyCells))]
		m.board[cell[0]][cell[1]] = 2
		if rand.Float32() < 0.1 {
			m.board[cell[0]][cell[1]] = 4
		}
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("q", "ctrl+c"))):
			return m, tea.Quit
		case key.Matches(msg, key.NewBinding(key.WithKeys("r"))):
			if m.gameOver {
				newModel, _ := initialModel()
				return newModel, nil
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("up", "down", "left", "right"))):
			if !m.gameOver {
				moved := m.move(msg.String())
				if moved {
					m.addTile()
					m.gameOver = m.isGameOver()
					playMoveSound()
					if m.gameOver {
						playEndingSound()
					}
				}
			}
		}
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
	}
	return m, nil
}

func (m model) View() string {
	maxBoardWidth := int(float64(m.termWidth) * 0.8)
	maxBoardHeight := int(float64(m.termHeight) * 0.8)

	cellWidth := (maxBoardWidth - size - 1) / size
	cellHeight := (maxBoardHeight - size - 1) / size

	if cellWidth < 3 {
		cellWidth = 3
	}
	if cellHeight < 1 {
		cellHeight = 1
	}

	boardHeight := cellHeight*size + size + 1

	board := make([]string, boardHeight)

	board[0] = "┌" + strings.Repeat("─", cellWidth) + strings.Repeat("┬"+strings.Repeat("─", cellWidth), size-1) + "┐"

	for i := 0; i < size; i++ {
		for row := 0; row < cellHeight; row++ {
			cellContent := ""
			for j := 0; j < size; j++ {
				value := m.board[i][j]
				cell := ""
				if value == 0 {
					cell = strings.Repeat(" ", cellWidth)
				} else {
					style := lipgloss.NewStyle().
						Bold(true).
						Foreground(lipgloss.Color("#000000")).
						Background(lipgloss.Color(cellColors[value])).
						Align(lipgloss.Center).
						Width(cellWidth)
					if row == cellHeight/2 {
						cell = style.Render(fmt.Sprintf("%d", value))
					} else {
						cell = style.Render(strings.Repeat(" ", cellWidth))
					}
				}
				cellContent += "│" + cell
			}
			cellContent += "│"
			board[i*(cellHeight+1)+row+1] = cellContent
		}

		if i < size-1 {
			board[i*(cellHeight+1)+cellHeight+1] = "├" + strings.Repeat("─", cellWidth) + strings.Repeat("┼"+strings.Repeat("─", cellWidth), size-1) + "┤"
		}
	}

	board[boardHeight-1] = "└" + strings.Repeat("─", cellWidth) + strings.Repeat("┴"+strings.Repeat("─", cellWidth), size-1) + "┘"

	s := "\n  2048 Game\n\n" + strings.Join(board, "\n")

	s += fmt.Sprintf("\n\nScore: %d\n", m.score)

	if m.gameOver {
		s += "Game Over! Press 'r' to restart or 'q' to quit.\n"
	} else {
		s += "Use arrow keys to move. Press 'q' to quit.\n"
	}

	return lipgloss.NewStyle().Width(m.termWidth).Align(lipgloss.Center).Render(s)
}

func (m *model) move(direction string) bool {
	rotations := map[string]int{"up": 3, "right": 2, "down": 1, "left": 0}
	m.rotateBoard(rotations[direction])
	moved := m.shiftAndMergeLeft()
	m.rotateBoard(4 - rotations[direction])
	return moved
}

func (m *model) rotateBoard(times int) {
	for i := 0; i < times; i++ {
		m.board = rotate90(m.board)
	}
}

func rotate90(b [size][size]int) [size][size]int {
	var rotated [size][size]int
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			rotated[j][size-1-i] = b[i][j]
		}
	}
	return rotated
}

func (m *model) shiftAndMergeLeft() bool {
	moved := false
	for row := 0; row < size; row++ {
		newRow, rowMoved, rowScore := shiftAndMergeRow(m.board[row])
		if rowMoved {
			moved = true
			m.board[row] = newRow
			m.score += rowScore
		}
	}
	return moved
}

func shiftAndMergeRow(row [size]int) ([size]int, bool, int) {
	var newRow [size]int
	idx := 0
	moved := false
	score := 0

	for i := 0; i < size; i++ {
		if row[i] != 0 {
			if idx > 0 && newRow[idx-1] == row[i] {
				newRow[idx-1] *= 2
				score += newRow[idx-1]
				moved = true
			} else {
				newRow[idx] = row[i]
				if idx != i {
					moved = true
				}
				idx++
			}
		}
	}
	return newRow, moved, score
}

func (m *model) isGameOver() bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if m.board[i][j] == 0 {
				return false
			}
			if (i < size-1 && m.board[i][j] == m.board[i+1][j]) ||
				(j < size-1 && m.board[i][j] == m.board[i][j+1]) {
				return false
			}
		}
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())
	m, err := initialModel()
	if err != nil {
		fmt.Printf("Error initializing model: %v\n", err)
		os.Exit(1)
	}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

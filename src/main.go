package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aasanchez/unixy2k-cli/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles (defined globally for reuse)
var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#87CEFA")). // LightSkyBlue
			PaddingBottom(1)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("242")). // Light grey
			Width(15).                          // Fixed width for alignment
			PaddingRight(1)

	valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")) // White

	binaryValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF69B4")) // HotPink

	countdownLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("220")). // Gold
				PaddingTop(1).
				PaddingBottom(1)

	countdownValueStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#32CD32")) // LimeGreen

	errorStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF0000")) // Red
			// Border(lipgloss.RoundedBorder()).
			// BorderForeground(lipgloss.Color("#FF0000")).
			// Padding(1, 2)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")). // Darker grey
			PaddingTop(1)
)

// model holds the application's state.
type model struct {
	timerStatus timer.Status
	err         error
	width       int
	height      int
}

// tickMsg is a message sent periodically to update the timer.
type tickMsg struct{}

// initialModel creates the initial state of the application.
func initialModel() model {
	ts, err := timer.GetStatus()
	if err != nil {
		return model{err: err}
	}
	return model{
		timerStatus: *ts,
		err:         nil,
	}
}

// Init is called when the program starts.
func (m model) Init() tea.Cmd {
	return tickCmd()
}

// tickCmd creates a command that sends a tickMsg after a delay.
func tickCmd() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

// Update handles messages and updates the model.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tickMsg:
		newStatus, err := timer.GetStatus()
		if err != nil {
			m.err = err
			return m, tickCmd()
		}
		m.timerStatus = *newStatus
		m.err = nil
		return m, tickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the UI.
func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		// Use a simple styled message for initializing
		initStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("242")).SetString("Initializing display...")
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, initStyle.String())
	}

	if m.err != nil {
		errorMsg := errorStyle.Render(fmt.Sprintf("Error: %v", m.err))
		// Add help text below error for consistency
		help := helpStyle.Render("Press 'q', 'esc', or 'ctrl+c' to quit.")
		fullError := lipgloss.JoinVertical(lipgloss.Center, errorMsg, help)
		return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, fullError)
	}

	// Content formatting
	title := titleStyle.Render("🕰️  Unixy2K CLI Countdown 🕰️") // Added emojis for fun

	utcRow := lipgloss.JoinHorizontal(lipgloss.Left,
		labelStyle.Render("UTC Date:"),
		valueStyle.Render(m.timerStatus.UTCString),
	)

	epochRow := lipgloss.JoinHorizontal(lipgloss.Left,
		labelStyle.Render("Epoch Time:"),
		valueStyle.Render(fmt.Sprintf("%d", m.timerStatus.Epoch)),
	)

	binaryChunkStr := ""
	if m.timerStatus.BinaryChunks != nil {
		binaryChunkStr = strings.Join(m.timerStatus.BinaryChunks, " ")
	}
	binaryRow := lipgloss.JoinHorizontal(lipgloss.Left,
		labelStyle.Render("Epoch Binary:"),
		binaryValueStyle.Render(binaryChunkStr),
	)

	countdownTextLabel := countdownLabelStyle.Render("⏳ Remaining Time until 2038-01-19 03:14:07 UTC:")

	remainingTimeStr := fmt.Sprintf("%d years, %d months, %d days, %d hours, %d minutes, %d seconds",
		m.timerStatus.Years, m.timerStatus.Months, m.timerStatus.Days,
		m.timerStatus.Hours, m.timerStatus.Minutes, m.timerStatus.Seconds,
	)
	countdownTextValue := countdownValueStyle.Render(remainingTimeStr)

	// Main content block
	mainContent := lipgloss.JoinVertical(lipgloss.Center, // Centering text within this block
		title,
		utcRow,
		epochRow,
		binaryRow,
		countdownTextLabel,
		countdownTextValue,
	)

	// Help text
	help := helpStyle.Render("Press 'q', 'esc', or 'ctrl+c' to quit.")

	// Combine main content and help text
	fullScreenContent := lipgloss.JoinVertical(lipgloss.Center,
		mainContent,
		help, // Place help text below the main content block
	)

	// Center the entire block on screen
	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		fullScreenContent,
	)
}

func main() {
	// It's good practice to initialize the model once
	m := initialModel()
	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())

	if _, err := p.Run(); err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
	}
}

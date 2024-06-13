package main

import (
	"fluent_ai/gemini"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

func init() {

	erro := godotenv.Load(".env")

	if erro != nil {
		fmt.Println("Error loading .env file:", erro)
		os.Exit(1)
	}
}

func main() {

	arguments := os.Args

	var idiomaSaida string
	if len(arguments) > 1 {
		idiomaSaida = arguments[1]
	} else {
		idiomaSaida = "português brasileiro"
	}

	p := tea.NewProgram(initialModel(idiomaSaida))

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}

type (
	errMsg error
)

type model struct {
	viewport    viewport.Model
	messages    []string
	idiomaSaida string
	textarea    textarea.Model
	senderStyle lipgloss.Style
	err         error
}

func initialModel(idioma string) model {
	ta := textarea.New()
	ta.Placeholder = "Insira o texto a ser traduzido aqui..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 1000

	ta.SetWidth(70)
	ta.SetHeight(4)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))

	ta.ShowLineNumbers = false

	vp := viewport.New(70, 2)
	vp.SetContent(`Digite abaixo a frase ou texto que você deseja traduzir.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:    ta,
		messages:    []string{},
		idiomaSaida: idioma,
		viewport:    vp,
		senderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		err:         nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			geminiAnswer := gemini.GeneratePrompt(m.textarea.Value(), m.idiomaSaida)
			m.messages = append(m.messages, m.senderStyle.Render("\nGemini: ")+geminiAnswer)
			m.viewport.SetContent(strings.Join(m.messages, "\n"))
			var height int
			if len(geminiAnswer) <= 100 {
				height = 3
			} else {
				height = len(geminiAnswer) / 45
			}
			m.viewport.Height = height
			m.textarea.Reset()
			m.viewport.GotoBottom()
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"

}

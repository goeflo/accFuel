package fuel

import (
	"fmt"
	"goAccFuel/acc"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO detect width
const width = 70
const mainWindowWidth = 68 // smaller size because of the border

var (
	highlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

	// window header and footer

	headerStyle = lipgloss.NewStyle().
			BorderForeground(highlightColor).
			Padding(0, 0).Align(lipgloss.Left).
			Border(lipgloss.NormalBorder()).UnsetBorderBottom()

	windowStyle = lipgloss.NewStyle().
			BorderForeground(highlightColor).
			Padding(0, 0).Align(lipgloss.Left).
			Border(lipgloss.NormalBorder()).UnsetBorderTop()

	// fuel status

	fuelLineStyle = lipgloss.NewStyle()

	// status bar

	statusNugget = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Padding(0, 1)

	statusStyle = lipgloss.NewStyle().
			Inherit(statusBarStyle).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1).
			MarginRight(1)

	statusBarStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
			Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusText = lipgloss.NewStyle().Inherit(statusBarStyle)

	clockStyle = statusNugget.Copy().Background(lipgloss.Color("#6124DF"))
)

type AccMsg struct {
	data acc.AccData
}

type FuelModel struct {
	help     help.Model
	extraLap int
	data     *acc.AccData
}

func NewFuelModel(extraLap int) *FuelModel {
	help := help.New()
	help.ShowAll = true

	return &FuelModel{
		help:     help,
		extraLap: extraLap,
	}
}

func (m *FuelModel) View() string {

	//fmt.Printf("view model %+v\n", m.data)
	doc := strings.Builder{}
	w := lipgloss.Width

	// some default values until acc is running
	fuelUntilEnd := "123.45"
	boxOpen := "67"
	track := "Hometrack"
	sessionState := "unknwon"

	// reset default values with current acc data
	if m.data != nil {
		fuelUntilEnd = strconv.Itoa(m.data.SessionLiter)
		boxOpen = strconv.Itoa(m.data.BoxLap)
		track = m.data.Track
		sessionState = m.data.SessionType.String()
	}

	// main window

	{
		doc.WriteString(headerStyle.Width(mainWindowWidth).Render(fmt.Sprintf("tack ................ %v", track)))
		doc.WriteString("\n")

		sessionEnd := fuelLineStyle.Render("until session end ... ")
		between := fuelLineStyle.Render("   box opens in ... ")
		laps := fuelLineStyle.Render(" laps")
		boxOpenVal := fuelLineStyle.Render(boxOpen)

		literUntilEnd := fuelLineStyle.Copy().
			Width(width - w(sessionEnd) - w(between) - w(boxOpenVal) - w(laps) - 7).
			Render(fuelUntilEnd + " liter ")

		fuelText := lipgloss.JoinHorizontal(lipgloss.Top,
			sessionEnd,
			literUntilEnd,
			between,
			boxOpenVal,
			laps,
		)

		fuelText += fmt.Sprintf("\nextra laps .......... %v", m.extraLap)

		doc.WriteString(windowStyle.Width(mainWindowWidth).Render(fuelText))
		doc.WriteString("\n")
	}

	// status bar

	{
		statusKey := statusStyle.Render("session state")
		clock := clockStyle.Render(time.Now().Format("15:04:05"))
		statusVal := statusText.Copy().
			Width(width - w(statusKey) - w(clock)).
			Render(sessionState)

		bar := lipgloss.JoinHorizontal(lipgloss.Top,
			statusKey,
			statusVal,
			clock,
		)

		doc.WriteString(statusBarStyle.Width(width).Render(bar))
	}

	doc.WriteString("\n")
	doc.WriteString(m.help.View(keys))
	return doc.String()
}

func (m *FuelModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// acc update msg

	case AccMsg:
		d := AccMsg(msg)
		m.data = &d.data
		return m, nil

	// key press msg

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		case key.Matches(msg, keys.Up):
			m.extraLap++
			return m, nil
		case key.Matches(msg, keys.Down):
			m.extraLap--
			if m.extraLap < 0 {
				m.extraLap = 0
			}
			return m, nil
		case key.Matches(msg, keys.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *FuelModel) Init() tea.Cmd {
	return nil
}

func UpdateAcc(data acc.AccData) tea.Msg {
	return AccMsg{
		data: data,
	}
}

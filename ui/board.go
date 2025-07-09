package ui

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/winslowb/onebill-kanban/data"
	"github.com/winslowb/onebill-kanban/model"
)

var columnOrder = []model.TaskStatus{
	model.StatusBacklog,
	model.StatusTodo,
	model.StatusInProgress,
	model.StatusTest,
	model.StatusDone,
}

type mode int

const (
	modeNormal mode = iota
	modeDetails
	modeNewItem
	modeAssignSprint
	modeDashboard  // dashboard mode toggle
	modeFilter     // filter mode
)

type boardModel struct {
	filterInput string
	filterTerm string
	items        map[model.TaskStatus][]model.WorkItem
	columnIndex  int
	rowIndex     int
	columnWidths int
	height       int
	width        int
	sprintInput  string
	mode         mode
	newTitle     string
	typingPos    int
}

func (m boardModel) Init() tea.Cmd {
	return nil
}

func NewBoardModel() (tea.Model, error) {
	all, err := data.LoadAllWorkItems()
	if err != nil {
		return nil, err
	}

	grouped := map[model.TaskStatus][]model.WorkItem{}
	for _, item := range all {
		grouped[item.Status] = append(grouped[item.Status], item)
	}

	return boardModel{
		items:        grouped,
		columnIndex:  0,
		rowIndex:     0,
		columnWidths: 20,
		mode:         modeNormal,
	}, nil
}

func (m boardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch m.mode {
		case modeNormal:
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "esc":
				m.mode = modeNormal
			case "h":
				if m.columnIndex > 0 {
					m.columnIndex--
					m.rowIndex = 0
				}
			case "l":
				if m.columnIndex < len(columnOrder)-1 {
					m.columnIndex++
					m.rowIndex = 0
				}
			case "j":
				column := columnOrder[m.columnIndex]
				if m.rowIndex < len(m.items[column])-1 {
					m.rowIndex++
				}
			case "k":
				if m.rowIndex > 0 {
					m.rowIndex--
				}
			case "enter", "d":
				m.mode = modeDetails
			case "n":
				m.mode = modeNewItem
				m.newTitle = ""
				m.typingPos = 0
			case "s":
				m.mode = modeFilter
				m.filterInput = ""
			case "v":
				m.mode = modeDashboard
			case "/", "f":
					m.mode = modeFilter
					m.filterInput = ""
			case "L":
				column := columnOrder[m.columnIndex]
				items := m.items[column]
				if len(items) > 0 && m.columnIndex < len(columnOrder)-1 {
					item := &items[m.rowIndex]
					item.Status = columnOrder[m.columnIndex+1]
					_ = data.SaveWorkItem(item)
					m.mode = modeNormal
					all, _ := data.LoadAllWorkItems()
					grouped := map[model.TaskStatus][]model.WorkItem{}
					for _, i := range all {
						grouped[i.Status] = append(grouped[i.Status], i)
					}
					m.items = grouped
					m.columnIndex++
					m.rowIndex = 0
				}
			case "H":
				column := columnOrder[m.columnIndex]
				items := m.items[column]
				if len(items) > 0 && m.columnIndex > 0 {
					item := &items[m.rowIndex]
					item.Status = columnOrder[m.columnIndex-1]
					_ = data.SaveWorkItem(item)
					m.mode = modeNormal
					all, _ := data.LoadAllWorkItems()
					grouped := map[model.TaskStatus][]model.WorkItem{}
					for _, i := range all {
						grouped[i.Status] = append(grouped[i.Status], i)
					}
					m.items = grouped
					m.columnIndex--
					m.rowIndex = 0
				}
			}
		case modeDetails:
			if msg.String() == "q" || msg.String() == "esc" {
				m.mode = modeNormal
			}
		case modeNewItem:
			switch msg.String() {
			case "enter":
				newTask := &model.WorkItem{
					Title:     m.newTitle,
					Type:      model.TypeTask,
					Status:    model.StatusBacklog,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}
				_ = data.SaveWorkItem(newTask)
				all, _ := data.LoadAllWorkItems()
				grouped := map[model.TaskStatus][]model.WorkItem{}
				for _, item := range all {
					grouped[item.Status] = append(grouped[item.Status], item)
				}
				m.items = grouped
				m.mode = modeNormal
			case "esc":
				m.mode = modeNormal
			case "backspace":
				if m.typingPos > 0 && len(m.newTitle) > 0 {
					m.newTitle = m.newTitle[:len(m.newTitle)-1]
					m.typingPos--
				}
			default:
				if len(msg.String()) == 1 {
					m.newTitle += msg.String()
					m.typingPos++
				}
			}
		case modeAssignSprint:
			switch msg.String() {
			case "enter":
				column := columnOrder[m.columnIndex]
				items := m.items[column]
				if len(items) > 0 {
					item := items[m.rowIndex]
					item.SprintID = m.sprintInput
					item.UpdatedAt = time.Now()
					_ = data.SaveWorkItem(&item)
				}
				all, _ := data.LoadAllWorkItems()
				grouped := map[model.TaskStatus][]model.WorkItem{}
				for _, i := range all {
					grouped[i.Status] = append(grouped[i.Status], i)
				}
				m.items = grouped
				m.mode = modeNormal
			case "esc":
				m.mode = modeNormal
			case "backspace":
				if len(m.sprintInput) > 0 {
					m.sprintInput = m.sprintInput[:len(m.sprintInput)-1]
				}
			default:
				if len(msg.String()) == 1 {
					m.sprintInput += msg.String()
				}
			}

  		case modeDashboard:
				switch msg.String() {
				case "q", "esc":
					m.mode = modeNormal
				}

		}
	}

	return m, nil
}

func (m boardModel) View() string {
	switch m.mode {
	case modeDetails:
		column := columnOrder[m.columnIndex]
		items := m.items[column]
		if len(items) == 0 {
			return "No item selected."
		}
		item := items[m.rowIndex]
		return renderDetails(item)

	case modeNewItem:
		return renderNewTaskForm(m.newTitle)

	case modeAssignSprint:
		return renderSprintInput(m.sprintInput)

	case modeFilter:
		return renderFilterInput(m.filterInput)

	case modeDashboard:
		return renderDashboard(m)

		default:
			return renderKanbanBoard(m)
	} 
}

func renderSprintInput(input string) string {
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("141")).
		Padding(1, 2).
		Width(50)
	content := fmt.Sprintf("Assign to Sprint\n\nSprint ID: %s\n\n[Enter to Save, Esc to Cancel]", input)
	return box.Render(content)
}



func renderDetails(item model.WorkItem) string {
	return fmt.Sprintf("ID: %s\nTitle: %s\nStatus: %s\nSprint: %s\n", item.ID, item.Title, item.Status, item.SprintID)
}

func renderNewTaskForm(input string) string {
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("36")).
		Padding(1, 2).
		Width(50)
	content := fmt.Sprintf("Create New Task\n\nTitle: %s\n\n[Enter to Save, Esc to Cancel]", input)
	return box.Render(content)
}

func renderKanbanBoard(m boardModel) string {
	columnStyle := lipgloss.NewStyle().
		Width(m.columnWidths).
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240"))

	selectedItemStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("238")).
		Padding(0, 1)

	titleStyle := lipgloss.NewStyle().Bold(true).Underline(true)
	columns := []string{}

	for ci, status := range columnOrder {
		var col strings.Builder
		col.WriteString(titleStyle.Render(string(status)) + "\n\n")
		for ri, item := range m.items[status] {
			if m.filterTerm != "" && 
			!(strings.Contains(item.SprintID, m.filterTerm) || 
			strings.Contains(string(item.Priority), m.filterTerm) || 
			hasTag(item, m.filterTerm)) {
				continue
			}
			text := fmt.Sprintf("• %s", item.Title)
			if ci == m.columnIndex && ri == m.rowIndex {
				text = selectedItemStyle.Render(text)
			}
			col.WriteString(text + "\n")
		}
		columns = append(columns, columnStyle.Render(col.String()))
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, columns...)
}


func renderDashboard(m boardModel) string {
	total := 0
	points := map[model.TaskStatus]int{}
	counts := map[model.TaskStatus]int{}

	for _, status := range columnOrder {
		for _, item := range m.items[status] {
			total++
			counts[status]++
			points[status] += item.Points
		}
	}

	var b strings.Builder
	title := lipgloss.NewStyle().Bold(true).Underline(true).Render
	high := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("212")).Render

	b.WriteString(title("📊 Dashboard Metrics") + "\n\n")
	b.WriteString(fmt.Sprintf("Total Tasks: %s\n\n", high(fmt.Sprintf("%d", total))))

	for _, status := range columnOrder {
		b.WriteString(fmt.Sprintf("%-13s : %2d tasks, %2d pts\n", status, counts[status], points[status]))
	}

	b.WriteString("\n[Press q or esc to return]")
	return b.String()
}



func hasTag(item model.WorkItem, tag string) bool {
	for _, t := range item.Tags {
		if strings.Contains(t, tag) {
			return true
		}
	}
	return false
}



func renderFilterInput(input string) string {
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39")).
		Padding(1, 2).
		Width(50)

	content := fmt.Sprintf("Filter Board\n\nSearch: %s\n\n[Enter to Apply, Esc to Clear]", input)
	return box.Render(content)
}

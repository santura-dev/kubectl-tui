package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	dimStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("243"))
)

func (m Model) View() string {
	if !m.ready {
		return "Initializing..."
	}
	var content string
	switch m.page {
	case screenMenu:
		content = m.menu.View()
	case screenPods:
		content = m.pods.View()
	case screenNamespaces:
		content = m.names.View()
	case screenOutput:
		vp := m.outView
		vp.SetContent(m.output)
		content = vp.View()
	case screenLogs:
		vp := m.outView
		vp.SetContent(m.output)
		content = vp.View()
	}
	return content + "\n" + m.helpLine()
}

func (m Model) helpLine() string {
	switch m.page {
	case screenMenu:
		return dimStyle.Render("↑↓ navigate · enter select · q quit")
	case screenPods:
		return dimStyle.Render("↑↓ navigate · l logs · d describe · q back")
	case screenNamespaces:
		return dimStyle.Render("↑↓ navigate · q back")
	default:
		return dimStyle.Render("q back")
	}
}

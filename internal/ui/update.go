package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/santura-dev/kubectl-tui/internal/kube"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		h := msg.Height - 2
		m.menu.SetSize(msg.Width, h)
		m.pods.SetSize(msg.Width, h)
		m.names.SetSize(msg.Width, h)
		m.outView = viewport.New(msg.Width, h)
		m.ready = true
		return m, nil

	case tea.KeyMsg:
		return m.handleKey(msg)

	case kube.PodsMsg:
		if msg.Err != nil {
			m.output, m.page = msg.Err.Error(), screenOutput
			return m, nil
		}
		items := make([]list.Item, 0, len(msg.Rows))
		for _, r := range msg.Rows {
			items = append(items, podItem{r})
		}
		m.pods.SetItems(items)
		return m, nil

	case kube.NodesMsg:
		if msg.Err != nil {
			m.output, m.page = msg.Err.Error(), screenOutput
			return m, nil
		}
		m.page = screenOutput
		var b strings.Builder
		b.WriteString("NODE                          STATUS     VERSION   GPU\n")
		for _, r := range msg.Rows {
			fmt.Fprintf(&b, "%-28s %-10s %s  %s\n", r.Name, r.Status, r.Version, r.GPUs)
		}
		m.output = b.String()
		return m, nil

	case kube.NamespaceMsg:
		if msg.Err != nil {
			m.output, m.page = msg.Err.Error(), screenOutput
			return m, nil
		}
		items := make([]list.Item, 0, len(msg.Names))
		for _, n := range msg.Names {
			items = append(items, nsItem{n})
		}
		m.names.SetItems(items)
		return m, nil

	case kube.ResultMsg:
		if msg.Err != nil {
			m.output = msg.Err.Error()
		} else {
			m.output = msg.Output
		}
		m.page = screenOutput
		m.outView.SetContent(m.output)
		return m, nil
	}
	return m, m.forward(msg)
}

func (m Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "q", "esc":
		if m.page == screenMenu {
			return m, tea.Quit
		}
		m.page = screenMenu
		return m, nil
	case "enter":
		return m.handleEnter()
	case "l":
		if m.page == screenPods {
			if it, ok := m.pods.SelectedItem().(podItem); ok {
				m.selected = it.row
				return m, kube.Logs(it.row.Namespace, it.row.Name)
			}
		}
	case "d":
		if m.page == screenPods {
			if it, ok := m.pods.SelectedItem().(podItem); ok {
				return m, kube.Describe("pod", it.row.Namespace, it.row.Name)
			}
		}
	}
	return m, m.forward(msg)
}

func (m Model) handleEnter() (tea.Model, tea.Cmd) {
	switch m.page {
	case screenMenu:
		it, ok := m.menu.SelectedItem().(commandItem)
		if !ok {
			return m, nil
		}
		switch it.title {
		case "Pods":
			m.page = screenPods
			return m, kube.ListPods()
		case "Nodes":
			return m, kube.ListNodes()
		case "Namespaces":
			m.page = screenNamespaces
			return m, kube.ListNamespaces()
		default:
			return m, kube.ListGeneric(it.title)
		}
	}
	return m, nil
}

func (m *Model) forward(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	switch m.page {
	case screenMenu:
		m.menu, cmd = m.menu.Update(msg)
	case screenPods:
		m.pods, cmd = m.pods.Update(msg)
	case screenNamespaces:
		m.names, cmd = m.names.Update(msg)
	}
	return cmd
}

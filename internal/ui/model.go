package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/santura-dev/kubectl-tui/internal/kube"
)

// screen is the active page.
type screen int

const (
	screenMenu screen = iota
	screenPods
	screenNodes
	screenNamespaces
	screenOutput
	screenLogs
)

// Model is the bubbletea application model.
type Model struct {
	menu    list.Model
	pods    list.Model
	names   list.Model
	outView viewport.Model

	page     screen
	output   string
	selected kube.PodRow
	ready    bool
	width    int
	height   int
}

// commandItem is a menu entry that maps to a kubectl action.
type commandItem struct {
	title, desc string
}

func (i commandItem) Title() string       { return i.title }
func (i commandItem) Description() string { return i.desc }
func (i commandItem) FilterValue() string { return i.title }

// podItem is one pod row in the pod list.
type podItem struct{ row kube.PodRow }

func (i podItem) Title() string { return i.row.Namespace + "/" + i.row.Name }
func (i podItem) Description() string {
	return i.row.Status + " · node " + i.row.Node
}
func (i podItem) FilterValue() string { return i.row.Name }

// nsItem is one namespace row.
type nsItem struct{ name string }

func (i nsItem) Title() string       { return i.name }
func (i nsItem) Description() string { return "" }
func (i nsItem) FilterValue() string { return i.name }

// New builds the initial model.
func New() Model {
	menu := newList("kubectl", []list.Item{
		commandItem{"Pods", "list pods across namespaces (enter: logs, d: describe)"},
		commandItem{"Nodes", "list nodes with GPU allocation"},
		commandItem{"Namespaces", "list namespaces"},
		commandItem{"Deployments", "list deployments"},
		commandItem{"Services", "list services"},
		commandItem{"Jobs", "list jobs"},
	})
	return Model{
		menu:  menu,
		pods:  newList("Pods", nil),
		names: newList("Namespaces", nil),
		page:  screenMenu,
	}
}

func newList(title string, items []list.Item) list.Model {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = title
	l.SetShowHelp(false)
	l.SetShowStatusBar(false)
	return l
}

// Init starts nothing; the menu drives all fetching.
func (m Model) Init() tea.Cmd { return nil }

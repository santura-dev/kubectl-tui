package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list    list.Model
	output  string
	loading bool
}

func initialModel() model {
	items := []list.Item{
		item{title: "Get Nodes", desc: "List cluster nodes"},
		item{title: "Get Pods", desc: "List all pods"},
		item{title: "Get Deployments", desc: "List deployments"},
		item{title: "Get Services", desc: "List services"},
		item{title: "Get Namespaces", desc: "List namespaces"},
		item{title: "Describe Pod", desc: "Describe a pod (interactive)"},
		item{title: "Logs", desc: "Get logs from a pod"},
		item{title: "Apply YAML", desc: "Apply a YAML file"},
		item{title: "Delete Resource", desc: "Delete a resource"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "kubectl Commands"
	l.Styles.Title = lipgloss.NewStyle().Foreground(lipgloss.Color("170")).Bold(true)

	return model{
		list:    l,
		output:  "Select a command to run.",
		loading: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "esc" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			selected := m.list.SelectedItem().(item)
			m.loading = true
			m.output = "Running: " + selected.title
			return m, runKubectlCmd(selected.title)
		}
	case tea.WindowSizeMsg:
		h, v := lipgloss.NewStyle().GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case outputMsg:
		m.output = string(msg)
		m.loading = false
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.loading {
		return m.list.View() + "\n\nLoading..."
	}
	return m.list.View() + "\n\n" + m.output
}

type outputMsg string

func runKubectlCmd(cmd string) tea.Cmd {
	return func() tea.Msg {
		var command string
		switch cmd {
		case "Get Nodes":
			command = "kubectl get nodes"
		case "Get Pods":
			command = "kubectl get pods -A"
		case "Get Deployments":
			command = "kubectl get deployments -A"
		case "Get Services":
			command = "kubectl get services -A"
		case "Get Namespaces":
			command = "kubectl get namespaces"
		case "Describe Pod":
			command = "kubectl describe pod $(kubectl get pods -o jsonpath='{.items[0].metadata.name}')"
		case "Logs":
			command = "kubectl logs $(kubectl get pods -o jsonpath='{.items[0].metadata.name}')"
		case "Apply YAML":
			command = "kubectl apply -f /home/sandra/inference/nginx-k8s.yaml"
		case "Delete Resource":
			command = "kubectl delete -f /home/sandra/inference/nginx-k8s.yaml"
		default:
			return outputMsg("Command not implemented")
		}

		out, err := exec.Command("bash", "-c", command).Output()
		if err != nil {
			return outputMsg("Error: " + err.Error())
		}
		return outputMsg(string(out))
	}
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

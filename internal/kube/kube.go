// Package kube runs kubectl commands asynchronously for the TUI.
package kube

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const cmdTimeout = 15 * time.Second

func run(args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), cmdTimeout)
	defer cancel()
	out, err := exec.CommandContext(ctx, "kubectl", args...).CombinedOutput()
	if err != nil {
		msg := strings.TrimSpace(string(out))
		if msg == "" {
			msg = err.Error()
		}
		return "", fmt.Errorf("kubectl %s: %s", strings.Join(args, " "), msg)
	}
	return strings.TrimSpace(string(out)), nil
}

// ResultMsg carries output or error of a kubectl command.
type ResultMsg struct {
	Output string
	Err    error
}

// PodsMsg carries the pod list.
type PodsMsg struct {
	Rows []PodRow
	Err  error
}

// PodRow is one line of the pod table.
type PodRow struct {
	Namespace, Name, Status, Node string
}

// ListPods fetches pods across all namespaces.
func ListPods() tea.Cmd {
	return func() tea.Msg {
		out, err := run("get", "pods", "-A", "--no-headers")
		if err != nil {
			return PodsMsg{Err: err}
		}
		var rows []PodRow
		for _, line := range strings.Split(out, "\n") {
			f := strings.Fields(line)
			if len(f) >= 5 {
				rows = append(rows, PodRow{f[0], f[1], f[3], f[len(f)-1]})
			}
		}
		return PodsMsg{Rows: rows}
	}
}

// NodesMsg carries the node list with GPU info.
type NodesMsg struct {
	Rows []NodeRow
	Err  error
}

// NodeRow is one line of the node table.
type NodeRow struct {
	Name, Status, Version string
	GPUs                  string
}

// ListNodes fetches nodes with their GPU count.
func ListNodes() tea.Cmd {
	return func() tea.Msg {
		out, err := run("get", "nodes", "--no-headers",
			"-o", "custom-columns=NAME:.metadata.name,STATUS:.status.conditions[-1].type,VERSION:.status.nodeInfo.kubeletVersion,GPU:.status.allocatable['nvidia\\.com/gpu']")
		if err != nil {
			return NodesMsg{Err: err}
		}
		var rows []NodeRow
		for _, line := range strings.Split(out, "\n") {
			f := strings.Fields(line)
			if len(f) >= 3 {
				gpus := "-"
				if len(f) >= 4 {
					gpus = f[3]
				}
				rows = append(rows, NodeRow{f[0], f[1], f[2], gpus})
			}
		}
		return NodesMsg{Rows: rows}
	}
}

// Logs fetches the last 100 log lines of a pod.
func Logs(namespace, pod string) tea.Cmd {
	return func() tea.Msg {
		out, err := run("logs", "-n", namespace, pod, "--tail=100")
		return ResultMsg{Output: out, Err: err}
	}
}

// Describe fetches the description of a resource.
func Describe(kind, namespace, name string) tea.Cmd {
	args := []string{"describe", kind}
	if namespace != "" {
		args = append(args, "-n", namespace)
	}
	args = append(args, name)
	return func() tea.Msg {
		out, err := run(args...)
		return ResultMsg{Output: out, Err: err}
	}
}

// ListGeneric fetches any resource kind as a plain table.
func ListGeneric(kind string) tea.Cmd {
	return func() tea.Msg {
		out, err := run("get", kind, "-A")
		return ResultMsg{Output: out, Err: err}
	}
}

// NamespaceMsg carries the namespace list.
type NamespaceMsg struct {
	Names []string
	Err   error
}

// ListNamespaces fetches all namespace names.
func ListNamespaces() tea.Cmd {
	return func() tea.Msg {
		out, err := run("get", "namespaces", "--no-headers")
		if err != nil {
			return NamespaceMsg{Err: err}
		}
		var names []string
		for _, line := range strings.Split(out, "\n") {
			if f := strings.Fields(line); len(f) > 0 {
				names = append(names, f[0])
			}
		}
		return NamespaceMsg{Names: names}
	}
}

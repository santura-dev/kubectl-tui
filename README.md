# kubectl TUI

Interactive terminal user interface for Kubernetes operations built with Go and Bubbletea.

## Features

- **Menu-Driven Interface**: Navigate kubectl commands through intuitive menus
- **Real-Time Execution**: Execute commands and view streaming output
- **Multi-Level Navigation**: Drill down into deployments, pods, and services
- **Context-Sensitive Help**: Dynamic help text for each screen
- **Selection-Based Input**: Choose from available resources instead of typing

## Screenshots

[Add screenshots here]

## Prerequisites

- Go 1.19+
- kubectl configured for your cluster

## Installation

```bash
git clone https://github.com/yourusername/kubectl-tui.git
cd kubectl-tui
go mod tidy
go build -o kubectl-tui
```

## Usage

```bash
./kubectl-tui
```

Navigate through menus to:
- Get nodes, pods, deployments
- Describe resources
- View logs
- Execute common kubectl operations

## Controls

- **↑↓**: Navigate menu items
- **Enter**: Select item
- **b**: Go back
- **q/Ctrl+C**: Quit

## Architecture

Built with:
- [Bubbletea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - UI components
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling

## Development

```bash
# Run in development
go run main.go

# Build for production
go build -o kubectl-tui -ldflags="-s -w"
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

MIT License - see LICENSE file for details

## Related Projects

- [vllm-tui](https://github.com/yourusername/vllm-tui) - vLLM chat interface
- [local-inference-operator](https://github.com/yourusername/local-inference-operator) - K8s operator
- [operator-tui](https://github.com/yourusername/operator-tui) - Operator management
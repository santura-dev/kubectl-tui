# kubectl-tui

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white) ![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg) ![Kubernetes](https://img.shields.io/badge/kubernetes-compatible-blue?logo=kubernetes)

Terminal UI for Kubernetes cluster inspection with a focus on inference workloads. Like `kubectl get` but live, navigable, and GPU-aware.

## The problem

`kubectl` is the right tool for scripting and automation. It is the wrong tool for exploration and debugging. When a pod is crash-looping, you want the pod, its events, its logs, and the node it is running on, all in one view, updating live. `kubectl get pods && kubectl describe pod && kubectl logs` is three commands and three context switches. This is one.

## The idea

A TUI tuned for clusters running inference workloads. GPU resource visibility is first-class: how much VRAM each pod is using, which nodes have GPUs, whether they are allocated. Common inference patterns (vLLM pods, model servers, operator CRDs) get special handling in the UI.

## Features

- **resource viewing**: pods, deployments, services, jobs, namespaces — one navigable list
- **GPU visibility**: node view includes `nvidia.com/gpu` allocatable capacity
- **pod context**: logs (`l`) and full describe (`d`) on any pod, across namespaces
- **async everything**: kubectl runs off the UI thread with timeouts; the TUI never freezes

## Key bindings

| key | action |
|---|---|
| `↑`/`↓` | navigate |
| `enter` | open / run |
| `l` | pod logs |
| `d` | describe pod |
| `q` / `esc` | back / quit |

## Install

```bash
go install github.com/santura-dev/kubectl-tui@latest
```

## Related

- [inference-operator-tui](https://github.com/santura-dev/inference-operator-tui) - TUI for local-inference-operator CRDs specifically
- [vllm-logprob-tui](https://github.com/santura-dev/vllm-logprob-tui) - TUI for vLLM logprob inspection

## License

MIT

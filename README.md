# kubectl-tui

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=flat&logo=go&logoColor=white) ![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg) ![Kubernetes](https://img.shields.io/badge/kubernetes-compatible-blue?logo=kubernetes)

Terminal UI for Kubernetes cluster inspection with a focus on inference workloads. Like `kubectl get` but live, navigable, and GPU-aware.

## The problem

`kubectl` is the right tool for scripting and automation. It is the wrong tool for exploration and debugging. When a pod is crash-looping, you want the pod, its events, its logs, and the node it is running on, all in one view, updating live. `kubectl get pods && kubectl describe pod && kubectl logs` is three commands and three context switches. This is one.

## The idea

A TUI tuned for clusters running inference workloads. GPU resource visibility is first-class: how much VRAM each pod is using, which nodes have GPUs, whether they are allocated. Common inference patterns (vLLM pods, model servers, operator CRDs) get special handling in the UI.

## Features

- **resource viewing**: pods, jobs, deployments across namespaces, one navigable list
- **GPU visibility**: per-pod and per-node GPU allocation. Which pods consume VRAM, which nodes have capacity.
- **log streaming**: live log tailing without `kubectl logs -f`
- **port-forward shortcuts**: one keypress to port-forward to a pod's API
- **filtering**: by label, namespace, or node. Useful when you have 200 pods and care about the vLLM ones.
- **resource usage**: CPU, memory, and GPU metrics per pod

## Key bindings

| key | action |
|---|---|
| `j`/`k` | navigate |
| `n` | switch namespace |
| `f` | filter by label |
| `l` | logs |
| `p` | port-forward |
| `g` | GPU view |
| `q` | quit |

## Install

```bash
go install github.com/santura-dev/kubectl-tui@latest
```

## Related

- [inference-operator-tui](https://github.com/santura-dev/inference-operator-tui) - TUI for local-inference-operator CRDs specifically
- [vllm-logprob-tui](https://github.com/santura-dev/vllm-logprob-tui) - TUI for vLLM logprob inspection

## License

MIT

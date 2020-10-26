# k3ai-core

K3ai-core is the core library for the k3ai installer.
The Go installer will replace the current bash installer.

[![Github build status](https://github.com/kf5i/k3ai-core/workflows/build/badge.svg)](https://github.com/kf5i/k3ai-core/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/kf5i/k3ai-core)](https://goreportcard.com/report/github.com/kf5i/k3ai-core)
[![codecov](https://codecov.io/gh/kf5i/k3ai-core/branch/main/graph/badge.svg)](https://codecov.io/gh/kf5i/k3ai-core)

## Building from source

Setup the environment using Golang v1.15.3+. A Linux, Mac OS or a WSL2 environment is recommended.

To build the project, run

```bash
make build-cli
```

To run the test suite, use

```bash
make lint
make test
```

Please feel free to open a Github issue or send a PR. Looking forward to your contribution.

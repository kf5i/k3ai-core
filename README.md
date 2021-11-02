# WE RELEASED THE NEW VERSION CHECK THE NEW ORG AT: [https://github.com/k3ai](https://github.com/k3ai) AND OUR NEW WEBSITE/DOCS AT: [https://k3ai.in](https://k3ai.in)

# k3ai-core

K3ai-core is the core library for the k3ai installer.
The Go installer will replace the current bash installer.

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/kf5i/k3ai-core/build?style=for-the-badge)
![Codecov](https://img.shields.io/codecov/c/github/kf5i/k3ai-core?style=for-the-badge)
![GitHub all releases](https://img.shields.io/github/downloads/kf5i/k3ai-core/total?style=for-the-badge)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/kf5i/k3ai-core?style=for-the-badge)
![GitHub contributors](https://img.shields.io/github/contributors/kf5i/k3ai-core?style=for-the-badge)


## Install k3ai-cli(Latest Version)

```bash
#Set a variable to grab latest version
Version=$(curl -s "https://api.github.com/repos/kf5i/k3ai-core/releases/latest" | awk -F '"' '/tag_name/{print $4}' | cut -c 2-6) 
# get the binaries
wget https://github.com/kf5i/k3ai-core/releases/download/v$Version/k3ai-core_${Version}_linux_amd64.tar.gz
```


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

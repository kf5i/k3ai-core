#!/bin/bash
nohup k3s server --write-kubeconfig-mode 644 --disable-agent > /dev/null 2>&1 &

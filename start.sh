#!/bin/bash
sudo bash -ic 'k3s server --write-kubeconfig-mode 644 > /dev/null 2>&1 &'

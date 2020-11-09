package kctl

import (
	"log"
	"time"
)

// Wait processes requests for resource readiness
type Wait interface {
	Process(config Config, namespace string, labels []string) error
}

// CliWait waits when a user is interacting directly with the client CLI.
type CliWait struct {
	// Todo implement timeout
	timeout time.Duration
}

const waitCmd = "wait"

// Process runs the wait command for each label
func (w *CliWait) Process(config Config, namespace string, labels []string) error {
	for _, label := range labels {
		cmd, args := prepareCommand(config, waitCmd, "-n", namespace, "--for", "condition=ready", "po", "-l", label)
		log.Printf("Waiting for %s", label)
		if err := execute(config, cmd, args...); err != nil {
			return err
		}
	}
	return nil
}

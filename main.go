package main

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"

	"github.com/opencontainers/runtime-spec/specs-go"
)

func main() {
	logFile, err := os.OpenFile("/home/tesla59/Code/podman-apparmor-hook/apparmor.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(1)
	}
	defer logFile.Close()

	logger := slog.New(slog.NewTextHandler(logFile, nil))
	slog.SetDefault(logger)

	slog.Info("Starting OCI hook execution")

	configJSON, err := io.ReadAll(os.Stdin)
	if err != nil {
		slog.Error("Error reading from stdin", "error", err)
		os.Exit(1)
	}

	var config specs.Spec
	if err := json.Unmarshal(configJSON, &config); err != nil {
		slog.Error("Error unmarshaling config JSON", "error", err)
		os.Exit(1)
	}

	slog.Info("Container configuration received",
		"rootfs", config.Root.Path,
		"process", config.Process != nil,
		"AppArmor", config.Process.ApparmorProfile)

	config.Process.ApparmorProfile = "kubearmor-aa-profile"
	
	if err := json.NewEncoder(os.Stdout).Encode(config); err != nil {
		slog.Error("Error encoding config JSON", "error", err)
		os.Exit(1)
	}

	slog.Info("OCI hook execution completed successfully")
}

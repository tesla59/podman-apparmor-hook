package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/opencontainers/runtime-spec/specs-go"
)

func main() {
	configJSON, err := io.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}
	var config specs.Spec
	if err := json.Unmarshal(configJSON, &config); err != nil {
		os.Exit(1)
	}
	config.Process.ApparmorProfile = "kubearmor-aa-profile"
	if err := json.NewEncoder(os.Stdout).Encode(config); err != nil {
		os.Exit(1)
	}
}

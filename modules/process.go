package modules

import (
	"os/exec"
)

type ProcessSpawner struct{}

func init() {
	Register(&ProcessSpawner{})
}

func (p *ProcessSpawner) Name() string        { return "proc_spawn" }
func (p *ProcessSpawner) Description() string { return "Spawns a suspicious shell command chain" }

func (p *ProcessSpawner) Generate(target string, port string) error {
	cmd := exec.Command("sh", "-c", "echo 'Telemetry Payload Executed'")
	output, err := cmd.CombinedOutput()
	if err == nil {
		println(string(output))
	}
	return err
}

func (p *ProcessSpawner) Cleanup() error { return nil }

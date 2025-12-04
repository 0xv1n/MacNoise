package modules

import (
	"fmt"
	"sort"
)

type Generator interface {
	Name() string
	Description() string
	Generate(target string, port string) error
	Cleanup() error
}

var registry = make(map[string]Generator)

func Register(g Generator) {
	registry[g.Name()] = g
}

func GetGenerator(name string) (Generator, bool) {
	g, ok := registry[name]
	return g, ok
}

func ListModules() {
	keys := make([]string, 0, len(registry))
	for k := range registry {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Println("Available Telemetry Modules:")
	for _, k := range keys {
		fmt.Printf("  %-20s : %s\n", k, registry[k].Description())
	}
}

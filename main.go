package main

import (
	"flag"
	"fmt"
	"os"

	"0xv1n/macnoise/modules"
)

func main() {
	moduleFlag := flag.String("module", "", "The specific module to run")
	listFlag := flag.Bool("list", false, "List available modules")
	targetFlag := flag.String("target", "127.0.0.1", "Target IP/Host for network modules")
	portFlag := flag.String("port", "8080", "Target Port for network modules")

	flag.Parse()

	if *listFlag {
		modules.ListModules()
		os.Exit(0)
	}

	if *moduleFlag == "" {
		fmt.Println("Error: Please provide a -module name")
		fmt.Println("Run with -list to see options")
		os.Exit(1)
	}

	gen, exists := modules.GetGenerator(*moduleFlag)
	if !exists {
		fmt.Printf("Error: Module '%s' not found.\n", *moduleFlag)
		os.Exit(1)
	}

	fmt.Printf("--- executing %s ---\n", gen.Name())

	if err := gen.Generate(*targetFlag, *portFlag); err != nil {
		fmt.Printf("[!] Error generating telemetry: %v\n", err)
	} else {
		fmt.Printf("[+] Telemetry generation complete.\n")
	}

	// TODO: Cleanup tasks kinda like ART
	// defer gen.Cleanup()
}

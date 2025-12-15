# gsh

![go](https://img.shields.io/badge/-go_1.25.1-00ADD8?style=flat&logo=go&logoColor=white)
[![License](https://img.shields.io/badge/license-GPL-green.svg)](LICENSE)

**gsh** is a lightweight, custom shell implemented from scratch in Go. It provides a simple REPL (Read-Eval-Print Loop) to execute commands, handle input/output, and spawn processes, aiming to give developers a clear understanding of shell internals while offering a functional CLI environment.

## Features

- **Custom Command Parser**: Built-in tokenizer and parser for command processing
- **Built-in Commands**: Essential shell commands implemented natively
  - `echo` - Display text
  - `pwd` - Print working directory
  - `cd` - Change directory
  - `ls` - List directory contents
  - `clear` - Clear the terminal screen
  - `whoami` - Display current user
  - `help` - Show available commands
  - `exit` - Exit the shell
- **External Command Execution**: Execute system commands from `/bin/`
- **Configurable Environment**: Customizable prompt, theme, and color mode
- **Command History**: Track command history throughout the session
- **Environment Variables**: Manage shell environment variables
- **Color Support**: Optional colored output for enhanced user experience

## Prerequisites

- Go 1.25.1 or higher
- Linux/Unix-based operating system

## Installation

### From Source

```bash
# Clone the repository
git clone <repository-url>
cd gsh

# Build the project
go build -o bin/gsh cmd/gsh/main.go

# Run the shell
./bin/gsh
```

### Using Go Install

```bash
go install cmd/gsh/main.go
```

## Usage

Start the shell by running:

```bash
./bin/gsh
```

You'll be greeted with a prompt that shows your current directory:

```
/home/user/workspace
gsh> 
```

## Project Structure

```
gsh/
├── cmd/
│   └── gsh/
│       └── main.go           # Entry point
├── internal/
│   ├── builtins/             # Built-in command implementations
│   │   ├── builtins.go       # Command registry
│   │   ├── cd.go
│   │   ├── clear.go
│   │   ├── echo.go
│   │   ├── exit.go
│   │   ├── help.go
│   │   ├── ls.go
│   │   ├── pwd.go
│   │   └── whoami.go
│   ├── command/              # Command parsing
│   │   ├── command.go
│   │   └── parser.go
│   ├── config/               # Configuration management
│   │   ├── config.go
│   │   └── deafult.go
│   └── shell/                # Core shell implementation
│       ├── deafult.go
│       ├── repl.go           # REPL loop
│       └── shell.go          # Shell structure
├── plugins/                  # Plugin system (extensible)
├── bin/                      # Compiled binaries
├── go.mod                    # Go module definition
└── README.md
```

## Configuration

The shell can be configured through the `Config` structure:

```go
type Config struct {
    Prompt      string    // Custom prompt string
    Theme       string    // Color theme
    ColorMode   bool      // Enable/disable colors
    HistoryFile string    // Command history file path
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/Feature`)
3. Commit your changes (`git commit -m 'Add some Feature'`)
4. Push to the branch (`git push origin feature/Feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Roadmap

- [ ] Plugin system implementation
- [ ] Command piping support (`|`)
- [ ] Input/Output redirection (`>`, `<`, `>>`)
- [ ] Background job execution (`&`)
- [ ] Tab completion
- [ ] Advanced history management with search
- [ ] Configuration file support
- [ ] Alias support
- [ ] Scripting capabilities


## Acknowledgments

- Inspired by traditional Unix shells (bash, zsh, sh)
- Built as an educational project to understand shell internals


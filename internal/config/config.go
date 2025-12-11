package config

import "os"

type Config struct {
	Prompt    string
	Theme     string
	ColorMode bool
	s         uintptr
	stdin     *os.File
	stdout    *os.File
	stderr    *os.File

	HistoryFile string
}

// getters
func (cnf *Config) Stdin() *os.File {
	return cnf.stdin
}

func (cnf *Config) Stdout() *os.File {
	return cnf.stdout
}

func (cnf *Config) Stderr() *os.File {
	return cnf.stderr
}

func SetDefault[typ any](field *typ, val *typ, def typ) {
	if val != nil {
		*field = *val
	} else {
		*field = def
	}
}

func (cfg *Config) NewConfig(
	prompt *string,
	theme *string,
	colorMode *bool,
	stdin *os.File,
	stdout *os.File,
	stderr *os.File,
	historyFile *string,
) {
	// Theme
	SetDefault(&cfg.Prompt, prompt, DefaultPrompt)
	SetDefault(&cfg.Theme, theme, "")
	SetDefault(&cfg.ColorMode, colorMode, true)

	// File descriptors
	SetDefault(&cfg.stdin, &stdin, os.Stdin)
	SetDefault(&cfg.stdout, &stdout, os.Stdout)
	SetDefault(&cfg.stderr, &stderr, os.Stderr)

	// History
	SetDefault(&cfg.HistoryFile, historyFile, "")
}

func NewDeafultConfig() *Config {
	return &Config{
		Prompt:      DefaultPrompt,
		Theme:       "",
		ColorMode:   true,
		stdin:       os.NewFile(DefaultStdin, "/dev/stdin"),
		stdout:      os.NewFile(DefaultStdout, "/dev/stdout"),
		stderr:      os.NewFile(DefaultStderr, "/dev/stderr"),
		HistoryFile: DefaultHistoryFile,
	}

}

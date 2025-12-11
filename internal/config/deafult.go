package config

const (
	//  Limits
	DefaultHistorySize = 1000
	MaxHistorySize     = 100000
	DefaultMaxJobs     = 10
	DefaultTimeout     = 30 // seconds

	//  Prompts
	DefaultPrompt = "gsh> "

	//  Colors (ANSI codes)
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"

	//  Config Filenames
	DefaultConfigFile  = ".gshrc"
	DefaultHistoryFile = "history"
)

var (
	DefaultStdin  = uintptr(0)
	DefaultStdout = uintptr(1)
	DefaultStderr = uintptr(2)
)

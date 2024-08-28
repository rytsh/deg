package config

// Build contains the build information
var Build = struct {
	Version string
	Commit  string
	Date    string
}{
	Version: "v0.0.0",
	Commit:  "?",
	Date:    "",
}

var Prog = struct {
	LogLevel string
}{
	LogLevel: "info",
}

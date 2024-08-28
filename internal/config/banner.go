package config

import "fmt"

var banner = `_________
______  /_____ ______
_  __  /_  __\_  __  /
/ /_/ / /  __/  /_/ /
\____/  \___/ \__  /
             /____/`

func Banner() string {
	return banner
}

func Info() string {
	return fmt.Sprintf("[%s] buildCommit=[%s] buildDate=[%s]",
		Build.Version, Build.Commit, Build.Date,
	)
}

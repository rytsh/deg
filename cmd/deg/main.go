package main

import (
	"context"

	"github.com/rytsh/deg/internal/config"

	"github.com/rakunlabs/into"
	"github.com/rakunlabs/logi"
	"github.com/rytsh/deg/cmd/deg/args"
)

var (
	version = "v0.0.0"
	commit  = "?"
	date    = ""
)

func main() {
	config.Build.Version = version
	config.Build.Date = date
	config.Build.Commit = commit

	into.Init(
		run,
		into.WithStartFn(nil), into.WithStopFn(nil),
		into.WithLogger(logi.InitializeLog(logi.WithCaller(false))),
	)
}

func run(ctx context.Context) error {
	return args.Execute(ctx)
}

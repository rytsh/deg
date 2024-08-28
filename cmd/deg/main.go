package main

import (
	"context"
	"sync"

	"github.com/rytsh/deg/internal/config"
	"github.com/worldline-go/initializer"

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

	initializer.Init(
		run,
		initializer.WithInitLog(false),
		initializer.WithLogger(initializer.Slog),
		initializer.WithOptionsLogi(logi.WithCaller(false)),
	)
}

func run(ctx context.Context, _ *sync.WaitGroup) error {
	return args.Execute(ctx)
}

package args

import (
	"context"
	"fmt"
	"os"

	"github.com/rytsh/deg/internal/config"
	"github.com/rytsh/deg/internal/repl"
	"github.com/spf13/cobra"
)

var replCmd = &cobra.Command{
	Use:   "repl",
	Short: "repl mode",
	Long:  "Provides a REPL (Read-Eval-Print-Loop) interface for deg.",
	RunE: func(cmd *cobra.Command, _ []string) error {
		return runRepl(cmd.Context())
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

var replFlags = struct {
	NoBanner bool
}{}

func init() {
	replCmd.Flags().BoolVarP(&replFlags.NoBanner, "no-banner", "b", false, "hide banner")
}

func runRepl(_ context.Context) error {
	if !replFlags.NoBanner {
		fmt.Println("deg " + config.Info())
	}

	repl.Start(os.Stdin, os.Stdout)

	return nil
}

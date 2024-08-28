package args

import (
	"context"
	"fmt"

	"github.com/rakunlabs/logi"
	"github.com/rytsh/deg/internal/config"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "deg",
	Short: "programing language",
	Long:  config.Banner() + "\ndeg programing language",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		if err := logi.SetLogLevel(config.Prog.LogLevel); err != nil {
			return err //nolint:wrapcheck // no need
		}

		return nil
	},
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, _ []string) error {
		if err := runRoot(cmd.Context()); err != nil {
			return err
		}

		return nil
	},
}

func Execute(ctx context.Context) error {
	setFlags()

	rootCmd.Version = config.Build.Version
	rootCmd.Long = fmt.Sprintf(
		"%s\nversion:[%s] commit:[%s] buildDate:[%s]",
		rootCmd.Long, config.Build.Version, config.Build.Commit, config.Build.Date,
	)

	rootCmd.AddCommand(replCmd)

	return rootCmd.ExecuteContext(ctx) //nolint:wrapcheck // no need
}

func setFlags() {
	rootCmd.PersistentFlags().StringVarP(&config.Prog.LogLevel, "log-level", "l", config.Prog.LogLevel, "log level")
}

func runRoot(ctx context.Context) error {
	fmt.Println("deg " + config.Info())

	return nil
}

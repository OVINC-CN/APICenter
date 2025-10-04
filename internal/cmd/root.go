package cmd

import (
	"log"
	"os"

	"github.com/ovinc-cn/apicenter/v2/internal/crond"
	"github.com/ovinc-cn/apicenter/v2/internal/version"
	"github.com/ovinc-cn/apicenter/v2/pkg/trace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "api-center",
	Short:   "API Center",
	Version: version.Version,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && cmd.CalledAs() == "api-center" {
			appMode := args[0]
			args = args[1:]
			switch appMode {
			case "api":
				apiCmd.Run(apiCmd, args)
			case "worker":
				workerCmd.Run(workerCmd, args)
			case "beat":
				beatCmd.Run(beatCmd, args)
			case "migrate":
				migrateCmd.Run(migrateCmd, args)
			default:
				log.Fatal("[CMD] unknown app mode")
			}
		} else {
			log.Fatal("[CMD] please specify app mode")
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		_ = crond.AsyncClient.Close()
		trace.OnExit()
		os.Exit(0)
	},
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetVersionTemplate(`{{printf "API Center "}}{{printf "%s" .Version}}`)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("[CMD] execute failed; %s", err)
	}
}

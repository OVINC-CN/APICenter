package cmd

import (
	"github.com/ovinc-cn/apicenter/v2/internal/crond/worker"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Api Center Worker",
	Run: func(cmd *cobra.Command, args []string) {
		worker.Serve()
	},
}

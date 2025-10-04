package cmd

import (
	"github.com/ovinc-cn/apicenter/v2/internal/crond/beat"
	"github.com/spf13/cobra"
)

var beatCmd = &cobra.Command{
	Use:   "beat",
	Short: "Api Center Scheduler",
	Run: func(cmd *cobra.Command, args []string) {
		beat.Serve()
	},
}

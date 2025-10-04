package cmd

import (
	"github.com/ovinc-cn/apicenter/v2/internal/httpd"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Api Center HttpServer",
	Run: func(cmd *cobra.Command, args []string) {
		httpd.Serve()
	},
}

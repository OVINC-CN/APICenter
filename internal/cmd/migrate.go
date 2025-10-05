package cmd

import (
	"log"

	"github.com/ovinc-cn/apicenter/v2/internal/apps/account"
	"github.com/ovinc-cn/apicenter/v2/pkg/mysql"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrator",
	Short: "Api Center Migrator",
	Run: func(cmd *cobra.Command, args []string) {
		if err := mysql.DB().AutoMigrate(&account.User{}); err != nil {
			log.Fatalf("[Migrate] auto migrate failed; %s", err)
		}
		log.Println("[Migrate] auto migrate success")
	},
}

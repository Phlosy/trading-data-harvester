// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trading-data-harvester",
	Short: "Trading Data Harvester 工具",
	Long:  `Trading Data Harvester 是一个用于获取交易数据并存储到数据库的命令行工具`,
}

// Execute 执行根命令
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

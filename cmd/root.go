package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iam-saml",
	Short: "IAM saml implementation in Golang",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Wooot")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

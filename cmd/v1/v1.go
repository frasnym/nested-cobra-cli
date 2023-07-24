/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package v1

import (
	"fmt"

	"github.com/frasnym/nested-cobra-cli/cmd/v1/transaction"
	"github.com/spf13/cobra"
)

// V1Cmd represents the v1 command
var V1Cmd = &cobra.Command{
	Use:   "v1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1 called")
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// v1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// v1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	V1Cmd.AddCommand(transaction.TransactionCmd)

}

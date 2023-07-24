/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package transaction

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	transactionReportDate string
)

// reportCmd represents the report command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("report called")

		date, _ := time.Parse("2006-01-02", transactionReportDate)
		generateTransactionReport(date)
	},
}

func init() {
	reportCmd.Flags().StringVarP(&transactionReportDate, "date", "d", "", "report date")
	reportCmd.MarkFlagRequired("date")

	TransactionCmd.AddCommand(reportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// reportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateTransactionReport(date time.Time) {
	fmt.Printf("generateTransactionReport called: %v\n", date)
}

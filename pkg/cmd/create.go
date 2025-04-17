package cmd

import (
	"github.com/spf13/cobra"
)

var (
	disputeID  string
	txnID      string
	merchantID string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new payment dispute",
	Run: func(cmd *cobra.Command, args []string) {
		DisputeSvc.CreateDispute(disputeID, txnID, merchantID)
	},
}

func init() {
	createCmd.Flags().StringVar(&disputeID, "id", "", "Dispute ID (required)")
	createCmd.Flags().StringVar(&txnID, "txn", "", "Transaction ID (required)")
	createCmd.Flags().StringVar(&merchantID, "merchant", "", "Merchant ID (required)")

	createCmd.MarkFlagRequired("id")
	createCmd.MarkFlagRequired("txn")
	createCmd.MarkFlagRequired("merchant")
}

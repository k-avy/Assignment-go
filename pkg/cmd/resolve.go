package cmd

import (
	"github.com/spf13/cobra"
)

var (
	resolveID string
	evidence  string
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve an existing dispute",
	Run: func(cmd *cobra.Command, args []string) {
		DisputeSvc.ResolveDispute(resolveID, evidence)
	},
}

func init() {
	resolveCmd.Flags().StringVar(&resolveID, "id", "", "Dispute ID to resolve (required)")
	resolveCmd.Flags().StringVar(&evidence, "evidence", "", "Path to evidence file (required)")

	resolveCmd.MarkFlagRequired("id")
	resolveCmd.MarkFlagRequired("evidence")
}

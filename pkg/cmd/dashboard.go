package cmd

import (
	"github.com/k-avy/Assignment-go/pkg/services"
	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "View dispute resolution analytics and agent performance",
	Run: func(cmd *cobra.Command, args []string) {
		services.ShowDashboard(AnalyticsSvc, DisputeSvc)
	},
}

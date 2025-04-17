package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/k-avy/Assignment-go/pkg/auth"
	"github.com/k-avy/Assignment-go/pkg/models"
	"github.com/k-avy/Assignment-go/pkg/services"
	"github.com/spf13/cobra"
)

var (
	DisputeSvc   *services.DisputeService
	AnalyticsSvc *services.AnalyticsService
	AgentSvc     *services.AgentService
	username     string
	password     string
)

var rootCmd = &cobra.Command{
	Use:   "dispute-cli",
	Short: "A CLI tool to manage payment disputes",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !auth.Authenticate(username, password) {
			fmt.Println("Exiting due to failed authentication.")
			os.Exit(1)
		}
	},
}

func InitializeServices() {
	// Initialize once and share across commands
	AgentSvc = services.NewAgentService()
	DisputeSvc = services.NewDisputeService()
	AnalyticsSvc = &services.AnalyticsService{
		DisputeSvc: DisputeSvc,
		AgentSvc:   AgentSvc,
	}
}

func Execute() {
	// Initialize services once
	InitializeServices()

	// Add commands
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(resolveCmd)
	rootCmd.AddCommand(dashboardCmd)

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		erro := models.ErrorFound{
			Message:   "Can not execute the command",
			Error:     err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		}
		fmt.Println(erro)
		os.Exit(1)
	}
}

func init() {
	// Set up persistent flags for authentication
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username for authentication (required)")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Password for authentication (required)")
	rootCmd.MarkPersistentFlagRequired("username")
	rootCmd.MarkPersistentFlagRequired("password")
}

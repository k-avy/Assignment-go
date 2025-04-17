package services

import (
	"fmt"
	"time"

	"github.com/k-avy/Assignment-go/pkg/models"
)

// ShowDashboard displays the current analytics on disputes and agent resolution performance
func ShowDashboard(as *AnalyticsService, ds *DisputeService) {
	as.mu.Lock()
	defer as.mu.Unlock()

	total := len(ds.Disputes)
	open := 0
	for _, d := range ds.Disputes {
		if d.Status == models.Open {
			open++
		}
	}

	resolved := total - open
	avgResolution := time.Duration(0)
	if len(as.ResolutionDurations) > 0 {
		for _, d := range as.ResolutionDurations {
			avgResolution += d
		}
		avgResolution /= time.Duration(len(as.ResolutionDurations))
	}

	fmt.Println("🧾 Dispute Resolution Dashboard")
	fmt.Println("--------------------------------------------------")
	fmt.Printf("📌 Total Disputes   : %d\n", total)
	fmt.Printf("✅ Resolved         : %d\n", resolved)
	fmt.Printf("🕐 Open             : %d\n", open)
	fmt.Printf("📊 Avg Resolution   : %s\n", avgResolution)
	fmt.Println("--------------------------------------------------")
	fmt.Println("👩‍💼 Agent Performance:")

	for agentID, count := range as.AgentSuccess {
		fmt.Printf("👤 Agent %s → Resolved: %d\n", agentID, count)
	}
	fmt.Println("--------------------------------------------------")
}

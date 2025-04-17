package services

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/k-avy/Assignment-go/pkg/models"
)

const analyticsFile = "analytics.json"

type AnalyticsService struct {
	mu                  sync.Mutex
	DisputeSvc          *DisputeService
	AgentSvc            *AgentService
	ResolvedCount       int
	ResolutionDurations []time.Duration
	AgentSuccess        map[string]int
}

func NewAnalyticsService(ds *DisputeService, ag *AgentService) *AnalyticsService {
	as := &AnalyticsService{
		DisputeSvc:          ds,
		AgentSvc:            ag,
		ResolvedCount:       0,
		ResolutionDurations: []time.Duration{},
		AgentSuccess:        make(map[string]int),
	}
	as.loadFromFile()
	return as
}

func (as *AnalyticsService) loadFromFile() {
	file, err := os.ReadFile(analyticsFile)
	if err != nil {
		fmt.Println("No analytics data found. Starting fresh.")
		return
	}
	_ = json.Unmarshal(file, as)
}

func (as *AnalyticsService) saveToFile() {
	data, err := json.MarshalIndent(as, "", "  ")
	if err != nil {
		fmt.Println("Failed to save analytics:", err)
		return
	}
	_ = os.WriteFile(analyticsFile, data, 0644)
}

func (as *AnalyticsService) UpdateOnResolve(dispute models.Dispute) {
	as.mu.Lock()
	defer as.mu.Unlock()

	if dispute.ResolvedAt == nil {
		return
	}

	duration := dispute.ResolvedAt.Sub(dispute.CreatedAt)
	as.ResolvedCount++
	as.ResolutionDurations = append(as.ResolutionDurations, duration)

	if dispute.AssignedAgent != "" {
		as.AgentSuccess[dispute.AssignedAgent]++
	}
	as.saveToFile()
}

func (as *AnalyticsService) PrintDashboard() {
	as.mu.Lock()
	defer as.mu.Unlock()

	total := len(as.DisputeSvc.Disputes)
	open := 0
	for _, d := range as.DisputeSvc.Disputes {
		if d.Status == models.Open {
			open++
		}
	}

	avgDuration := time.Duration(0)
	if len(as.ResolutionDurations) > 0 {
		for _, d := range as.ResolutionDurations {
			avgDuration += d
		}
		avgDuration /= time.Duration(len(as.ResolutionDurations))
	}

	fmt.Println("Dispute Dashboard")
	fmt.Printf("Total Disputes: %d\n", total)
	fmt.Printf("Resolved: %d | Open: %d\n", as.ResolvedCount, open)
	fmt.Printf("Avg Resolution Time: %s\n", avgDuration)
	fmt.Println("Agent Performance:")
	for agent, count := range as.AgentSuccess {
		fmt.Printf("  %s → %d resolved\n", agent, count)
	}
}

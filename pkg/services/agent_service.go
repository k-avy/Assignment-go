package services

import (
	"github.com/k-avy/Assignment-go/pkg/models"
)

type AgentService struct {
	Agents map[string]*models.Agent
}

func NewAgentService() *AgentService {
	return &AgentService{
		Agents: map[string]*models.Agent{
			"A1": {ID: "A1", Name: "Alice", Expertise: "Retail"},
			"A2": {ID: "A2", Name: "Bob", Expertise: "Travel"},
		},
	}
}

func (as *AgentService) AssignAgent() *models.Agent {
	var selected *models.Agent
	for _, agent := range as.Agents {
		if selected == nil || agent.Load < selected.Load {
			selected = agent
		}
	}
	return selected
}

func (as *AgentService) GetAgent(id string) *models.Agent {
	return as.Agents[id]
}

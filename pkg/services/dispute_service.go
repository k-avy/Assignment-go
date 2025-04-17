package services

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/k-avy/Assignment-go/pkg/models"
)

const disputeFile = "disputes.json"

type DisputeService struct {
	Disputes map[string]models.Dispute
	mu       sync.Mutex
}

func NewDisputeService() *DisputeService {
	ds := &DisputeService{
		Disputes: make(map[string]models.Dispute),
	}
	ds.loadFromFile()
	return ds
}

func (ds *DisputeService) loadFromFile() {
	file, err := os.ReadFile(disputeFile)
	if err != nil {
		fmt.Println("No existing dispute data found. Starting fresh.")
		return
	}
	json.Unmarshal(file, &ds.Disputes)
}

func (ds *DisputeService) saveToFile() {
	data, err := json.MarshalIndent(ds.Disputes, "", "  ")
	if err != nil {
		fmt.Println("Failed to serialize disputes:", err)
		return
	}
	err = os.WriteFile(disputeFile, data, 0644)
	if err != nil {
		fmt.Println("Failed to save disputes to file:", err)
	}
}

func (ds *DisputeService) CreateDispute(id, txnID, merchantID string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, exists := ds.Disputes[id]; exists {
		fmt.Println("Dispute already exists with this ID.")
		return
	}

	dispute := models.Dispute{
		ID:            id,
		TransactionID: txnID,
		MerchantID:    merchantID,
		CreatedAt:     time.Now(),
		Status:        models.Open,
	}

	dispute.AssignedAgent = "Auto-Assigned"
	ds.Disputes[id] = dispute
	ds.saveToFile()
	fmt.Printf("Dispute %s created.\n", id)
}

func (ds *DisputeService) ResolveDispute(id, evidenceFile string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	dispute, exists := ds.Disputes[id]
	if !exists {
		fmt.Println("Dispute not found.")
		return
	}

	if dispute.Status != models.Open {
		fmt.Println("Dispute already resolved or marked duplicate.")
		return
	}

	if _, err := os.Stat(evidenceFile); os.IsNotExist(err) {
		fmt.Println("Evidence file not found.")
		return
	}

	now := time.Now()
	dispute.Status = models.Resolved
	dispute.ResolvedAt = &now
	dispute.EvidencePath = evidenceFile

	ds.Disputes[id] = dispute
	ds.saveToFile()

	fmt.Printf("Dispute %s resolved.\n", id)
}

func (ds *DisputeService) GetDispute(id string) (models.Dispute, bool) {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	dispute, exists := ds.Disputes[id]
	return dispute, exists
}

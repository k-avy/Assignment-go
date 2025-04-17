package models

import "time"

type DisputeStatus string

const (
	Open      DisputeStatus = "Open"
	Resolved  DisputeStatus = "Resolved"
	Duplicate DisputeStatus = "Duplicate"
)

type Dispute struct {
	ID            string        `json:"id"`
	TransactionID string        `json:"transaction_id"`
	MerchantID    string        `json:"merchant_id"`
	CreatedAt     time.Time     `json:"created_at"`
	Status        DisputeStatus `json:"status"`
	AssignedAgent string        `json:"assigned_agent"`
	EvidencePath  string        `json:"evidence_path"`
	ResolvedAt    *time.Time    `json:"resolved_at,omitempty"`
}

package interfaces

import "ccusage-rainbow/internal/domain/entities"

// CostService defines the interface for fetching cost data
type CostService interface {
	// FetchCostData fetches cost data from ccusage command
	FetchCostData() (*entities.CostResponse, error)
}

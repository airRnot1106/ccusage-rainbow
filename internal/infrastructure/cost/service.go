package cost

import (
	"ccusage-rainbow/internal/domain/entities"
	"encoding/json"
	"os/exec"
)

// Service implements the CostService interface
type Service struct{}

// NewService creates a new cost service
func NewService() *Service {
	return &Service{}
}

// FetchCostData fetches cost data from ccusage command
func (s *Service) FetchCostData() (*entities.CostResponse, error) {
	// Execute the ccusage command with -j flag
	cmd := exec.Command("npx", "ccusage@latest", "-j")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Parse the JSON output
	var costResponse entities.CostResponse
	if err := json.Unmarshal(output, &costResponse); err != nil {
		return nil, err
	}

	return &costResponse, nil
}

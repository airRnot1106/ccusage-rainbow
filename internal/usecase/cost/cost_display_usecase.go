package cost

import (
	"ccusage-rainbow/internal/domain/entities"
	"ccusage-rainbow/internal/domain/interfaces"
)

// CostDisplayUseCase handles the business logic for fetching and displaying cost data
type CostDisplayUseCase struct {
	costService interfaces.CostService
}

// NewCostDisplayUseCase creates a new CostDisplayUseCase
func NewCostDisplayUseCase(costService interfaces.CostService) *CostDisplayUseCase {
	return &CostDisplayUseCase{
		costService: costService,
	}
}

// GetCostText fetches cost data and returns formatted text for display
func (uc *CostDisplayUseCase) GetCostText() (*entities.Text, error) {
	costData, err := uc.costService.FetchCostData()
	if err != nil {
		// Return error text if fetching fails
		return entities.NewText("ERROR"), err
	}

	// Format the total cost as text
	costText := costData.Totals.FormatCost()

	return entities.NewText(costText), nil
}

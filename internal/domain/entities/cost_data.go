package entities

import "fmt"

// CostResponse represents the response from ccusage API
type CostResponse struct {
	Daily  []DailyUsage `json:"daily"`
	Totals Totals       `json:"totals"`
}

// DailyUsage represents daily usage statistics
type DailyUsage struct {
	Date                string           `json:"date"`
	InputTokens         int              `json:"inputTokens"`
	OutputTokens        int              `json:"outputTokens"`
	CacheCreationTokens int              `json:"cacheCreationTokens"`
	CacheReadTokens     int              `json:"cacheReadTokens"`
	TotalTokens         int              `json:"totalTokens"`
	TotalCost           float64          `json:"totalCost"`
	ModelsUsed          []string         `json:"modelsUsed"`
	ModelBreakdowns     []ModelBreakdown `json:"modelBreakdowns"`
}

// ModelBreakdown represents model-specific usage breakdown
type ModelBreakdown struct {
	ModelName           string  `json:"modelName"`
	InputTokens         int     `json:"inputTokens"`
	OutputTokens        int     `json:"outputTokens"`
	CacheCreationTokens int     `json:"cacheCreationTokens"`
	CacheReadTokens     int     `json:"cacheReadTokens"`
	Cost                float64 `json:"cost"`
}

// Totals represents total usage statistics
type Totals struct {
	InputTokens         int     `json:"inputTokens"`
	OutputTokens        int     `json:"outputTokens"`
	CacheCreationTokens int     `json:"cacheCreationTokens"`
	CacheReadTokens     int     `json:"cacheReadTokens"`
	TotalTokens         int     `json:"totalTokens"`
	TotalCost           float64 `json:"totalCost"`
}

// FormatCost formats the total cost as a string for display
func (t *Totals) FormatCost() string {
	return fmt.Sprintf("$%.2f", t.TotalCost)
}

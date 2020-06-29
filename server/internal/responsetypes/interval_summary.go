package responsetypes

// IntervalSum provides a structure for the statistics that get graphed in Phobos
type IntervalSum struct {
	Interval         string  `json:"interval"`
	Duration         float64 `json:"duration"`
	Miles            float64 `json:"miles"`
	SortIndex        int     `json:"sortIndex"`
	PercentageActive float64 `json:"percentage_active"`
}

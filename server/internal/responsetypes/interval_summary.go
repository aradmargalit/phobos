package responsetypes

// IntervalSum provides a structure for the statistics that get graphed in Phobos
type IntervalSum struct {
	Interval    string  `json:"interval"`
	SortIndex   int     `json:"sort_index"`
	Duration    float64 `json:"duration"`
	Miles       float64 `json:"miles"`
	DaysSkipped float64 `json:"days_skipped"`
}

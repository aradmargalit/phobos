package responsetypes

// Stats is a grouping of statistics about a user and their activities
type Stats struct {
	Workouts      int            `json:"workouts"`
	Hours         float64        `json:"hours"`
	Miles         float64        `json:"miles"`
	LastTen       []float64      `json:"last_ten"`
	TypeBreakdown []TypePortion  `json:"type_breakdown"`
	DayBreakdown  []DayBreakdown `json:"day_breakdown"`
}

// TypePortion represents an activity type and it's portion relative to others
type TypePortion struct {
	Name    string `json:"name"`
	Portion int    `json:"portion"`
}

// DayBreakdown represents the day of the week and how many activities have fallen on that day
type DayBreakdown struct {
	DOW   string `json:"day_of_week"`
	Count int    `json:"count"`
}

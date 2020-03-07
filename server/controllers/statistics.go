package controllers

// GetUserStatistics returns some fun user statistics for the frontend
func (e *Env) GetUserStatistics(c *gin.Context) {
	// Pull user out of context to figure out which activities to grab
	uid, ok := c.Get("user")
	if !ok {
		panic("No user id in cookie!")
	}

	a, err := e.DB.GetActivitiesByUser(uid.(int))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	monthMap := map[string]float64{}

	for _, activity := range a {
		m, _ := time.Parse("2006-01-02 15:04:05", activity.ActivityDate)
		month := fmt.Sprintf("%v %v", m.Month(), m.Year())
		_, ok := monthMap[month]
		if !ok {
			monthMap[month] = 0
		}
		monthMap[month] += activity.Duration
	}

	type monthlySum struct {
		Month    string  `json:"month"`
		Duration float64 `json:"duration"`
	}

	response := []monthlySum{}
	for k, v := range monthMap {
		response = append(response, monthlySum{k, v})
	}
	c.JSON(http.StatusOK, response)
}
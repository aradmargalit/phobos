package models

// StravaWebhookEvent represents the payload we get from Strava when the webhook is triggered
// Technically, they send more than this, but as of today, we don't care
type StravaWebhookEvent struct {
	ObjectType     string `json:"object_type"`
	ObjectID       int    `json:"object_id"`
	AspectType     string `json:"aspect_type"`
	OwnerID        int    `json:"owner_id"`
	EventTime      int    `json:"event_time"`
	SubscriptionID int    `json:"subscription_id"`
}

// StravaActivity represents the activity in Strava's terms
type StravaActivity struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Distance    float64 `json:"distance"`
	ElapsedTime int     `json:"elapsed_time"`
	MovingTime  int     `json:"moving_time"`
	Type        string  `json:"type"`
	StartDate   string  `json:"start_date"`
	Timezone    string  `json:"timezone"`
}

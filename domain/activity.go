package domain

type Activity struct {
	Page         string                 `json:"page" dynamo:"page"`
	Timestamp    int64                  `json:"timestamp" dynamo:"timestamp"`
	ActivityType string                 `json:"activityType" dynamo:"activityType"`
	IpAddress    string                 `json:"ipAddress" dynamo:"ipAddress"`
	UserAgent    string                 `json:"userAgent" dynamo:"userAgent"`
	Data         map[string]interface{} `json:"data" dynamo:"data"`
}

package domain

type Activity struct {
	Page         string                 `json:"page" dynamodbav:"page"`
	Timestamp    int64                  `json:"timestamp" dynamodbav:"timestamp"`
	ActivityType string                 `json:"activityType" dynamodbav:"activityType"`
	IpAddress    string                 `json:"ipAddress" dynamodbav:"ipAddress"`
	UserAgent    string                 `json:"userAgent" dynamodbav:"userAgent"`
	Data         map[string]interface{} `json:"data" dynamodbav:"data"`
}

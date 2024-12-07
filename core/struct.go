package core

import "time"

type RequestLog struct {
	Id                 string    `json:"id"`
	JobId              string    `json:"jobId"`
	Status             string    `json:"status"`
	RequestUrl         string    `json:"requestUrl"`
	RequestMethod      string    `json:"requestMethod"`
	ResponseStatusCode int       `json:"responseStatusCode"`
	CreatedAt          time.Time `json:"createdAt"`
}

type FullRequestLog struct {
	RequestLog
	RequestHeaders  map[string][]string `json:"requestHeaders"`
	RequestBody     string              `json:"requestBody"`
	ResponseHeaders map[string][]string `json:"responseHeaders"`
	ResponseBody    string              `json:"responseBody"`
}

type RequestTypePayload struct {
	Status             string
	RequestUrl         string
	RequestMethod      string
	RequestHeaders     map[string][]string
	RequestBody        string
	ResponseHeaders    map[string][]string
	ResponseBody       string
	ResponseStatusCode int
}

type Trigger struct {
	Name    string            `json:"name"`
	Payload map[string]string `json:"payload"`
}

type FullTrigger struct {
	Trigger
	JobId string `json:"jobId"`
}

type Action struct {
	Name    string                 `json:"name"`
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

type FullAction struct {
	Action
	JobId string `json:"jobId"`
}

type Job struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Author      string   `json:"author"`
	Members     []string `json:"members"`
}

type FullJob struct {
	Job
	JobID     string    `json:"jobId"`
	CreatedAt time.Time `json:"createdAt"`
}

type FullAuthorization struct {
	Name    string      `json:"name"`
	Payload interface{} `json:"payload"`
}

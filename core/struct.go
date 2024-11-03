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

type FullTrigger struct {
	JobId   string            `json:"jobId"`
	Name    string            `json:"name"`
	Payload map[string]string `json:"payload"`
}

type FullAction struct {
	JobId   string                 `json:"jobId"`
	Name    string                 `json:"name"`
	Payload map[string]interface{} `json:"payload"`
}

type FullJob struct {
	JobID       string    `json:"jobId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Members     []string  `json:"members"`
	CreatedAt   time.Time `json:"createdAt"`
}

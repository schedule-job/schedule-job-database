package core

type Database interface {
	InsertAction(job_id, name string, payload map[string]interface{}) error
	UpdateAction(job_id, name string, payload map[string]interface{}) error
	DeleteAction(job_id string) error
	SelectAction(job_id string) (*FullAction, error)
	SelectActions() (*[]FullAction, error)
	SelectIdsByAction() ([]string, error)

	InsertJob(name, description, author string, members []string) (string, error)
	UpdateJob(job_id, name, description, author string, members []string) error
	DeleteJob(job_id string) error
	SelectJob(job_id string) (*FullJob, error)
	SelectJobs(user, last_id string, limit int) (*[]FullJob, error)

	InsertRequestLog(job_id string, payload interface{}) error
	SelectRequestLog(id, job_id string) (*FullRequestLog, error)
	SelectRequestLogs(job_id, last_id string, limit int) ([]RequestLog, error)

	InsertTrigger(job_id, name string, payload map[string]interface{}) error
	UpdateTrigger(job_id, name string, payload map[string]interface{}) error
	DeleteTrigger(job_id string) error
	SelectTrigger(job_id string) (*FullTrigger, error)
}

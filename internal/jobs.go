package internal

import (
	"go.uber.org/zap"
)

// FindJobs interface provides basic methods needed to be able to write jobs to a file
type FindJobs interface {
	GetURL() string
	GetJobs(*zap.Logger) []Job
	GetPath() string
}

// Job contains key information about the job
type Job struct {
	Title        string
	Type         string
	Salary       string
	Location     string
	URL          string
	Requirements []string
}

// JobSource contains the location to scrape, a list of jobs from that source and a file path to write the jobs
// TODO: Consider renaming/splitting this, its a bit convoluted
type JobSource struct {
	URL      string
	Jobs     []Job
	FilePath string
}

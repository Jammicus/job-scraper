package internal

import "go.uber.org/zap"

// Think about how to reduce the size of this interface
type FindJobs interface {
	GetURL() string
	GetJobs(*zap.Logger) []Job
	GetPath() string
}

type Job struct {
	Title        string
	Type         string
	Salary       string
	Location     string
	URL          string
	Requirements []string
}

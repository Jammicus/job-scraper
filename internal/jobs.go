package internal

import "go.uber.org/zap"

// Each site should implement this
// *Logger
type FindJobs interface {
	Scrape(*zap.Logger) ([]Job, error)
}

type Job struct {
	Title        string
	Type         string
	Salary       string
	Location     string
	URL          string
	Requirements []string
}

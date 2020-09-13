package companies

import (
	"encoding/json"
	"io/ioutil"
	jobs "job-scraper/internal"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

type amazonAPI struct {
	// Number of results returned
	Hits int `json:"hits"`
	Jobs []amazonAPIJob
}

type amazonAPIJob struct {
	BasicQuals string `json:"basic_qualifications"`
	// Eg full-time
	Schedule     string `json:"job_schedule_type"`
	Location     string `json:"location"`
	PerferedQual string `json:"preferred_qualifications"`
	Title        string `json:"title"`
	JobPath      string `json:"job_path"`
}

// Amazon implements Jobsource
type Amazon jobs.JobSource

// GetJobs gets all jobs if it has not already found the jobs
func (a Amazon) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(a.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		a.findJobs(logger)
	}
	return a.Jobs
}

// GetPath returns the file path specified for a given reciever
func (a Amazon) GetPath() string {
	return a.FilePath
}

func (a *Amazon) findJobs(logger *zap.Logger) {
	var aAPI amazonAPI

	sugar := logger.Sugar()
	sugar.Infof("Querying Amazon API for jobs")
	resp, err := http.Get(a.URL)

	if err != nil {
		sugar.Errorf("Error querying Amazon API for jobs %v", zap.Error(err))
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(responseData, &aAPI)

	a.Jobs = a.gatherSpecs(aAPI, logger)

	if err != nil {
		sugar.Errorf("Error getting amazon specs", zap.Error(err))
	}
}

func (a Amazon) gatherSpecs(aAPI amazonAPI, logger *zap.Logger) []jobs.Job {
	foundJobs := []jobs.Job{}
	for _, item := range aAPI.Jobs {
		job := jobs.Job{}
		job.Title = a.getJobTitle(item)
		job.Type = a.getJobType(item)
		job.URL = a.getJobURL(item)
		job.Requirements = a.getJobRequirements(item)
		job.Salary = a.getJobSalary(item)
		job.Location = a.getJobLocation(item)

		foundJobs = append(foundJobs, job)
	}

	return foundJobs
}

func (a Amazon) getJobURL(job amazonAPIJob) string {
	return "www.amazon.jobs" + job.JobPath
}

func (a Amazon) getJobTitle(job amazonAPIJob) string {
	return job.Title
}

func (a Amazon) getJobType(job amazonAPIJob) string {
	return job.Schedule
}

func (a Amazon) getJobLocation(job amazonAPIJob) string {
	return job.Location
}

func (a Amazon) getJobSalary(job amazonAPIJob) string {
	return "N/A"
}

func (a Amazon) getJobRequirements(job amazonAPIJob) []string {
	// TODO: Remove black elements
	return append(strings.Split(job.BasicQuals, "<br/>"), strings.Split(job.PerferedQual, "<br/>")...)
}

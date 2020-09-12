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
	Jobs []struct {
		BasicQuals string `json:"basic_qualifications"`
		// Eg full-time
		Schedule     string `json:"job_schedule_type"`
		Location     string `json:"location"`
		PerferedQual string `json:"preferred_qualifications"`
		Title        string `json:"title"`
		JobPath      string `json:"job_path"`
	}
}

type Amazon jobs.JobSource

func (a Amazon) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(a.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		a.findJobs(logger)
	}
	return a.Jobs
}

func (a Amazon) GetURL() string {
	return a.URL
}

func (a Amazon) GetPath() string {
	return a.FilePath
}

func (a *Amazon) findJobs(logger *zap.Logger) {
	var aAPI amazonAPI

	sugar := logger.Sugar()
	sugar.Infof("Querying Amazon API for jobs")
	resp, err := http.Get(a.GetURL())

	if err != nil {
		sugar.Errorf("Error querying Amazon API for jobs %v", zap.Error(err))
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(responseData, &aAPI)

	a.Jobs, err = a.gatherSpecs(aAPI, logger)

	if err != nil {
		sugar.Errorf("Error getting amazon specs", zap.Error(err))
	}
}

func (a Amazon) gatherSpecs(aAPI amazonAPI, logger *zap.Logger) ([]jobs.Job, error) {
	foundJobs := []jobs.Job{}
	for _, item := range aAPI.Jobs {
		job := jobs.Job{}
		job.Title = item.Title
		job.Type = item.Schedule
		job.URL = "www.amazon.jobs" + item.JobPath
		// TODO split into own method.
		job.Requirements = append(strings.Split(item.BasicQuals, "<br/>"), strings.Split(item.PerferedQual, "<br/>")...)
		job.Salary = "N/A"
		job.Location = item.Location

		foundJobs = append(foundJobs, job)

	}

	return foundJobs, nil
}

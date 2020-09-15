package recruiters

import (
	"fmt"
	jobs "job-scraper/internal"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

// Understanding is a JobSource
type Understanding jobs.JobSource

// GetJobs returns all jobs  for a given receiver
func (u Understanding) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(u.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		u.findJobs(logger)
	}
	return u.Jobs
}

// GetPath returns the filepath to write the CSV to for a given receiver
func (u Understanding) GetPath() string {
	return u.FilePath
}

func (u *Understanding) findJobs(logger *zap.Logger) {
	sugar := logger.Sugar()
	foundJobs := []jobs.Job{}

	c := colly.NewCollector(
		colly.Async(true),
	)

	err := jobs.IsUp(u.URL)
	if err != nil {
		sugar.Fatal(err)
	}

	c.OnHTML("li.job-result-item", func(e *colly.HTMLElement) {
		e.ForEach("span.job-read-more-link", func(_ int, el *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
			sugar.Infof("Looking for jobs at: %v", link)
			job, err := u.gatherSpecs(link, logger)
			if err != nil {
				sugar.Error(zap.Error(err))
			}

			sugar.Infof("Job successfully scraped with title: %v", job.Title)
			foundJobs = append(foundJobs, job)

		})
	})
	c.OnHTML(".next", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Debugf("Next page link found: %v", link)
		e.Request.Visit(link)

	})

	c.Visit(u.URL)
	c.Wait()

	u.Jobs = foundJobs

}

func (u Understanding) gatherSpecs(url string, logger *zap.Logger) (jobs.Job, error) {

	sugar := logger.Sugar()

	job := jobs.Job{}

	d := colly.NewCollector(
		colly.Async(true),
	)

	d.Visit(url)
	d.OnRequest(func(r *colly.Request) {
		url := u.getJobURL(r)
		sugar.Debugf("Visiting page %v", url)
		job.URL = url
	})

	d.OnHTML("h1", func(e *colly.HTMLElement) {
		job.Title = u.getJobTitle(e)

	})

	d.OnHTML("ul.job-details-list.clearfix", func(e *colly.HTMLElement) {
		var err error

		job.Location, err = u.getJobLocation(e)
		if err != nil {
			sugar.Errorf("Could not get job location %v", zap.Error(err))
		}

		job.Salary, err = u.getJobSalary(e)
		if err != nil {
			sugar.Errorf("Could not get job salary %v", zap.Error(err))
		}

		job.Type, err = u.getJobType(e)
		if err != nil {
			sugar.Errorf("Could not get job type %v", zap.Error(err))
		}
	})

	d.OnHTML("article.clearfix.mb20", func(e *colly.HTMLElement) {
		var err error

		job.Requirements, err = u.getJobRequirements(e)
		if err != nil {
			log.Fatal(err)
		}
	})

	d.Wait()

	if len(job.Requirements) == 0 {
		return jobs.Job{}, fmt.Errorf("No requirements found for job %v", job.URL)
	}

	sugar.Debugf("Job details found %v", job)
	return job, nil
}

func (u Understanding) getJobType(e *colly.HTMLElement) (string, error) {
	jobTypeReplacer := strings.NewReplacer("Job type:", "", "\n", "", "\t", "")
	jobType := e.ChildText("li.job_type")

	return jobTypeReplacer.Replace(jobType), nil
}

func (u Understanding) getJobLocation(e *colly.HTMLElement) (string, error) {
	locationReplacer := strings.NewReplacer("Location", "", "\n", "", "\t", "")
	location := e.ChildText("li.location")

	return locationReplacer.Replace(location), nil
}

func (u Understanding) getJobSalary(e *colly.HTMLElement) (string, error) {

	salaryReplacer := strings.NewReplacer("Salary:", "", "\n", "", "\t", "")
	salary := e.ChildText("li.job_salary")
	return salaryReplacer.Replace(salary), nil
}

func (u Understanding) getJobRequirements(e *colly.HTMLElement) ([]string, error) {
	requirements := []string{}
	e.ForEach("li", func(_ int, el *colly.HTMLElement) {
		// requirements one at a time
		requirements = append(requirements, el.Text)
	})
	return requirements, nil
}

func (u Understanding) getJobTitle(e *colly.HTMLElement) string {

	return e.Text
}

func (u Understanding) getJobURL(r *colly.Request) string {

	return r.URL.String()
}

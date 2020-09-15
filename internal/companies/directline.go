package companies

import (
	"fmt"
	jobs "job-scraper/internal"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

// DirectLine is a JobSource
type DirectLine jobs.JobSource

// GetJobs returns all jobs  for a given receiver
func (dr DirectLine) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(dr.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		dr.findJobs(logger)
	}
	return dr.Jobs
}

// GetPath returns the filepath to write the CSV to for a given receiver
func (dr DirectLine) GetPath() string {
	return dr.FilePath
}

func (dr *DirectLine) findJobs(logger *zap.Logger) {

	sugar := logger.Sugar()
	foundJobs := []jobs.Job{}

	c := colly.NewCollector(
		colly.Async(true),
	)
	// As there is duplicate a.icon-right-ico.pagination__right-nav on the same page, we only want to expand it once.
	c.AllowURLRevisit = false

	err := jobs.IsUp(dr.URL)
	if err != nil {
		sugar.Fatal(err)
	}

	c.OnHTML("div.card.grid__item", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Infof("Looking for jobs at: %v", link)
		job, err := dr.gatherSpecs(link, logger)

		if err != nil {
			sugar.Error(zap.Error(err))
		}
		sugar.Infof("Job successfully scraped with title: %v", job.Title)
		foundJobs = append(foundJobs, job)
	})

	c.OnHTML("a.icon-right-ico.pagination__right-nav", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		sugar.Debugf("Next page link found: %v", link)
		e.Request.Visit(link)
	})

	c.Visit(dr.URL)
	c.Wait()

	dr.Jobs = foundJobs

}

func (dr DirectLine) gatherSpecs(url string, logger *zap.Logger) (jobs.Job, error) {
	var err error
	sugar := logger.Sugar()

	job := jobs.Job{}

	d := colly.NewCollector(
		colly.Async(true),
	)
	d.Limit(&colly.LimitRule{
		Parallelism: 1,
		RandomDelay: 5 * time.Second,
	})
	d.Visit(url)

	d.OnRequest(func(r *colly.Request) {

		url := dr.getJobURL(r)
		sugar.Debugf("Visiting page %v", url)
		job.URL = url
	})

	d.OnHTML("div.job-container", func(e *colly.HTMLElement) {

		job.Requirements, err = dr.getJobRequirements(e)
		if err != nil {
			sugar.Errorf("Error geting job requiremnts %v", zap.Error(err))
		}

	})
	d.OnHTML("div.location.map-ico-black", func(e *colly.HTMLElement) {
		job.Location = dr.getJobLocation(e)

	})

	d.OnHTML("div.top", func(e *colly.HTMLElement) {

		job.Title = dr.getJobTitle(e)

	})

	// Not returning job type
	d.OnHTML("div.hero-box__details", func(e *colly.HTMLElement) {
		job.Type, err = dr.getJobType(e)
		if err != nil {
			sugar.Errorf("Error geting job type %v", zap.Error(err))
		}
	})

	// Doesnt tell us salary on their job pages.
	job.Salary = dr.getJobSalary(nil)
	d.Wait()

	if len(job.Requirements) == 0 {
		return jobs.Job{}, fmt.Errorf("No requirements found for job %v", job.URL)
	}

	sugar.Debugf("Job details found %v", job)

	return job, nil
}

func (dr DirectLine) getJobRequirements(e *colly.HTMLElement) ([]string, error) {
	requirements := []string{}
	e.ForEach("li", func(_ int, el *colly.HTMLElement) {
		requirements = append(requirements, el.Text)
	})

	return requirements, nil
}

func (dr DirectLine) getJobLocation(e *colly.HTMLElement) string {

	return e.ChildText("span")
}

func (dr DirectLine) getJobType(e *colly.HTMLElement) (string, error) {

	jobType := ""
	e.ForEach("div.copy", func(i int, el *colly.HTMLElement) {
		if i < 1 {

			jobType = jobType + el.Text
		}
	})

	return jobType, nil
}

func (dr DirectLine) getJobURL(r *colly.Request) string {
	return r.URL.String()
}

func (dr DirectLine) getJobTitle(e *colly.HTMLElement) string {
	return strings.TrimSpace(e.Text)
}

func (dr DirectLine) getJobSalary(e *colly.HTMLElement) string {
	return "N/A"
}

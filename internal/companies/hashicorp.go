package companies

import (
	jobs "job-scraper/internal"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

type Hashicorp jobs.JobSource

func (h Hashicorp) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(h.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		h.findJobs(logger)
	}
	return h.Jobs
}

func (h Hashicorp) GetPath() string {
	return h.FilePath
}

func (h *Hashicorp) findJobs(logger *zap.Logger) {
	sugar := logger.Sugar()
	foundJobs := []jobs.Job{}

	c := colly.NewCollector(
		colly.Async(true),
	)

	err := jobs.IsUp(h.URL)
	if err != nil {
		sugar.Fatal(err)
	}

	c.OnHTML(".item", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Infof("Looking for jobs at: %v", link)
		job, err := h.gatherSpecs(link, logger)

		if err != nil {
			sugar.Error(zap.Error(err))
		}
		sugar.Infof("Job successfully scraped with title: %v", job.Title)
		foundJobs = append(foundJobs, job)
	})

	c.Visit(h.URL)
	c.Wait()

	h.Jobs = foundJobs
}

func (h Hashicorp) gatherSpecs(url string, logger *zap.Logger) (jobs.Job, error) {

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
		url := h.getJobURL(r)
		sugar.Debugf("Visiting page %v", url)
		job.URL = url
	})

	d.OnHTML("h1", func(e *colly.HTMLElement) {
		job.Title = h.getJobTitle(e)
	})

	d.OnHTML(".g-type-label.location", func(e *colly.HTMLElement) {
		job.Location = h.getJobLocation(e)

	})

	d.OnHTML(".application-details", func(e *colly.HTMLElement) {

		job.Requirements = h.getJobRequirements(e)

	})

	// Not specified on the pages.
	job.Salary = h.getJobSalary(nil)
	job.Type = h.getJobType(nil)
	d.Wait()

	sugar.Debugf("Job details found %v", job)

	return job, nil
}

func (h Hashicorp) getJobRequirements(e *colly.HTMLElement) []string {
	requirements := []string{}
	e.ForEach("li", func(_ int, el *colly.HTMLElement) {
		requirements = append(requirements, el.Text)
	})

	return requirements
}

func (h Hashicorp) getJobLocation(e *colly.HTMLElement) string {

	return strings.TrimSpace(e.Text)
}

func (h Hashicorp) getJobType(e *colly.HTMLElement) string {

	return "Permanent"
}

func (h Hashicorp) getJobURL(r *colly.Request) string {
	return r.URL.String()
}

func (h Hashicorp) getJobTitle(e *colly.HTMLElement) string {
	return strings.TrimSpace(e.Text)
}

func (h Hashicorp) getJobSalary(e *colly.HTMLElement) string {
	return "N/A"
}

package companies

import (
	"fmt"
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

	sugar.Debugf("Site %v is accessible", h.URL)

	c.OnHTML(".item", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Infof("Looking for jobs at: %v", link)
		job, err := h.gatherSpecs(link, logger)

		if err != nil {
			sugar.Error(zap.Error(err))
		}
		sugar.Infof("Job successfully scraped at: %v", link)
		foundJobs = append(foundJobs, job)
	})

	c.Visit(h.URL)
	c.Wait()

	h.Jobs = foundJobs
}

func (h Hashicorp) gatherSpecs(url string, logger *zap.Logger) (jobs.Job, error) {

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
		sugar.Debugf("Visiting page %v", r.URL.String())
		job.URL = r.URL.String()
	})

	d.OnHTML("h1", func(e *colly.HTMLElement) {
		job.Title = strings.TrimSpace(e.Text)
	})

	d.OnHTML(".g-type-label.location", func(e *colly.HTMLElement) {
		job.Location = strings.TrimSpace(e.Text)

	})

	d.OnHTML(".application-details", func(e *colly.HTMLElement) {

		job.Requirements, err = h.getRequirements(e)
		if err != nil {
			sugar.Errorf("Error geting job requiremnts %v", zap.Error(err))
		}

	})

	// Not specified on the pages.
	job.Salary = "N/A"
	job.Type = "Permanent"
	d.Wait()
	return job, nil
}

func (h Hashicorp) getRequirements(e *colly.HTMLElement) ([]string, error) {
	requirements := []string{}
	e.ForEach("li", func(_ int, el *colly.HTMLElement) {
		requirements = append(requirements, el.Text)
	})

	fmt.Println(requirements)

	return requirements, nil
}

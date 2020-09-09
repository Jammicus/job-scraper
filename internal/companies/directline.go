package companies

import (
	jobs "job-scraper/internal"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

type DirectLine jobs.JobSource

func (dr DirectLine) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(dr.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		dr.findJobs(logger)
	}
	return dr.Jobs
}

func (dr DirectLine) GetURL() string {
	return dr.URL
}

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

	sugar.Debugf("Site %v is accessible", dr.URL)

	c.OnHTML("div.card.grid__item", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Infof("Looking for jobs at: %v", link)
		job, err := dr.gatherSpecs(link, logger)

		if err != nil {
			sugar.Error(zap.Error(err))
		}
		sugar.Infof("Job successfully scraped at: %v", link)
		foundJobs = append(foundJobs, job)
	})

	c.OnHTML("a.icon-right-ico.pagination__right-nav", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		sugar.Infof("Next page link found: %v", link)
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
		sugar.Debugf("Visiting page %v", r.URL.String())
		job.URL = r.URL.String()
	})

	d.OnHTML("h1", func(e *colly.HTMLElement) {
		job.Title = strings.TrimSpace(e.Text)

	})

	d.OnHTML("div.job-container ", func(e *colly.HTMLElement) {

		job.Requirements, err = dr.getRequirements(e)

	})
	d.OnHTML("div.location.map-ico-black", func(e *colly.HTMLElement) {
		job.Location, err = dr.getJobLocation(e)
	})

	// Not returning job type
	d.OnHTML("div.hero-box__details", func(e *colly.HTMLElement) {
		job.Type, err = dr.getJobType(e)
	})

	// Doesnt tell use salary
	job.Salary = "N/A"
	d.Wait()

	return job, nil
}

func (dr DirectLine) getRequirements(e *colly.HTMLElement) ([]string, error) {
	requirements := []string{}
	e.ForEach("li", func(_ int, el *colly.HTMLElement) {
		requirements = append(requirements, el.Text)
	})

	return requirements, nil
}

func (dr DirectLine) getJobLocation(e *colly.HTMLElement) (string, error) {

	return e.ChildText("span"), nil
}

func (dr DirectLine) getJobType(e *colly.HTMLElement) (string, error) {

	jobType := ""
	e.ForEach("div.copy", func(i int, el *colly.HTMLElement) {
		if i < 2 {

			jobType = jobType + el.Text
		}
	})

	return jobType, nil
}

package companies

import (
	jobs "job-scraper/internal"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

type Apple jobs.JobSource

//https://jobs.apple.com/en-gb/search?location=united-kingdom-GBR&team=apps-and-frameworks-SFTWR-AF+cloud-and-infrastructure-SFTWR-CLD+core-operating-systems-SFTWR-COS+devops-and-site-reliability-SFTWR-DSR+engineering-project-management-SFTWR-EPM+information-systems-and-technology-SFTWR-ISTECH+machine-learning-and-ai-SFTWR-MCHLN+security-and-privacy-SFTWR-SEC+software-quality-automation-and-tools-SFTWR-SQAT+wireless-software-SFTWR-WSFT

func (a Apple) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(a.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		a.findJobs(logger)
	}
	return a.Jobs
}

func (a Apple) GetPath() string {
	return a.FilePath
}

func (a *Apple) findJobs(logger *zap.Logger) {
	pageNum := 2
	sugar := logger.Sugar()

	foundJobs := []jobs.Job{}

	c := colly.NewCollector(
		colly.Async(true),
	)

	err := jobs.IsUp(a.URL)
	if err != nil {
		sugar.Fatal(err)
	}

	sugar.Debugf("Site %v is accessible", a.URL)

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Infof("Looking for jobs at: %v", link)

		job, err := a.gatherSpecs(link, logger)
		sugar.Infof("Job successfully scraped at: %v", link)

		if err != nil {
			sugar.Error(zap.Error(err))
		}

		foundJobs = append(foundJobs, job)

	})
	c.OnHTML("li.pagination__next", func(e *colly.HTMLElement) {
		// Apple has a function that manipulates the URL taking from the link in the  ">" button.
		// To work around this, we'll append "page=x" to the end, incrementing as we go through until we do not see a ">" button
		// TODO: Convert this to use the net/url package
		var link string
		if strings.Contains(a.URL, "?") {
			link = a.URL + "&page=" + strconv.Itoa(pageNum)
		}

		if !strings.Contains(a.URL, "?") {
			link = a.URL + "?page=" + strconv.Itoa(pageNum)
		}

		pageNum = pageNum + 1
		sugar.Infof("Next page link found: %v", link)
		e.Request.Visit(link)

	})

	c.Visit(a.URL)
	c.Wait()

	a.Jobs = foundJobs

}

func (a Apple) gatherSpecs(url string, logger *zap.Logger) (jobs.Job, error) {
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
		url := a.getJobURL(r)
		sugar.Debugf("Visiting page %v", url)
		job.URL = url
	})

	d.OnHTML("h1", func(e *colly.HTMLElement) {
		job.Title = a.getJobTitle(e)

	})

	d.OnHTML("#accordion_keyqualifications", func(e *colly.HTMLElement) {
		var err error

		job.Requirements, err = a.getJobRequirements(e)

		if err != nil {
			log.Fatal(err)
		}

	})

	d.OnHTML("#job-location-name", func(e *colly.HTMLElement) {
		var err error

		job.Location, err = a.getJobLocation(e)

		if err != nil {
			log.Fatal(err)
		}

	})

	// Type and salary not provided.
	job.Type = a.getJobType(nil)
	job.Salary = a.getJobSalary(nil)

	d.Wait()
	sugar.Debugf("Job details found %v", job)
	return job, nil
}

func (a Apple) getJobLocation(e *colly.HTMLElement) (string, error) {
	location := ""
	e.ForEach("span", func(_ int, el *colly.HTMLElement) {
		location = location + el.Text
	})
	return location, nil
}

func (a Apple) getJobRequirements(e *colly.HTMLElement) ([]string, error) {

	requirements := []string{}
	e.ForEach("li", func(_ int, el *colly.HTMLElement) {
		requirements = append(requirements, el.Text)
	})

	return requirements, nil
}

func (a Apple) getJobURL(r *colly.Request) string {
	return r.URL.String()
}

func (a Apple) getJobTitle(e *colly.HTMLElement) string {
	return e.Text
}

func (a Apple) getJobType(e *colly.HTMLElement) string {
	return "N/A"
}

func (a Apple) getJobSalary(e *colly.HTMLElement) string {
	return "N/A"
}

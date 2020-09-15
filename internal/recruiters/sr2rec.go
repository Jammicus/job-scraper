package recruiters

import (
	"fmt"
	jobs "job-scraper/internal"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"go.uber.org/zap"
)

var srURL = "https://www.sr2rec.co.uk/jobs/?jf_what=&jf_where=London"

type SR2 jobs.JobSource

func (sr SR2) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(sr.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		sr.findJobs(logger)
	}
	return sr.Jobs
}

func (sr SR2) GetPath() string {
	return sr.FilePath
}

func (sr *SR2) findJobs(logger *zap.Logger) {
	sugar := logger.Sugar()

	foundJobs := []jobs.Job{}

	c := colly.NewCollector(
		colly.Async(true),
	)

	err := jobs.IsUp(sr.URL)
	if err != nil {
		sugar.Fatal(err)
	}

	c.OnHTML("article", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Infof("Looking for jobs at: %v", link)

		job, err := sr.gatherSpecs(link, logger)
		sugar.Infof("Job successfully scraped with title: %v", job.Title)

		if err != nil {
			sugar.Error(zap.Error(err))
		}

		foundJobs = append(foundJobs, job)

	})

	c.OnHTML("div.sr2-jobs-pagination.sr2-jobs-pagination-bottom", func(e *colly.HTMLElement) {

		e.ForEach("a.next.page-numbers", func(_ int, e *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(e.Attr("href"))
			sugar.Debugf("Next page link found: %v", link)
			e.Request.Visit(link)

		})

	})

	c.Visit(sr.URL)
	c.Wait()

	sr.Jobs = foundJobs

}

func (sr SR2) gatherSpecs(url string, logger *zap.Logger) (jobs.Job, error) {
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
		url := sr.getJobURL(r)
		sugar.Debugf("Visiting page %v", url)
		job.URL = url
	})

	d.OnHTML("h1", func(e *colly.HTMLElement) {
		job.Title = sr.getJobTitle(e)

	})
	d.OnHTML("article.single-wpbb_job-content.entry.clr", func(e *colly.HTMLElement) {
		job.Requirements = sr.getJobRequirements(e)
	})

	d.OnHTML("div.wpbb-job-data__wrapper", func(e *colly.HTMLElement) {
		var err error

		job.Salary, err = sr.getJobSalary(e)
		if err != nil {
			sugar.Errorf("Could not get job salary %v", zap.Error(err))
		}

		job.Location, err = sr.getJobLocation(e)
		if err != nil {
			sugar.Errorf("Could not get job location %v", zap.Error(err))
		}

		job.Type, err = sr.getJobType(e)
		if err != nil {
			sugar.Errorf("Could not get job type %v", zap.Error(err))
		}
	})

	d.Wait()
	sugar.Debugf("Job details found %v", job)
	return job, nil
}

func (sr SR2) getJobType(e *colly.HTMLElement) (string, error) {
	var jobType string

	e.ForEach("div.wpbb-job-data__field.wpbb-job-data__field--job-type", func(_ int, el *colly.HTMLElement) {
		jobType = el.ChildText("span.value")
	})
	if jobType == "" {
		return "", fmt.Errorf("Could not get job type")
	}

	return jobType, nil
}

func (sr SR2) getJobLocation(e *colly.HTMLElement) (string, error) {
	var location string

	e.ForEach("div.wpbb-job-data__field.wpbb-job-data__field--job-location", func(_ int, el *colly.HTMLElement) {
		location = el.ChildText("span.value")
	})

	if strings.EqualFold(location, "") {
		return "", fmt.Errorf("Could not get location")
	}

	return location, nil
}

// TODO: Rethink errors
func (sr SR2) getJobSalary(e *colly.HTMLElement) (string, error) {
	var salary string
	var salaryTo string
	var salaryFrom string

	e.ForEach("div.wpbb-job-data__field.wpbb-job-data__field--salary-display", func(_ int, el *colly.HTMLElement) {
		if el.ChildText("span.value") != "" {
			salary = el.ChildText("span.value")
		}
	})
	e.ForEach("div.wpbb-job-data__field.wpbb-job-data__field--salary-from", func(_ int, el *colly.HTMLElement) {
		salaryTo = el.ChildText("span.value")
	})
	e.ForEach("div.wpbb-job-data__field.wpbb-job-data__field--salary-to", func(_ int, el *colly.HTMLElement) {
		salaryFrom = el.ChildText("span.value")
	})

	if !strings.EqualFold(salary, "") {
		return salary, nil
	}

	if !strings.EqualFold(salaryTo, "") && !strings.EqualFold(salaryFrom, "") {
		return salaryTo + " - " + salaryFrom, nil
	}

	return "", fmt.Errorf("Could not find salary information")
}

func (sr SR2) getJobRequirements(e *colly.HTMLElement) []string {

	requirements := []string{}
	e.ForEach("li:not(.meta-date):not(.share-twitter):not(.share-facebook):not(.share-googleplus):not(.share-pinterest)", func(_ int, el *colly.HTMLElement) {
		requirements = append(requirements, el.Text)
	})

	return requirements
}

func (sr SR2) getJobTitle(e *colly.HTMLElement) string {
	jobTitle := regexp.MustCompile(`^.*?[^|-]*`)
	return jobTitle.FindStringSubmatch(e.Text)[0]
}

func (sr SR2) getJobURL(r *colly.Request) string {
	return r.URL.String()
}

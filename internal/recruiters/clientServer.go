package recruiters

import (
	"fmt"
	"job-scraper/internal"
	jobs "job-scraper/internal"
	"regexp"
	"strings"

	"go.uber.org/zap"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// CLient servers page is a mass of text only.

var jobType = regexp.MustCompile(`^Job Type:`)
var jobDetails = regexp.MustCompile(`£[0-9]+ - £[0-9]+`)
var jobRequirements = regexp.MustCompile(`Requirements:\s*?[*]`)

type ClientServer struct {
	URL      string
	Jobs     []jobs.Job
	FilePath string
}

func (cs ClientServer) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(cs.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		cs.findJobs(logger)
	}
	cs.findJobs(logger)
	return cs.Jobs
}

func (cs ClientServer) GetURL() string {
	return cs.URL
}

func (cs ClientServer) GetPath() string {
	return cs.FilePath
}

func (cs *ClientServer) findJobs(logger *zap.Logger) {
	sugar := logger.Sugar()
	foundJobs := []jobs.Job{}

	c := colly.NewCollector(
		colly.Async(true),
	)

	err := internal.IsUp(cs.URL)
	if err != nil {
		sugar.Fatal(err)
	}

	sugar.Debugf("Site %v is accessible", cs.URL)

	c.OnHTML("li.job-result-item", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Infof("Looking for jobs at: %v", link)
		job, err := cs.gatherSpecs(link, logger)

		if err != nil {
			sugar.Error(zap.Error(err))
		}
		sugar.Infof("Job successfully scraped at: %v", link)
		foundJobs = append(foundJobs, job)
	})

	c.OnHTML(".next", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Infof("Next page link found: %v", link)
		e.Request.Visit(link)
	})

	c.Visit(cs.URL)
	c.Wait()

	cs.Jobs = foundJobs
}

func (cs ClientServer) gatherSpecs(url string, logger *zap.Logger) (jobs.Job, error) {
	sugar := logger.Sugar()

	jobType := regexp.MustCompile(`^Job Type:`)
	jobDetails := regexp.MustCompile(`£[0-9]+ - £[0-9]+`)
	jobRequirements := regexp.MustCompile(`Requirements:\s*?[*]`)

	job := jobs.Job{}

	d := colly.NewCollector(
		colly.Async(true),
	)

	d.Visit(url)
	d.OnRequest(func(r *colly.Request) {
		sugar.Debug("Visinting page")
		job.URL = r.URL.String()
	})

	d.OnHTML(".job.col-md-8.col-sm-12.clearfix", func(e *colly.HTMLElement) {

		title := e.ChildText("h1")

		job.URL = url
		job.Title = title

		e.ForEach("p", func(_ int, el *colly.HTMLElement) {

			switch {
			case jobRequirements.MatchString(el.Text):

				job.Requirements, _ = cs.getRequirements(el)

			case jobDetails.MatchString(el.Text):

				job.Location, _ = cs.getJobLocation(el.Text)
				job.Salary, _ = cs.getSalary(el.Text)

			case jobType.MatchString(el.Text):

				job.Type, _ = cs.getJobType(el.Text)
			}

		})
	})

	d.Wait()

	if len(job.Requirements) == 0 {
		return jobs.Job{}, fmt.Errorf("No requirements found for job %v", job.URL)
	}

	sugar.Debugf("Job details found %v", job)
	return job, nil
}

func (cs ClientServer) getJobType(s string) (string, error) {
	jobTypeAgain := jobType.Split(s, 2)

	x := strings.TrimSpace(jobTypeAgain[len(jobTypeAgain)-1])
	return x, nil
}

func (cs ClientServer) getJobLocation(s string) (string, error) {
	location := jobDetails.Split(s, -1)[0]

	return location, nil
}

func (cs ClientServer) getSalary(s string) (string, error) {
	salary := jobDetails.FindStringSubmatch(s)[0]

	return salary, nil
}

func (cs ClientServer) getRequirements(el *colly.HTMLElement) ([]string, error) {
	jobRequirement := regexp.MustCompile(`[^<br\/>][^<]*`)

	requirements := []string{}
	el.DOM.Each(func(_ int, s *goquery.Selection) {
		// Use <BR> To identify between requirements
		h, _ := s.Html()
		test := jobRequirement.FindAllStringSubmatch(h, -1)

		for _, item := range test {
			// IF its in a list, its requirements for the job, remove the asterix and print
			if strings.HasPrefix(item[0], "*") {
				requirements = append(requirements, strings.Trim(item[0], "*"))
			}

			if strings.HasPrefix(item[0], "\n*") {
				requirements = append(requirements, strings.Trim(item[0], "\n*"))
			}
		}
	})
	return requirements, nil
}

package recruiters

import (
	"fmt"
	jobs "job-scraper/internal"
	"regexp"
	"strings"

	"go.uber.org/zap"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

// Client servers page is a mass of text only.

var jobType = regexp.MustCompile(`^Job Type:`)
var jobDetails = regexp.MustCompile(`£[0-9]+ - £[0-9]+`)
var jobRequirements = regexp.MustCompile(`Requirements:\s*?[*]`)


// ClientServer is a JobSource
type ClientServer jobs.JobSource

// GetJobs returns all jobs  for a given receiver
func (cs ClientServer) GetJobs(logger *zap.Logger) []jobs.Job {
	if len(cs.Jobs) == 0 {
		sugar := logger.Sugar()
		sugar.Info("Jobs have not previously been found, finding jobs.")
		cs.findJobs(logger)
	}
	return cs.Jobs
}

// GetPath returns the filepath to write the CSV to for a given receiver
func (cs ClientServer) GetPath() string {
	return cs.FilePath
}

func (cs *ClientServer) findJobs(logger *zap.Logger) {
	sugar := logger.Sugar()
	foundJobs := []jobs.Job{}

	c := colly.NewCollector(
		colly.Async(true),
	)

	err := jobs.IsUp(cs.URL)
	if err != nil {
		sugar.Fatal(err)
	}

	c.OnHTML("li.job-result-item", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Infof("Looking for jobs at: %v", link)
		job, err := cs.gatherSpecs(link, logger)

		if err != nil {
			sugar.Error(zap.Error(err))
		}
		sugar.Infof("Job successfully scraped with title: %v", job.Title)
		foundJobs = append(foundJobs, job)
	})

	c.OnHTML(".next", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		sugar.Debugf("Next page link found: %v", link)
		e.Request.Visit(link)
	})

	c.Visit(cs.URL)
	c.Wait()

	cs.Jobs = foundJobs
}

func (cs ClientServer) gatherSpecs(url string, logger *zap.Logger) (jobs.Job, error) {
	sugar := logger.Sugar()

	job := jobs.Job{}

	d := colly.NewCollector(
		colly.Async(true),
	)

	d.Visit(url)
	d.OnRequest(func(r *colly.Request) {
		sugar.Debugf("Visiting page %v", r.URL.String())
		job.URL = r.URL.String()
	})

	// Refactor this, it is gross
	d.OnHTML(".job.col-md-8.col-sm-12.clearfix", func(e *colly.HTMLElement) {

		e.ForEach("h1", func(_ int, el *colly.HTMLElement) {
			job.Title = cs.getJobTitle(el)
		})

		e.ForEach("p", func(_ int, el *colly.HTMLElement) {

			switch {
			case jobRequirements.MatchString(el.Text):

				job.Requirements = cs.getJobRequirements(el)

			case jobDetails.MatchString(el.Text):

				job.Location = cs.getJobLocation(el.Text)
				job.Salary = cs.getJobSalary(el.Text)

			case jobType.MatchString(el.Text):

				job.Type = cs.getJobType(el.Text)
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

func (cs ClientServer) getJobTitle(e *colly.HTMLElement) string {

	return strings.TrimSpace(e.Text)
}

func (cs ClientServer) getJobType(s string) string {
	jobTypeAgain := jobType.Split(s, 2)

	x := strings.TrimSpace(jobTypeAgain[len(jobTypeAgain)-1])
	return x
}

func (cs ClientServer) getJobLocation(s string) string {
	location := strings.TrimSpace(jobDetails.Split(s, -1)[0])

	if strings.HasSuffix(location, ",") {
		return strings.TrimSuffix(location, ",")
	}

	return strings.TrimSpace(location)
}

func (cs ClientServer) getJobSalary(s string) string {
	salary := jobDetails.FindStringSubmatch(s)[0]

	return strings.TrimSpace(salary)
}

func (cs ClientServer) getJobRequirements(el *colly.HTMLElement) []string {
	jobRequirement := regexp.MustCompile(`[^<br\/>][^<]*`)

	fmt.Println("@@@@@@@@")
	fmt.Println(el.Text)
	fmt.Println("@@@@@@@@")

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
	return requirements
}

func (cs ClientServer) getJobURL(r *colly.Request) string {
	return r.URL.String()
}

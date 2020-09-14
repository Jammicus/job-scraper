package recruiters

import (
	jobs "job-scraper/internal"
	"log"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGatherSpecsUnderstanding(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	ts := jobs.StartTestServer("../../testdata/recruiters/understanding-job.html")
	defer ts.Close()

	understanding := Understanding{
		URL: ts.URL + "/job",
	}
	expected := jobs.Job{
		Title:    "DevOps Engineer (Green Energy)",
		Type:     "Permanent",
		Salary:   "£80000 - £95000 per annum + benefits",
		Location: "London, England",
		URL:      ts.URL + "/job",
		Requirements: []string{
			"Strong technical capabilities and an ability to work within GCP or AWS cloud",
			"Experience with production container tools including Kubernetes and Docker",
			"Experience in an infrastructure as code environment using tools likes Terraform or Cloudformation",
			"Python or development experience is very important (Python a strong preference in language)",
		},
	}

	result, err := understanding.gatherSpecs(understanding.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestGetJobTypeUnderstanding(t *testing.T) {
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	understanding := Understanding{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/understanding-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "ul.job-details-list.clearfix"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := understanding.getJobType(elements[0])

	assert.Equal(t, "Permanent", result)

}

func TestGetJobLocationUnderstanding(t *testing.T) {
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	understanding := Understanding{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/understanding-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "ul.job-details-list.clearfix"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := understanding.getJobLocation(elements[0])

	assert.Equal(t, "London, England", result)

}

func TestGetJobSalaryUnderstanding(t *testing.T) {
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	understanding := Understanding{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/understanding-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "ul.job-details-list.clearfix"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := understanding.getJobSalary(elements[0])

	assert.Equal(t, "£80000 - £95000 per annum + benefits", result)

}

func TestGetJobRequirementsUnderstanding(t *testing.T) {

	expected := []string{
		"Strong technical capabilities and an ability to work within GCP or AWS cloud",
		"Experience with production container tools including Kubernetes and Docker",
		"Experience in an infrastructure as code environment using tools likes Terraform or Cloudformation",
		"Python or development experience is very important (Python a strong preference in language)",
	}
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	understanding := Understanding{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/understanding-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "article.clearfix.mb20"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := understanding.getJobRequirements(elements[0])

	assert.Equal(t, expected, result)

}

func TestGetJobTitleUnderstanding(t *testing.T) {
	expected := "DevOps Engineer (Green Energy)"
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	u := Understanding{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/understanding-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "h1"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result := u.getJobTitle(elements[0])

	assert.Equal(t, expected, result)
}

// Test via the gatherSpecs method for now.
// TODO: Refactor test to use colly.Request
func TestGetJobURLUnderstanding(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	ts := jobs.StartTestServer("../../testdata/recruiters/understanding-job.html")
	defer ts.Close()

	u := Understanding{
		URL: ts.URL + "/job",
	}
	expected := jobs.Job{
		Title:    "DevOps Engineer (Green Energy)",
		Type:     "Permanent",
		Salary:   "£80000 - £95000 per annum + benefits",
		Location: "London, England",
		URL:      ts.URL + "/job",
		Requirements: []string{
			"Strong technical capabilities and an ability to work within GCP or AWS cloud",
			"Experience with production container tools including Kubernetes and Docker",
			"Experience in an infrastructure as code environment using tools likes Terraform or Cloudformation",
			"Python or development experience is very important (Python a strong preference in language)",
		},
	}

	result, err := u.gatherSpecs(u.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected.URL, result.URL)
}

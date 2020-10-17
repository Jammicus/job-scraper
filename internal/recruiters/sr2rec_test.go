package recruiters

import (
	"log"
	"os"
	"testing"

	jobs "job-scraper/internal"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

var testServerURLSR2 string

func init() {
	testServer := jobs.StartTestServer("../../testdata/recruiters/sr2rec-job.html")
	testServerURLSR2 = testServer.URL + "/job"
}

func TestGatherSpecsSR2(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	sr := SR2{
		URL: testServerURLSR2,
	}
	expected := jobs.Job{
		Title:    "Lead Java Developer – SC Cleared",
		Type:     "Contract",
		Salary:   "£400 - £500",
		Location: "London",
		URL:      testServerURLSR2,
		Requirements: []string{
			"Java 8+",
			"RESTful APIs",
			"Spring",
			"CI/CD",
			"Gradle or Maven",
			"AWS or Azure",
		},
	}

	result, err := sr.gatherSpecs(sr.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestGetJobTypeSR2(t *testing.T) {

	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	sr := SR2{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/sr2rec-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "div.wpbb-job-data__wrapper"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := sr.getJobType(elements[0])

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "Contract", result)
}

func TestGetJobLocationSR2(t *testing.T) {
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	sr := SR2{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/sr2rec-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "div.wpbb-job-data__wrapper"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := sr.getJobLocation(elements[0])

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "London", result)
}

func TestGetJobSalarySR2(t *testing.T) {
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	sr := SR2{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/sr2rec-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "div.wpbb-job-data__wrapper"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := sr.getJobSalary(elements[0])

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "£400 - £500", result)
}

func TestGetJobRequirementsSR2(t *testing.T) {
	expected := []string{
		"Java 8+",
		"RESTful APIs",
		"Spring",
		"CI/CD",
		"Gradle or Maven",
		"AWS or Azure",
	}
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	sr := SR2{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/sr2rec-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "article.single-wpbb_job-content.entry.clr"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result := sr.getJobRequirements(elements[0])

	assert.Equal(t, expected, result)
}

func TestGetJobTitleSR2(t *testing.T) {
	expected := "Lead Java Developer – SC Cleared"
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	sr := SR2{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/sr2rec-job.html")

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

	result := sr.getJobTitle(elements[0])

	assert.Equal(t, expected, result)
}

// Test via the gatherSpecs method for now.
// TODO: Refactor test to use colly.Request
func TestGetJobURLSR2(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	sr := SR2{
		URL: testServerURLSR2,
	}
	expected := jobs.Job{
		Title:    "Lead Java Developer – SC Cleared",
		Type:     "Contract",
		Salary:   "£400 - £500",
		Location: "London",
		URL:      testServerURLSR2,
		Requirements: []string{
			"Java 8+",
			"RESTful APIs",
			"Spring",
			"CI/CD",
			"Gradle or Maven",
			"AWS or Azure",
		},
	}

	result, err := sr.gatherSpecs(sr.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected.URL, result.URL)
}

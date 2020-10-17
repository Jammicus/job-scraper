package companies

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

var testServerURLApple string

func init() {
	testServer := jobs.StartTestServer("../../testdata/companies/apple-job.html")
	testServerURLApple = testServer.URL + "/job"
}

func TestGatherSpecsApple(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	a := Apple{
		URL: testServerURLApple,
	}
	expected := jobs.Job{
		Title:    "AI/ML - Software Engineer (Python, Spark) - ML Platform & Technologies",
		Type:     "N/A",
		Salary:   "N/A",
		Location: "Cambridge, Cambridgeshire, United Kingdom",
		URL:      testServerURLApple,
		Requirements: []string{
			"Experience building applications/services with Python / Java / Scala",
			"Strong background in building scalable and fault-tolerant distributed systems, particularly in realtime environments.",
			"Experience in building data pipelines, data caching/storage systems, and/or RPC services.",
			"Data technologies, eg Spark, Hadoop, Kafka",
			"Microservices architecture",
			"SQL / NoSQL databases",
			"Strong understanding of data structures & algorithms",
			"Excellent problem solving and debugging skills",
			"Strong written and verbal communication skills",
			"Experience with Javascript and UI frameworks/libraries like Ember, Angular, React, D3",
		},
	}

	result, err := a.gatherSpecs(a.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestGetJobLocationApple(t *testing.T) {
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	a := Apple{
		URL: "",
	}

	file, err := os.Open("../../testdata/companies/apple-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "#job-location-name"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := a.getJobLocation(elements[0])

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "Cambridge, Cambridgeshire, United Kingdom", result)
}

func TestGetJobRequirementsApple(t *testing.T) {
	expected := []string{
		"Experience building applications/services with Python / Java / Scala",
		"Strong background in building scalable and fault-tolerant distributed systems, particularly in realtime environments.",
		"Experience in building data pipelines, data caching/storage systems, and/or RPC services.",
		"Data technologies, eg Spark, Hadoop, Kafka",
		"Microservices architecture",
		"SQL / NoSQL databases",
		"Strong understanding of data structures & algorithms",
		"Excellent problem solving and debugging skills",
		"Strong written and verbal communication skills",
		"Experience with Javascript and UI frameworks/libraries like Ember, Angular, React, D3",
	}
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	a := Apple{
		URL: "",
	}

	file, err := os.Open("../../testdata/companies/apple-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "#accordion_keyqualifications"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := a.getJobRequirements(elements[0])

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

// Test via the gatherSpecs method for now.
// TODO: Refactor test to use colly.Request
func TestGetJobURLApple(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	a := Apple{
		URL: testServerURLApple,
	}

	result, err := a.gatherSpecs(a.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, testServerURLApple, result.URL)
}

func TestGetJobTitleApple(t *testing.T) {
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	a := Apple{
		URL: "",
	}

	file, err := os.Open("../../testdata/companies/apple-job.html")

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

	result := a.getJobTitle(elements[0])

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "AI/ML - Software Engineer (Python, Spark) - ML Platform & Technologies", result)
}

func TestGetJobTypeApple(t *testing.T) {

	a := Apple{
		URL: "",
	}

	result := a.getJobType(nil)

	assert.Equal(t, "N/A", result)
}

func TestGetJobSalaryApple(t *testing.T) {

	a := Apple{
		URL: "",
	}

	result := a.getJobSalary(nil)

	assert.Equal(t, "N/A", result)
}

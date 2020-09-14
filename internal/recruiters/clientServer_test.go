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

func TestGatherSpecsClientServer(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	ts := jobs.StartTestServer("../../testdata/recruiters/clientserver-job.html")
	defer ts.Close()

	cs := ClientServer{
		URL: ts.URL + "/job",
	}
	expected := jobs.Job{
		Title:    "Lead JavaScript Developer - React TypeScript",
		Type:     "Permanent",
		Salary:   "£70000 - £95000",
		Location: "London, England",
		URL:      ts.URL + "/job",
		Requirements: []string{
			"You have strong commercial experience as a Front End Developer with all of the following: JavaScript, React, Redux and TypeScript ",
			"You have experience of building user interfaces on top of RESTful APIs",
			"You&#39;re experienced with building responsive sites for desktop, tablet and mobile",
			"You&#39;re collaborative with good communication skills, keen to be part of a diverse and creative team",
			"You&#39;re degree educated in a STEM subject",
		}}

	result, err := cs.gatherSpecs(ts.URL+"/job", logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestGetJobTitleClientServer(t *testing.T) {
	expected := "Lead JavaScript Developer - React TypeScript"
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	cs := ClientServer{
		URL: "",
	}

	file, err := os.Open("../../testdata/recruiters/clientserver-job.html")

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

	result := cs.getJobTitle(elements[0])

	assert.Equal(t, expected, result)
}

func TestGetJobTypeClientServer(t *testing.T) {
	input := "Job Type: Permanent"
	expected := "Permanent"

	cs := ClientServer{}
	result := cs.getJobType(input)

	assert.Equal(t, expected, result)
}

func TestGetJobLocationClientServer(t *testing.T) {

	// With comma after country and without.
	testCases := []struct {
		input    string
		expected string
	}{
		{"London, England, £70000 - £95000 per annum + benefits", "London, England"},
		{"London, England £70000 - £95000 per annum + benefits", "London, England"},
	}
	cs := ClientServer{}

	for _, test := range testCases {
		result := cs.getJobLocation(test.input)

		assert.Equal(t, test.expected, result)
	}
}

func TestGetJobSalaryClientServer(t *testing.T) {
	input := "London, England, £70000 - £95000 per annum + benefits"
	expected := "£70000 - £95000"

	cs := ClientServer{}
	result := cs.getJobSalary(input)

	assert.Equal(t, expected, result)
}

// Test via the gatherSpecs method for now.
// TODO: Refactor test once gatherSpecs has been refactored
func TestGetJobRequirementsClientServer(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	ts := jobs.StartTestServer("../../testdata/recruiters/clientserver-job.html")
	defer ts.Close()

	cs := ClientServer{
		URL: ts.URL + "/job",
	}
	expected := jobs.Job{
		Title:    "Lead JavaScript Developer - React TypeScript",
		Type:     "Permanent",
		Salary:   "£70000 - £95000",
		Location: "London, England",
		URL:      ts.URL + "/job",
		Requirements: []string{
			"You have strong commercial experience as a Front End Developer with all of the following: JavaScript, React, Redux and TypeScript ",
			"You have experience of building user interfaces on top of RESTful APIs",
			"You&#39;re experienced with building responsive sites for desktop, tablet and mobile",
			"You&#39;re collaborative with good communication skills, keen to be part of a diverse and creative team",
			"You&#39;re degree educated in a STEM subject",
		}}

	result, err := cs.gatherSpecs(cs.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected.Requirements, result.Requirements)
}

// Test via the gatherSpecs method for now.
// TODO: Refactor test to use colly.Request
func TestGetJobURLClientServer(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	ts := jobs.StartTestServer("../../testdata/recruiters/clientserver-job.html")
	defer ts.Close()

	cs := ClientServer{
		URL: ts.URL + "/job",
	}
	expected := jobs.Job{
		Title:    "Lead Java Developer – SC Cleared",
		Type:     "Contract",
		Salary:   "£400 - £500",
		Location: "London",
		URL:      ts.URL + "/job",
		Requirements: []string{
			"Java 8+",
			"RESTful APIs",
			"Spring",
			"CI/CD",
			"Gradle or Maven",
			"AWS or Azure",
		},
	}

	result, err := cs.gatherSpecs(cs.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected.URL, result.URL)
}

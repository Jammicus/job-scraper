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

func TestGatherSpecsHashicorp(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	ts := jobs.StartTestServer("../../testdata/companies/hashicorp-job.html")
	defer ts.Close()

	h := Hashicorp{
		URL: ts.URL + "/job",
	}
	expected := jobs.Job{
		Title:    "Site Reliability Engineer - Cloud Services",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "United States (Remote)",
		URL:      ts.URL + "/job",
		Requirements: []string{
			"Design, implement, and maintain a secure and scalable infrastructure platform for delivering Cloud Services’ applications",
			"Own and ensure that internal and external SLA’s meet and exceed expectations, System centric KPIs are continuously monitored and improved",
			"Create tools for automating deployment, monitoring and operations of the overall platform",
			"Participate in on-call rotation to provide application support, incident management, and troubleshooting",
			"Provide ongoing maintenance and support of internal tools, improve system health and reliability",
			"Program mostly in Golang, learning from and contributing to a team committed to continually improving their skills",
			"Familiarity with infrastructure management and operations lifecycle concepts and ecosystem",
			"Experience operating and maintaining production systems in a Linux and public cloud environment",
			"You have prior experience working in high performance or distributed systems; while we strive to hire at a variety of experience levels, this particular opening is not well-suited for recent graduates",
			"Working knowledge of industry best practices with regard to information security",
			"You have built or operated a large scale Cloud service",
			"Comfortable with Go or another low-level programming language",
		},
	}

	result, err := h.gatherSpecs(h.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestGetRequirementsHashicorp(t *testing.T) {
	expected := []string{
		"Design, implement, and maintain a secure and scalable infrastructure platform for delivering Cloud Services’ applications",
		"Own and ensure that internal and external SLA’s meet and exceed expectations, System centric KPIs are continuously monitored and improved",
		"Create tools for automating deployment, monitoring and operations of the overall platform",
		"Participate in on-call rotation to provide application support, incident management, and troubleshooting",
		"Provide ongoing maintenance and support of internal tools, improve system health and reliability",
		"Program mostly in Golang, learning from and contributing to a team committed to continually improving their skills",
		"Familiarity with infrastructure management and operations lifecycle concepts and ecosystem",
		"Experience operating and maintaining production systems in a Linux and public cloud environment",
		"You have prior experience working in high performance or distributed systems; while we strive to hire at a variety of experience levels, this particular opening is not well-suited for recent graduates",
		"Working knowledge of industry best practices with regard to information security",
		"You have built or operated a large scale Cloud service",
		"Comfortable with Go or another low-level programming language",
	}
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	h := Hashicorp{
		URL: "",
	}

	file, err := os.Open("../../testdata/companies/hashicorp-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := ".application-details"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := h.getRequirements(elements[0])

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

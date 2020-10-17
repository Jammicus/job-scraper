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

var testServerURLDL string

func init() {
	testServer := jobs.StartTestServer("../../testdata/companies/directline-job.html")
	testServerURLDL = testServer.URL + "/job"
}

func TestGatherSpecsDirectline(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	dr := DirectLine{
		URL: testServerURLDL,
	}
	expected := jobs.Job{
		Title:    "Senior DevOps Engineer (Cloud Provisioning & Access Management)",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "Churchill Court, Bromley, BR1 1DP",
		URL:      testServerURLDL,
		Requirements: []string{
			"Providing your expertise to build core services to be consumed by Projects (e.g. landing zones, Security controls, governance automation & Core networking) to cloud best practice.",
			"Liaising with projects to help them consume services in an appropriate and governed manner",
			"Working with internal and external suppliers to design and build the configuration management, release, deployment and operations cycle to meet business requirements.",
			"Amplifying feedback loops - and increase the frequency - through the automation",
			"Working with Architecture, Security and Other key teams to create pattern based best practice approaches to delivery.",
			"Identifying and implementing opportunities for innovation and continuous improvement in the development and continuous deployment of applications.",
			"Terraform, Cloud formation and Python - essential",
			"Excellent AWS skills – both depth and breadth of services – Certified preferred",
			"Clear and opinionated understanding of cloud (AWS) best practice",
			"Understanding of security and governance in the Cloud using native services – preferred",
			"Usage and implementation of CI/CD pipelines to deliver continuous improvement - essential",
			"Excellent written, and verbal communication skills",
		},
	}

	result, err := dr.gatherSpecs(dr.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestGetJobRequirementsDirectline(t *testing.T) {
	expected := []string{
		"Providing your expertise to build core services to be consumed by Projects (e.g. landing zones, Security controls, governance automation & Core networking) to cloud best practice.",
		"Liaising with projects to help them consume services in an appropriate and governed manner",
		"Working with internal and external suppliers to design and build the configuration management, release, deployment and operations cycle to meet business requirements.",
		"Amplifying feedback loops - and increase the frequency - through the automation",
		"Working with Architecture, Security and Other key teams to create pattern based best practice approaches to delivery.",
		"Identifying and implementing opportunities for innovation and continuous improvement in the development and continuous deployment of applications.",
		"Terraform, Cloud formation and Python - essential",
		"Excellent AWS skills – both depth and breadth of services – Certified preferred",
		"Clear and opinionated understanding of cloud (AWS) best practice",
		"Understanding of security and governance in the Cloud using native services – preferred",
		"Usage and implementation of CI/CD pipelines to deliver continuous improvement - essential",
		"Excellent written, and verbal communication skills",
	}
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	dr := DirectLine{
		URL: "",
	}

	file, err := os.Open("../../testdata/companies/directline-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "div.job-container"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := dr.getJobRequirements(elements[0])

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestGetJobLocationDirectline(t *testing.T) {
	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	dr := DirectLine{
		URL: "",
	}

	file, err := os.Open("../../testdata/companies/directline-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "div.location.map-ico-black"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result := dr.getJobLocation(elements[0])

	assert.Equal(t, "Churchill Court, Bromley, BR1 1DP", result)
}

func TestGetJobTypeDirectline(t *testing.T) {

	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	dr := DirectLine{
		URL: "",
	}

	file, err := os.Open("../../testdata/companies/directline-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "div.hero-box__details"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result, err := dr.getJobType(elements[0])

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "Permanent", result)
}

// Test via the gatherSpecs method for now.
// TODO: Refactor test to use colly.Request
func TestGetJobURLDirectline(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	dr := DirectLine{
		URL: testServerURLDL,
	}
	expected := jobs.Job{
		Title:    "Senior DevOps Engineer (Cloud Provisioning & Access Management)",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "Churchill Court, Bromley, BR1 1DP",
		URL:      testServerURLDL,
		Requirements: []string{
			"Providing your expertise to build core services to be consumed by Projects (e.g. landing zones, Security controls, governance automation & Core networking) to cloud best practice.",
			"Liaising with projects to help them consume services in an appropriate and governed manner",
			"Working with internal and external suppliers to design and build the configuration management, release, deployment and operations cycle to meet business requirements.",
			"Amplifying feedback loops - and increase the frequency - through the automation",
			"Working with Architecture, Security and Other key teams to create pattern based best practice approaches to delivery.",
			"Identifying and implementing opportunities for innovation and continuous improvement in the development and continuous deployment of applications.",
			"Terraform, Cloud formation and Python - essential",
			"Excellent AWS skills – both depth and breadth of services – Certified preferred",
			"Clear and opinionated understanding of cloud (AWS) best practice",
			"Understanding of security and governance in the Cloud using native services – preferred",
			"Usage and implementation of CI/CD pipelines to deliver continuous improvement - essential",
			"Excellent written, and verbal communication skills",
		},
	}

	result, err := dr.gatherSpecs(dr.URL, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected.URL, result.URL)
}

func TestGetJobTitleDirectline(t *testing.T) {

	ctx := &colly.Context{}
	resp := &colly.Response{
		Request: &colly.Request{
			Ctx: ctx,
		},
		Ctx: ctx,
	}
	dr := DirectLine{
		URL: "",
	}

	file, err := os.Open("../../testdata/companies/directline-job.html")

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(file)
	sel := "div.top"
	elements := []*colly.HTMLElement{}
	i := 0
	doc.Find(sel).Each(func(_ int, s *goquery.Selection) {
		for _, n := range s.Nodes {
			elements = append(elements, colly.NewHTMLElementFromSelectionNode(resp, s, n, i))
			i++
		}
	})

	result := dr.getJobTitle(elements[0])

	assert.Equal(t, "Senior DevOps Engineer (Cloud Provisioning & Access Management)", result)
}

func TestGetJobSalaryDirectline(t *testing.T) {

	dr := DirectLine{
		URL: "",
	}

	result := dr.getJobSalary(nil)

	assert.Equal(t, "N/A", result)
}

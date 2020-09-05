package recruiters

import (
	jobs "job-scraper/internal"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

// TODO: Trim the spaces
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
		Location: "London, England,",
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

package companies

import (
	jobs "job-scraper/internal"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestFindJobsAmazon(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	ts := jobs.StartTestServer("../../testdata/companies/amazon-job.json")
	defer ts.Close()

	a := Amazon{
		URL: ts.URL + "/job",
	}
	expected := []jobs.Job{{
		Title:    "Senior Consultant, Data Lake \u0026 Analytics",
		Type:     "full-time",
		Salary:   "N/A",
		Location: "UK, London",
		URL:      "www.amazon.jobs/en-gb/jobs/994927/senior-consultant-data-lake-analytics",
		Requirements: []string{
			"Bachelor’s degree, or equivalent experience, in Computer Science, Engineering, Mathematics or a related field",
			"8+ years of experience of IT platform implementation in a highly technical and analytical role.",
			"5+ years’ experience of Data Lake/Hadoop platform implementation, including 3+ years of hands-on experience in implementation and performance tuning Hadoop/Spark implementations",
			"Ability to think strategically about business, product, and technical challenges in an enterprise environment.",
			"Experience with analytic solutions applied to the Marketing or Risk needs of enterprises",
			"Highly technical and analytical, possessing 5 or more years of IT platform implementation experience.",
			"Understanding of Apache Hadoop and the Hadoop ecosystem. Experience with one or more relevant tools (Sqoop, Flume, Kafka, Oozie, Hue, Zookeeper, HCatalog, Solr, Avro).",
			"Familiarity with one or more SQL-on-Hadoop technology (Hive, Impala, Spark SQL, Presto).", "Experience developing software code in one or more programming languages (Java, JavaScript, Python, etc).", "", "", "Experience in a Chief Architect role or similar", "Masters or PhD in Computer Science, Physics, Engineering or Math.", "Hands on experience leading large-scale global data warehousing and analytics projects.", "Ability to lead effectively across organizations.", "Understanding of database and analytical technologies in the industry including MPP and NoSQL databases, Data Warehouse design, BI reporting and Dashboard development.", "Demonstrated industry leadership in the fields of database, data warehousing or data sciences.", "Implementation and tuning experience specifically using Amazon Elastic Map Reduce (EMR).", "Implementing AWS services in a variety of distributed computing, enterprise environments.", "Computer Science or Math background preferred.", "Customer facing skills to represent AWS well within the customer’s environment and drive discussions with senior personnel regarding trade-offs, best practices, project management and risk mitigation. Should be able to interact with Chief Marketing Officers, Chief Risk Officers, Chief Technology Officers, and Chief Information Officers, as well as the people within their organizations.",
			"",
			"",
		},
	}}

	a.findJobs(logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, a.Jobs)
}

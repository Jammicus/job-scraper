package main

import (
	"flag"
	"log"

	csv "job-scraper/internal"
	recuriters "job-scraper/internal/recruiters"

	"go.uber.org/zap"
)

func main() {
	file := flag.String("file", "jobs.csv", "Path to file to store results in")
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Fatal("Error starting logger", zap.Error(err))
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("Starting program")

	flag.Parse()
	sugar.Info("Parsing Flags")
	understanding := recuriters.Understanding{
		URL: "https://www.understandingrecruitment.co.uk/jobs/england/?&search%5Bradius%5D=50.0&selected_locations=a-6269131",
	}

	exampleAgainAgain, err := understanding.FindJobs(logger)

	if err != nil {
		log.Fatal(err)
	}

	sr := recuriters.SR2{
		URL: "https://www.sr2rec.co.uk/jobs/?jf_what=&jf_where=London",
	}

	exampleAgain, err := sr.FindJobs(logger)

	if err != nil {
		log.Fatal(err)
	}

	cs := recuriters.ClientServer{
		URL: "https://www.client-server.com/jobs/london/",
	}

	example, err := cs.FindJobs(logger)
	if err != nil {
		log.Fatal(err)
	}

	sugar.Infof("Number of Jobs found %v", len(exampleAgainAgain)+len(exampleAgain)+len(example))

	csv.WriteToCSV(*file, exampleAgainAgain, logger)
	csv.WriteToCSV(*file, exampleAgain, logger)
	csv.WriteToCSV(*file, example, logger)

}

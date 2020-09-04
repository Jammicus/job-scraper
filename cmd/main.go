package main

import (
	"fmt"
	"log"

	"go.uber.org/zap"

	recuriters "job-scraper/internal/recruiters"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		logger.Fatal("Error starting logger", zap.Error(err))
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	sugar.Info("Starting program")

	understanding := recuriters.Understanding{
		URL: "https://www.understandingrecruitment.co.uk/jobs/england/?&search%5Bradius%5D=50.0&selected_locations=a-6269131",
	}

	exampleAgainAgain, err := understanding.FindJobs(logger)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(exampleAgainAgain)

	sr := recuriters.SR2{
		URL: "https://www.sr2rec.co.uk/jobs/?jf_what=&jf_where=London",
	}

	exampleAgain, err := sr.FindJobs(logger)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(exampleAgain)

	cs := recuriters.ClientServer{
		URL: "https://www.client-server.com/jobs/london/",
	}

	example, err := cs.FindJobs(logger)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(example)
	sugar.Infof("Starting program: %v", len(exampleAgainAgain)+len(exampleAgain)+len(example))
}

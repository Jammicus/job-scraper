package main

import (
	csv "job-scraper/internal"
	recuriters "job-scraper/internal/recruiters"

	"go.uber.org/zap"
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
		URL:      "https://www.understandingrecruitment.co.uk/jobs/england/?&search%5Bradius%5D=50.0&selected_locations=a-6269131",
		FilePath: "understanding.csv",
	}

	sr := recuriters.SR2{
		URL:      "https://www.sr2rec.co.uk/jobs/?jf_what=&jf_where=London",
		FilePath: "sr2rec.csv",
	}

	cs := recuriters.ClientServer{
		URL:      "https://www.client-server.com/jobs/london/",
		FilePath: "clientServer.csv",
	}

	csv.WriteToCSV(understanding, logger)
	csv.WriteToCSV(sr, logger)
	csv.WriteToCSV(cs, logger)

}

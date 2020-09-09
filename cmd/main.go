package main

import (
	csv "job-scraper/internal"
	companies "job-scraper/internal/companies"
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

	dr := companies.DirectLine{
		URL:      "https://www.directlinegroupcareers.com/search-apply/?jobcategory=ca77d4d22fb7013e872107cf18d65905&location=&keywords=&sortby=fj&distance=50&resultsperpage=12",
		FilePath: "directLine.csv",
	}

	csv.WriteToCSV(dr, logger)
	csv.WriteToCSV(understanding, logger)
	csv.WriteToCSV(sr, logger)
	csv.WriteToCSV(cs, logger)

}

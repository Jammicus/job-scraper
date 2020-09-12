package main

import (
	csv "job-scraper/internal"
	companies "job-scraper/internal/companies"
	recruiters "job-scraper/internal/recruiters"

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

	understanding := recruiters.Understanding{
		URL:      "https://www.understandingrecruitment.co.uk/jobs/england/?&search%5Bradius%5D=50.0&selected_locations=a-6269131",
		FilePath: "understanding.csv",
	}

	sr := recruiters.SR2{
		URL:      "https://www.sr2rec.co.uk/jobs/?jf_what=&jf_where=London",
		FilePath: "sr2rec.csv",
	}

	cs := recruiters.ClientServer{
		URL:      "https://www.client-server.com/jobs/london/",
		FilePath: "clientServer.csv",
	}

	dr := companies.DirectLine{
		URL:      "https://www.directlinegroupcareers.com/search-apply/?jobcategory=ca77d4d22fb7013e872107cf18d65905&location=&keywords=&sortby=fj&distance=50&resultsperpage=12",
		FilePath: "directLine.csv",
	}

	hs := companies.Hashicorp{
		URL:      "https://www.hashicorp.com/jobs/engineering",
		FilePath: "hashiCorp.csv",
	}

	a := companies.Amazon{
		URL:      "https://www.amazon.jobs/en-gb/search.json?schedule_type_id[]=Full-Time&radius=24km&facets[]=location&facets[]=business_category&facets[]=category&facets[]=schedule_type_id&facets[]=employee_class&facets[]=normalized_location&facets[]=job_function_id&offset=0&result_limit=10000&sort=relevant&latitude=&longitude=&loc_group_id=&loc_query=&base_query=&city=&country=&region=&county=&query_options=&location[]=london&",
		FilePath: "amazon.csv",
	}

	ap := companies.Apple{
		URL:      "https://jobs.apple.com/en-gb/search?location=united-kingdom-GBR&team=apps-and-frameworks-SFTWR-AF+cloud-and-infrastructure-SFTWR-CLD+core-operating-systems-SFTWR-COS+devops-and-site-reliability-SFTWR-DSR+engineering-project-management-SFTWR-EPM+information-systems-and-technology-SFTWR-ISTECH+machine-learning-and-ai-SFTWR-MCHLN+security-and-privacy-SFTWR-SEC+software-quality-automation-and-tools-SFTWR-SQAT+wireless-software-SFTWR-WSFT",
		FilePath: "apple.csv",
	}

	g := companies.Google{
		URL:      "https://careers.google.com/api/v2/jobs/search/?company=Google&company=Google%20Fiber&company=YouTube&employment_type=FULL_TIME&hl=en_US&jlo=en_US&location=London%2C%20UK&q=&sort_by=relevance",
		FilePath: "google.csv",
	}

	csv.WriteToCSV(g, logger)

	csv.WriteToCSV(ap, logger)
	csv.WriteToCSV(a, logger)
	csv.WriteToCSV(hs, logger)
	csv.WriteToCSV(dr, logger)
	csv.WriteToCSV(understanding, logger)
	csv.WriteToCSV(sr, logger)
	csv.WriteToCSV(cs, logger)

}

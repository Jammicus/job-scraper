package internal

import (
	"encoding/csv"
	"os"
	"strings"

	"go.uber.org/zap"
)

// WriteToCSV takes a file path, a list of jobs and a logging instance and writes those jobs to the file at the path provided
func WriteToCSV(j FindJobs, logger *zap.Logger) error {
	sugar := logger.Sugar()

	// create if does not exist, otherwise append
	file, err := os.OpenFile(j.GetPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		sugar.Fatal(err)
	}

	fInfo, err := file.Stat()

	if err != nil {
		sugar.Fatal(err)
	}
	defer file.Close()

	sugar.Infof("Using file")

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if fInfo.Size() == 0 {
		headings := [][]string{{"Title", "Job Type", "Salary", "Location", "URL", "Requirements"}}
		sugar.Infof("Writing headings to file %v", fInfo.Name())
		if err = writer.WriteAll(headings); err != nil {
			sugar.Fatal(err)

		}
		sugar.Infof("Finished writing headings to file %v", fInfo.Name())
	}

	content := createContent(j.GetJobs(logger), logger)
	sugar.Infof("Writing content to file %v", fInfo.Name())
	if err = writer.WriteAll(content); err != nil {
		sugar.Fatal(err)

	}
	sugar.Infof("Finished writing content to file %v", fInfo.Name())

	return nil
}

func createContent(jobs []Job, logger *zap.Logger) [][]string {
	content := [][]string{}

	sugar := logger.Sugar()
	for _, job := range jobs {
		values := []string{job.Title, job.Type, job.Salary, job.Location, job.URL, strings.Join(job.Requirements[:], ",")}
		sugar.Debugf("Entry created with values %v", values)
		content = append(content, values)
	}

	return content
}

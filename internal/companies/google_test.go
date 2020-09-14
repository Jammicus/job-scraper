package companies

import (
	jobs "job-scraper/internal"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGatherSpecs(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	g := Google{}

	input := googleJob{
		Title:        "Software Engineer, PhD University Graduate, Infrastructure",
		Requirements: "<p>Minimum qualifications:</p><ul>\n<li>PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.\n</li>\n<li>Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.</li>\n<li>Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.</li>\n</ul><br><p>Preferred qualifications:</p><ul>\n<li>Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  </li>\n<li>Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.</li>\n<li>Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.</li>\n<li>Experience with a large scale systems design in Unix/Linux.</li>\n<li>Authorization to legally work in the US.</li>\n<li>Ability to start in 2020 or 2021.</li>\n</ul>",
		Education:    []string{"DOCTORAL_OR_EQUIVALENT"},
		ID:           "jobs/136853555093873350",
		Locations: []struct {
			// Needs the json otherwise it complains
			Display string `json:"display"`
		}{

			{"Mountain View, CA, USA"},
			{"Sunnyvale, CA, USA"},
			{"Madison, WI, USA"},
			{"Seattle, WA, USA"},
			{"Kirkland, WA, USA"},
		},
	}

	expected := jobs.Job{
		Title:    "Software Engineer, PhD University Graduate, Infrastructure",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "Mountain View, CA, USA Sunnyvale, CA, USA Madison, WI, USA Seattle, WA, USA Kirkland, WA, USA ",
		URL:      "https://careers.google.com/jobs/results/136853555093873350",
		Requirements: []string{
			"PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.",
			"Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.",
			"Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.",
			"Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  ",
			"Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.",
			"Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.",
			"Experience with a large scale systems design in Unix/Linux.",
			"Authorization to legally work in the US.",
			"Ability to start in 2020 or 2021.",
		},
	}

	result, err := g.gatherSpecs(input, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected, result)
}

func TestGetJobURLGoogle(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	g := Google{}

	input := googleJob{
		Title:        "Software Engineer, PhD University Graduate, Infrastructure",
		Requirements: "<p>Minimum qualifications:</p><ul>\n<li>PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.\n</li>\n<li>Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.</li>\n<li>Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.</li>\n</ul><br><p>Preferred qualifications:</p><ul>\n<li>Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  </li>\n<li>Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.</li>\n<li>Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.</li>\n<li>Experience with a large scale systems design in Unix/Linux.</li>\n<li>Authorization to legally work in the US.</li>\n<li>Ability to start in 2020 or 2021.</li>\n</ul>",
		Education:    []string{"DOCTORAL_OR_EQUIVALENT"},
		ID:           "jobs/136853555093873350",
		Locations: []struct {
			// Needs the json otherwise it complains
			Display string `json:"display"`
		}{

			{"Mountain View, CA, USA"},
			{"Sunnyvale, CA, USA"},
			{"Madison, WI, USA"},
			{"Seattle, WA, USA"},
			{"Kirkland, WA, USA"},
		},
	}

	expected := jobs.Job{
		Title:    "Software Engineer, PhD University Graduate, Infrastructure",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "Mountain View, CA, USA Sunnyvale, CA, USA Madison, WI, USA Seattle, WA, USA Kirkland, WA, USA ",
		URL:      "https://careers.google.com/jobs/results/136853555093873350",
		Requirements: []string{
			"PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.",
			"Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.",
			"Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.",
			"Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  ",
			"Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.",
			"Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.",
			"Experience with a large scale systems design in Unix/Linux.",
			"Authorization to legally work in the US.",
			"Ability to start in 2020 or 2021.",
		},
	}

	result, err := g.gatherSpecs(input, logger)

	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, expected.URL, result.URL)
}

func TestGetJobTitleGoogle(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	g := Google{}

	input := googleJob{
		Title:        "Software Engineer, PhD University Graduate, Infrastructure",
		Requirements: "<p>Minimum qualifications:</p><ul>\n<li>PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.\n</li>\n<li>Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.</li>\n<li>Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.</li>\n</ul><br><p>Preferred qualifications:</p><ul>\n<li>Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  </li>\n<li>Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.</li>\n<li>Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.</li>\n<li>Experience with a large scale systems design in Unix/Linux.</li>\n<li>Authorization to legally work in the US.</li>\n<li>Ability to start in 2020 or 2021.</li>\n</ul>",
		Education:    []string{"DOCTORAL_OR_EQUIVALENT"},
		ID:           "jobs/136853555093873350",
		Locations: []struct {
			// Needs the json otherwise it complains
			Display string `json:"display"`
		}{

			{"Mountain View, CA, USA"},
			{"Sunnyvale, CA, USA"},
			{"Madison, WI, USA"},
			{"Seattle, WA, USA"},
			{"Kirkland, WA, USA"},
		},
	}

	expected := jobs.Job{
		Title:    "Software Engineer, PhD University Graduate, Infrastructure",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "Mountain View, CA, USA Sunnyvale, CA, USA Madison, WI, USA Seattle, WA, USA Kirkland, WA, USA ",
		URL:      "https://careers.google.com/jobs/results/136853555093873350",
		Requirements: []string{
			"PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.",
			"Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.",
			"Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.",
			"Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  ",
			"Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.",
			"Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.",
			"Experience with a large scale systems design in Unix/Linux.",
			"Authorization to legally work in the US.",
			"Ability to start in 2020 or 2021.",
		},
	}

	result := g.getJobTitle(input)

	assert.Equal(t, expected.Title, result)
}

func TestGetJobTypeGoogle(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	g := Google{}

	input := googleJob{
		Title:        "Software Engineer, PhD University Graduate, Infrastructure",
		Requirements: "<p>Minimum qualifications:</p><ul>\n<li>PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.\n</li>\n<li>Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.</li>\n<li>Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.</li>\n</ul><br><p>Preferred qualifications:</p><ul>\n<li>Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  </li>\n<li>Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.</li>\n<li>Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.</li>\n<li>Experience with a large scale systems design in Unix/Linux.</li>\n<li>Authorization to legally work in the US.</li>\n<li>Ability to start in 2020 or 2021.</li>\n</ul>",
		Education:    []string{"DOCTORAL_OR_EQUIVALENT"},
		ID:           "jobs/136853555093873350",
		Locations: []struct {
			// Needs the json otherwise it complains
			Display string `json:"display"`
		}{

			{"Mountain View, CA, USA"},
			{"Sunnyvale, CA, USA"},
			{"Madison, WI, USA"},
			{"Seattle, WA, USA"},
			{"Kirkland, WA, USA"},
		},
	}

	expected := jobs.Job{
		Title:    "Software Engineer, PhD University Graduate, Infrastructure",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "Mountain View, CA, USA Sunnyvale, CA, USA Madison, WI, USA Seattle, WA, USA Kirkland, WA, USA ",
		URL:      "https://careers.google.com/jobs/results/136853555093873350",
		Requirements: []string{
			"PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.",
			"Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.",
			"Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.",
			"Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  ",
			"Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.",
			"Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.",
			"Experience with a large scale systems design in Unix/Linux.",
			"Authorization to legally work in the US.",
			"Ability to start in 2020 or 2021.",
		},
	}

	result := g.getJobType(input)

	assert.Equal(t, expected.Type, result)
}

func TestGetJobLocationGoogle(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	g := Google{}

	input := googleJob{
		Title:        "Software Engineer, PhD University Graduate, Infrastructure",
		Requirements: "<p>Minimum qualifications:</p><ul>\n<li>PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.\n</li>\n<li>Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.</li>\n<li>Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.</li>\n</ul><br><p>Preferred qualifications:</p><ul>\n<li>Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  </li>\n<li>Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.</li>\n<li>Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.</li>\n<li>Experience with a large scale systems design in Unix/Linux.</li>\n<li>Authorization to legally work in the US.</li>\n<li>Ability to start in 2020 or 2021.</li>\n</ul>",
		Education:    []string{"DOCTORAL_OR_EQUIVALENT"},
		ID:           "jobs/136853555093873350",
		Locations: []struct {
			// Needs the json otherwise it complains
			Display string `json:"display"`
		}{

			{"Mountain View, CA, USA"},
			{"Sunnyvale, CA, USA"},
			{"Madison, WI, USA"},
			{"Seattle, WA, USA"},
			{"Kirkland, WA, USA"},
		},
	}

	expected := jobs.Job{
		Title:    "Software Engineer, PhD University Graduate, Infrastructure",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "Mountain View, CA, USA Sunnyvale, CA, USA Madison, WI, USA Seattle, WA, USA Kirkland, WA, USA ",
		URL:      "https://careers.google.com/jobs/results/136853555093873350",
		Requirements: []string{
			"PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.",
			"Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.",
			"Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.",
			"Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  ",
			"Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.",
			"Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.",
			"Experience with a large scale systems design in Unix/Linux.",
			"Authorization to legally work in the US.",
			"Ability to start in 2020 or 2021.",
		},
	}

	result := g.getJobLocation(input)

	assert.Equal(t, expected.Location, result)
}

func TestGetJobSalaryGoogle(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	g := Google{}

	input := googleJob{
		Title:        "Software Engineer, PhD University Graduate, Infrastructure",
		Requirements: "<p>Minimum qualifications:</p><ul>\n<li>PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.\n</li>\n<li>Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.</li>\n<li>Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.</li>\n</ul><br><p>Preferred qualifications:</p><ul>\n<li>Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  </li>\n<li>Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.</li>\n<li>Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.</li>\n<li>Experience with a large scale systems design in Unix/Linux.</li>\n<li>Authorization to legally work in the US.</li>\n<li>Ability to start in 2020 or 2021.</li>\n</ul>",
		Education:    []string{"DOCTORAL_OR_EQUIVALENT"},
		ID:           "jobs/136853555093873350",
		Locations: []struct {
			// Needs the json otherwise it complains
			Display string `json:"display"`
		}{

			{"Mountain View, CA, USA"},
			{"Sunnyvale, CA, USA"},
			{"Madison, WI, USA"},
			{"Seattle, WA, USA"},
			{"Kirkland, WA, USA"},
		},
	}

	expected := jobs.Job{
		Title:    "Software Engineer, PhD University Graduate, Infrastructure",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "Mountain View, CA, USA Sunnyvale, CA, USA Madison, WI, USA Seattle, WA, USA Kirkland, WA, USA ",
		URL:      "https://careers.google.com/jobs/results/136853555093873350",
		Requirements: []string{
			"PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.",
			"Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.",
			"Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.",
			"Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  ",
			"Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.",
			"Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.",
			"Experience with a large scale systems design in Unix/Linux.",
			"Authorization to legally work in the US.",
			"Ability to start in 2020 or 2021.",
		},
	}

	result := g.getJobSalary(input)

	assert.Equal(t, expected.Salary, result)
}

func TestGetJobRequirementsGoogle(t *testing.T) {
	logger, err := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	if err != nil {
		sugar.Fatalf("Unable to create logger")
	}

	g := Google{}

	input := googleJob{
		Title:        "Software Engineer, PhD University Graduate, Infrastructure",
		Requirements: "<p>Minimum qualifications:</p><ul>\n<li>PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.\n</li>\n<li>Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.</li>\n<li>Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.</li>\n</ul><br><p>Preferred qualifications:</p><ul>\n<li>Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  </li>\n<li>Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.</li>\n<li>Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.</li>\n<li>Experience with a large scale systems design in Unix/Linux.</li>\n<li>Authorization to legally work in the US.</li>\n<li>Ability to start in 2020 or 2021.</li>\n</ul>",
		Education:    []string{"DOCTORAL_OR_EQUIVALENT"},
		ID:           "jobs/136853555093873350",
		Locations: []struct {
			// Needs the json otherwise it complains
			Display string `json:"display"`
		}{

			{"Mountain View, CA, USA"},
			{"Sunnyvale, CA, USA"},
			{"Madison, WI, USA"},
			{"Seattle, WA, USA"},
			{"Kirkland, WA, USA"},
		},
	}

	expected := jobs.Job{
		Title:    "Software Engineer, PhD University Graduate, Infrastructure",
		Type:     "Permanent",
		Salary:   "N/A",
		Location: "Mountain View, CA, USA Sunnyvale, CA, USA Madison, WI, USA Seattle, WA, USA Kirkland, WA, USA ",
		URL:      "https://careers.google.com/jobs/results/136853555093873350",
		Requirements: []string{
			"PhD degree in Computer Science, Engineering, Mathematics, Information Technology, or equivalent practical experience.",
			"Examples of coding in one of the following programming languages including but not limited to: C, C++, Java, Python.",
			"Experience in one or more of the following: architecting and/or developing large scale distributed systems, concurrency, multithreading or synchronization.",
			"Experience with TCP/IP and network programming.  Also with database internals, database language theories, database design, SQL and database programming.  ",
			"Understanding of technologies such as virtualization and global infrastructure, load balancing, networking, massive data storage, Hadoop, MapReduce and security.",
			"Interest in or exposure to networking technologies/concepts such as Software Defined Networking (SDN) and OpenFlow.",
			"Experience with a large scale systems design in Unix/Linux.",
			"Authorization to legally work in the US.",
			"Ability to start in 2020 or 2021.",
		},
	}

	result := g.getJobRequirements(input)

	assert.Equal(t, expected.Requirements, result)
}

package helpers

import (
	"regexp"
	"strings"
)

type AppInstanceRow struct {
	Index  string
	State  string
	Since  string
	CPU    string
	Memory string
	Disk   string
}

type AppProcessTable struct {
	Title     string
	Instances []AppInstanceRow
}

type AppTable struct {
	Processes []AppProcessTable
}

func ParseV3AppTable(input []byte) AppTable {
	appTable := AppTable{}

	rows := strings.Split(string(input), "\n")
	for i, row := range rows {
		if i < 10 {
			// this is where the processes start
			continue
		}

		if row == "" {
			continue
		}

		if strings.HasPrefix(row, "#") {
			// instance row
			columns := splitColumns(row)
			instanceRow := AppInstanceRow{
				Index:  columns[0],
				State:  columns[1],
				Since:  columns[2],
				CPU:    columns[3],
				Memory: columns[4],
				Disk:   columns[5],
			}
			lastProcessIndex := len(appTable.Processes) - 1
			appTable.Processes[lastProcessIndex].Instances = append(
				appTable.Processes[lastProcessIndex].Instances,
				instanceRow,
			)

		} else if !strings.HasPrefix(row, " ") {
			// process title
			appTable.Processes = append(appTable.Processes, AppProcessTable{
				Title: row,
			})
		} else {
			// column headers
			continue
		}

	}

	return appTable
}

func splitColumns(row string) []string {
	// uses 3 spaces between columns
	return regexp.MustCompile(`\s{3,}`).Split(strings.TrimSpace(row), -1)
}

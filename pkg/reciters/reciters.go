package reciters

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
)

//go:embed reciters.csv
var reciterFS embed.FS

type Reciter struct {
	FullName string
	Name     string
	ID       string
}

func GetReciters() ([]Reciter, error) {
	f, err := reciterFS.Open("reciters.csv")
	if err != nil {
		return nil, fmt.Errorf("failed to open reciters.csv: %w", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	var reciters []Reciter

	// Skip header
	if _, err := csvReader.Read(); err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV record: %w", err)
		}
		reciters = append(reciters, Reciter{
			FullName: record[0],
			Name:     record[1],
			ID:       record[2],
		})
	}

	sort.Slice(reciters, func(i, j int) bool {
		return strings.ToLower(reciters[i].FullName) < strings.ToLower(reciters[j].FullName)
	})

	return reciters, nil
}

func DisplayRecitersTable(w io.Writer) error {
	reciters, err := GetReciters()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"ID", "Full Name", "Short Name"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	for _, r := range reciters {
		table.Append([]string{r.ID, r.FullName, r.Name})
	}

	table.Render()
	return nil
}

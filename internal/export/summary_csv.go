package export

import (
	"encoding/csv"
	"os"
	"strconv"

	"sofiasoft/internal/summary"
)

func WriteEmotionSummaryCSV(filename string, summaries []summary.EmotionSummary) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"type",
		"name",
		"count",
		"average_err",
	}

	if err := writer.Write(header); err != nil {
		return err
	}

	for _, item := range summaries {
		row := []string{
			item.Type,
			item.Name,
			strconv.Itoa(item.Count),
			strconv.FormatFloat(item.AverageERR, 'f', 2, 64),
		}

		if err := writer.Write(row); err != nil {
			return err
		}
	}

	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}

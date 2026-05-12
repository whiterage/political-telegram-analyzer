package export

import (
	"encoding/csv"
	"os"
	"strconv"

	"sofiasoft/internal/summary"
)

func WriteChannelSummaryCSV(filename string, summaries []summary.ChannelSummary) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{
		"channel_name",
		"count",
		"average_views",
		"average_reactions",
		"average_err",
	}

	if err := writer.Write(header); err != nil {
		return err
	}

	for _, item := range summaries {
		row := []string{
			item.ChannelName,
			strconv.Itoa(item.Count),
			strconv.FormatFloat(item.AverageViews, 'f', 2, 64),
			strconv.FormatFloat(item.AverageReactions, 'f', 2, 64),
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

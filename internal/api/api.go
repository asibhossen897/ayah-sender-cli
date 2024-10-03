package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

type Chapter struct {
	Chapter struct {
		NameSimple  string `json:"name_simple"`
		VersesCount int    `json:"verses_count"`
	} `json:"chapter"`
}

func GetReciterName(reciterID string) string {
	recitersFile := filepath.Join("..", "..", "ayah_sender", "reciters.csv") // Updated path
	file, err := os.Open(recitersFile)
	if err != nil {
		fmt.Printf("Error opening reciters file: %v\n", err)
		return ""
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading reciters file: %v\n", err)
		return ""
	}

	for _, record := range records {
		if record[2] == reciterID {
			return record[0]
		}
	}

	return ""
}

func GetChapterName(chapterNum string) string {
	url := fmt.Sprintf("https://api.quran.com/api/v4/chapters/%s", chapterNum)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching chapter info: %v\n", err)
		return ""
	}
	defer resp.Body.Close()

	var chapter Chapter
	if err := json.NewDecoder(resp.Body).Decode(&chapter); err != nil {
		fmt.Printf("Error decoding chapter info: %v\n", err)
		return ""
	}

	return chapter.Chapter.NameSimple
}

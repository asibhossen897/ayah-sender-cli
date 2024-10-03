package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asibhossen897/ayah-sender-cli/pkg/logger"
	"github.com/asibhossen897/ayah-sender-cli/pkg/reciters"
)

type Chapter struct {
	Chapter struct {
		NameSimple  string `json:"name_simple"`
		VersesCount int    `json:"verses_count"`
	} `json:"chapter"`
}

func GetReciterName(reciterID string) string {
	reciters, err := reciters.GetReciters()
	if err != nil {
		logger.Error("Error getting reciters", "error", err)
		return reciterID // Return reciterID as fallback
	}

	for _, r := range reciters {
		if r.ID == reciterID {
			return r.FullName // Return FullName instead of Name
		}
	}

	logger.Info("Reciter not found, using ID", "reciterID", reciterID)
	return reciterID // Return reciterID if not found
}

func GetChapterName(chapterNum string) string {
	url := fmt.Sprintf("https://api.quran.com/api/v4/chapters/%s", chapterNum)
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error fetching chapter info", "error", err)
		return ""
	}
	defer resp.Body.Close()

	var chapter Chapter
	if err := json.NewDecoder(resp.Body).Decode(&chapter); err != nil {
		logger.Error("Error decoding chapter info", "error", err)
		return ""
	}

	return chapter.Chapter.NameSimple
}

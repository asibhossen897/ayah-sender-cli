package audio

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/asibhossen897/ayah-sender-cli/internal/api"
)

func DownloadAudio(reciterID, chapterNum, startVerse, endVerse string) {
	baseURL := "https://everyayah.com/data"
	reciterName := api.GetReciterName(reciterID)
	chapterName := api.GetChapterName(chapterNum)

	startVerseInt, _ := strconv.Atoi(startVerse)
	endVerseInt, _ := strconv.Atoi(endVerse)

	for verse := startVerseInt; verse <= endVerseInt; verse++ {
		verseStr := fmt.Sprintf("%03d", verse)
		url := fmt.Sprintf("%s/%s/%s%s.mp3", baseURL, reciterName, chapterNum, verseStr)
		fileName := fmt.Sprintf("%s_Surah_%s(%s)_Ayah%s.mp3", reciterName, chapterName, chapterNum, verseStr)

		fmt.Printf("Downloading %s...\n", fileName)
		if err := downloadFile(url, fileName); err != nil {
			fmt.Printf("Error downloading %s: %v\n", fileName, err)
		} else {
			fmt.Printf("%s downloaded successfully!\n", fileName)
		}
	}
}

func downloadFile(url, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

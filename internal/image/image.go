package image

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/asibhossen897/ayah-sender-cli/internal/api"
)

func DownloadImage(chapterNum, verseNum string) {
	baseURL := "https://everyayah.com/data/images_png"
	chapterName := api.GetChapterName(chapterNum)

	url := fmt.Sprintf("%s/%s_%s.png", baseURL, chapterNum, verseNum)
	fileName := fmt.Sprintf("Surah_%s(%s)_%s.png", chapterName, chapterNum, verseNum)

	fmt.Printf("Downloading %s...\n", fileName)
	if err := downloadFile(url, fileName); err != nil {
		fmt.Printf("Error downloading %s: %v\n", fileName, err)
	} else {
		fmt.Printf("%s downloaded successfully!\n", fileName)
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

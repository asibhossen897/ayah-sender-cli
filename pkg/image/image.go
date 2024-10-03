package image

import (
	"fmt"
	"path/filepath"

	"github.com/asibhossen897/ayah-sender-cli/pkg/api"
	"github.com/asibhossen897/ayah-sender-cli/pkg/downloader"
	"github.com/asibhossen897/ayah-sender-cli/pkg/logger"
)

func DownloadImage(chapterNum, verseNum, downloadPath string) error {
	baseURL := "https://everyayah.com/data/images_png"
	chapterName := api.GetChapterName(chapterNum)

	url := fmt.Sprintf("%s/%s_%s.png", baseURL, chapterNum, verseNum)
	fileName := fmt.Sprintf("Surah_%s(%s)_%s.png", chapterName, chapterNum, verseNum)
	filePath := filepath.Join(downloadPath, fileName)

	logger.Info("Downloading image", "file", fileName)
	if err := downloader.DownloadFile(url, filePath); err != nil {
		logger.Error("Failed to download image", "file", fileName, "error", err)
		return err
	}
	logger.Info("Image downloaded successfully", "file", fileName)
	return nil
}

package audio

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/asibhossen897/ayah-sender-cli/pkg/api"
	"github.com/asibhossen897/ayah-sender-cli/pkg/downloader"
	"github.com/asibhossen897/ayah-sender-cli/pkg/logger"
)

func DownloadAudio(reciterID, chapterNum, startVerse, endVerse string) error {
	baseURL := "https://everyayah.com/data"
	reciterName := api.GetReciterName(reciterID)
	chapterName := api.GetChapterName(chapterNum)

	startVerseInt, err := strconv.Atoi(startVerse)
	if err != nil {
		return fmt.Errorf("invalid start verse: %w", err)
	}
	endVerseInt, err := strconv.Atoi(endVerse)
	if err != nil {
		return fmt.Errorf("invalid end verse: %w", err)
	}

	// Convert chapterNum to int and format with leading zeros
	chapterInt, err := strconv.Atoi(chapterNum)
	if err != nil {
		return fmt.Errorf("invalid chapter number: %w", err)
	}
	formattedChapter := fmt.Sprintf("%03d", chapterInt)

	for verse := startVerseInt; verse <= endVerseInt; verse++ {
		formattedVerse := fmt.Sprintf("%03d", verse)
		url := fmt.Sprintf("%s/%s/%s%s.mp3", baseURL, reciterName, formattedChapter, formattedVerse)
		fileName := fmt.Sprintf("%s_Surah_%s(%s)_Ayah%s.mp3", reciterName, chapterName, chapterNum, formattedVerse)

		logger.Info("Downloading audio", "file", fileName, "url", url)
		if err := downloader.DownloadFile(url, fileName); err != nil {
			logger.Error("Failed to download audio", "file", fileName, "error", err)
		} else {
			logger.Info("Audio downloaded successfully", "file", fileName)
		}
	}

	return nil
}

func DownloadAndMergeAudio(reciterID, chapterNum, startVerse, endVerse string) error {
	baseURL := "https://everyayah.com/data"
	reciterName := api.GetReciterName(reciterID)
	chapterName := api.GetChapterName(chapterNum)

	startVerseInt, err := strconv.Atoi(startVerse)
	if err != nil {
		return fmt.Errorf("invalid start verse: %w", err)
	}
	endVerseInt, err := strconv.Atoi(endVerse)
	if err != nil {
		return fmt.Errorf("invalid end verse: %w", err)
	}

	chapterInt, err := strconv.Atoi(chapterNum)
	if err != nil {
		return fmt.Errorf("invalid chapter number: %w", err)
	}
	formattedChapter := fmt.Sprintf("%03d", chapterInt)

	tempDir, err := os.MkdirTemp("", "quran-audio")
	if err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	var audioData [][]byte

	for verse := startVerseInt; verse <= endVerseInt; verse++ {
		formattedVerse := fmt.Sprintf("%03d", verse)
		url := fmt.Sprintf("%s/%s/%s%s.mp3", baseURL, reciterName, formattedChapter, formattedVerse)
		fileName := filepath.Join(tempDir, fmt.Sprintf("%s%s.mp3", formattedChapter, formattedVerse))

		logger.Info("Downloading audio", "file", fileName, "url", url)
		if err := downloader.DownloadFile(url, fileName); err != nil {
			logger.Error("Failed to download audio", "file", fileName, "error", err)
			continue
		}

		data, err := os.ReadFile(fileName)
		if err != nil {
			logger.Error("Failed to read audio file", "file", fileName, "error", err)
			continue
		}
		audioData = append(audioData, data)
	}

	if len(audioData) == 0 {
		return fmt.Errorf("no audio files were downloaded")
	}

	outputFileName := fmt.Sprintf("%s_Surah_%s(%s)_Ayah%s-%s.mp3", reciterName, chapterName, chapterNum, startVerse, endVerse)
	err = mergeAudioFiles(audioData, outputFileName)
	if err != nil {
		return fmt.Errorf("failed to merge audio files: %w", err)
	}

	logger.Info("Audio files merged successfully", "output", outputFileName)
	return nil
}

func mergeAudioFiles(audioData [][]byte, outputFile string) error {
	out, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer out.Close()

	for _, data := range audioData {
		_, err := io.Copy(out, bytes.NewReader(data))
		if err != nil {
			return fmt.Errorf("failed to write audio data: %w", err)
		}
	}

	return nil
}

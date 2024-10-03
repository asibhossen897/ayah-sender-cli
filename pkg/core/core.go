package core

import (
	"github.com/asibhossen897/ayah-sender-cli/pkg/audio"
	"github.com/asibhossen897/ayah-sender-cli/pkg/image"
	"github.com/asibhossen897/ayah-sender-cli/pkg/reciters"
)

type AyahSender struct {
	// Add any shared state here if needed
}

func NewAyahSender() *AyahSender {
	return &AyahSender{}
}

func (as *AyahSender) DownloadAudio(reciterID, chapterNum, startVerse, endVerse, downloadPath string) error {
	return audio.DownloadAudio(reciterID, chapterNum, startVerse, endVerse, downloadPath)
}

func (as *AyahSender) DownloadAndMergeAudio(reciterID, chapterNum, startVerse, endVerse, downloadPath string) error {
	return audio.DownloadAndMergeAudio(reciterID, chapterNum, startVerse, endVerse, downloadPath)
}

func (as *AyahSender) DownloadImage(chapterNum, verseNum, downloadPath string) error {
	return image.DownloadImage(chapterNum, verseNum, downloadPath)
}

func (as *AyahSender) GetReciters() ([]reciters.Reciter, error) {
	return reciters.GetReciters()
}

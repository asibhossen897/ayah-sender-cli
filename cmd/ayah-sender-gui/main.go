package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"github.com/asibhossen897/ayah-sender-cli/pkg/api"
	"github.com/asibhossen897/ayah-sender-cli/pkg/core"
)

func main() {
	ayahSender := core.NewAyahSender()
	a := app.New()
	w := a.NewWindow("Ayah Sender")

	// Create input fields
	reciterSelect := widget.NewSelect([]string{}, func(value string) {})
	chapterEntry := widget.NewEntry()
	startVerseEntry := widget.NewEntry()
	endVerseEntry := widget.NewEntry()
	downloadPathEntry := widget.NewEntry()
	downloadPathEntry.SetPlaceHolder("Select download path...")

	// Populate reciter select
	recitersData, err := ayahSender.GetReciters()
	if err != nil {
		fmt.Println("Error loading reciters:", err)
	} else {
		reciterOptions := make([]string, len(recitersData))
		for i, r := range recitersData {
			reciterOptions[i] = fmt.Sprintf("%s (%s)", r.FullName, r.ID)
		}
		reciterSelect.Options = reciterOptions
	}

	// Function to update verse entries based on chapter
	updateVerseEntries := func() {
		chapterNum := chapterEntry.Text
		if chapterNum == "" {
			return
		}

		chapterInfo, err := api.GetChapterInfo(chapterNum)
		if err != nil {
			fmt.Println("Error fetching chapter info:", err)
			return
		}

		startVerseEntry.SetText("1")
		endVerseEntry.SetText(fmt.Sprintf("%d", chapterInfo.Chapter.VersesCount))
	}

	// Add onChange listener to chapterEntry
	chapterEntry.OnChanged = func(s string) {
		updateVerseEntries()
	}

	// Create buttons
	downloadAudioButton := widget.NewButton("Download Audio", func() {
		reciterID := reciterSelect.Selected
		chapterNum := chapterEntry.Text
		startVerse := startVerseEntry.Text
		endVerse := endVerseEntry.Text
		downloadPath := downloadPathEntry.Text

		// Extract reciter ID from the selected option
		if reciterID != "" {
			reciterID = reciterID[len(reciterID)-3 : len(reciterID)-1]
		}

		go func() {
			err := ayahSender.DownloadAudio(reciterID, chapterNum, startVerse, endVerse, downloadPath)
			if err != nil {
				fmt.Println("Error downloading audio:", err)
			} else {
				fmt.Println("Audio downloaded successfully")
			}
		}()
	})

	mergeAudioButton := widget.NewButton("Merge Audio", func() {
		reciterID := reciterSelect.Selected
		chapterNum := chapterEntry.Text
		startVerse := startVerseEntry.Text
		endVerse := endVerseEntry.Text
		downloadPath := downloadPathEntry.Text

		// Extract reciter ID from the selected option
		if reciterID != "" {
			reciterID = reciterID[len(reciterID)-3 : len(reciterID)-1]
		}

		go func() {
			err := ayahSender.DownloadAndMergeAudio(reciterID, chapterNum, startVerse, endVerse, downloadPath)
			if err != nil {
				fmt.Println("Error merging audio:", err)
			} else {
				fmt.Println("Audio merged successfully")
			}
		}()
	})

	downloadImageButton := widget.NewButton("Download Image", func() {
		chapterNum := chapterEntry.Text
		verseNum := startVerseEntry.Text
		downloadPath := downloadPathEntry.Text

		go func() {
			err := ayahSender.DownloadImage(chapterNum, verseNum, downloadPath)
			if err != nil {
				fmt.Println("Error downloading image:", err)
			} else {
				fmt.Println("Image downloaded successfully")
			}
		}()
	})

	selectPathButton := widget.NewButton("Select Path", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil {
				fmt.Println("Error selecting folder:", err)
				return
			}
			if uri == nil {
				return
			}
			downloadPathEntry.SetText(uri.Path())
		}, w)
	})

	// Create responsive layout
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Reciter", Widget: reciterSelect},
			{Text: "Chapter Number", Widget: chapterEntry},
			{Text: "Start Verse", Widget: startVerseEntry},
			{Text: "End Verse", Widget: endVerseEntry},
			{Text: "Download Path", Widget: container.NewBorder(nil, nil, nil, selectPathButton, downloadPathEntry)},
		},
	}

	buttons := container.NewHBox(
		layout.NewSpacer(),
		downloadAudioButton,
		mergeAudioButton,
		downloadImageButton,
		layout.NewSpacer(),
	)

	content := container.NewVBox(
		form,
		buttons,
	)

	// Use a padding container to add some space around the edges
	paddedContent := container.NewPadded(content)

	// Use a max size to ensure the content doesn't stretch too wide on large screens
	maxContent := container.New(layout.NewMaxLayout(),
		paddedContent,
		widget.NewLabel(""), // This empty label sets a minimum width
	)

	w.SetContent(maxContent)
	w.Resize(fyne.NewSize(400, 350))
	w.SetMaster()
	w.ShowAndRun()
}

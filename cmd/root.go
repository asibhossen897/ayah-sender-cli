package cmd

import (
	"fmt"
	"os"

	"github.com/asibhossen897/ayah-sender-cli/internal/audio"
	"github.com/asibhossen897/ayah-sender-cli/internal/image"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quran-cli",
	Short: "A CLI application for Quranic audio and images",
	Long:  `quran-cli is a command-line interface for downloading Quranic audio recitations and verse images.`,
}

var audioCmd = &cobra.Command{
	Use:   "audio [reciter_id] [chapter_num] [start_verse] [end_verse]",
	Short: "Download Quranic audio",
	Long:  `Download Quranic audio recitations for a specific range of verses.`,
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		reciterID := args[0]
		chapterNum := args[1]
		startVerse := args[2]
		endVerse := args[3]
		audio.DownloadAudio(reciterID, chapterNum, startVerse, endVerse)
	},
}

var imageCmd = &cobra.Command{
	Use:   "image [chapter_num] [verse_num]",
	Short: "Download Quranic verse image",
	Long:  `Download an image of a specific Quranic verse.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		chapterNum := args[0]
		verseNum := args[1]
		image.DownloadImage(chapterNum, verseNum)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(audioCmd)
	rootCmd.AddCommand(imageCmd)
}

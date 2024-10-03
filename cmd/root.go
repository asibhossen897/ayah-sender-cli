package cmd

import (
	"os"

	"github.com/asibhossen897/ayah-sender-cli/pkg/audio"
	"github.com/asibhossen897/ayah-sender-cli/pkg/config"
	"github.com/asibhossen897/ayah-sender-cli/pkg/image"
	"github.com/asibhossen897/ayah-sender-cli/pkg/logger"
	"github.com/asibhossen897/ayah-sender-cli/pkg/reciters"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ayah-sender",
	Short: "A CLI application for Quranic audio and images",
	Long:  `ayah-sender is a command-line interface for downloading Quranic audio recitations and verse images.`,
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
		if err := audio.DownloadAudio(reciterID, chapterNum, startVerse, endVerse); err != nil {
			logger.Error("Failed to download audio", "error", err)
		}
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
		if err := image.DownloadImage(chapterNum, verseNum); err != nil {
			logger.Error("Failed to download image", "error", err)
		}
	},
}

var mergeAudioCmd = &cobra.Command{
	Use:   "merge-audio [reciter_id] [chapter_num] [start_verse] [end_verse]",
	Short: "Download and merge Quranic audio",
	Long:  `Download Quranic audio recitations for a specific range of verses and merge them into a single file.`,
	Args:  cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		reciterID := args[0]
		chapterNum := args[1]
		startVerse := args[2]
		endVerse := args[3]
		if err := audio.DownloadAndMergeAudio(reciterID, chapterNum, startVerse, endVerse); err != nil {
			logger.Error("Failed to download and merge audio", "error", err)
		}
	},
}

var listRecitersCmd = &cobra.Command{
	Use:   "list-reciters",
	Short: "List all available reciters",
	Long:  `Display a table of all available reciters with their IDs and names.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := reciters.DisplayRecitersTable(os.Stdout); err != nil {
			logger.Error("Failed to display reciters", "error", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Error("Failed to execute root command", "error", err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	rootCmd.AddCommand(audioCmd)
	rootCmd.AddCommand(imageCmd)
	rootCmd.AddCommand(mergeAudioCmd)
	rootCmd.AddCommand(listRecitersCmd)
}

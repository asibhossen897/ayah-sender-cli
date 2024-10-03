Ayah Sender CLI Application
===========================

This CLI application allows you to download single or multiple Quranic audio recitations and also download verse images.

Setup:
------
1. Ensure you have Go installed on your system.
2. Clone the repository or navigate to the project directory.
3. Run `go mod tidy` to ensure all dependencies are correctly installed.

Building the application (optional):
------------------------------------
- Run `go build -o ayah-sender-cli cmd/ayah-sender-cli/main.go`
- This will create an executable named `ayah-sender-cli` in your current directory.

Running the application:
------------------------
- If you built the executable: `./ayah-sender-cli [command] [arguments]`
- If you didn't build it: `go run cmd/ayah-sender-cli/main.go [command] [arguments]`

Available Commands:
-------------------
1. Download individual audio files:
   ./ayah-sender-cli audio [reciter_id] [chapter_num] [start_verse] [end_verse]
   Example: ./ayah-sender-cli audio 3 1 1 7

2. Download and merge audio files:
   ./ayah-sender-cli merge-audio [reciter_id] [chapter_num] [start_verse] [end_verse]
   Example: ./ayah-sender-cli merge-audio 3 1 1 7

3. Download verse image:
   ./ayah-sender-cli image [chapter_num] [verse_num]
   Example: ./ayah-sender-cli image 1 1

4. List all available reciters:
   ./ayah-sender-cli list-reciters
   This command displays a table of all available reciters with their IDs and names.

Reciter IDs:
------------
- The application now includes a built-in list of reciters.
- Use the `list-reciters` command to see all available reciters and their IDs.

Output:
-------
- Audio files: ReciterName_Surah_ChapterName(ChapterNumber)_AyahNumber.mp3
- Merged audio: ReciterName_Surah_ChapterName(ChapterNumber)_AyahStartNumber-EndNumber.mp3
- Images: Surah_ChapterName(ChapterNumber)_VerseNumber.png

Troubleshooting:
----------------
- For 404 errors, verify the reciter ID, chapter number, and verse numbers.
- Check the application logs for any issues.

Examples:
---------
1. List all available reciters:
   ./ayah-sender-cli list-reciters

2. Download verses 1-3 of Surah Al-Baqarah (Chapter 2) recited by Mishary Alafasy (ID 12):
   ./ayah-sender-cli audio 12 2 1 3

3. Download and merge verses 1-10 of Surah Yasin (Chapter 36) recited by Abdul Basit Murattal (ID 3):
   ./ayah-sender-cli merge-audio 3 36 1 10

4. Download the image for verse 255 of Surah Al-Baqarah (Ayat Al-Kursi):
   ./ayah-sender-cli image 2 255

Note: Replace `./ayah-sender-cli` with `go run cmd/ayah-sender-cli/main.go` if you haven't built the executable.

Enjoy using this tool for your Quranic studies or projects!
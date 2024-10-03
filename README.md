Ayah Sender CLI Application
===========================

This CLI application allows you to download single or multiple Quranic audio recitations and also download verse images.

Setup:
------
1. Download the executable from the releases page for your operating system: https://github.com/abdullah-al-aqsa/ayah-sender-cli/releases
2. Make it executable if you're on linux: `chmod +x ayah-sender*`
3. Optionally add the executable to your PATH by running `sudo mv ayah-sender* /usr/local/bin/`
4. If you're on windows, add it to your PATH
5. Run `ayah-sender --help` to see the available commands

Building the application from source (optional):
------------------------------------
- Ensure you have Go installed on your system.
- Clone the repository or navigate to the project directory.
- Run `go mod tidy` to ensure all dependencies are correctly installed.
- Run `go build -o ayah-sender cmd/ayah-sender-cli/main.go`
- This will create an executable named `ayah-sender` in your current directory.

Running the application:
------------------------
- Open a terminal in the directory in which you want to download the files
- You can run the application by running `ayah-sender [command] [arguments]`

Available Commands:
-------------------
1. Download individual audio files:
   **ayah-sender audio [reciter_id] [chapter_num] [start_verse] [end_verse]**
   Example: 
   ```bash
   ayah-sender audio 3 1 1 7
   ```

2. Download and merge audio files:
   **ayah-sender merge-audio [reciter_id] [chapter_num] [start_verse] [end_verse]**
   Example: 
   ```bash
   ayah-sender merge-audio 3 1 1 7
   ```

3. Download verse image:
   ayah-sender image [chapter_num] [verse_num]
   Example: 
   ```bash
   ayah-sender image 1 1
   ```

4. List all available reciters:
   ayah-sender list-reciters
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
   ```bash
   ayah-sender list-reciters
   ```

2. Download verses 1-3 of Surah Al-Baqarah (Chapter 2) recited by Mishary Alafasy (ID 12):
   ```bash
   ayah-sender audio 12 2 1 3
   ```

3. Download and merge verses 1-10 of Surah Yasin (Chapter 36) recited by Abdul Basit Murattal (ID 3):
   ```bash
   ayah-sender merge-audio 3 36 1 10
   ```

4. Download the image for verse 255 of Surah Al-Baqarah (Ayat Al-Kursi):
   ```bash
   ayah-sender image 2 255
   ```

Enjoy using this tool for your Quranic studies or projects!
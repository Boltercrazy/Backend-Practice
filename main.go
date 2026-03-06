package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	for {
		noteText, err := promptForNote()
		if err != nil {
			fmt.Println("error reading note:", err)
			return
		}

		if noteText == "" {
			fmt.Println("Goodbye. All entrys saved.")
			return
		}

		entry := formatEntry(noteText)

		err = appendToJournal(entry)
		if err != nil {
			fmt.Println("error writing journal:", err)
			return
		}

		fmt.Println("Entry saved successfully at " + time.Now().Format("15:04:05"))

	}
}

// asks user fir note
func promptForNote() (string, error) {
	fmt.Println("enter today's note: ")

	reader := bufio.NewReader(os.Stdin)

	noteText, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	noteText = strings.TrimSpace(noteText)
	return noteText, nil
}

// adds timestamp to note
func formatEntry(noteText string) string {
	formattedTime := time.Now().Format("2006-01-02 15:04:05")
	return formattedTime + " - " + noteText + "\n"

}

// append entry to journal.txt
func appendToJournal(entry string) error {
	file, err := os.OpenFile("journal.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(entry)
	if err != nil {
		return err
	}

	return nil
}

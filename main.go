package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	wd, _ := os.Getwd()
	fmt.Println("Working directory:", wd)

	fmt.Println("enter today's note: ")

	reader := bufio.NewReader(os.Stdin)

	note, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error reading input:", err)
		return
	}

	// add time to note
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	entry := formattedTime + " - " + note

	// open or create journal.txt for appending
	file, err := os.OpenFile("journal.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(entry)
	if err != nil {
		fmt.Println("error writing to file:", err)
		return
	}

	fmt.Println("You Wrote:")
	fmt.Println(note)
}

package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"unicode/utf16"
)

type Entry struct {
	Name        string
	Description string
}

func loadCSV() ([]Entry, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("unable to get executable path: %v", err)
	}

	csvPath := filepath.Join(filepath.Dir(exePath), "snippets.csv")

	file, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV at %s: %v", csvPath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %v", err)
	}

	var entries []Entry
	for _, record := range records {
		if len(record) >= 2 {
			entries = append(entries, Entry{
				Name:        record[0],
				Description: record[1],
			})
		}
	}
	return entries, nil
}

func filterEntries(entries []Entry, query string) []Entry {
	var filtered []Entry

	// Create a regex pattern to match the query characters in order, case-insensitive
	pattern := "(?i)" // case-insensitive flag
	for _, char := range query {
		pattern += string(char) + ".*" // Match the character and allow for any characters in between
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return filtered
	}

	for _, entry := range entries {
		// Match the name against the regular expression
		if re.MatchString(entry.Name) {
			filtered = append(filtered, entry)
			if len(filtered) >= 5 {
				break
			}
		}
	}
	return filtered
}

// copyToClipboard encodes text to UTF-16LE and copies it via the clipboard
func copyToClipboard(text string) error {
	// Convert to UTF-16LE
	utf16Encoded := utf16.Encode([]rune(text))
	buf := new(bytes.Buffer)
	for _, r := range utf16Encoded {
		buf.WriteByte(byte(r))
		buf.WriteByte(byte(r >> 8))
	}

	// Send to Windows clipboard
	cmd := exec.Command("cmd", "/c", "clip")
	cmd.Stdin = buf
	return cmd.Run()
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func searchAndCopy(entries []Entry) {
	reader := bufio.NewReader(os.Stdin)
	var lastMatches []Entry // Store the results of the most recent search

	for {
		// Prompt for input
		fmt.Print("Enter search query: ")
		query, _ := reader.ReadString('\n')

		// Remove the newline character at the end of the query
		query = strings.TrimSpace(query)

		// If the query is empty, continue to prompt for input
		if query == "" {
			continue
		}

		// Filter entries based on the query
		lastMatches = filterEntries(entries, query)

		// Clear the screen after the search has been entered
		clearScreen()

		// Display search results
		if len(lastMatches) > 0 {
			fmt.Println("Matches:")
			for i, m := range lastMatches {
				// Display index as single digits (1, 2, 3, ...)
				fmt.Printf("%d: %s\n", i+1, m.Name)
			}

			// Prompt for selection or a new search
			fmt.Print("\nSelect option (1-5) or press enter for new search: ")
			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice) //remove newline char

			switch choice {
			case "1", "2", "3", "4", "5":
				index := int(choice[0] - '1') // Convert '1', '2', or '3' into 0, 1, or 2
				if index < len(lastMatches) {
					err := copyToClipboard(lastMatches[index].Description)
					if err != nil {
						fmt.Println("Error copying to clipboard:", err)
					} else {
						fmt.Println("Description copied to clipboard.")
					}
					// Exit after successful selection
					os.Exit(0)
				} else {
					fmt.Println("Invalid selection.")
				}
			default:
				// If the user enters a string (new search), start a new search
				if choice != "" {
					fmt.Println("Invalid input.")
					continue // Start a new search cycle
				} else {
					fmt.Println("\nPerforming a new search...")
				}
			}
		} else {
			fmt.Println("No matches found.")
		}
	}
}

func main() {
	entries, err := loadCSV()
	if err != nil {
		fmt.Println("Failed to load CSV:", err)
		return
	}
	searchAndCopy(entries)
}

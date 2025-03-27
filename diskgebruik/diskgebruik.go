package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	FileTypes map[string][]string `json:"fileTypes"`
}

var (
	config         Config
	sizeByCategory = make(map[string]int64) 
)

func loadConfig(filename string) error {
	data, err := os.ReadFile(filename) 
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &config)
}

func categorizeFile(filePath string, size int64) {
	extension := strings.ToLower(filepath.Ext(filePath))

	for category, extensions := range config.FileTypes {
		for _, ext := range extensions {
			if extension == ext {
				sizeByCategory[category] += size
				return
			}
		}
	}

	sizeByCategory["Overig"] += size
}

func scanDirectory(path string) error {
	return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			categorizeFile(filePath, info.Size())
		}
		return nil
	})
}

func formatSize(size int64) string {
	const gb = 1 << 30
	const mb = 1 << 20

	if size > gb {
		return fmt.Sprintf("%.1f GB", float64(size)/float64(gb))
	}
	return fmt.Sprintf("%.1f MB", float64(size)/float64(mb))
}

func getPathFromUser() string {
	fmt.Print("Voer het pad in van de map die je wilt scannen: ")
	var path string
	fmt.Scanln(&path) 
	return strings.TrimSpace(path)
}

func printResults() {
	fmt.Println("\nDiskgebruik per categorie:")
	for category, size := range sizeByCategory {
		fmt.Printf("%s: %s\n", category, formatSize(size))
	}
}

func main() {
	if err := loadConfig("config.json"); err != nil {
		fmt.Println("Fout bij het laden van de configuratie:", err)
		return
	}

	fmt.Println("Configuratie geladen.")

	for {
		path := getPathFromUser()

		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("Fout: De map bestaat niet. Probeer opnieuw.")
			continue
		}

		if err := scanDirectory(path); err != nil {
			fmt.Println("Fout bij het scannen:", err)
			continue
		}

		printResults()

		fmt.Print("\nWil je nog een map scannen? (ja/nee): ")
		var antwoord string
		fmt.Scanln(&antwoord)

		if strings.ToLower(antwoord) != "ja" {
			break
		}
	}
}

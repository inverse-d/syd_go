package main

import ( 
	"fmt"
	"os"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	content := ImportConfig("~/test")
	fmt.Println(content)
}
func ImportConfig(filename string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(filename, "~") {
		fileContent, err := os.ReadFile(filepath.Join(filename))
		if err != nil {
			log.Fatal(err)	
		}
		return string(fileContent)
	}
	fileContent, err := os.ReadFile(strings.Replace(filename, "~", homeDir, 1))
	if err != nil {
		log.Fatal(err)
	}
	return string(fileContent)
}

func BackupDotfiles() {

}

func RestoreDotfiles() {

}

func CreateLocalRepo() {

}

func PushToGit() {

}


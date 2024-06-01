package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"path/filepath"
)

func main() {
	config := "~/.config/syd/syd.conf"
	content, err := ReadConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(content); i++ {
		fmt.Println(content[i])
	}
}
func ReadConfig(filename string) ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	if !strings.Contains(filename, "~") {
		file, err := os.Open(filepath.Join(filename))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		var dotFiles []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			dotFiles = append(dotFiles, scanner.Text())
		}
		return dotFiles, scanner.Err()
	}
	file, err := os.Open(strings.Replace(filename, "~", homeDir, 1))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		var dotFiles []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			dotFiles = append(dotFiles, scanner.Text())
		}
		return dotFiles, scanner.Err()
	
}

func BackupDotfiles() {

}

func RestoreDotfiles() {

}

func CreateLocalRepo() {

}

func PushToGit() {

}


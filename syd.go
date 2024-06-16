package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	configFile := FormatPath("~/.config/syd/syd.conf")
	backupFolder := FormatPath("~/syd")
	dotFiles, err := ReadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	CreateBackupFolder(backupFolder)
	BackupDotfiles(dotFiles, FormatPath(backupFolder))
}

func FormatPath(homePath string) string {
	if !strings.Contains(homePath, "~") {
		homePathFormated := filepath.Join(homePath)
		return homePathFormated
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	homePathFormated := strings.Replace(homePath, "~", homeDir, 1)
	return homePathFormated
}

func ReadConfig(configFile string) ([]string, error) {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var dotFiles []string
	for scanner.Scan() {
		dotFiles = append(dotFiles, scanner.Text())
	}
	return dotFiles, scanner.Err()
}

func CreateBackupFolder(backupFolder string) bool {
	if _, err := os.Stat(backupFolder); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(backupFolder, 0755)
			fmt.Println("Created folder:", backupFolder)
			return true
		} else {
			log.Fatal(err)
		}
	}
	return true
}

func BackupDotfiles(dotFiles []string, backupFolder string) {
	for _,file := range dotFiles {
		file = FormatPath(file)
		source, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer source.Close()
		fileName := filepath.Base(file)
		destination := filepath.Join(backupFolder, fileName)
		destinationFile, err := os.Create(destination)
		if err != nil {
			log.Fatal(err)
		}
		defer destinationFile.Close()
		_, err = io.Copy(destinationFile, source)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Copied file: ", file)
	}
}

func RestoreDotfiles() {

}

func CreateLocalRepo() {

}

func PushToGit() {

}


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
	configFile := "~/.config/syd/syd.conf"
	backupFolder := "~/syd"
	dotFiles, err := ReadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	CreateBackupFolder(backupFolder)
	fmt.Println(dotFiles)
}
func FormatePath(homePath string) string {
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
	configFile = FormatePath(configFile)
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
	backupFolder = FormatePath(backupFolder)
	if _, err := os.Stat(backupFolder); err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(backupFolder, 0755)
			fmt.Println("Created folder:", backupFolder)
			return true
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("Folder", backupFolder, "already exists.")
	return true
}

func BackupDotfiles(dotFiles []string, backupFolder string) {

}

func RestoreDotfiles() {

}

func CreateLocalRepo() {

}

func PushToGit() {

}


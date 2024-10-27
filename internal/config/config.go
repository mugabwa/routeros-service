package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func Load() {
	cwd, err := os.Getwd()
    if err != nil {
        log.Fatalf("Error getting current working directory: %v", err)
    }
	dir := cwd
	files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatalf("Error reading directory: %v", err)
    }
    fmt.Println("Current working directory:", cwd)
	for _, file := range files {
        // Print file or directory name
        fmt.Println(file.Name())

        // You can also check if it's a directory or a file
        if file.IsDir() {
            fmt.Println(file.Name(), "is a directory")
        } else {
            fmt.Println(file.Name(), "is a file")
        }
    }
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file, %v", err)
	}
}

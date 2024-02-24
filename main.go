package main

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

func main() {
	var path string

	if len(os.Args) == 2 {
		path = os.Args[1]
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		path = pwd
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	configPaths := []string{"glow.json", "package.json"}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		var content []byte
		for _, configPath := range configPaths {
			glowConfigPath := fmt.Sprintf("%s/%s/%s", path, entry.Name(), configPath)
			fileContent, err := os.ReadFile(glowConfigPath)
			if err != nil {
				continue
			}

			content = fileContent
			break
		}

		if content == nil {
			continue
		}

		var project Project
		if err := json.Unmarshal(content, &project); err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("┌─ %s - v%s\n%s\n\n", project.Name, project.Version, cmp.Or(project.Description, "No description..."))
	}
}

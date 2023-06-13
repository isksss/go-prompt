package config

import (
	"encoding/json"
	"log"
	"os"
)

func CreateDefaultConfig() {
	filepath := GetConfigFilePath()

	prompt := Prompt{
		FgColor: "black",
		BgColor: "white",
	}

	config := Config{
		Prompt: prompt,
	}

	file, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		log.Fatal(err)
	}
}

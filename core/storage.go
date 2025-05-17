package core

import (
	"encoding/json"
	"fmt"
	"os"
	"Tubes-IdeaManagerCLI/model"
)

const dataFile = "data/data.json"

func SaveToFile(arr [model.MaxIdeas]model.Idea, count int) {
	temp := arr[:count]
	file, err := os.Create(dataFile)
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(temp)
}

func LoadFromFile(arr *[model.MaxIdeas]model.Idea, count *int) {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("Data kosong atau gagal dibuka.")
		return
	}
	defer file.Close()

	var temp []model.Idea
	err = json.NewDecoder(file).Decode(&temp)
	if err != nil {
		fmt.Println("Gagal membaca data:", err)
		return
	}

	for i := 0; i < len(temp) && i < model.MaxIdeas; i++ {
		arr[i] = temp[i]
	}
	*count = len(temp)
}
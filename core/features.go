package core

import (
	"Tubes-IdeaManagerCLI/model"
	"fmt"
	"time"
)

func AddIdea(arr *[model.MaxIdeas]model.Idea, count *int, title, category string) {
	if *count < model.MaxIdeas {
		used := make(map[int]bool)
		for i := 0; i < *count; i++ {
			used[arr[i].ID] = true
		}
		//Nemuin ID kecil yang belum terpakai
		newID := 1
		for used[newID] {
			newID++
		}

		arr[*count] = model.Idea{
			ID:        newID,
			Title:     title,
			Category:  category,
			Upvotes:   0,
			CreatedAt: time.Now(),
		}
		*count++
		SortIdeasByID(arr, *count)
	}
}

func EditIdea(arr *[model.MaxIdeas]model.Idea, count int, id int, newTitle, newCategory string) bool {
	for i := 0; i < count; i++ {
		if arr[i].ID == id {
			arr[i].Title = newTitle
			arr[i].Category = newCategory
			return true
		}
	}
	return false
}

func DeleteIdea(arr *[model.MaxIdeas]model.Idea, count *int, id int) bool {
	for i := 0; i < *count; i++ {
		if arr[i].ID == id {
			for j := i; j < *count-1; j++ {
				arr[j] = arr[j+1]
			}
			*count--
			return true
		}
	}
	return false
}

func UpvoteIdea(arr *[model.MaxIdeas]model.Idea, count int, id int) bool {
	for i := 0; i < count; i++ {
		if arr[i].ID == id {
			arr[i].Upvotes++
			return true
		}
	}
	return false
}

func DisplayIdeas(arr *[model.MaxIdeas]model.Idea, count int) {
	SortIdeasByID(arr, count)
	for i := 0; i < count; i++ {
		idea := arr[i]
		println("ID:", idea.ID, "| Judul:", idea.Title, "| Kategori:", idea.Category, "| Upvotes:", idea.Upvotes)
	}
}

func DisplayIdeaTitlesOnly(arr [model.MaxIdeas]model.Idea, count int) {
	if count == 0 {
		println("Tidak ada ide yang tersedia.")
	}
	println("=== Daftar Ide ===")
	for i := 0; i < count; i++ {
		idea := arr[i]
		println("ID:", idea.ID, "| Judul:", idea.Title, "| Upvotes:", idea.Upvotes)
	}
	println("===================")
}

func ShowMostPopularIdea(arr [model.MaxIdeas]model.Idea, count int) {
	if count == 0 {
		println("Belum ada ide yang tersedia.")
		return
	}

	popular := arr[0]
	for i := 1; i < count; i++ {
		if arr[i].Upvotes > popular.Upvotes {
			popular = arr[i]
		}
	}

	println("=== Ide Terpopuler ===")
	println("ID:", popular.ID)
	println("Judul:", popular.Title)
	println("Kategori:", popular.Category)
	println("Upvotes:", popular.Upvotes)
	println("Dibuat pada:", popular.CreatedAt.Format("02 Jan 2006 15:04:05"))
	println("======================")
}

func ShowMostRecentIdea(arr [model.MaxIdeas]model.Idea, count int) {
	if count == 0 {
		println("Belum ada ide yang tersedia.")
		return
	}

	recent := arr[0]
	for i := 1; i < count; i++ {
		if arr[i].CreatedAt.After(recent.CreatedAt) {
			recent = arr[i]
		}
	}

	println("=== Ide Terbaru ===")
	println("ID:", recent.ID)
	println("Judul:", recent.Title)
	println("Kategori:", recent.Category)
	println("Upvotes:", recent.Upvotes)
	println("Dibuat pada:", recent.CreatedAt.Format("02 Jan 2006 15:04:05"))
	println("===================")
}

func PromptReturnToMenu(ideas *[model.MaxIdeas]model.Idea, count int) bool {
	var choice int
	for {
		fmt.Println()
		fmt.Println("Apa yang ingin kamu lakukan selanjutnya?")
		fmt.Println("[1] Kembali ke Menu")
		fmt.Println("[2] Logout")
		fmt.Print("Pilihanmu: ")
		fmt.Scanln(&choice)

		if choice == 1 {
			SaveToFile(*ideas, count)
			return true 
		} else if choice == 2 {
			SaveToFile(*ideas, count)
			fmt.Println("Data disimpan. Terima kasih sudah menggunakan Idea Manager CLI!")
			return false 
		} else {
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}
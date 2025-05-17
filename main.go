package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"Tubes-IdeaManagerCLI/core"
	"Tubes-IdeaManagerCLI/model"
)

var ideas [model.MaxIdeas]model.Idea
var count int

func runIdeaManager() {
	core.LoadFromFile(&ideas, &count)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== APLIKASI IDEA MANAGER ===")
		fmt.Println("[1] Lihat Semua Ide")
		fmt.Println("[2] Tambah Ide Baru")
		fmt.Println("[3] Edit Ide")
		fmt.Println("[4] Hapus Ide")
		fmt.Println("[5] Upvote Ide")
		fmt.Println("[6] Tampilkan Ide Paling Populer")
		fmt.Println("[7] Tampilkan Ide Paling Baru")
		fmt.Println("[8] Cari Ide")
		fmt.Println("[9] Lihat Vote Terbanyak")
		fmt.Println("[0] Logout")
		fmt.Print("Pilih menu> ")

		scanner.Scan()
		menu := scanner.Text()

		switch menu {
		case "1":
			core.DisplayIdeas(&ideas, count)

		case "2":
			fmt.Println("Daftar kategori ide(tulis nama kategori saat input):")
			fmt.Println("[1] Teknologi")
			fmt.Println("[2] Kesehatan")
			fmt.Println("[3] Pendidikan")
			fmt.Println("[4] Lingkungan")
			fmt.Println("[5] Sosial")
			fmt.Print("Judul: ")
			scanner.Scan()
			title := scanner.Text()
			fmt.Print("Kategori: ")
			scanner.Scan()
			category := scanner.Text()
			core.AddIdea(&ideas, &count, title, category)
			fmt.Println("Ide berhasil ditambahkan.")

		case "3":
			core.DisplayIdeaTitlesOnly(ideas, count)
			fmt.Print("ID yang ingin diedit: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Judul baru: ")
			scanner.Scan()
			newTitle := scanner.Text()
			fmt.Print("Kategori baru: ")
			scanner.Scan()
			newCategory := scanner.Text()
			if core.EditIdea(&ideas, count, id, newTitle, newCategory) {
				fmt.Println("Ide berhasil diedit.")
			} else {
				fmt.Println("ID tidak ditemukan.")
			}

		case "4":
			core.DisplayIdeaTitlesOnly(ideas, count)
			fmt.Print("ID yang ingin dihapus: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			if core.DeleteIdea(&ideas, &count, id) {
				fmt.Println("Ide berhasil dihapus.")
			} else {
				fmt.Println("ID tidak ditemukan.")
			}

		case "5":
			core.DisplayIdeaTitlesOnly(ideas, count)
			fmt.Print("ID yang ingin di-upvote: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			if core.UpvoteIdea(&ideas, count, id) {
				fmt.Println("Upvote berhasil.")
			} else {
				fmt.Println("ID tidak ditemukan.")
			}

		case "6":
			core.ShowMostPopularIdea(ideas, count)

		case "7":
			core.ShowMostRecentIdea(ideas, count)

		case "8":
			fmt.Print("Judul yang dicari: ")
			scanner.Scan()
			keyword := scanner.Text()
			index := core.SequentialSearch(ideas, count, keyword)
			if index != -1 {
				fmt.Println("Ide ditemukan dengan ID dan Judul: ", ideas[index].ID, ideas[index].Title)
			} else {
				fmt.Println("Ide tidak ditemukan.")
			}

		case "9":
			core.SelectionSortByUpvotes(&ideas, count, false)
			for i := 0; i < count; i++ {
				fmt.Printf("ID: %d | Judul: %s | Upvotes: %d\n", ideas[i].ID, ideas[i].Title, ideas[i].Upvotes)
			}
			fmt.Println("Daftar ide berhasil diurutkan berdasarkan vote terbanyak")

		case "0":
			core.SaveToFile(ideas, count)
			fmt.Println("Data disimpan. Terima kasih sudah menggunakan Idea Manager CLI!")
			return

		default:
			fmt.Println("Menu tidak valid.")
		}
		if !core.PromptReturnToMenu(&ideas, count) {
		break
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n===== SELAMAT DATANG DI IDEA MANAGER CLI =====")
		fmt.Println("[1] Login")
		fmt.Println("[2] Register")
		fmt.Println("[0] Keluar")
		fmt.Print("Pilih menu> ")

		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			fmt.Print("Username: ")
			scanner.Scan()
			username := strings.TrimSpace(scanner.Text())

			fmt.Print("Password: ")
			scanner.Scan()
			password := strings.TrimSpace(scanner.Text())

			if core.LoginUser(username, password) {
				fmt.Println("Login berhasil, selamat datang", username+"!")
				runIdeaManager()
			} else {
				fmt.Println("Username atau password salah.")
			}

		case "2":
			fmt.Print("Buat Username: ")
			scanner.Scan()
			username := strings.TrimSpace(scanner.Text())

			fmt.Print("Buat Password: ")
			scanner.Scan()
			password := strings.TrimSpace(scanner.Text())

			err := core.RegisterUser(username, password)
			if err != nil {
				fmt.Println("Gagal register:", err)
			} else {
				fmt.Println("Akun berhasil dibuat! Silakan login.")
			}

		case "0":
			fmt.Println("Sampai jumpa!")
			return

		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}
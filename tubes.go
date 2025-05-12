package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Idea struct {
	Judul    string
	Kategori string
	Tanggal  time.Time
	Upvotes  int
}

type Kategori struct {
	Nama string
}

type User struct {
	Username string
	Password string
}

var ideas []Idea
var kategoris []Kategori
var users []User
var currentUser *User
var votedIde map[int]bool

func tampilkanMenuLogin() {
	fmt.Println("\n=== Menu Login dan Register ===")
	fmt.Println("[1] Login")
	fmt.Println("[2] Register")
	fmt.Println("[0] Keluar")
	fmt.Println("---------------------------------")
	fmt.Print("Pilih aksi> ")
}

func tampilkanMenu() {
	fmt.Println("\n=== APLIKASI PENGELOLAAN IDE STARTUP ===")
	fmt.Println("[1] Daftar ide proyek")
	fmt.Println("[2] Daftar kategori proyek")
	fmt.Println("[3] Dashboard Voting")
	fmt.Println("[4] Pencarian dan Pengurutan ide")
	fmt.Println("[0] Logout")
	fmt.Println("----------------------------------------")
}

func daftarIdeProyek() {
	for {
		fmt.Println("\n=== Daftar Ide Proyek ===")
		fmt.Println("[1] Lihat ide")
		fmt.Println("[2] Tambah ide")
		fmt.Println("[3] Edit ide")
		fmt.Println("[4] Hapus ide")
		fmt.Println("[0] Kembali ke menu utama")
		fmt.Println("---------------------------")
		fmt.Print("Pilih aksi> ")
		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			listIde()
		case 2:
			if kembali := tambahIde(); kembali {
				return 
			}
		case 3:
			editIde()
		case 4:
			hapusIde()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func daftarKategoriProyek() {
	for {
		fmt.Println("\n=== Daftar Kategori Proyek ===")
		fmt.Println("[1] Lihat kategori")
		fmt.Println("[2] Tambah kategori")
		fmt.Println("[3] Edit kategori")  
		fmt.Println("[4] Hapus kategori")
		fmt.Println("[0] Kembali ke menu utama")
		fmt.Println("--------------------------------")
		fmt.Print("Pilih aksi> ")
		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			listKategori()
		case 2:
			tambahKategori()
		case 3:
			editKategori()  
		case 4:
			hapusKategori()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func listIde() {
	if len(ideas) == 0 {
		fmt.Println("Belum ada ide.")
		return
	}
	fmt.Println("\nDaftar Ide:")
	for i, ide := range ideas {
		fmt.Printf("[%d] %s | Kategori: %s | Tanggal: %s | Upvotes: %d\n", i+1, ide.Judul, ide.Kategori, ide.Tanggal.Format("2006-01-02"), ide.Upvotes)
	}
}

func listKategori() {
	if len(kategoris) == 0 {
		fmt.Println("Belum ada kategori.")
		return
	}
	fmt.Println("\nDaftar Kategori:")
	for i, k := range kategoris {
		fmt.Printf("[%d] %s\n", i+1, k.Nama)
	}
}

func tambahIde() bool {
	if len(kategoris) == 0 {
		fmt.Println("Tidak ada kategori. Kembali ke menu utama")
		return true 
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Judul ide: ")
	judul, _ := reader.ReadString('\n')
	judul = strings.TrimSpace(judul)

	listKategori()
	fmt.Print("Pilih nomor kategori dari daftar di atas: ")
	var pilih int
	fmt.Scanln(&pilih)

	if pilih < 1 || pilih > len(kategoris) {
		fmt.Println("Nomor kategori tidak valid.")
		return false
	}
	kategori := kategoris[pilih-1].Nama

	ideas = append(ideas, Idea{
		Judul:    judul,
		Kategori: kategori,
		Tanggal:  time.Now(),
		Upvotes:  0,
	})
	fmt.Println("Ide berhasil ditambahkan.")
	return false
}

func tambahKategori() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama kategori: ")
	nama, _ := reader.ReadString('\n')
	nama = strings.TrimSpace(nama)
	kategoris = append(kategoris, Kategori{Nama: nama})
	fmt.Println("Kategori berhasil ditambahkan.")
}

func editKategori() {
	listKategori()
	fmt.Print("Pilih nomor kategori yang ingin diubah: ")
	var pilih int
	fmt.Scanln(&pilih)
	if pilih < 1 || pilih > len(kategoris) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama kategori baru: ")
	namaBaru, _ := reader.ReadString('\n')
	namaBaru = strings.TrimSpace(namaBaru)

	kategoris[pilih-1].Nama = namaBaru
	fmt.Println("Kategori berhasil diubah.")
}

func editIde() {
	listIde()
	fmt.Print("Pilih nomor ide yang ingin diubah: ")
	var pilih int
	fmt.Scanln(&pilih)
	if pilih < 1 || pilih > len(ideas) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Judul baru: ")
	judul, _ := reader.ReadString('\n')
	judul = strings.TrimSpace(judul)
	listKategori()
	fmt.Print("Kategori baru: ")
	kategori, _ := reader.ReadString('\n')
	kategori = strings.TrimSpace(kategori)

	ideas[pilih-1].Judul = judul
	ideas[pilih-1].Kategori = kategori
	fmt.Println("Ide berhasil diubah.")
}

func hapusIde() {
	listIde()
	fmt.Print("Pilih nomor ide yang ingin dihapus: ")
	var pilih int
	fmt.Scanln(&pilih)
	if pilih < 1 || pilih > len(ideas) {
		fmt.Println("Nomor tidak valid.")
		return
	}
	ideas = append(ideas[:pilih-1], ideas[pilih:]...)
	fmt.Println("Ide berhasil dihapus.")
}

func hapusKategori() {
	listKategori()
	fmt.Print("Pilih nomor kategori yang ingin dihapus: ")
	var pilih int
	fmt.Scanln(&pilih)
	if pilih < 1 || pilih > len(kategoris) {
		fmt.Println("Nomor tidak valid.")
		return
	}
	kategoris = append(kategoris[:pilih-1], kategoris[pilih:]...)
	fmt.Println("Kategori berhasil dihapus.")
}

func voteIde() {
	if len(ideas) == 0 {
		fmt.Println("Belum ada ide untuk divote.")
		return
	}

	if len(votedIde) > 0 {
		fmt.Println("Anda hanya bisa memberikan voting ke satu ide.")
		return
	}

	listIde()
	fmt.Print("Pilih nomor ide untuk upvote: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(ideas) {
		fmt.Println("Nomor tidak valid.")
		return
	}

	ideas[index-1].Upvotes++
	votedIde[index] = true
	fmt.Println("Upvote berhasil. Anda tidak dapat voting lagi.")
}




func lihatVotingTerbanyak() {
	if len(ideas) == 0 {
		fmt.Println("Belum ada ide untuk dilihat.")
		return
	}

	sort.Slice(ideas, func(i, j int) bool {
		return ideas[i].Upvotes > ideas[j].Upvotes
	})

	fmt.Println("\nIde dengan Voting Terbanyak:")
	for i, ide := range ideas {
		fmt.Printf("[%d] %s | Kategori: %s | Upvotes: %d\n", i+1, ide.Judul, ide.Kategori, ide.Upvotes)
	}
}

func cariDanUrutkanIde() {
	if len(ideas) == 0 {
		fmt.Println("Belum ada ide.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nPencarian Ide:")
	fmt.Print("Masukkan kata kunci (judul/kategori) atau tekan Enter untuk urutkan semua ide: ")
	kataKunci, _ := reader.ReadString('\n')
	kataKunci = strings.TrimSpace(strings.ToLower(kataKunci))

	var hasil []Idea
	if kataKunci == "" {
		hasil = append(hasil, ideas...)
		sort.Slice(hasil, func(i, j int) bool {
			return hasil[i].Upvotes > hasil[j].Upvotes
		})
	} else {
		for _, ide := range ideas {
			if strings.Contains(strings.ToLower(ide.Judul), kataKunci) || strings.Contains(strings.ToLower(ide.Kategori), kataKunci) {
				hasil = append(hasil, ide)
			}
		}
	}

	if len(hasil) == 0 {
		fmt.Println("Tidak ada ide yang cocok dengan pencarian.")
	} else {
		fmt.Println("\nHasil pencarian:")
		for i, ide := range hasil {
			fmt.Printf("[%d] %s | Kategori: %s | Upvotes: %d\n", i+1, ide.Judul, ide.Kategori, ide.Upvotes)
		}
	}
}

func dashboardVoting() {
	if currentUser == nil {
		fmt.Println("Silakan login terlebih dahulu.")
		return
	}

	for {
		fmt.Println("\n=== Dashboard Voting ===")
		fmt.Println("[1] Lihat voting terbanyak")
		fmt.Println("[2] Lakukan voting")
		fmt.Println("[0] Kembali ke menu utama")
		fmt.Println("--------------------------")
		fmt.Print("Pilih aksi> ")
		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			lihatVotingTerbanyak()
		case 2:
			voteIde()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func register() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Cek apakah username sudah ada
	for _, user := range users {
		if user.Username == username {
			fmt.Println("Username sudah terdaftar.")
			return
		}
	}

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Simpan user baru
	users = append(users, User{
		Username: username,
		Password: password,
	})

	fmt.Println("Registrasi berhasil.")
}

func login() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	for _, user := range users {
		if user.Username == username && user.Password == password {
			currentUser = &user
			votedIde = make(map[int]bool) 
			fmt.Println("Login berhasil.")
			return
		}
	}

	fmt.Println("Username atau password salah.")
}

func main() {
	for {
		tampilkanMenuLogin() 
		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			login()
			if currentUser != nil {
				menuUtama()
			}
		case 2:
			register()
		case 0:
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func menuUtama() {
	for {
		tampilkanMenu()
		fmt.Print("Pilih menu> ")
		var pilih int
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			daftarIdeProyek()
		case 2:
			daftarKategoriProyek()
		case 3:
			dashboardVoting() 
		case 4:
			cariDanUrutkanIde()
		case 0:
			fmt.Println("Keluar dari aplikasi...")
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
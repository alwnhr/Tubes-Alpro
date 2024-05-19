package main

import "fmt"

const NMAX int = 100

type users struct {
	userId   int
	fullName string
	username string
	password string
	email    string
	phone    string
}

type date struct {
	tgl, bln, thn int
}

type events struct {
	eventId      int
	namaEvent    string
	deskripsi    string
	tanggal      date
	creatorId    int
	peserta      [NMAX]int
	pesertaCount int
}

type statusAcara struct {
	usersList   [NMAX]users
	eventslist  [NMAX]events
	totalUsers  int
	totalEvents int
	currentUser *users
}

var status statusAcara

func main() {
	var menu int
	selamatDatang()
	for {
		menu_registrasi(&menu)
		if menu == 3 {
			return // keluar dari fungsi main dan program selesai
		}
	}
	terimaKasih()
}

func selamatDatang() {
	clearScreen()
	fmt.Println("Selamat datang di aplikasi Manajemen Acara!")
}

func terimaKasih() {
	clearScreen()
	fmt.Println("Terima kasih telah menggunakan aplikasi.")
}

func clearScreen() {
	for i := 0; i < 1; i++ { // cetak 1 baris kosong
		fmt.Println()
	}
}

func menu_registrasi(m *int) {
	clearScreen()
	fmt.Println("=========================")
	fmt.Println("     MENU REGISTRASI     ")
	fmt.Println("-------------------------")
	fmt.Println("1. Signing Up")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Println("-------------------------")
	fmt.Print("Menu yang dipilih (1/2/3): ")
	fmt.Println("=========================")
	fmt.Scan(m)
	menuRegis(*m)
}

func menuRegis(m int) {
	switch m {
	case 1:
		signingUp()
	case 2:
		if userLogin() {
			userDashboard()
		}
	case 3:
		return
	default:
		fmt.Println("Pilihan tidak valid! Silahkan coba lagi.")
	}
}

// Fungsi signing up
func signingUp() {
	var fullName, username, password, email, phone string

	clearScreen()
	fmt.Println("=========================")
	fmt.Println("        SIGNING UP       ")
	fmt.Println("=========================")
	fmt.Print("Full Name: ")
	fmt.Scan(&fullName)
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)
	fmt.Print("E-Mail: ")
	fmt.Scan(&email)
	fmt.Print("Phone Number: ")
	fmt.Scan(&phone)

	if userSigning(&status, fullName, username, password, email, phone) {
		clearScreen()
		fmt.Println("-----------------")
		fmt.Println("Signing berhasil!")
		fmt.Println("-----------------")
	} else {
		clearScreen()
		fmt.Println("-------------------------")
		fmt.Println("Username sudah digunakan.")
		fmt.Println("-------------------------")
	}
}

// Fungsi user signing untuk menyimpan data pengguna baru dari fungsi sebelumnya
func userSigning(status *statusAcara, fullName, username, password, email, phone string) bool {
	if status.totalUsers >= NMAX {
		fmt.Println("------------------")
		fmt.Println("User limit reached!")
		fmt.Println("------------------")
		return false
	}

	// Check jika username sudah ada
	for i := 0; i < status.totalUsers; i++ {
		if status.usersList[i].username == username {
			return false
		}
	}

	// Berhasil menambahkan pengguna baru
	status.usersList[status.totalUsers].userId = status.totalUsers + 1
	status.usersList[status.totalUsers].fullName = fullName
	status.usersList[status.totalUsers].username = username
	status.usersList[status.totalUsers].password = password
	status.usersList[status.totalUsers].email = email
	status.usersList[status.totalUsers].phone = phone

	status.totalUsers++
	return true
}

// Fungsi login
func userLogin() bool {
	var username, password string

	clearScreen()
	fmt.Println("=========================")
	fmt.Println("          LOGIN          ")
	fmt.Println("=========================")
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	// Login berhasil
	for i := 0; i < status.totalUsers; i++ {
		if status.usersList[i].username == username && status.usersList[i].password == password {
			status.currentUser = &status.usersList[i]
			fmt.Println("-------------------------")
			fmt.Println("     Login berhasil!     ")
			fmt.Printf("Selamat datang, %s.\n", status.currentUser.username)
			fmt.Println("-------------------------")
			return true
		}
	}

	// Login gagal
	fmt.Println("-----------------------------")
	fmt.Println("         Login gagal!        ")
	fmt.Println("Username atau password salah.")
	fmt.Println("-----------------------------")
	return false
}

// Menu dashboard
func userDashboard() {
	clearScreen()
	fmt.Println("-------------------------")
	fmt.Println("        Dashboard        ")
	fmt.Println("-------------------------")
	fmt.Printf("Selamat datang, %s!\n", status.currentUser.fullName)
	var tgl, bln, thn int
	fmt.Println("Masukkan tanggal untuk melihat acara:")
	fmt.Scan(&tgl, &bln, &thn)
	today := date{tgl: tgl, bln: bln, thn: thn}
	fmt.Println("Daftar acara yang sedang berlangsung:")
	displayEvents("current")
	fmt.Println("Daftar acara yang akan datang:")
	displayEvents("upcoming")
}

func displayEvents(today date, eventType string) {
	for i := 0; i < status.totalEvents; i++ {
		event := status.eventslist[i]

		switch eventType {
		case "current":
			if event.tanggal == today {
				printEvent(event)
			}
		case "upcoming":
			if (event.tanggal.thn > today.thn) || (event.tanggal.thn == today.thn && event.tanggal.bln == today.bln && event.tanggal.tgl > today.tgl) {
				printEvent(event)
			}
		}
	}
}

func printEvent(event events) {
	fmt.Printf("- %s (%s)\n", event.namaEvent, formatDate(event.tanggal))
	fmt.Printf("	Deskripsi: %s\n", event.deskripsi)
	fmt.Printf("	Tanggal: %s\n", formatDate(event.tanggal))
	fmt.Printf("	Pembuat Acara: %s\n", event.pembuatAcara)
	fmt.Println("-------------------------")
}

func formatDate(d date) string {
	return fmt.Printf("%02d-%02d-%04d", d.tgl, d.bln, d.thn)
}

package main

import "fmt"

const NMAX int = 100

type users struct {
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
	namaEvent    string
	deskripsi    string
	tanggal      date
	pembuatAcara string
	peserta      [50]int
	pesertaCount int
}

type statusAcara struct {
	usersList   [100]users
	eventslist  [100]events
	totalUsers  int
	totalEvents int
	currentUser *users
}

var menu int

func main() {
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
	fmt.Println("-------------------------")
	fmt.Println("     MENU REGISTRASI     ")
	fmt.Println("-------------------------")
	fmt.Println("1. Signing Up")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Println("-------------------------")
	fmt.Print("Menu yang dipilih (1/2/3): ")
	fmt.Scan(m)
	menuRegis(*m)
}

var fullName, username, password, email, phone string

func menuRegis(m int) {
	switch m {
	case 1:
		userSigning(fullName, username, password, email, phone)
	case 2:
		userLogin(username, password)
	case 3:
		return
	default:
		fmt.Println("Pilihan tidak valid")
	}
}

var status statusAcara
var user users
var currentUser *users

// Menu signing up
func userSigning(fullName, username, password, email, phone string) bool {
	if status.totalUsers >= NMAX {
		fmt.Println("==================")
		fmt.Println("User limit reached")
		fmt.Println("==================")
		return false
	}
	clearScreen()
	for {
		fmt.Println("-------------------------")
		fmt.Println("        Signing Up       ")
		fmt.Println("-------------------------")
		fmt.Print("Full Name: ")
		fmt.Scan(&user.fullName)
		fmt.Print("Username: ")
		fmt.Scan(&user.username)
		fmt.Print("Password: ")
		fmt.Scan(&user.password)
		fmt.Print("E-Mail: ")
		fmt.Scan(&user.email)
		fmt.Print("Phone Number: ")
		fmt.Scan(&user.phone)

		// Check jika username sudah ada
		for i := 0; i < status.totalUsers; i++ {
			if status.usersList[i].username == username {
				fmt.Println("=========================")
				fmt.Println("Username sudah digunakan!")
				fmt.Println("=========================")
				return false
			}
		}

		// Berhasil menambahkan pengguna baru
		status.usersList[status.totalUsers].fullName = fullName
		status.usersList[status.totalUsers].username = username
		status.usersList[status.totalUsers].password = password
		status.usersList[status.totalUsers].email = email
		status.usersList[status.totalUsers].phone = phone

		status.totalUsers++
		fmt.Println("=================================================")
		fmt.Println("User registered successfully!")
		fmt.Println("=================================================")
		return true
	}
}

// Menu login
func userLogin(username, password string) bool {
	clearScreen()
	for {
		fmt.Println("-------------------------")
		fmt.Println("          Login          ")
		fmt.Println("-------------------------")
		fmt.Print("Username: ")
		fmt.Scan(&user.username)
		fmt.Print("Password: ")
		fmt.Scan(&user.password)

		// Login berhasil
		for i := 0; i < status.totalUsers; i++ {
			if status.usersList[i].username == username && status.usersList[i].password == password {
				status.currentUser = &status.usersList[i]
				fmt.Println("============================")
				fmt.Println("Login berhasil!")
				fmt.Printf("Selamat datang, %s.\n", status.currentUser.username)
				fmt.Println("============================")
				return true
			}
		}

		// Login gagal
		fmt.Println("=============================")
		fmt.Println("Login gagal!")
		fmt.Println("Username atau password salah.")
		fmt.Println("=============================")
		return false
	}
}

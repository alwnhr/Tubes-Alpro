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

type events struct {
	namaEvent    string
	deskripsi    string
	tanggal      date
	pembuatAcara string
	peserta      [50]int
}

type date struct {
	tgl, bln, thn int
}

type statusAcara struct {
	usersList   [100]users
	eventList   [100]events
	totalUsers  int
	totalEvents int
}

var status statusAcara
var currentUser *users

// ini isinya subprogram belum diisi tapi

// Menu signing up
func userSigning(fullName, username, password, email, phone string) bool {
	if status.totalUsers >= NMAX {
		fmt.Println("User limit reached")
		return false
	}
	// Check jika username sudah ada
	for i := 0; i < status.totalUsers; i++ {
		if status.usersList[i].username == username {
			fmt.Println("Username sudah digunakan!")
			return false
		}
	}

}

// Menu login
func userLogin(username, password string) bool {
	for i := 0; i < status.totalUsers; i++ {
		if status.usersList[i].username == username && status.usersList[i].password == password {
			currentUser = &status.usersList[i]
			return true
		}
	}
	return false
}

func main() {
	ClearScreen()
	menu_registrasi()

	var status statusAcara
	var menu int // MENU disini buat nentuin nanti mau signing up, login, exit
	var fullName, username, password, email, phone string
}

// aku pake switch case ya

func menu_registrasi() {
	for {
		fmt.Println()
		fmt.Println("-------------------------")
		fmt.Println("     MENU REGISTRASI     ")
		fmt.Println("-------------------------")
		fmt.Println("1. Signing Up")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println("-------------------------")
		fmt.Print("Menu yang dipilih: ")
		fmt.Scan(&menu)
		ClearScreen()

		switch menu {
		case 1:
			fmt.Print("Nama Lengkap: ")
			fmt.Scan(&fullName)
			fmt.Print("Username: ")
			fmt.Scan(&username)
			fmt.Print("Password: ")
			fmt.Scan(&password)
			fmt.Print("E-Mail: ")
			fmt.Scan(&email)
			fmt.Print("Phone Number: ")
			fmt.Scan(&phone)

			if userSigning() { //belum dibuat subprogramnya, tapi ini pengandaian dulu
				fmt.Println("Signing Up Berhasil!")
			} else {
				fmt.Println("Signing Up gagal! Pengguna mungkin sudah ada.")
			}
			ClearScreen()

		case 2:
			fmt.Print("Username: ")
			fmt.Scan(&username)
			fmt.Print("Password: ")
			fmt.Scan(&password)

			if userLogin() { //belum dibuat
				fmt.Println("Login Berhasil!")
				userDashboard() //belum dibuat, tapi ini nanti ngescan status
			} else {
				fmt.Println("Login gagal! Pengguna dengan data di atas tidak dapat ditemukan.")
			}

		case 3:
			//wait belum
		}
	}
}

func userDashboard(status *statusAcara) {
	var menu int
	//on progress
}

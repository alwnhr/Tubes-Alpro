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

type listUser [1000]users

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

var menu int

func main() {
	var listUser listUser
	var nUser int = 0
	selamatDatang()
	for {
		menu_registrasi()
		var menuRegis int
		fmt.Scan(&menuRegis)
		if menuRegis == 1 {
			userSigning(&listUser, &nUser)
			for i := 0; i < nUser; i++ {
				fmt.Print(listUser[i].fullName)
			}
		}
		if menuRegis == 2 {
			userLogin(&listUser, &nUser)
		}
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

func menu_registrasi() {
	clearScreen()
	fmt.Println("=========================")
	fmt.Println("     MENU REGISTRASI     ")
	fmt.Println("-------------------------")
	fmt.Println("1. Signing Up")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Println("-------------------------")
	fmt.Print("Menu yang dipilih (1/2/3): ")
}

// func menuRegis(m int, lu *listUser) {
// 	switch m {
// 	case 1:
// 		userSigning(fullName, usename, password, email, phone)
// 	case 2:
// 		userLogin(username, password)
// 	case 3:
// 		return
// 	default:
// 		fmt.Println("Pilihan tidak valid")
// 	}
// }

// var status statusAcara
// var user users
// var currentUser *users

// Menu signing up
func userSigning(status *listUser, n *int) bool {
	var fullName, username, password, email, phone string
	// if status.totalUsers >= NMAX {
	// 	fmt.Println("==================")
	// 	fmt.Println("User limit reached")
	// 	fmt.Println("==================")
	// 	return false
	// }
	// clearScreen()
	for {
		fmt.Println("-------------------------")
		fmt.Println("Username sudah digunakan.")
		fmt.Println("-------------------------")
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

		// Check jika username sudah ada
		for i := 0; i < *n; i++ {
			if status[i].username == username {
				fmt.Println("=========================")
				fmt.Println("Username sudah digunakan!")
				fmt.Println("=========================")
				return false
			}
		}

		// Berhasil menambahkan pengguna baru
		status[*n].fullName = fullName
		status[*n].username = username
		status[*n].password = password
		status[*n].email = email
		status[*n].phone = phone

		*n++
		fmt.Println("=================================================")
		fmt.Println("User registered successfully!")
		fmt.Println("=================================================")
		return true
	}
}

// Menu login
func userLogin(status *listUser, n *int) bool {
	var username, password string
	clearScreen()
	for {
		fmt.Println("-------------------------")
		fmt.Println("          Login          ")
		fmt.Println("-------------------------")
		fmt.Print("Username: ")
		fmt.Scan(&username)
		fmt.Print("Password: ")
		fmt.Scan(&password)

		// Login berhasil
		for i := 0; i < *n; i++ {
			if status[i].username == username && status[i].password == password {
				// status.currentUser = &status.usersList[i]
				fmt.Println("============================")
				fmt.Println("Login berhasil!")
				// fmt.Printf("Selamat datang, %s.\n", status.currentUser.username)
				fmt.Println("============================")
				return true
			}
		}

		// Login gagal
		// fmt.Println("=============================")
		// fmt.Println("Login gagal!")
		// fmt.Println("Username atau password salah.")
		// fmt.Println("=============================")
		// return false
	}
}

// Menu dashboard
// func userDashboard() {
// 	clearScreen()
// 	fmt.Println("-------------------------")
// 	fmt.Println("        Dashboard        ")
// 	fmt.Println("-------------------------")
// 	fmt.Printf("Selamat datang, %s!\n", status.currentUser.fullName)
// 	var tgl, bln, thn int
// 	fmt.Println("Masukkan tanggal untuk melihat acara:")
// 	fmt.Scan(&tgl, &bln, &thn)
// 	today := date{tgl: tgl, bln: bln, thn: thn}
// 	fmt.Println("Daftar acara yang sedang berlangsung:")
// 	displayEvents("current")
// 	fmt.Println("Daftar acara yang akan datang:")
// 	displayEvents("upcoming")
// }

// func displayEvents(today date, eventType string) {
// 	for i := 0; i < status.totalEvents; i++ {
// 		event := status.eventslist[i]

// 		switch eventType {
// 		case "current":
// 			if event.tanggal == today {
// 				printEvent(event)
// 			}
// 		case "upcoming":
// 			if (event.tanggal.thn > today.thn) || (event.tanggal.thn == today.thn && event.tanggal.bln == today.bln && event.tanggal.tgl > today.tgl) {
// 				printEvent(event)
// 			}
// 		}
// 	}
// }

// func printEvent(event events) {
// 	fmt.Printf("- %s (%s)\n", event.namaEvent, formatDate(event.tanggal))
// 	fmt.Printf("	Deskripsi: %s\n", event.deskripsi)
// 	fmt.Printf("	Tanggal: %s\n", formatDate(event.tanggal))
// 	fmt.Printf("	Pembuat Acara: %s\n", event.pembuatAcara)
// 	fmt.Println("-------------------------")
// }

// func formatDate(d date) string {
// 	return fmt.Printf("%02d-%02d-%04d", d.tgl, d.bln, d.thn)
// }

// Membuat acara baru
// func createEvent() {
// 	if status.totalEvents >= NMAX {
// 		fmt.Println("Event limit reached")
// 		return
// 	}

// 	var newEvent events
// 	fmt.Println("Masukkan acara baru:")
// 	fmt.Print("Nama Event: ")
// 	fmt.Scan(&newEvent.namaEvent)
// 	fmt.Print("Deskripsi: ")
// 	fmt.Scan(&newEvent.deskripsi)
// 	fmt.Print("Tanggal: ")
// 	fmt.Scan(&newEvent.tanggal.tgl, &newEvent.tanggal.bln, &newEvent.tanggal.thn)
// 	newEvent.pembuatAcara = status.currentUser.username

// 	status.eventslist[status.totalEvents] = newEvent
// 	status.totalEvents++
// 	fmt.Println("Acara berhasil dibuat!")
// }

// func editEvent() {
// 	fmt.Print("Masukkan nama event yang ingin diedit: ")
// 	var eventName string
// 	fmt.Scan(&eventName)

// 	for i := 0; i < status.totalEvents; i++ {
// 		if status.eventslist[i].namaEvent == eventName && status.eventslist[i].pembuatAcara == status.currentUser.username {
// 			fmt.Println("Silahkan edit detail acara: ")
// 			var input string

// 			fmt.Print("Nama event: ")
// 			fmt.Scan(&input)
// 			status.eventslist[i].namaEvent = input

// 			fmt.Print("Deskripsi: ")
// 			fmt.Scan(&input)
// 		}
// 	}
// }

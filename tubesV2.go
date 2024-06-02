package main

import (
	"fmt"
)

type User struct {
	Username string
	Password string
	FullName string
	Email    string
	Phone    string
}

type Event struct {
	Name        string
	Description string
	Date        string
}

const maxUsers = 100
const maxEvents = 1000

var users [maxUsers]User
var events [maxEvents]Event
var userCount int
var eventCount int

var ongoingEvents [maxEvents]Event
var ongoingEventCount int
var upcomingEvents [maxEvents]Event
var upcomingEventCount int

func main() {
	// Dummy data for users
	users[0] = User{Username: "alwn", Password: "123", FullName: "Alwan Hutama", Email: "alwan@gmail.com", Phone: "123456789"}
	users[1] = User{Username: "gis", Password: "456", FullName: "Gisel", Email: "gisel@gmail.com", Phone: "987654321"}
	userCount = 2

	// Dummy data for events
	dummyEvents := [5]Event{
		{Name: "Seminar Teknologi", Description: "Diskusi tentang tren terbaru dan masa depan teknologi", Date: "2024-06-01"},
		{Name: "Pameran Seni Rupa", Description: "Pameran karya seni terbaru dari seniman lokal", Date: "2024-06-05"},
		{Name: "Konferensi Kesehatan Global", Description: "Forum untuk membahas isu kesehatan global", Date: "2024-06-10"},
		{Name: "Bazar Kreatif dan UMKM", Description: "Pameran produk kreatif dan UMKM lokal", Date: "2024-06-08"},
		{Name: "Pelatihan Keterampilan Digital", Description: "Workshop untuk mempelajari keterampilan digital", Date: "2024-06-23"},
	}

	for i := 0; i < len(dummyEvents); i++ {
		events[eventCount] = dummyEvents[i]
		if dummyEvents[i].Date == "2024-06-01" {
			ongoingEvents[ongoingEventCount] = dummyEvents[i]
			ongoingEventCount++
		} else {
			upcomingEvents[upcomingEventCount] = dummyEvents[i]
			upcomingEventCount++
		}
		eventCount++
	}

	for {
		fmt.Println("Selamat datang di aplikasi Manajemen Acara!")
		fmt.Println("=========================")
		fmt.Println("     MENU REGISTRASI     ")
		fmt.Println("-------------------------")
		fmt.Println("1. Signing Up")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")
		fmt.Println("-------------------------")
		fmt.Print("Menu yang dipilih (1/2/3): ")

		var pilih int
		fmt.Scan(&pilih)

		if pilih == 1 {
			register()
		} else if pilih == 2 {
			login()
		} else if pilih == 3 {
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		} else {
			fmt.Println("Opsi tidak valid, harap coba lagi.")
		}
	}
}

func register() {
	if userCount >= maxUsers {
		fmt.Println("==================")
		fmt.Println("User limit reached")
		fmt.Println("==================")
		return
	}

	var username, password, fullName, email, phone string

	fmt.Println("-------------------------")
	fmt.Println("         Register        ")
	fmt.Println("-------------------------")

	fmt.Print("Username: ")
	fmt.Scan(&username)

	for i := 0; i < userCount; i++ {
		if users[i].Username == username {
			fmt.Println("Username sudah digunakan. Silakan coba yang lain.")
			return
		}
	}

	fmt.Print("Password: ")
	fmt.Scan(&password)

	fmt.Print("Full Name: ")
	fmt.Scan(&fullName)

	fmt.Print("Email: ")
	fmt.Scan(&email)

	validPhone := false
	for !validPhone {
		fmt.Print("Phone Number: ")
		fmt.Scan(&phone)
		if isNumeric(phone) {
			validPhone = true
		} else {
			fmt.Println("Nomor telepon tidak valid. Silakan masukkan nomor telepon valid yang hanya terdiri dari angka.")
		}
	}

	users[userCount] = User{Username: username, Password: password, FullName: fullName, Email: email, Phone: phone}
	userCount++
	fmt.Println("=================================================")
	fmt.Println("					User registered successfully!						")
	fmt.Println("=================================================")
}

func isNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func login() {
	var username, password string
	fmt.Println("-------------------------")
	fmt.Println("          Login          ")
	fmt.Println("-------------------------")

	fmt.Print("Username: ")
	fmt.Scan(&username)

	fmt.Print("Password: ")
	fmt.Scan(&password)

	for i := 0; i < userCount; i++ {
		if users[i].Username == username && users[i].Password == password {
			fmt.Println("Login successful!")
			dashboardMenu(users[i]) // Mengarahkan ke menu dashboard setelah login
			return
		}
	}
	fmt.Println("Login failed. Username or password incorrect.")
}

func dashboardMenu(user User) {
	for {
		fmt.Println("\nSelamat Datang,", user.FullName)
		fmt.Println("Dashboard Menu:")
		fmt.Println("Acara yang sedang berlangsung:")
		showOngoingEvents()
		fmt.Println()
		fmt.Println("Acara yang akan datang:")
		showUpcomingEvents()
		fmt.Println("\nMenu:")
		fmt.Println("1. Acara")
		fmt.Println("2. Peserta")
		fmt.Println("3. Profil")
		fmt.Println("4. Logout")
		fmt.Print("Menu yang dipilih (1/2/3/4): ")

		var pilih int
		fmt.Scan(&pilih)

		if pilih == 1 {
			menuEvents(user)
		} else if pilih == 2 {
			// Tambahkan fungsi peserta di sini
		} else if pilih == 3 {
			// Tambahkan fungsi profil di sini
		} else if pilih == 4 {
			fmt.Println("Logging out...")
			return
		} else {
			fmt.Println("Opsi tidak valid, harap coba lagi.")
		}
	}
}

func showOngoingEvents() {
	if ongoingEventCount == 0 {
		fmt.Println("Tidak ada acara yang sedang berlangsung.")
		return
	}

	for i := 0; i < ongoingEventCount; i++ {
		fmt.Printf("%d. %s - %s\nDeskripsi: %s\n", i+1, ongoingEvents[i].Name, ongoingEvents[i].Date, ongoingEvents[i].Description)
	}
}

func showUpcomingEvents() {
	if upcomingEventCount == 0 {
		fmt.Println("Tidak ada acara mendatang.")
		return
	}

	for i := 0; i < upcomingEventCount; i++ {
		fmt.Printf("%d. %s - %s\nDeskripsi: %s\n", i+1, upcomingEvents[i].Name, upcomingEvents[i].Date, upcomingEvents[i].Description)
	}
}

func menuEvents(user User) {
	for {
		fmt.Println("\nMenu Acara:")
		fmt.Println("1. Buat Acara Baru")
		fmt.Println("2. Edit Detail Acara")
		fmt.Println("3. Cari Acara")
		fmt.Println("4. Urutkan Acara")
		fmt.Println("5. Kembali ke Dashboard")
		fmt.Print("Menu yang dipilih (1/2/3/4/5): ")

		var pilih int
		fmt.Scan(&pilih)

		if pilih == 1 {
			createEvent(user)
		} else if pilih == 2 {
			editEvent(user)
		} else if pilih == 3 {
			searchEvent()
		} else if pilih == 4 {
			sortEvent()
		} else if pilih == 5 {
			fmt.Println("Kembali ke dashboard...")
			return
		} else {
			fmt.Println("Opsi tidak valid, harap coba lagi.")
		}
	}
}

func createEvent(user User) {
	if eventCount >= maxEvents {
		fmt.Println("===================")
		fmt.Println("Event limit reached")
		fmt.Println("===================")
		return
	}

	var name, description, date string

	fmt.Print("Masukkan nama acara: ")
	fmt.Scan(&name)

	fmt.Print("Masukkan deskripsi acara: ")
	fmt.Scan(&description)

	fmt.Print("Masukkan tanggal acara (YYYY-MM-DD): ")
	fmt.Scan(&date)

	newEvent := Event{Name: name, Description: description, Date: date}

	if date == "2024-06-01" {
		ongoingEvents[ongoingEventCount] = newEvent
		ongoingEventCount++
	} else {
		upcomingEvents[upcomingEventCount] = newEvent
		upcomingEventCount++
	}

	events[eventCount] = newEvent
	eventCount++

	fmt.Println("Acara berhasil dibuat!")
}

func editEvent(user User) {
	fmt.Println("Daftar Acara Anda:")
	for i := 0; i < eventCount; i++ {
		fmt.Printf("%d. %s - %s\n", i+1, events[i].Name, events[i].Date)
	}

	fmt.Print("Masukkan nomor acara yang ingin Anda edit: ")
	var eventNum int
	fmt.Scan(&eventNum)

	if eventNum < 1 || eventNum > eventCount {
		fmt.Println("Nomor acara tidak valid.")
		return
	}

	var name, description, date string

	fmt.Print("Masukkan nama acara baru: ")
	fmt.Scan(&name)

	fmt.Print("Masukkan deskripsi acara baru: ")
	fmt.Scan(&description)

	fmt.Print("Masukkan tanggal acara baru (YYYY-MM-DD): ")
	fmt.Scan(&date)

	events[eventNum-1] = Event{Name: name, Description: description, Date: date}

	// Update ongoing or upcoming events lists
	if date == "2024-06-01" {
		ongoingEvents[ongoingEventCount] = events[eventNum-1]
		ongoingEventCount++
		// Remove from upcomingEvents if present
		for i := 0; i < upcomingEventCount; i++ {
			if upcomingEvents[i].Name == name {
				// Shift elements left to fill the gap
				for j := i; j < upcomingEventCount-1; j++ {
					upcomingEvents[j] = upcomingEvents[j+1]
				}
				upcomingEventCount--
				i--
			}
		}
	} else {
		upcomingEvents[upcomingEventCount] = events[eventNum-1]
		upcomingEventCount++
		// Remove from ongoingEvents if present
		for i := 0; i < ongoingEventCount; i++ {
			if ongoingEvents[i].Name == name {
				// Shift elements left to fill the gap
				for j := i; j < ongoingEventCount-1; j++ {
					ongoingEvents[j] = ongoingEvents[j+1]
				}
				ongoingEventCount--
				i--
			}
		}
	}

	fmt.Println("Detail acara berhasil diperbarui!")
}

func searchEvent() {
	fmt.Println("\nCari acara")
	fmt.Println("1. Cari berdasarkan nama")
	fmt.Println("2. Cari berdasarkan Tanggal")
	fmt.Print("Opsi yang dipilih (1/2): ")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		var name string
		fmt.Println("Masukkan nama acara: ")
		fmt.Scan(&name)
		searchEventByName(name)
	} else if choice == 2 {
		var date string
		fmt.Print("Masukkan tanggal acara (YYYY-MM-DD): ")
		fmt.Scan(&date)
		searchEventByDate(date)
	} else {
		fmt.Println("Opsi tidak valid, harap coba lagi.")
	}
}

func searchEventByName(name string) {
	fmt.Println("\nHasil Pencarian untuk Nama Acara:", name)
	found := false
	for i := 0; i < eventCount; i++ {
		if events[i].Name == name {
			fmt.Printf("Acara %d:\nNama Acara: %s\nDeskripsi: %s\nTanggal Acara: %s\n", i+1, events[i].Name, events[i].Description, events[i].Date)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan acara dengan nama yang diberikan.")
	}
}

func searchEventByDate(date string) {
	fmt.Println("\nHasil Pencarian untuk Tanggal Acara:", date)
	found := false
	for i := 0; i < eventCount; i++ {
		if events[i].Date == date {
			fmt.Printf("Acara %d:\nNama Acara: %s\nDeskripsi: %s\nTanggal Acara: %s\n", i+1, events[i].Name, events[i].Description, events[i].Date)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada acara yang ditemukan pada tanggal yang diberikan.")
	}
}

func sortEvent() {
	fmt.Println("\nUrutkan Acara")
	fmt.Println("1. Urutkan berdasarkan Tanggal (Ascending)")
	fmt.Println("2. Urutkan berdasarkan Tanggal (Descending)")
	fmt.Print("Opsi yang dipilih (1/2):")

	var pilih int
	fmt.Scan(&pilih)

	if pilih == 1 {
		sortEventsByDate(true)
	} else if pilih == 2 {
		sortEventsByDate(false)
	} else {
		fmt.Println("Opsi tidak valid, harap coba lagi.")
	}
}

func sortEventsByDate(ascending bool) {
	for i := 0; i < eventCount-1; i++ {
		for j := 0; j < eventCount-i-1; j++ {
			dateI := events[j].Date
			dateJ := events[j+1].Date

			if (ascending && dateI > dateJ) || (!ascending && dateI < dateJ) {
				events[j], events[j+1] = events[j+1], events[j]
			}
		}
	}

	fmt.Println("\nAcara yang Diurutkan:")
	for i := 0; i < eventCount; i++ {
		fmt.Printf("Acara %d:\nNama Acara: %s\nDeskripsi: %s\nTanggal Acara: %s\n", i+1, events[i].Name, events[i].Description, events[i].Date)
	}
}

package main

import "fmt"

const NMAX int = 100

type users struct {
	userId      int
	fullName    string
	username    string
	password    string
	email       string
	phone       string
	totalEvents int
	active      bool // Untuk menentukan pengguna sedang login atau tidak
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
	eventsList  [NMAX]events
	totalUsers  int
	totalEvents int
	currentUser *users
}

var status statusAcara

func main() {
	var menu int
	var signedUp bool
	selamatDatang()

	for {
		if signedUp || menu == 2 {
			userDashboard()
			if !status.currentUser.active {
				signedUp = false
			}
		} else {
			menu_registrasi()
			if menu == 3 {
				return // keluar dari fungsi main dan program selesai
			}
			if menu == 1 {
				signedUp = signingUp()
			}
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

func menu_registrasi() int {
	var menu int
	clearScreen()
	fmt.Println("=========================")
	fmt.Println("     MENU REGISTRASI     ")
	fmt.Println("-------------------------")
	fmt.Println("1. Signing Up")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	fmt.Println("-------------------------")
	fmt.Print("Menu yang dipilih (1/2/3): ")
	fmt.Scan(&menu)
	fmt.Println("=========================")
	menuRegis(menu)
	return menu
}

func menuRegis(m int) {
	switch m {
	case 1:
		if signingUp() {
			userDashboard()
		}
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
func signingUp() bool {
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

	if userSigningUp(&status, fullName, username, password, email, phone) {
		clearScreen()
		fmt.Println("-----------------")
		fmt.Println("Signing berhasil!")
		fmt.Println("-----------------")

		// Bagi pengguna yang baru mendaftar dan berhasil, jadikan current user
		for i := 0; i < status.totalUsers; i++ {
			if status.usersList[i].username == username && status.usersList[i].password == password {
				status.currentUser = &status.usersList[i]
			}
		}
		return true

	} else {
		clearScreen()
		fmt.Println("-------------------------")
		fmt.Println("Username sudah digunakan.")
		fmt.Println("-------------------------")
		return false
	}
}

// Fungsi user signing up untuk menyimpan data pengguna baru dari fungsi sebelumnya
func userSigningUp(status *statusAcara, fullName, username, password, email, phone string) bool {
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
			fmt.Println("Login berhasil!")
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

// Dashboard yang ditampilkan
func userDashboard() {
	var menu int
	status.currentUser.active = true
	for status.currentUser.active {
		clearScreen()
		fmt.Println("=====================================")
		fmt.Println("              DASHBOARD              ")
		fmt.Println("-------------------------------------")
		fmt.Printf("Selamat datang, %s!\n", status.currentUser.username)
		fmt.Println("-------------------------------------")
		fmt.Println("1. Lihat Daftar Acara")
		fmt.Println("2. Buat Acara")
		fmt.Println("3. Ubah Acara")
		fmt.Println("4. Tambah Acara")
		fmt.Println("5. Hapus Peserta")
		fmt.Println("6. Ubah Profil")
		fmt.Println("7. Cari User")
		fmt.Println("8. Logout")
		fmt.Println("-------------------------------------")
		fmt.Println("Menu yang dipilih (1/2/3/4/5/6/7/8): ")
		fmt.Scan(&menu)
		fmt.Println("=====================================")

		switch menu {
		case 1:
			var today date
			var eventType string
			displayEvents(today, eventType)
		case 2:
			createEvent()
		case 3:
			editEvent()
		case 4:
			tambahPeserta()
		case 5:
			hapusPeserta()
		case 6:
			editProfile()
		case 7:
			cariUser()
		case 8:
			status.currentUser.active = false
		default:
			fmt.Println("Pilihan tidak valid! Silahkan coba lagi.")
		}
	}
}

// Fungsi melihat daftar acara
func displayEvents(today date, eventType string) {
	if status.totalEvents == 0 {
		fmt.Println("Tidak ada acara.")
		return
	}

	fmt.Println("Acara saat ini dan mendatang:")
	sortEventsByDate()
	for i := 0; i < status.totalEvents; i++ {
		fmt.Printf("ID: %d, Name: &s, Date: %02d-%02d-%04d, Creator ID: %d\n",
			status.eventsList[i].eventId, status.eventsList[i].namaEvent,
			status.eventsList[i].tanggal.tgl, status.eventsList[i].tanggal.bln, status.eventsList[i].tanggal.thn,
			status.eventsList[i].creatorId)
	}
}

func printEvent(event events) {
	fmt.Printf("- %s (%s)\n", event.namaEvent)
	fmt.Printf("	Deskripsi: %s\n", event.deskripsi)
	fmt.Printf("	Tanggal: %s\n", formatDate(event.tanggal))
	fmt.Printf("	Pembuat Acara: %s\n", event.creatorId)
	fmt.Println("-------------------------")
}

func formatDate(d date) string {
	return fmt.Sprintf("Tanggal: %02d-%02d-%04d", d.tgl, d.bln, d.thn)
	// Sprintf: mengembalikan string yang diformat tanpa mencetaknya langsung ke output
}

// Fungsi untuk membuat acara baru
func createEvent() {
	var namaEvent, deskripsi string
	var tgl, bln, thn int

	fmt.Print("Masukkan nama acara: ")
	fmt.Scan(&namaEvent)
	fmt.Println("Masukkan deskripsi acara: ")
	fmt.Scan(&deskripsi)
	fmt.Println("Masukkan tanggal acara (tanggal, bulan, tahun): ")
	fmt.Scan(&tgl, &bln, &thn)

	if status.totalEvents >= NMAX {
		fmt.Println("Event limit reached!")
		return
	}

	status.eventsList[status.totalEvents].eventId = status.totalEvents + 1
	status.eventsList[status.totalEvents].namaEvent = namaEvent
	status.eventsList[status.totalEvents].deskripsi = deskripsi
	status.eventsList[status.totalEvents].tanggal = date{tgl: tgl, bln: bln, thn: thn} // Bentar bingung
	status.eventsList[status.totalEvents].creatorId = status.currentUser.userId

	status.totalEvents++
	fmt.Println("Acara berhasil ditambahkan!")
}

// Fungsi ubah atau edit acara
func editEvent() {
	var idAcara int
	fmt.Print("Masukkan ID acara untuk mengubah: ")
	fmt.Scan(&idAcara)

	eventIndex := temukanAcaraBerdasarkanID(idAcara)
	if eventIndex == -1 {
		fmt.Println("Acara tidak ditemukan!")
		return
	}

	if status.eventsList[eventIndex].creatorId != status.currentUser.userId {
		fmt.Println("Anda tidak memiliki izin untuk mengubah acara ini.")
		return
	}

	var namaEvent, deskripsi string
	var tgl, bln, thn int
	fmt.Printf("Masukkan nama acara yang baru: ")
	fmt.Scan(&namaEvent)
	fmt.Print("Masukkan deskripsi acara yang baru: ")
	fmt.Scan(&deskripsi)
	fmt.Print("Masukkan tanggal acara yang baru (format: dd mm yyyy): ")
	fmt.Scan(&tgl, &bln, &thn)

	status.eventsList[eventIndex].namaEvent = namaEvent
	status.eventsList[eventIndex].deskripsi = deskripsi
	status.eventsList[eventIndex].tanggal = date{tgl: tgl, bln: bln, thn: thn}
	fmt.Println("Acara berhasil diperbarui!")
}

// Fungsi untuk menambah peserta kedalam acara
func tambahPeserta() {
	var eventId, userId int
	fmt.Print("Masukkan ID Acara: ")
	fmt.Scan(&eventId)
	fmt.Print("Masukkan ID Pengguna/Peserta untuk ditambahkan: ")
	fmt.Scan(&userId)

	eventIndex := temukanAcaraBerdasarkanID(eventId)
	if eventIndex == -1 {
		fmt.Println("Acara tidak ditemukan!")
		return
	}

	if status.eventsList[eventIndex].pesertaCount >= NMAX {
		fmt.Println("Telah mencapai limit peserta pada acara ini!")
		return
	}

	userIndex := temukanUserBerdasarkanID(userId)
	if userIndex == -1 {
		fmt.Println("Pengguna tidak ditemukan!")
		return
	}

	status.eventsList[eventIndex].peserta[status.eventsList[eventIndex].pesertaCount] = userId
	status.eventsList[eventIndex].pesertaCount++
	fmt.Println("Pengguna berhasil ditambahkan kedalam acara!")
}

// Fungsi untuk menghapus peserta dari acara terkait
func hapusPeserta() {
	var eventId, userId int
	fmt.Print("Masukkan ID Acara: ")
	fmt.Scan(&eventId)
	fmt.Print("Masukkan ID Pengguna/Peserta yang akan dihapus: ")
	fmt.Scan(&userId)

	eventIndex := temukanAcaraBerdasarkanID(eventId)
	if eventIndex == -1 {
		fmt.Println("Acara tidak ditemukan!")
		return
	}

	for i := 0; i < status.eventsList[eventIndex].pesertaCount; i++ {
		if status.eventsList[eventIndex].peserta[i] == userId {
			status.eventsList[eventIndex].peserta[i] = status.eventsList[eventIndex].peserta[status.eventsList[eventIndex].pesertaCount-1]
			status.eventsList[eventIndex].pesertaCount--
			fmt.Println("Peserta berhasil dihapus dari acara!")
			return
		}
	}
	fmt.Println("Peserta tidak dapat ditemukan pada acara ini!")
}

// Fungsi untuk mengubah profil pengguna (mau apa aja yang diubah disini?)
func editProfile() {
	fmt.Print("Masukkan username yang baru: ")
	fmt.Scan(&status.currentUser.username)
	fmt.Print("Masukkan password yang baru: ")
	fmt.Scan(&status.currentUser.password)
	fmt.Print("Masukkan email yang baru: ")
	fmt.Scan(&status.currentUser.email)
	fmt.Print("Masukkan nomor ponsel yang baru: ")
	fmt.Scan(&status.currentUser.phone)
	fmt.Println("Profil berhasil diperbaharui!")
}

// Fungsi untuk mencari pengguna berdasarkan username
func cariUser() {
	var username string
	fmt.Print("Masukkan username yang dicari: ")
	fmt.Scan(&username)

	for i := 0; i < status.totalUsers; i++ {
		if status.usersList[i].username == username {
			fmt.Printf("Pengguna ditemukan: %s (ID: %d)\n", status.usersList[i].fullName, status.usersList[i].userId)
			return
		}
	}
	fmt.Println("Pengguna tidak ditemukan!")
}

// Fungsi untuk menemukan pengguna berdasarkan ID
func temukanUserBerdasarkanID(userId int) int {
	for i := 0; i < status.totalUsers; i++ {
		if status.usersList[i].userId == userId {
			return i
		}
	}
	return -1
}

// Fungsi untuk menemukan acara berdasarkan ID
func temukanAcaraBerdasarkanID(eventId int) int {
	for i := 0; i < status.totalEvents; i++ {
		if status.eventsList[i].eventId == eventId {
			return i
		}
	}
	return -1
}

// Fungsi untuk mengurutkan daftar pengguna berdasarkan jumlah acara yang diikuti
func sortUsersByEventCount() {
	for i := 0; i < status.totalUsers-1; i++ {
		for j := i + 1; j < status.totalUsers; j++ {
			if status.usersList[i].totalEvents < status.usersList[j].totalEvents {
				status.usersList[i], status.usersList[j] = status.usersList[j], status.usersList[i]
			}
		}
	}
}

// Fungsi untuk mengurutkan daftar acara berdasarkan tanggal
func sortEventsByDate() {
	for i := 0; i < status.totalEvents-1; i++ {
		for j := i + 1; j < status.totalEvents; j++ {
			if compareDates(status.eventsList[i].tanggal, status.eventsList[j].tanggal) > 0 {
				status.eventsList[i], status.eventsList[j] = status.eventsList[j], status.eventsList[i]
			}
		}
	}
}

// Fungsi compare tanggal
func compareDates(d1, d2 date) int {
	if d1.thn != d2.thn {
		return d1.thn - d2.thn
	}
	if d1.bln != d2.bln {
		return d1.bln - d2.bln
	}
	return d1.tgl - d2.tgl
}

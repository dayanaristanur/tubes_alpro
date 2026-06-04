package main

import "fmt"

const NMAX = 100

type Kandidat struct {
	Nama  string
	Kelas string
	Suara int
}

var data [NMAX]Kandidat
var n int

func tambahKandidat() {
	if n >= NMAX {
		fmt.Println("Data penuh!")
		return
	}

	fmt.Print("Nama Kandidat : ")
	fmt.Scan(&data[n].Nama)

	fmt.Print("Kelas : ")
	fmt.Scan(&data[n].Kelas)

	data[n].Suara = 0
	n++

	fmt.Println("Kandidat berhasil ditambahkan!")
}

func tampilkanKandidat() {
	if n == 0 {
		fmt.Println("Belum ada kandidat.")
		return
	}

	fmt.Println("\n=== DAFTAR KANDIDAT ===")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s (%s) - %d suara\n",
			i+1,
			data[i].Nama,
			data[i].Kelas,
			data[i].Suara)
	}
}

func cariKandidat(nama string) int {
	for i := 0; i < n; i++ {
		if data[i].Nama == nama {
			return i
		}
	}
	return -1
}

func ubahKandidat() {
	var nama string

	fmt.Print("Masukkan nama kandidat yang ingin diubah: ")
	fmt.Scan(&nama)

	idx := cariKandidat(nama)

	if idx == -1 {
		fmt.Println("Kandidat tidak ditemukan.")
		return
	}

	fmt.Print("Nama baru : ")
	fmt.Scan(&data[idx].Nama)

	fmt.Print("Kelas baru : ")
	fmt.Scan(&data[idx].Kelas)

	fmt.Println("Data berhasil diubah.")
}

func hapusKandidat() {
	var nama string

	fmt.Print("Masukkan nama kandidat yang ingin dihapus: ")
	fmt.Scan(&nama)

	idx := cariKandidat(nama)

	if idx == -1 {
		fmt.Println("Kandidat tidak ditemukan.")
		return
	}

	for i := idx; i < n-1; i++ {
		data[i] = data[i+1]
	}

	n--

	fmt.Println("Data berhasil dihapus.")
}

func voting() {
	if n == 0 {
		fmt.Println("Belum ada kandidat.")
		return
	}

	tampilkanKandidat()

	var pilih int

	fmt.Print("Pilih nomor kandidat : ")
	fmt.Scan(&pilih)

	if pilih >= 1 && pilih <= n {
		data[pilih-1].Suara++
		fmt.Println("Voting berhasil!")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func pencarian() {
	var nama string

	fmt.Print("Masukkan nama kandidat yang dicari : ")
	fmt.Scan(&nama)

	idx := cariKandidat(nama)

	if idx == -1 {
		fmt.Println("Kandidat tidak ditemukan.")
	} else {
		fmt.Printf("Ditemukan: %s (%s) - %d suara\n",
			data[idx].Nama,
			data[idx].Kelas,
			data[idx].Suara)
	}
}

func urutkanSuara() {
	var temp Kandidat

	for i := 0; i < n-1; i++ {
		max := i

		for j := i + 1; j < n; j++ {
			if data[j].Suara > data[max].Suara {
				max = j
			}
		}

		temp = data[i]
		data[i] = data[max]
		data[max] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan suara terbanyak.")
}

func hasilVoting() {
	if n == 0 {
		fmt.Println("Belum ada kandidat.")
		return
	}

	fmt.Println("\n=== HASIL VOTING ===")
	for i := 0; i < n; i++ {
		fmt.Printf("%s : %d suara\n",
			data[i].Nama,
			data[i].Suara)
	}

	fmt.Printf("\nPemenang sementara: %s dengan %d suara\n",
		data[0].Nama,
		data[0].Suara)
}

func main() {
	var pilihan int

	for {
		fmt.Println("\n===== APLIKASI VOTING =====")
		fmt.Println("1. Tambah Kandidat")
		fmt.Println("2. Tampilkan Kandidat")
		fmt.Println("3. Ubah Kandidat")
		fmt.Println("4. Hapus Kandidat")
		fmt.Println("5. Voting")
		fmt.Println("6. Cari Kandidat")
		fmt.Println("7. Urutkan Berdasarkan Suara")
		fmt.Println("8. Hasil Voting")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahKandidat()
		case 2:
			tampilkanKandidat()
		case 3:
			ubahKandidat()
		case 4:
			hapusKandidat()
		case 5:
			voting()
		case 6:
			pencarian()
		case 7:
			urutkanSuara()
		case 8:
			hasilVoting()
		case 9:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

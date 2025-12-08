package main

import "fmt"

const NMAX int = 100

type buku_komik struct {
	nama_buku, genre, penulis string
	id_buku, stok, tahun      int
}

type member_toko struct {
	nama_member                      string
	npinjam, id_member, bayar, denda int
	pinjam                           [NMAX]string
}

type tabbuku [NMAX]buku_komik
type tabmember [NMAX]member_toko

func main() {
	var buku tabbuku
	var member tabmember
	var banyak_buku, banyak_member, total int
	banyak_buku = 0
	banyak_member = 0
	var n int
	var command bool = true
	dataawal(&buku, &banyak_buku, &member, &banyak_member, &total)
	fmt.Println("======================================================================")
	fmt.Println("||          Selamat Datang Di Aplikasi Peminjaman Komik             ||")
	fmt.Println("||                        Kelompok 7                                ||")
	fmt.Println("======================================================================")
	for command == true {
		n = 0
		fmt.Println("======================================================================")
		fmt.Println("1. Data")
		fmt.Println("2. Cetak Data")
		fmt.Println("3. Stop")
		fmt.Println("======================================================================")
		fmt.Scan(&n)
		if n == 1 {
			fmt.Println("======================================================================")
			fmt.Println("1. Input / Penambahan Data")
			fmt.Println("2. Edit Data")
			fmt.Println("3. Hapus / Pemulangan Data")
			fmt.Println("4. Shorting / Cari Data")
			fmt.Println("======================================================================")
			fmt.Scan(&n)
			if n == 1 {
				fmt.Println("======================================================================")
				fmt.Println("1. Input / Penambahan Data Komik")
				fmt.Println("2. Input / Penambahan Data Member")
				fmt.Println("3. Input / Penambahan Data Peminjaman")
				fmt.Println("======================================================================")
				fmt.Scan(&n)
				if n == 1 {
					inputdatabuku(&buku, &banyak_buku)
				} else if n == 2 {
					inputdatamember(&member, &banyak_member, &total)
				} else if n == 3 {
					inputdatapeminjaman(&member, banyak_member, &buku, banyak_buku)
				}
			} else if n == 2 {
				fmt.Println("======================================================================")
				fmt.Println("1. Edit Data Komik")
				fmt.Println("2. Edit Data Member")
				fmt.Println("3. Edit Data Peminjaman")
				fmt.Println("======================================================================")
				fmt.Scan(&n)
				if n == 1 {
					editdatabuku(&buku, banyak_buku)
				} else if n == 2 {
					editdatamember(&member, &banyak_member)
				} else if n == 3 {
					editdatapeminjaman(&member, banyak_member, &buku, banyak_buku)
				}
			} else if n == 3 {
				fmt.Println("======================================================================")
				fmt.Println("1. Hapus Data Komik")
				fmt.Println("2. Hapus Data Member")
				fmt.Println("3. Pemulangan Peminjaman")
				fmt.Println("======================================================================")
				fmt.Scan(&n)
				if n == 1 {
					hapusdatabuku(&buku, &banyak_buku)
				} else if n == 2 {
					hapusdatamember(&member, &banyak_member)
				} else if n == 3 {
					pemulanganpeminjaman(&member, banyak_member, &buku, banyak_buku, &total)
				}
			} else if n == 4 {
				fmt.Println("======================================================================")
				fmt.Println("1. Shorting buku berdasarkan tahun")
				fmt.Println("2. Shorting buku berdasarkan banyak stok")
				fmt.Println("3. Cetak berdasarkan genre")
				fmt.Println("4. Cari Buku berdasarkan id")
				fmt.Println("======================================================================")
				fmt.Scan(&n)
				if n == 1 {
					shorttahun(&buku, banyak_buku)
				} else if n == 2 {
					shortstok(&buku, banyak_buku)
				} else if n == 3 {
					cetakgenre(buku, banyak_buku)
				} else if n == 4 {
					Shortid(&buku, banyak_buku)
					Cariid(buku, banyak_buku)
				}
			}
		} else if n == 2 {
			fmt.Println("======================================================================")
			fmt.Println("Berikut merupakan list data member")
			fmt.Println("======================================================================")
			for i := 0; i < banyak_member; i++ {
				fmt.Println(member[i].id_member, member[i].nama_member)
			}
			fmt.Println("======================================================================")
			fmt.Println("Berikut merupakan list data komik")
			fmt.Println("======================================================================")
			for j := 0; j < banyak_buku; j++ {
				fmt.Println(buku[j].id_buku, buku[j].nama_buku, buku[j].genre, buku[j].penulis, buku[j].tahun)
				fmt.Print("Sisa stok: ")
				fmt.Print(buku[j].stok)
				fmt.Println("")
			}
			fmt.Println("======================================================================")
			fmt.Println("Berikut merupakan list data peminjaman")
			fmt.Println("======================================================================")
			for l := 0; l < banyak_member; l++ {
				fmt.Print(member[l].id_member, " ", member[l].nama_member, " ")
				if member[l].npinjam != 0 {
					fmt.Print("Meminjam Komik: ")
					for ll := 0; ll < member[l].npinjam; ll++ {
						fmt.Print(member[l].pinjam[ll])
					}
				} else {
					fmt.Print("Meminjam Komik: ", "tidak ada")
				}
				fmt.Println("")
			}
			fmt.Println("======================================================================")
			fmt.Println("Berikut merupakan keuangan berasal dari peminjaman")
			fmt.Println("======================================================================")
			for i := 0; i < banyak_member; i++ {
				fmt.Println(member[i].nama_member, "Bayar", member[i].bayar, "serta denda: ", member[i].denda)
			}
			fmt.Print("Berikut merupakan total pendapatan dari denda dan langganan: ")
			fmt.Print(total)
			fmt.Println("")
		} else if n == 3 {
			command = false
		}
	}
}

func Cariid(A tabbuku, n int) {
	var id, idx int
	fmt.Println("masukan id yang ingin dicari:")
	fmt.Scan(&id)
	idx = binersearch(A, n, id)
	if idx != -1 {
		fmt.Println("buku dengab id ", id, "adalah sebagai berikut: ")
		fmt.Println(A[idx].id_buku, A[idx].nama_buku, A[idx].genre, A[idx].penulis, A[idx].tahun)
	} else {
		fmt.Println("tidak ada buku dengan id ", id)
	}
}

func binersearch(A tabbuku, n, id int) int {
	var L, R, M, i, idx int
	idx = -1
	L = 0
	R = n - 1
	for i < n && idx == -1 {
		M = (L + R) / 2
		if A[M].id_buku == id {
			idx = M
		} else if A[M].id_buku > id {
			R = M - 1
		} else if A[M].id_buku < id {
			L = M + 1
		}
		i++
	}
	return idx
}

func Shortid(A *tabbuku, n int) {
	var i, pass int
	var temp buku_komik
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[i]
		for i > 0 && temp.id_buku < A[i-1].id_buku {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func cetakgenre(A tabbuku, n int) {
	fmt.Println("======================================================================")
	fmt.Println("Berikut merupakan list komik terurut berdasarkan genre")
	fmt.Println("======================================================================")
	fmt.Println("Genre Action")
	fmt.Println("======================================================================")
	for i := 0; i < n; i++ {
		if A[i].genre == "Action" {
			fmt.Println(A[i].id_buku, A[i].nama_buku, A[i].genre, A[i].penulis, A[i].tahun)
			fmt.Print("Sisa stok: ")
			fmt.Print(A[i].stok)
			fmt.Println("")
		}
	}
	fmt.Println("======================================================================")
	fmt.Println("Genre Adventure")
	fmt.Println("======================================================================")
	for i := 0; i < n; i++ {
		if A[i].genre == "Adventure" {
			fmt.Println(A[i].id_buku, A[i].nama_buku, A[i].genre, A[i].penulis, A[i].tahun)
			fmt.Print("Sisa stok: ")
			fmt.Print(A[i].stok)
			fmt.Println("")
		}
	}
	fmt.Println("======================================================================")
	fmt.Println("Genre Fantacy")
	fmt.Println("======================================================================")
	for i := 0; i < n; i++ {
		if A[i].genre == "Fantasy" {
			fmt.Println(A[i].id_buku, A[i].nama_buku, A[i].genre, A[i].penulis, A[i].tahun)
			fmt.Print("Sisa stok: ")
			fmt.Print(A[i].stok)
			fmt.Println("")
		}
	}
	fmt.Println("======================================================================")
	fmt.Println("Genre Mystery")
	fmt.Println("======================================================================")
	for i := 0; i < n; i++ {
		if A[i].genre == "Mystery" {
			fmt.Println(A[i].id_buku, A[i].nama_buku, A[i].genre, A[i].penulis, A[i].tahun)
			fmt.Print("Sisa stok: ")
			fmt.Print(A[i].stok)
			fmt.Println("")
		}
	}
	fmt.Println("======================================================================")
	fmt.Println("Genre Comedy")
	fmt.Println("======================================================================")
	for i := 0; i < n; i++ {
		if A[i].genre == "Comedy" {
			fmt.Println(A[i].id_buku, A[i].nama_buku, A[i].genre, A[i].penulis, A[i].tahun)
			fmt.Print("Sisa stok: ")
			fmt.Print(A[i].stok)
			fmt.Println("")
		}
	}
}

func shortstok(A *tabbuku, n int) {
	var i, idx, pass int
	var temp buku_komik
	pass = 1
	for pass <= n-1 {
		i = pass
		idx = pass - 1
		for i < n {
			if A[idx].stok < A[i].stok {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func shorttahun(A *tabbuku, n int) {
	var i, pass int
	var temp buku_komik
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = A[i]
		for i > 0 && temp.tahun > A[i-1].tahun {
			A[i] = A[i-1]
			i = i - 1
		}
		A[i] = temp
		pass = pass + 1
	}
}

func pemulanganpeminjaman(member *tabmember, banyakm int, buku *tabbuku, banyakb int, total *int) {
	var n, idx, durasi, index, bayar, denda int
	var nama, kondisi string
	bayar = 0
	denda = 0
	index = -1
	fmt.Println("Masukan ID member: ")
	fmt.Scan(&n)
	idx = caridatamember(*member, banyakm, n)
	if idx != -1 {
		if member[idx].npinjam != 0 {
			fmt.Println("Berikut adalah komik yang", member[idx].nama_member, "pinjam: ")
			for i := 0; i < member[idx].npinjam; i++ {
				fmt.Println(member[idx].pinjam[i])
			}
			fmt.Println("Masukan nama komik yang dikembalikan")
			fmt.Scan(&nama)
			for i := 0; i < member[idx].npinjam; i++ {
				if member[idx].pinjam[i] == nama {
					index = i
				}
			}
			if index != -1 {
				for i := 0; i < member[idx].npinjam; i++ {
					member[idx].pinjam[i] = member[idx].pinjam[i+1]
				}
				member[idx].npinjam--
				for i := 0; i < banyakb; i++ {
					if buku[i].nama_buku == nama {
						buku[i].stok++
					}
				}
				fmt.Println("Kondisi komik saat pengembalian Sama / Rusak / Kotor")
				fmt.Scan(&kondisi)
				if kondisi == "Rusak" {
					bayar = bayar + (5000 * (durasi - 7))
					denda = denda + (5000 * (durasi - 7))bayar = bayar + 20000
					denda = denda + 20000
				} else if kondisi == "Kotor" {
					bayar = bayar + 10000
					denda = denda + 10000
				}
				fmt.Println("Masukan lama durasi", member[idx].nama_member, "meminjam komik: ")
				fmt.Scan(&durasi)
				if durasi > 7 {
					
				}
				member[idx].bayar = member[idx].bayar + bayar
				member[idx].denda = member[idx].denda + denda
			} else {
				fmt.Println(member[idx].nama_member, "Tidak meminjam komik berjudul", nama)
			}
		} else {
			fmt.Println(member[idx].nama_member, "Tidak sedang meminjam komik")
		}
		*total = *total + member[idx].bayar
	} else {
		fmt.Println("Member tidak ada")
	}
}

func editdatapeminjaman(member *tabmember, banyakm int, buku *tabbuku, banyakb int) {
	var id, idx, n, index, jumlah int
	var command string
	fmt.Println("Masukan ID member: ")
	fmt.Scan(&n)
	idx = caridatamember(*member, banyakm, n)
	if idx != -1 {
		jumlah = 0
		for command != "t" {
			fmt.Println("Masukan komik yang dipinjam: ")
			fmt.Scan(&id)
			index = caridatabuku(*buku, banyakb, id)
			if index != -1 {
				member[idx].pinjam[jumlah] = buku[index].nama_buku
				buku[index].stok--
				jumlah++
			} else {
				fmt.Println("bukutidak ada")
			}
			if buku[index].stok != 0 {
				fmt.Println("lanjut? y/t")
				fmt.Scan(&command)
			} else {
				command = "t"
			}
		}
		member[idx].npinjam = jumlah
	} else {
		fmt.Println("member tidak ada")
	}
}

func inputdatapeminjaman(member *tabmember, banyakm int, buku *tabbuku, banyakb int) {
	var n, idx, index, jumlah, id int
	var command string
	fmt.Println("masukan id member yang akan meminjam: ")
	fmt.Scan(&n)
	idx = caridatamember(*member, banyakm, n)
	if idx != -1 {
		jumlah = member[idx].npinjam
		for command != "t" {
			fmt.Println("masukan id Komik yang dipinjam: ")
			fmt.Scan(&id)
			index = caridatabuku(*buku, banyakb, id)
			if index != -1 {
				member[idx].pinjam[jumlah] = buku[index].nama_buku
				buku[index].stok--
				jumlah++
			} else {
				fmt.Println("Komik tidak ada")
			}
			if buku[index].stok != 0 {
				fmt.Println("Lanjut? y/t")
				fmt.Scan(&command)
			} else {
				command = "t"
			}
		}
		member[idx].npinjam = jumlah
	} else {
		fmt.Println("member tidak ada")
	}
}

func inputdatabuku(buku *tabbuku, banyak *int) {
	var banyak1 int
	fmt.Println("======================================================================")
	fmt.Println("masukan banyaknya data")
	fmt.Println("======================================================================")
	fmt.Scan(&banyak1)
	fmt.Println("======================================================================")
	fmt.Println("masukan data")
	fmt.Println("======================================================================")
	for i := *banyak; i < (*banyak + banyak1); i++ {
		fmt.Println("masukan id komik")
		fmt.Scan(&buku[i].id_buku)
		fmt.Println("masukan judul komik")
		fmt.Scan(&buku[i].nama_buku)
		fmt.Println("masukan genre komik")
		fmt.Scan(&buku[i].genre)
		fmt.Println("masukan penulis komik")
		fmt.Scan(&buku[i].penulis)
		fmt.Println("masukan banyaknya stok komik")
		fmt.Scan(&buku[i].stok)
		fmt.Println("masukan banyaknya tahun release komik")
		fmt.Scan(&buku[i].tahun)
	}
	*banyak = *banyak + banyak1
}

func inputdatamember(member *tabmember, banyak, total *int) {
	var banyak1 int
	fmt.Println("======================================================================")
	fmt.Println("masukan banyaknya data")
	fmt.Println("======================================================================")
	fmt.Scan(&banyak1)
	for i := *banyak; i < (*banyak + banyak1); i++ {
		fmt.Println("masukan id member")
		fmt.Scan(&member[i].id_member)
		fmt.Println("masukan nama member")
		fmt.Scan(&member[i].nama_member)
		member[i].bayar = 25000
		*total = *total + member[i].bayar
	}
	*banyak++
}

func editdatabuku(buku *tabbuku, banyak int) {
	var idx int
	var n int
	fmt.Println("======================================================================")
	fmt.Println("masukan id komik untuk di edit")
	fmt.Println("======================================================================")
	fmt.Scan(&n)
	idx = caridatabuku(*buku, banyak, n)
	if idx != -1 {
		fmt.Println("masukan judul komik")
		fmt.Scan(&buku[idx].nama_buku)
		fmt.Println("masukan genre komik")
		fmt.Scan(&buku[idx].genre)
		fmt.Println("masukan penulis komik")
		fmt.Scan(&buku[idx].penulis)
		fmt.Println("masukan banyaknya stok komik")
		fmt.Scan(&buku[idx].stok)
	} else if idx == -1 {
		fmt.Println("Tidak Ada")
	}
}

func caridatabuku(buku tabbuku, banyak int, n int) int {
	var idx, i int
	idx = -1
	i = 0
	for i < banyak && idx == -1 {
		if buku[i].id_buku == n {
			idx = i
		}
		i = i + 1
	}
	return idx
}

func editdatamember(member *tabmember, banyak *int) {
	var idx int
	var n int
	fmt.Println("======================================================================")
	fmt.Println("masukan id member untuk di edit")
	fmt.Println("======================================================================")
	fmt.Scan(&n)
	idx = caridatamember(*member, *banyak, n)
	if idx != -1 {
		fmt.Scan(&member[idx].nama_member)
	} else if idx == -1 {
		fmt.Println("Member tidak ada")
	}
}

func caridatamember(member tabmember, banyak, n int) int {
	var idx, i int
	i = 0
	idx = -1
	for i < banyak && idx == -1 {
		if member[i].id_member == n {
			idx = i
		}
		i++
	}
	return idx
}

func hapusdatabuku(buku *tabbuku, banyak *int) {
	var idx int
	var n int
	fmt.Println("======================================================================")
	fmt.Println("masukan id komik untuk di hapus")
	fmt.Println("======================================================================")
	fmt.Scan(&n)
	idx = caridatabuku(*buku, *banyak, n)
	if idx != -1 {
		for idx < *banyak {
			buku[idx] = buku[idx+1]
			idx = idx + 1
		}
	} else {
	}
	*banyak = *banyak - 1
}

func hapusdatamember(member *tabmember, banyak *int) {
	var idx int
	var n int
	fmt.Println("======================================================================")
	fmt.Println("masukan id member untuk di hapus")
	fmt.Println("======================================================================")
	fmt.Scan(&n)
	idx = caridatamember(*member, *banyak, n)
	fmt.Println(idx)
	if idx != -1 {
		for idx < *banyak {
			member[idx].id_member = member[idx+1].id_member
			member[idx].nama_member = member[idx+1].nama_member
			member[idx].npinjam = member[idx+1].npinjam
			for i := 0; i < member[idx].npinjam; i++ {
				member[idx].pinjam[i] = member[idx+1].pinjam[i+1]
			}
			idx = idx + 1
		}
	}
	*banyak = *banyak - 1
}

func selectionSort(buku *tabbuku, banyak int) {
	var i, idx, j int
	var t buku_komik
	i = 1
	for i <= banyak-1 {
		idx = i - 1
		j = i
		for j < banyak {
			if buku[idx].stok > buku[j].stok {
				idx = j
			}
			j++
		}
		t = buku[idx]
		buku[idx] = buku[i-1]
		buku[i-1] = t
		i++
	}
}

func dataawal(A *tabbuku, Nb *int, B *tabmember, Nm, total *int) {
	A[0].id_buku = 1
	A[0].nama_buku = "Watchmen"
	A[0].genre = "Action"
	A[0].penulis = "Moore"
	A[0].stok = 19
	A[0].tahun = 1986

	A[1].id_buku = 2
	A[1].nama_buku = "Y"
	A[1].genre = "Action"
	A[1].penulis = "Vaughan"
	A[1].stok = 14
	A[1].tahun = 2002

	A[2].id_buku = 11
	A[2].nama_buku = "Knight"
	A[2].genre = "Action"
	A[2].penulis = "Miller"
	A[2].stok = 50
	A[2].tahun = 1986

	A[3].id_buku = 3
	A[3].nama_buku = "Sandman"
	A[3].genre = "Fantasy"
	A[3].penulis = "Gaiman"
	A[3].stok = 45
	A[3].tahun = 1989

	A[4].id_buku = 4
	A[4].nama_buku = "Bone"
	A[4].genre = "Fantasy"
	A[4].penulis = "Smith"
	A[4].stok = 22
	A[4].tahun = 1991

	A[5].id_buku = 5
	A[5].nama_buku = "Fables"
	A[5].genre = "Fantasy"
	A[5].penulis = "Willingham"
	A[5].stok = 18
	A[5].tahun = 2002

	A[6].id_buku = 6
	A[6].nama_buku = "Blacksad"
	A[6].genre = "Mystery"
	A[6].penulis = "Canales"
	A[6].stok = 25
	A[6].tahun = 2000

	A[7].id_buku = 7
	A[7].nama_buku = "Joke"
	A[7].genre = "Mystery"
	A[7].penulis = "Moore"
	A[7].stok = 30
	A[7].tahun = 1988

	A[8].id_buku = 8
	A[8].nama_buku = "Vendetta"
	A[8].genre = "Mystery"
	A[8].penulis = "Moore"
	A[8].stok = 20
	A[8].tahun = 1988

	A[9].id_buku = 9
	A[9].nama_buku = "Persepolis"
	A[9].genre = "Adventure"
	A[9].penulis = "Satrapi"
	A[9].stok = 20
	A[9].tahun = 2000

	A[10].id_buku = 10
	A[10].nama_buku = "Maus"
	A[10].genre = "Adventure"
	A[10].penulis = "Spiegelman"
	A[10].stok = 30
	A[10].tahun = 1991

	A[11].id_buku = 12
	A[11].nama_buku = "Tintin"
	A[11].genre = "Adventure"
	A[11].penulis = "Hergé"
	A[11].stok = 15
	A[11].tahun = 1929

	A[12].id_buku = 13
	A[12].nama_buku = "Saga"
	A[12].genre = "Comedy"
	A[12].penulis = "Vaughan"
	A[12].stok = 88
	A[12].tahun = 2012

	A[13].id_buku = 14
	A[13].nama_buku = "Deadpool"
	A[13].genre = "Comedy"
	A[13].penulis = "Nicieza"
	A[13].stok = 40
	A[13].tahun = 1991

	A[14].id_buku = 15
	A[14].nama_buku = "Squirrel"
	A[14].genre = "Comedy"
	A[14].penulis = "North"
	A[14].stok = 25
	A[14].tahun = 2015

	*Nb = 15

	B[0].id_member = 1
	B[0].nama_member = "Rizqy"
	B[0].npinjam = 1
	B[0].pinjam[0] = "Watchmen"
	B[0].bayar = 25000

	B[1].id_member = 2
	B[1].nama_member = "Ikhsan"
	B[1].npinjam = 1
	B[1].pinjam[0] = "Saga"
	B[1].bayar = 25000

	B[2].id_member = 3
	B[2].nama_member = "Ladaku"
	B[2].npinjam = 1
	B[2].pinjam[0] = "Saga"
	B[2].bayar = 25000

	B[3].id_member = 4
	B[3].nama_member = "Ali"
	B[3].npinjam = 1
	B[3].pinjam[0] = "Maus"
	B[3].bayar = 25000

	B[4].id_member = 5
	B[4].nama_member = "Budi"
	B[4].npinjam = 1
	B[4].pinjam[0] = "Bone"
	B[4].bayar = 25000

	B[5].id_member = 6
	B[5].nama_member = "Cici"
	B[5].npinjam = 1
	B[5].pinjam[0] = "Y"
	B[5].bayar = 25000

	B[6].id_member = 7
	B[6].nama_member = "Dina"
	B[6].npinjam = 1
	B[6].pinjam[0] = "Persepolis"
	B[6].bayar = 25000

	B[7].id_member = 8
	B[7].nama_member = "Eko"
	B[7].npinjam = 1
	B[7].pinjam[0] = "Blacksad"
	B[7].bayar = 25000

	B[8].id_member = 9
	B[8].nama_member = "Fina"
	B[8].npinjam = 1
	B[8].pinjam[0] = "Fables"
	B[8].bayar = 25000

	B[9].id_member = 10
	B[9].nama_member = "Gita"
	B[9].npinjam = 1
	B[9].pinjam[0] = "Watchmen"
	B[9].bayar = 25000

	*Nm = 10
	*total = 250000
}
package main

import "fmt"

const NMAX int = 1000

type playlist struct {
	lagu     string
	penyanyi string
	durasi   int
}

type arrPlaylist [NMAX]playlist

func main() {
	var T arrPlaylist
	var i int

	IsiPlaylist(&T, &i)
	showInfo(&T, i)
	findFastDuration(&T, i)
}

func IsiPlaylist(T *arrPlaylist, i *int) {
	var n, durasi int
	var lagu, penyanyi string

	fmt.Println("Masukan jumlah data playlist:")
	fmt.Scanln(&n)

	for *i = 0; *i < n; *i++ {
		fmt.Println("Masukan lagu:")
		fmt.Scanln(&lagu)
		fmt.Println("Masukan penyanyi:")
		fmt.Scanln(&penyanyi)
		fmt.Println("Masukan durasi:")
		fmt.Scanln(&durasi)

		T[*i].lagu = lagu
		T[*i].penyanyi = penyanyi
		T[*i].durasi = durasi
	}
}

func showInfo(T *arrPlaylist, i int) {
	for n := 0; n < i; n++ {
		fmt.Println("Lagu :", T[n].lagu)
		fmt.Println("Penyanyi :", T[n].penyanyi)
		fmt.Println("Durasi :", T[n].durasi, "detik")
	}
}

func findFastDuration(T *arrPlaylist, i int) {
	if i == 0 {
		fmt.Println("Tidak ada data dalam playlist.")
		return
	}

	var min playlist
	min = T[0]

	for n := 1; n < i; n++ {
		if T[n].durasi < min.durasi {
			min = T[n]
		}
	}

	fmt.Println("Lagu dengan durasi tersingkat adalah:")
	fmt.Println("Lagu :", min.lagu)
	fmt.Println("Penyanyi :", min.penyanyi)
	fmt.Println("Durasi :", min.durasi)
}

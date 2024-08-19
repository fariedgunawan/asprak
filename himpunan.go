package main

import "fmt"

const NMAX int = 1000

type angka int

type arrAngka [NMAX]angka

func main() {
	var T arrAngka
	var array1, array2 int
	var T2 arrAngka

	fmt.Println("Masukkan jumlah angka untuk array 1:")
	fmt.Scanln(&array1)
	isiArray(&T, array1)
	
	fmt.Println("Masukkan jumlah angka untuk array 2:")
	fmt.Scanln(&array2)
	isiArray(&T2, array2)
	
	fmt.Println("Hasil:")
	if isSame(T, T2, array1, array2) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}

func isiArray(T *arrAngka, n int) {
	for i := 0; i < n; i++ {
		fmt.Scanln(&T[i])
	}
}

func SortAscending(T *arrAngka, n int) {
	for i := 1; i < n; i++ {
		temp := T[i]
		j := i - 1
		
		for j >= 0 && T[j] > temp {
			T[j+1] = T[j]
			j--
		}

		T[j+1] = temp
	}
}

func isSame(T arrAngka, S arrAngka, n int, t int) bool {
	if n != t {
		return false
	}
	
	SortAscending(&T, n)
	SortAscending(&S, t)
	
	for i := 0; i < n; i++ {
		if T[i] != S[i] {
			return false
		}
	}
	
	return true
}

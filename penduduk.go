package main

import "fmt"

const NMAX int = 100

type penduduk struct{
	nama string
	umur int
}

type arrPenduduk [NMAX]penduduk


func main(){
	var T arrPenduduk
	var jumlah int
	
	fmt.Println("Masukan jumlah data penduduk : ")
	fmt.Scan(&jumlah)
	
	inputPenduduk(&T,jumlah)
	showPenduduk(T,jumlah)
	fmt.Println("Setelah melakukan sorting dari umur")
	sortByAge(&T,jumlah)
	showPenduduk(T,jumlah)
	fmt.Println("Rerata umur penduduk : ")
	fmt.Print(avgYear(T,jumlah))
}

func inputPenduduk(T *arrPenduduk, i int){
	var nama string
	var umur int
	
	for n:= 0;n < i;n++{
		fmt.Print("Masukan Nama: ")
		fmt.Scan(&nama)
		fmt.Println("Masukan Umur: ")
		fmt.Scan(&umur)
		
		T[n].nama = nama
		T[n].umur = umur
	}
}

func showPenduduk(T arrPenduduk, i int){
	for n:=0; n < i; n++{
		fmt.Println("Nama penduduk ke ",n+1, " : ", T[n].nama)
		fmt.Println("Umur penduduk ke ",n+1, " : ", T[n].umur)
	}
}

func sortByAge(T *arrPenduduk, n int){
	var temp penduduk
	var i , j int
	
	i = 1
	for i <= n-1{
		j = i
		temp = T[j]
		for j > 0 && temp.umur > T[j-1].umur{
			T[j] = T[j-1]
			j--
		}
		T[j] = temp
		i++
	}
}

func avgYear(T arrPenduduk, i int) float64 {
	var sum int

	for n := 0; n < i; n++ {
		sum = sum + T[n].umur
	}

	average := float64(sum) / float64(i)

	return average
}





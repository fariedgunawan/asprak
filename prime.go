package main

import "fmt"

func main() {
	var n,f,d int
	
	fmt.Println("Masukkan nilai N:")
	fmt.Scanln(&n)
	
	fmt.Println("Bilangan prima dari 1 sampai", n, "adalah:")
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			if i % 2 == 0 {
				fmt.Print("Angka Genap: ")
				f++
			} else {
				fmt.Print("Angka Ganjil: ")
				d++
			}
			fmt.Println(i)
		}
	}
	
	fmt.Println("Banyaknya angka genap:",f)
	fmt.Println("Banyaknya angka ganjil:",d)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

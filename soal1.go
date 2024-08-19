package main

import "fmt"

const NMAX int = 100

type info int
type arrAngka [NMAX]info

func main() {
    var A arrAngka
    var n, i, angka int

    fmt.Println("Masukkan jumlah angka:")
    fmt.Scanln(&n)

    for i = 0; i < n; i++ {
        fmt.Printf("Masukkan angka ke-%d: ", i+1)
        fmt.Scanln(&angka)
        insertNumber(&A, &i, info(angka))
    }

    fmt.Println("Array yang diinput:")
    for i = 0; i < n; i++ {
        fmt.Println(A[i])
    }

    max := findMax(A, n)
    fmt.Println("Nilai maksimum dalam array adalah:", max)
}

func insertNumber(A *arrAngka, i *int, n info) {
    (*A)[*i] = n
    *i++
}

func findMax(A arrAngka, n int) int {
    var i int
    max := int(A[0]) 

    for i = 1; i < n; i++ {
        if int(A[i]) > max {
            max = int(A[i])
        }
    }

    return max
}

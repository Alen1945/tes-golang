package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(number int) string {
	numberString := strconv.Itoa(number)
	arrString := strings.Split(numberString, "")
	reverseString := strings.Split(numberString, "")
	for i := 0; i < len(arrString); i++ {
		reverseString[i] = arrString[(len(arrString)-1)-i]
	}
	return strings.Join(reverseString, "")
}
func isPalindrom(number int) bool {
	return strconv.Itoa(number) == reverse(number)
}
func main() {
	nPalindrom := 0
	var batasAwal int
	var batasAkhir int
	fmt.Print("input angka batas awal: ")
	fmt.Scan(&batasAwal)
	fmt.Print("input angka batas akhir: ")
	fmt.Scan(&batasAkhir)
	for i := batasAwal; i <= batasAkhir; i++ {
		if isPalindrom(i) {
			nPalindrom = nPalindrom + 1
		}
	}
	fmt.Println(nPalindrom)
}

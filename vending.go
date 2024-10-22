package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Drink struct {
	name  string
	price int
}

var drinks = []Drink{
	{"Coffee", 5000},
	{"Tea", 3000},
	{"Mineral Water", 4000},
	{"Soda", 6000},
}

func main() {
	fmt.Println("Welcome to Vending Machine!")
	fmt.Println("Minuman tersedia:")

	// Cetak header tabel
	fmt.Printf("%-5s %-15s %-10s\n", "No.", "Minuman", "Harga")
	fmt.Println("-----------------------------------")

	// Cetak isi tabel
	for i, drink := range drinks {
		fmt.Printf("%-5d %-15s Rp. %-10d\n", i+1, drink.name, drink.price)
	}

	var choices string
	fmt.Print("Pilih minumanmu (misal: 1,2,3): ")
	fmt.Scanln(&choices)

	// Pisahkan input berdasarkan koma dan buat array dari pilihan
	choiceList := strings.Split(choices, ",")

	var totalPrice int
	var selectedDrinks []string

	// Proses setiap pilihan
	for _, choice := range choiceList {
		// Konversi dari string ke int dan kurangi 1 karena index array dimulai dari 0
		choiceInt, err := strconv.Atoi(strings.TrimSpace(choice))
		if err != nil || choiceInt < 1 || choiceInt > len(drinks) {
			fmt.Println("Pilihan tidak valid, coba lagi.")
			return
		}

		chosenDrink := drinks[choiceInt-1]
		selectedDrinks = append(selectedDrinks, chosenDrink.name)
		totalPrice += chosenDrink.price
	}

	fmt.Printf("Pilihanmu: %s. Total harga: Rp. %d\n", strings.Join(selectedDrinks, ", "), totalPrice)

	var payment int
	fmt.Print("Masukkan uang (Rp. ): ")
	fmt.Scanln(&payment)

	if payment < totalPrice {
		fmt.Println("Pembayaran tidak mencukupi. Silakan coba lagi.")
		return
	}

	change := payment - totalPrice
	fmt.Printf("Silakan ambil minumanmu: %s. Selamat menikmati!\n", strings.Join(selectedDrinks, ", "))

	// Pengembalian uang jika ada kembalian
	if change > 0 {
		fmt.Printf("Kembalianmu: Rp. %d\n", change)
	}

	fmt.Println("Terima kasih telah menggunakan vending machine!")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadBytes('\n') // Waits for the Enter key
}

package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Flight struct {
	Airline    string
	Price      int
	Departure  string
	Arrival    string
	FlightTime string
}

type Booking struct {
	Name        string
	Flight      Flight
	Destination string
}

var bookings []Booking

func main() {
	for {
		var choice int
		fmt.Println("Menu:")
		fmt.Println("1. Penerbangan Domestik")
		fmt.Println("2. Penerbangan Internasional")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		if choice == 3 {
			fmt.Println("Terima kasih telah menggunakan layanan kami.")
			break
		}

		if choice != 1 && choice != 2 {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		var destination string
		if choice == 1 {
			fmt.Print("Masukkan tujuan penerbangan domestik: ")
		} else if choice == 2 {
			fmt.Print("Masukkan tujuan penerbangan internasional: ")
		}
		fmt.Scan(&destination)

		flights := getFlights(destination)
		if len(flights) == 0 {
			fmt.Println("Tidak ada penerbangan yang tersedia untuk tujuan ini.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Println("Daftar penerbangan:")
		printFlights(flights)

		var sortChoice int
		fmt.Println("Menu Sorting:")
		fmt.Println("1. Termurah ke Termahal")
		fmt.Println("2. Termahal ke Termurah")
		fmt.Println("3. Berdasarkan Waktu Keberangkatan")
		fmt.Print("Pilih menu sorting: ")
		fmt.Scan(&sortChoice)

		switch sortChoice {
		case 1:
			sort.Slice(flights, func(i, j int) bool {
				return flights[i].Price < flights[j].Price
			})
		case 2:
			sort.Slice(flights, func(i, j int) bool {
				return flights[i].Price > flights[j].Price
			})
		case 3:
			sort.Slice(flights, func(i, j int) bool {
				return flights[i].Departure < flights[j].Departure
			})
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Println("Daftar penerbangan setelah sorting:")
		printFlights(flights)

		var flightChoice int
		fmt.Print("Pilih penerbangan (nomor): ")
		fmt.Scan(&flightChoice)
		if flightChoice < 1 || flightChoice > len(flights) {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		flight := flights[flightChoice-1]
		fmt.Printf("Detail penerbangan yang dipilih:\nMaskapai: %s\nHarga: Rp%d\nKeberangkatan: %s\nKedatangan: %s\nDurasi: %s\n",
			flight.Airline, flight.Price, flight.Departure, flight.Arrival, flight.FlightTime)

		var confirm string
		fmt.Print("Apakah Anda ingin melanjutkan dengan penerbangan ini? (yes/no): ")
		fmt.Scan(&confirm)
		if strings.ToLower(confirm) != "yes" {
			fmt.Println("Pemesanan dibatalkan.")
			time.Sleep(2 * time.Second)
			continue
		}

		var name string
		fmt.Print("Masukkan nama Anda: ")
		fmt.Scan(&name)

		booking := Booking{
			Name:        name,
			Flight:      flight,
			Destination: destination,
		}
		bookings = append(bookings, booking)

		fmt.Printf("Anda telah sukses terdaftar pada penerbangan %s dengan tujuan %s.\n", flight.Airline, destination)
		time.Sleep(3 * time.Second)
	}
}

func getFlights(destination string) []Flight {
	// Dummy data, biasanya ini akan diambil dari database atau API
	return []Flight{
		{"Garuda Indonesia", 1500000, "08:00", "10:00", "2h"},
		{"Lion Air", 900000, "09:00", "11:00", "2h"},
		{"AirAsia", 1200000, "10:00", "12:00", "2h"},
		{"Singapore Airlines", 2500000, "11:00", "13:00", "2h"},
	}
}

func printFlights(flights []Flight) {
	for i, flight := range flights {
		fmt.Printf("%d. %s - Rp%d (Keberangkatan: %s, Kedatangan: %s)\n", i+1, flight.Airline, flight.Price, flight.Departure, flight.Arrival)
	}
}

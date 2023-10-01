package main

import (
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Id           int
	Name         string
	Address      string
	Occupation   string
	ReasonGolang string
}

var classmate = []Person{
	{1, "Fitri", "Jakarta", "Developer", "Suka dengan bahasa Go"},
	{2, "Cesar", "Jakarta", "Software Engineer", "Menambah Skill"},
	{3, "Budi", "Bandung", "Accounting", "Switch Career"},
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Silakan masukkan nama teman / ID teman yang ingin Anda cari.")
		return
	}

	searchParam := args[0]
	var indexOfFoundClassmate int

	if num, err := strconv.Atoi(searchParam); err == nil {
		indexOfFoundClassmate := findClassmateById(num)

		if indexOfFoundClassmate == -1 {
			fmt.Printf("Data Teman dengan Id '%d' tidak ditemukan.\n", num)
		}
	} else {
		indexOfFoundClassmate := findClassmateByName(searchParam)

		if indexOfFoundClassmate == -1 {
			fmt.Printf("Teman dengan nama '%s' tidak ditemukan.\n", searchParam)
		}
	}

	showClassmateData(classmate[indexOfFoundClassmate])
}

func findClassmateByName(name string) int {
	for i, person := range classmate {
		if person.Name == name {
			return i
		}
	}
	return -1
}

func findClassmateById(id int) int {
	for i, person := range classmate {
		if person.Id == id {
			return i
		}
	}

	return -1
}

func showClassmateData(person Person) {
	fmt.Println("Nama:", person.Name)
	fmt.Println("Alamat:", person.Address)
	fmt.Println("Pekerjaan:", person.Occupation)
	fmt.Println("Alasan memilih kelas Golang:", person.ReasonGolang)
}

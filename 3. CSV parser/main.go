package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type student struct {
	Name       string  `csv:"name"`
	Age        int64   `csv:"age"`
	Percentage float64 `csv:"percentage"`
	Address    string  `csv:"address"`
}

func readCsvFile(filePath string) []*student {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	students := []*student{}

	if err := gocsv.UnmarshalFile(f, &students); err != nil {
		panic(err)
	}

	return students
}

func main() {
	records := readCsvFile("student.csv")
	fmt.Println("Index     Name  Age Percentage  Address")

	for i, st := range records {
		// fmt.Println(st)
		// fmt.Println(strconv.Itoa(i) + " " + st.Name + " " + strconv.FormatInt(int64(st.Age), 10) + " " + fmt.Sprintf("%v", (st.Percentage)) + " " + st.Address)
		fmt.Printf("%-5d %8s %4d %8.2f %10s\n", i, st.Name, st.Age, st.Percentage, st.Address)
	}
}

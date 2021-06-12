package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/twpayne/go-polyline"
)

var coords [][]float64
var parts []float64

func main() {
	readFile()
	fmt.Println(string(polyline.EncodeCoords(coords)))
}

func readFile() {
	file, err := os.Open("./example.log")
	if err != nil {
		fmt.Println("Hubo un error: ", err)
	} else {
		scanner := bufio.NewScanner(file)
		i := 1
		for scanner.Scan() {
			line := scanner.Text()
			extractCoords(line, i)
		}
	}
	file.Close()
}

func extractCoords(line string, i int) {
	index := strings.Index(line, "Location arrived")
	//line position 59 is where the word "Location arrived" is found, on that line are the coordinates
	if index == 59 {
		sl1 := strings.Split(line, "<")
		sl2 := strings.Split(sl1[1], ">")
		sl3 := strings.Split(sl2[0], ",")
		storeCoords(sl3[0], sl3[1])
		i++
	}
}

func storeCoords(left string, right string) {
	val1, _ := strconv.ParseFloat(left, 64)
	val2, _ := strconv.ParseFloat(right, 64)

	parts = append(parts, val1)
	parts = append(parts, val2)
	coords = append(coords, [][]float64{parts}...)
	parts = nil
}

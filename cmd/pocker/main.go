package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/firstrow/pocker"
	"log"
	"os"
	"strings"
)

func init() {
	flag.Parse()
}

func main() {
	if len(flag.Args()) == 0 {
		log.Fatal("Please provide path to input file.")
	}
	filepath := flag.Args()[0]

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		hand := arr[0:5]
		deck := arr[5:10]

		result := pocker.BestHand(hand, deck)
		fmt.Println(result.ToString())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

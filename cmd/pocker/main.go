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
		codes := strings.Split(line, " ")
		hand := codes[0:5]
		deck := codes[5:10]

		best := pocker.BestHand(hand, deck)
		// TODO: Format as in task examples
		fmt.Println(best.ToString())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

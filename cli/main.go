package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rainbowkatz/split"
)

func main() {
	// cli flags
	tab := flag.Float64("total", 10.00, "Total amount to split into shares")
	shares := flag.Int("shares", 2, "Number of shares to split tab into")
	diff := flag.Float64("diff", 1.00, "Difference between consecutive shares")
	isV := flag.Bool("v", false, "Prints more verbose output")
	flag.Parse()

	// call tabs.SplitDiffPrint, which calls SplitDiff and prints results
	_, err := split.SplitDiffPrint(*tab, *shares, *diff, *isV)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		os.Exit(1)
	}
}

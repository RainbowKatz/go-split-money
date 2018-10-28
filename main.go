package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/RainbowKatz/go-split-money/tabs"
)

func main() {
	// tabs.SplitDiff command-line flags (defaults to tab=10, shares=2, diff=1)
	tab := flag.Float64("tab", 10.00, "Total amount to split into shares")
	shares := flag.Int("shares", 2, "Number of shares to split tab into")
	diff := flag.Float64("diff", 1.00, "Difference between consecutive shares")
	flag.Parse()

	// call tabs.SplitDiffPrint, which calls SplitDiff and prints results
	_, err := tabs.SplitDiffPrint(*tab, *shares, *diff)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		os.Exit(1)
	}
}

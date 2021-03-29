/*
	Package Tabs provides methods of splitting a typical tab between multiple parties.
*/
package split

import (
	"fmt"
	"math"
)

const stdErr float64 = 0.009

/* SplitDiff splits a tab into a given number of shares, with a diff between each consecutive share.
An error is returned along with approximate resultsif rounding error(s) occurred. */
func SplitDiff(tab float64, shares int, diff float64) ([]float64, error) {
	// output
	var splitTabs []float64
	var shareTotal float64

	// validate
	if err := validateSplitDiff(tab, shares, diff); err != nil {
		return nil, err
	}

	// calculate first/lowest share
	var lowestShare float64
	if shares%2 == 0 {
		lowestShare = tab/float64(shares) - (diff/float64(2))*(2*float64(shares/2)-1)
	} else {
		lowestShare = tab/float64(shares) - float64(shares/2)*diff
	}

	for w, currentShare := 0, 0.; w < shares; w++ {
		currentShare = lowestShare + float64(w)*diff

		// rount to nearest penny
		currentShare = math.Round(currentShare/0.01) * 0.01
		splitTabs = append(splitTabs, currentShare)
		shareTotal += currentShare
	}

	// if split not possible due to rounding to nearest cent, return results anyway along with error
	if math.Abs(shareTotal-tab) >= stdErr {
		return splitTabs, fmt.Errorf("Results not exact due to rounding to nearest cent")
	}

	// return output with no error
	return splitTabs, nil
}

// SplitDiffPrint calls SplitDiff and then prints the results
func SplitDiffPrint(tab float64, shares int, diff float64, isVerbose bool) ([]float64, error) {
	results, err := SplitDiff(tab, shares, diff)
	if err != nil && results == nil {
		return nil, err
	}

	//display results
	var indent string
	if isVerbose {
		fmt.Printf("\nSplitDiff(tab=%f, shares=%d, diff=%f) >>\n", tab, shares, diff)
		indent = "  "
	}
	for i, share := range results {
		fmt.Printf("%sShare %d: $%.2f\n", indent, i+1, share)
	}

	// return results
	return results, err
}

// validateSplitDiff performs validation for the SplitDiff function
func validateSplitDiff(tab float64, shares int, diff float64) error {
	// validate total, must be positive
	if tab <= 0. {
		return fmt.Errorf("tab must be positive")
	}

	// validate shares is at least 2
	if shares < 2 {
		return fmt.Errorf("shares must be at least 2")
	}

	// validate diff, must be non-negative
	if diff < 0. {
		return fmt.Errorf("diff must be non-negative")
	}

	// calculate the minLargestShare, minimum largest share to satisfy the diff and number of shares
	//   i.e. 4 shares and diff of 2 means minLargestShare of 6 >> 0, 2, 4, 6
	minLargestShare := float64(shares-1) * diff

	// calculate minTab, that is the minimum tab to perform the split given the diffs
	minTab := float64(shares/2) * minLargestShare
	// 1st term is 0, last is minLargestShare, formula @ http://www.mathwords.com/a/arithmetic_series.htm

	// validate input tab is not less than minTab
	if tab < minTab {
		return fmt.Errorf("cannot split tab %f into %d tabs with a diff of %f between each tab", tab, shares, diff)
	}

	// validated
	return nil
}

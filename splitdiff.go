package split

import (
	"fmt"
	"math"
)

const stdErr float64 = 0.009

/* SplitDiff splits a total into a given number of shares, with a diff between each consecutive
share.  An error is returned along with approximate results if rounding error(s) occurred. */
func SplitDiff(total float64, shares int, diff float64) ([]float64, error) {
	// output
	var splitTabs []float64
	var shareTotal float64

	// validate
	if err := validateSplitDiff(total, shares, diff); err != nil {
		return nil, err
	}

	// calculate first/lowest share
	var lowestShare float64
	if shares%2 == 0 {
		lowestShare = total/float64(shares) - (diff/float64(2))*(2*float64(shares/2)-1)
	} else {
		lowestShare = total/float64(shares) - float64(shares/2)*diff
	}

	for w, currentShare := 0, 0.; w < shares; w++ {
		currentShare = lowestShare + float64(w)*diff

		// rount to nearest penny
		currentShare = math.Round(currentShare/0.01) * 0.01
		splitTabs = append(splitTabs, currentShare)
		shareTotal += currentShare
	}

	// return approximate results along with error is exact split not possible
	if math.Abs(shareTotal-total) >= stdErr {
		return splitTabs, fmt.Errorf("Results not exact due to rounding to nearest cent")
	}

	// return output with no error
	return splitTabs, nil
}

// SplitDiffPrint is a wrapper for SplitDiff which also prints the results
func SplitDiffPrint(total float64, shares int, diff float64, isVerbose bool) ([]float64, error) {
	results, err := SplitDiff(total, shares, diff)
	if err != nil && results == nil {
		return nil, err
	}

	//display results
	var indent string
	if isVerbose {
		fmt.Printf("\nSplitDiff(total=%f, shares=%d, diff=%f) >>\n", total, shares, diff)
		indent = "  "
	}
	for i, share := range results {
		fmt.Printf("%sShare %d: $%.2f\n", indent, i+1, share)
	}

	// return results
	return results, err
}

// validateSplitDiff performs validation for the SplitDiff function
func validateSplitDiff(total float64, shares int, diff float64) error {
	// validate total, must be positive
	if total <= 0. {
		return fmt.Errorf("total must be positive")
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

	// calculate minTab, that is the minimum total to perform the split given the diffs
	minTab := float64(shares/2) * minLargestShare
	// 1st term is 0, last is minLargestShare, formula @ http://www.mathwords.com/a/arithmetic_series.htm

	// validate input total is not less than minTab
	if total < minTab {
		return fmt.Errorf("cannot split total %f into %d tabs with a diff of %f between each total", total, shares, diff)
	}

	// validated
	return nil
}

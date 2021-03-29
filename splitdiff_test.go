package split_test

import (
	"math"
	"testing"

	"github.com/rainbowkatz/split"
)

const stdErr float64 = 0.01

var (
	// these test cases will complete sucessfully, though some will have rounding errors
	positiveTestCases = []struct {
		tab             float64   // tab total to split
		shares          int       // number of shares to split the tab into
		diff            float64   // difference between consecutive shares
		expectedResults []float64 // expected result shares
		isExpectedErr   bool      // boolean that is true only if error expected
	}{
		{10., 2, 1., []float64{4.5, 5.5}, false},
		{10., 4, 0., []float64{2.5, 2.5, 2.5, 2.5}, false},
		{111., 3, 18., []float64{19., 37., 55.}, false},
		{220., 4, 20., []float64{25., 45., 65., 85.}, false},
		{17., 7, 0.18, []float64{1.89, 2.07, 2.25, 2.43, 2.61, 2.79, 2.97}, true}, // rounding error
		{7., 2, 2.33, []float64{2.34, 4.67}, true},                                // rounding error
		{22., 3, 5., []float64{2.33, 7.33, 12.33}, true},                          // rounding error
		{12., 4, 1.11, []float64{1.34, 2.45, 3.56, 4.67}, true},                   // rounding error
	}

	// these test cases will result in errors for various validation reasons (not rounding errors)
	negativeTestCases = []struct {
		tab    float64 // tab total to split
		shares int     // number of shares to split the tab into
		diff   float64 // difference between consecutive shares
	}{
		{10., 1, 1.},    // min shares
		{10., 0, 1.},    // min shares
		{-2., 4, 20.},   // positive tab
		{0., 4, 20.},    // positive tab
		{20., 4, -0.05}, // non-negative diff
		{10., 2, 11.},   // minTab
		{111., 13, 18.}, // minTab
	}
)

func TestSplitDiffPositiveFlows(t *testing.T) {
	for _, tt := range positiveTestCases {
		actualResults, err := split.SplitDiff(tt.tab, tt.shares, tt.diff)
		if err == nil && tt.isExpectedErr {
			t.Errorf("Error incorrect, got: nil, expected: err")
			break
		}
		if err != nil && !tt.isExpectedErr {
			t.Errorf("Error incorrect, got: err, expected: nil")
			break
		}

		// compare results
		for caseIdx := 0; caseIdx < len(tt.expectedResults); caseIdx++ {
			// wrong if off by a stdErr or more
			if math.Abs(actualResults[caseIdx]-tt.expectedResults[caseIdx]) >= stdErr {
				t.Errorf("Results incorrect, got: %v, expected: %v >> case %d: {%.2f, %d, %.2f}", actualResults, tt.expectedResults, caseIdx, tt.tab, tt.shares, tt.diff)
				break
			}
		}
	}
}

func TestSplitDiffNegativeFlows(t *testing.T) {
	for caseIdx, tt := range negativeTestCases {
		_, err := split.SplitDiff(tt.tab, tt.shares, tt.diff)
		if err == nil {
			t.Errorf("Error incorrect, got: nil, expected: err >> case %d: {%.2f, %d, %.2f}", caseIdx, tt.tab, tt.shares, tt.diff)
			break
		}
	}
}

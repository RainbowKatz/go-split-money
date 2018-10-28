<h1>go-split-money</h1>

An example project that holds a ```go get```able package to split a tab into shares.

<h2>Table of Contents</h2>

- [Package List](#package-list)
    - [tabs](#tabs)


# tabs
## GoDoc
For documentation, see [GoDoc](https://godoc.org/github.com/RainbowKatz/go-split-money/tabs)

## How to Run
The ```main.go``` file is set to run ```SplitDiffPrint```, a convenience func which calls the ```SplitDiff``` function, and then displays the formatted results.

To run with default values (from repo root):
```
$ go run main.go
```
output:
```
SplitDiff(tab=10.000000, shares=2, diff=1.000000) >>
  Share 1: $4.50
  Share 2: $5.50
```

To run with custom values, use command-line flags:

```
$ go run main.go -tab=28.00 -shares=5 -diff=2.00

SplitDiff(tab=28.000000, shares=5, diff=2.000000) >>
  Share 1: $1.60
  Share 2: $3.60
  Share 3: $5.60
  Share 4: $7.60
  Share 5: $9.60
```

Testing a case with rounding error (taken from [test cases](./tabs/splitdiff_test.go))
```
$ go run main.go -tab=7 -shares=2 -diff=2.33

SplitDiff(tab=7.000000, shares=2, diff=2.330000) >>
  Share 1: $2.34
  Share 2: $4.67

Error: Results not exact due to rounding to nearest cent
exit status 1
```

## How to Test
To run [tests](./tabs/splitdiff_test.go) (from repo root):
```
$ go test ./tabs
ok      github.com/RainbowKatz/go-split-money/tabs      0.346s
```

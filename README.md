<h1>split</h1>

A Go example module that splits an amount into equal shares.

<h2>Table of Contents</h2>

- [Docs](#docs)
- [Using](#using)
  - [CLI](#cli)
- [Testing](#testing)

<br/>

## Docs
Docs are [here](https://pkg.go.dev/github.com/rainbowkatz/split)

<br/>

## Using
Import as `github.com/rainbowkatz/split` for use in another module or use CLI for direct use (see below)

<br/>

### CLI
There is a CLI provided as a separate module: 
```bash
go get github.com/rainbowkatz/split/gosplit
```
If your $GOPATH is part of your system path, that's all you need to do.  

<br/>

To run with default values (from repo root):
```
$ gosplit
```
output:
```
SplitDiff(tab=10.000000, shares=2, diff=1.000000) >>
  Share 1: $4.50
  Share 2: $5.50
```

To run with custom values, use command-line flags:

```
$ gosplit -tab=28.00 -shares=5 -diff=2.00

SplitDiff(tab=28.000000, shares=5, diff=2.000000) >>
  Share 1: $1.60
  Share 2: $3.60
  Share 3: $5.60
  Share 4: $7.60
  Share 5: $9.60
```

Testing a case with rounding error (taken from [test cases](./tabs/splitdiff_test.go))
```
$ gosplit -tab=7 -shares=2 -diff=2.33

SplitDiff(tab=7.000000, shares=2, diff=2.330000) >>
  Share 1: $2.34
  Share 2: $4.67

Error: Results not exact due to rounding to nearest cent
exit status 1
```

<br/>

## Testing
To run [tests](./splitdiff_test.go) (from repo root):
```
$ go test .
ok      github.com/RainbowKatz/go-split-money/tabs      0.346s
```

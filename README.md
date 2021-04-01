# EV Charge Time Leaderboard

This application identifies the *n* employees who have used the carging stations the most over a specific period of time.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Output](#output)
- [Testing](#testing)
- [Benchmarking](#benchmarking)

## Installation

1. Install the [Golang](https://golang.org/doc/install) environment
2. Clone the repo
    ```sh
    git clone https://github.com/mtormento/evchargetime.git
    ```
3. Change directory
    ```sh
    cd evchargetime
    ```
4. Build executable
* Unix based
    ```sh
    go build -o bin/evchargetime go/charge_time.go
    ```
* Windows
    ```bat
    go build -o bin\evchargetime.exe go\charge_time.go
    ```

## Usage

```sh
evchargetime -d <datafile> -l <number_of_output_lines>
```

### Data file (-d)
The data file is a text file in this format:
```
employee_id start_time finish_time
PTTG8 1609498637564 1609505453995
JFFO9 1611847305439 1611856467760
LQOB1 1611681652028 1611690814349
PTTG8 1610439349208 1610445506538
ZEAY5 1610712347411 1610722014263
```

### Number of lines (-l)
The program will output *this* number of employees which have used the charging stations the most.  
**The default value is 3**.

## Output
The output will be in this format:
```
PTTG8 3h36m
ZEAY5 2h41m
JFFO9 2h32m
```

## Testing
```sh
go test -timeout 30s -cover github.com/mtormento/evchargetime/go/calc
go test -timeout 30s -cover github.com/mtormento/evchargetime/go/fmt
```

## Benchmarking
This benchmark will run the program on about 1 millions rows.
```sh
go test -bench . github.com/mtormento/evchargetime/go/calc
```

### Result on test system
This benchmark run has executed the program 4 times in a row on about 1 millions rows on a test system.
```
goos: linux
goarch: amd64
pkg: github.com/mtormento/evchargetime/go/calc
cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
BenchmarkBuildOrderedChargeInfoArray-12    	       4	 309407224 ns/op
PASS
ok  	github.com/mtormento/evchargetime/go/calc	2.399s
```
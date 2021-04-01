package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mtormento/evchargetime/go/calc"
	chargeTimeFmt "github.com/mtormento/evchargetime/go/fmt"
)

var (
	datafile = ""
	lines    = 0
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.StringVar(&datafile, "d", "", "Charging times data file")
	fs.IntVar(&lines, "l", 3, "Number of lines to output")
	fs.Parse(os.Args[1:])
}

func main() {

	if datafile == "" {
		log.Fatal("No datafile specified")
	}

	if lines <= 0 {
		log.Fatal("Number of lines must be > 0")
	}

	// Build ordered array of charge info structures, that contains employee id and charge time
	chargeInfoArray, err := calc.BuildSortedChargeInfoArray(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Print only <lines> number of lines
	for idx, chargeInfo := range chargeInfoArray {
		if idx >= lines {
			break
		}
		fmt.Printf("%s %s\n", chargeInfo.EmployeeId, chargeTimeFmt.FmtDuration(chargeInfo.Elapsed))
	}
}

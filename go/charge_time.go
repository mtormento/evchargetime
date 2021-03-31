package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"daze.com/evchargetime/go/calc"
	chargeTimeFmt "daze.com/evchargetime/go/fmt"
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

	arrayOfPlates, err := calc.BuildOrderedChargeInfoArray(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}

	for idx, chargeInfo := range arrayOfPlates {
		if idx >= lines {
			break
		}
		fmt.Printf("%s %s\n", chargeInfo.Plate, chargeTimeFmt.FmtDuration(chargeInfo.Elapsed))
	}
}

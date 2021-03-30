package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type plateInfo struct {
	plate   string
	elapsed time.Duration
}

type lessFunc func(p1, p2 *plateInfo) bool

type multiSorter struct {
	plateInfos []plateInfo
	less       []lessFunc
}

func (ms *multiSorter) Sort(plateInfos []plateInfo) {
	ms.plateInfos = plateInfos
	sort.Sort(ms)
}

func OrderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.plateInfos)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.plateInfos[i], ms.plateInfos[j] = ms.plateInfos[j], ms.plateInfos[i]
}

func (ms *multiSorter) Less(i, j int) bool {
	p, q := &ms.plateInfos[i], &ms.plateInfos[j]
	var k int
	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}
	return ms.less[k](p, q)
}

func main() {
	if len(os.Args) == 2 {
		filename := os.Args[1]

		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		platesMap := make(map[string]plateInfo)

		firstLine := true
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if firstLine {
				firstLine = false
				continue
			}
			line := scanner.Text()
			tokens := strings.Split(line, " ")
			if len(tokens) != 3 {
				panic("invalid data")
			} else {
				plate := tokens[0]
				start, err := strconv.ParseInt(tokens[1], 10, 64)
				if err != nil {
					panic(err.Error())
				}
				end, err := strconv.ParseInt(tokens[2], 10, 64)
				if err != nil {
					panic(err.Error())
				}
				elapsed := time.Duration(end-start) * time.Millisecond
				pi, ok := platesMap[plate]
				if ok {
					pi.elapsed += elapsed
					platesMap[plate] = pi
				} else {
					platesMap[plate] = plateInfo{plate, elapsed}
				}
			}
		}
		var arrayOfPlates []plateInfo
		for _, element := range platesMap {
			arrayOfPlates = append(arrayOfPlates, element)
		}

		byPlate := func(c1, c2 *plateInfo) bool {
			return c1.plate < c2.plate
		}
		byElapsed := func(c1, c2 *plateInfo) bool {
			return c1.elapsed > c2.elapsed
		}

		OrderedBy(byElapsed, byPlate).Sort(arrayOfPlates)
		for _, plateInfo := range arrayOfPlates {
			fmt.Printf("%s %s\n", plateInfo.plate, fmtDuration(plateInfo.elapsed))
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
	}
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%dh%02dm", h, m)
}

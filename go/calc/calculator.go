package calc

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ChargeInfo struct {
	Plate   string
	Elapsed time.Duration
}

func BuildOrderedChargeInfoArray(filename string) ([]ChargeInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	plateToChargeInfoMap := make(map[string]ChargeInfo)

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
			return nil, errors.New("invalid data")
		} else {
			plate, elapsed, err := decodeTokens(tokens)
			if err != nil {
				return nil, err
			}
			ci, ok := plateToChargeInfoMap[plate]
			if ok {
				ci.Elapsed += elapsed
				plateToChargeInfoMap[plate] = ci
			} else {
				plateToChargeInfoMap[plate] = ChargeInfo{plate, elapsed}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	arrayOfPlates := make([]ChargeInfo, len(plateToChargeInfoMap))
	idx := 0
	for _, element := range plateToChargeInfoMap {
		arrayOfPlates[idx] = element
		idx++
	}

	orderByChargeTimeAndPlate(arrayOfPlates)

	return arrayOfPlates, nil
}

func orderByChargeTimeAndPlate(arrayOfPlates []ChargeInfo) {
	sort.Slice(arrayOfPlates, func(i, j int) bool {
		if arrayOfPlates[i].Elapsed > arrayOfPlates[j].Elapsed {
			return true
		}
		if arrayOfPlates[i].Elapsed < arrayOfPlates[j].Elapsed {
			return false
		}
		return arrayOfPlates[i].Plate < arrayOfPlates[j].Plate
	})
}

func decodeTokens(tokens []string) (string, time.Duration, error) {
	plate := tokens[0]
	start, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return "", -1, err
	}
	end, err := strconv.ParseInt(tokens[2], 10, 64)
	if err != nil {
		return "", -1, err
	}
	elapsed := time.Duration(end-start) * time.Millisecond
	if elapsed < 0 {
		return "", -1, errors.New("elapsed < 0")
	}
	return plate, elapsed, nil
}

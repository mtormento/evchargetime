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
	EmployeeId string
	Elapsed    time.Duration
}

// BuildSortedChargeInfoArray builds an change info array, sorted by duration desc and employee asc
func BuildSortedChargeInfoArray(filename string) ([]ChargeInfo, error) {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	employeeIdToChargeInfoMap := make(map[string]ChargeInfo)

	firstLine := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Skip header line
		if firstLine {
			firstLine = false
			continue
		}
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		if len(tokens) != 3 {
			return nil, errors.New("invalid data")
		} else {
			employeeId, elapsed, err := decodeTokens(tokens)
			if err != nil {
				return nil, err
			}
			ci, ok := employeeIdToChargeInfoMap[employeeId]
			if ok {
				// If employee id already exists add duration
				ci.Elapsed += elapsed
				employeeIdToChargeInfoMap[employeeId] = ci
			} else {
				employeeIdToChargeInfoMap[employeeId] = ChargeInfo{employeeId, elapsed}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	// Dump charge infos to array
	chargeInfoArray := make([]ChargeInfo, len(employeeIdToChargeInfoMap))
	idx := 0
	for _, element := range employeeIdToChargeInfoMap {
		chargeInfoArray[idx] = element
		idx++
	}

	// Sort charge infos
	sortByChargeTimeAndEmployeeId(chargeInfoArray)

	return chargeInfoArray, nil
}

// Sort function
func sortByChargeTimeAndEmployeeId(chargeInfoArray []ChargeInfo) {
	sort.Slice(chargeInfoArray, func(i, j int) bool {
		if chargeInfoArray[i].Elapsed > chargeInfoArray[j].Elapsed {
			return true
		}
		if chargeInfoArray[i].Elapsed < chargeInfoArray[j].Elapsed {
			return false
		}
		// If duration is equal sort by employee id
		return chargeInfoArray[i].EmployeeId < chargeInfoArray[j].EmployeeId
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

package calc

import (
	"testing"
	"time"

	"github.com/mtormento/evchargetime/go/fmt"
	"github.com/stretchr/testify/assert"
)

func BenchmarkBuildSortedChargeInfoArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuildSortedChargeInfoArray("../../data/charging_data_big.txt")
	}
}

func TestBuildSortedChargeInfoArray(t *testing.T) {
	assert := assert.New(t)

	chargeInfoArray, err := BuildSortedChargeInfoArray("../../data/charging_data.txt")
	assert.Nil(err)
	assert.NotNil(chargeInfoArray)
	assert.Equal(15, len(chargeInfoArray))

	chargeInfoArray, err = BuildSortedChargeInfoArray("../../data/charging_data_example.txt")
	assert.Nil(err)
	assert.Equal("PTTG8", chargeInfoArray[0].EmployeeId)
	assert.Equal("3h36m", fmt.FmtDuration(chargeInfoArray[0].Elapsed))
	assert.Equal("ZEAY5", chargeInfoArray[1].EmployeeId)
	assert.Equal("2h41m", fmt.FmtDuration(chargeInfoArray[1].Elapsed))
	assert.Equal("JFFO9", chargeInfoArray[2].EmployeeId)
	assert.Equal("2h32m", fmt.FmtDuration(chargeInfoArray[2].Elapsed))

	_, err = BuildSortedChargeInfoArray("../../data/non_existent_file.txt")
	assert.NotNil(err)
	assert.Equal("open ../../data/non_existent_file.txt: no such file or directory", err.Error())

	_, err = BuildSortedChargeInfoArray("../../data/charging_data_error_1.txt")
	assert.NotNil(err)
	assert.Equal("elapsed < 0", err.Error())

	_, err = BuildSortedChargeInfoArray("../../data/charging_data_error_2.txt")
	assert.NotNil(err)
	assert.Equal("elapsed < 0", err.Error())

	_, err = BuildSortedChargeInfoArray("../../data/charging_data_error_3.txt")
	assert.NotNil(err)
	assert.Equal("strconv.ParseInt: parsing \"not_valid\": invalid syntax", err.Error())

	_, err = BuildSortedChargeInfoArray("../../data/charging_data_error_4.txt")
	assert.NotNil(err)
	assert.Equal("strconv.ParseInt: parsing \"not_valid\": invalid syntax", err.Error())

	_, err = BuildSortedChargeInfoArray("../../data/charging_data_error_5.txt")
	assert.NotNil(err)
	assert.Equal("invalid data", err.Error())
}

func TestSortByChargeTimeAndEmployeeId(t *testing.T) {
	assert := assert.New(t)

	ChargeInfoArray := make([]ChargeInfo, 4)
	ChargeInfoArray[0] = ChargeInfo{"BBBBB", 10 * time.Second}
	ChargeInfoArray[1] = ChargeInfo{"CCCCC", 10 * time.Second}
	ChargeInfoArray[2] = ChargeInfo{"AAAAA", 10 * time.Second}
	ChargeInfoArray[3] = ChargeInfo{"DDDDD", 10 * time.Second}

	sortByChargeTimeAndEmployeeId(ChargeInfoArray)
	assert.Equal("AAAAA", ChargeInfoArray[0].EmployeeId)
	assert.Equal(10*time.Second, ChargeInfoArray[0].Elapsed)
	assert.Equal("BBBBB", ChargeInfoArray[1].EmployeeId)
	assert.Equal(10*time.Second, ChargeInfoArray[1].Elapsed)
	assert.Equal("CCCCC", ChargeInfoArray[2].EmployeeId)
	assert.Equal(10*time.Second, ChargeInfoArray[2].Elapsed)
	assert.Equal("DDDDD", ChargeInfoArray[3].EmployeeId)
	assert.Equal(10*time.Second, ChargeInfoArray[3].Elapsed)

	ChargeInfoArray[0] = ChargeInfo{"BBBBB", 40 * time.Second}
	ChargeInfoArray[1] = ChargeInfo{"CCCCC", 20 * time.Second}
	ChargeInfoArray[2] = ChargeInfo{"AAAAA", 30 * time.Second}
	ChargeInfoArray[3] = ChargeInfo{"DDDDD", 10 * time.Second}

	sortByChargeTimeAndEmployeeId(ChargeInfoArray)
	assert.Equal("BBBBB", ChargeInfoArray[0].EmployeeId)
	assert.Equal(40*time.Second, ChargeInfoArray[0].Elapsed)
	assert.Equal("AAAAA", ChargeInfoArray[1].EmployeeId)
	assert.Equal(30*time.Second, ChargeInfoArray[1].Elapsed)
	assert.Equal("CCCCC", ChargeInfoArray[2].EmployeeId)
	assert.Equal(20*time.Second, ChargeInfoArray[2].Elapsed)
	assert.Equal("DDDDD", ChargeInfoArray[3].EmployeeId)
	assert.Equal(10*time.Second, ChargeInfoArray[3].Elapsed)
}

func TestDecodeTokens(t *testing.T) {
	assert := assert.New(t)

	tokens := make([]string, 3)
	tokens[0] = "AAAAA"
	tokens[1] = "10"
	tokens[2] = "200"

	Plate, Elapsed, err := decodeTokens(tokens)
	assert.Nil(err)
	assert.Equal("AAAAA", Plate)
	assert.Equal(190*time.Millisecond, Elapsed)

	tokens[0] = "AAAAA"
	tokens[1] = "not_valid"
	tokens[2] = "200"

	_, _, err = decodeTokens(tokens)
	assert.NotNil(err)

	tokens[0] = "AAAAA"
	tokens[1] = "10"
	tokens[2] = "not_valid"

	_, _, err = decodeTokens(tokens)
	assert.NotNil(err)

	tokens[0] = "AAAAA"
	tokens[1] = "10"
	tokens[2] = "1"

	_, _, err = decodeTokens(tokens)
	assert.NotNil(err)
	assert.Equal("elapsed < 0", err.Error())
}

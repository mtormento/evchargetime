package fmt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFmtDuration(t *testing.T) {
	assert := assert.New(t)

	duration := time.Millisecond * 10000
	formattedDuration := FmtDuration(duration)
	assert.Equal("0h00m", formattedDuration)

	duration = time.Second * 60
	formattedDuration = FmtDuration(duration)
	assert.Equal("0h01m", formattedDuration)

	duration = time.Minute * 60
	formattedDuration = FmtDuration(duration)
	assert.Equal("1h00m", formattedDuration)

	duration = time.Hour * 1000
	formattedDuration = FmtDuration(duration)
	assert.Equal("1000h00m", formattedDuration)

	duration = -time.Hour * 1000
	formattedDuration = FmtDuration(duration)
	assert.Equal("0h00m", formattedDuration)
}

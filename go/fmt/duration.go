package fmt

import (
	"fmt"
	"time"
)

// FmtDuration formats a duration in 0h00m format
func FmtDuration(d time.Duration) string {
	if d < 0 {
		return "0h00m"
	}
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%dh%02dm", h, m)
}

package autoupdate

import (
	"fmt"
	"time"
)

//DutyUpdate events for duty and email sending
func DutyUpdate() {
	tick := time.Tick(24 * time.Hour)
	for range tick {
		fmt.Println("Test")
	}
}

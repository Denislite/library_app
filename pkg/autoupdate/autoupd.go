package autoupdate

import (
	"fmt"
	"github.com/Denislite/library_app/pkg/service"
	"time"
)

func UserChecker() {
	tick := time.Tick(24 * time.Hour)
	for range tick {
		err := service.UpdateDuty()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Upd")
	}
}

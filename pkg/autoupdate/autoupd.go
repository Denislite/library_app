package autoupdate

import (
	"fmt"
	"github.com/Denislite/library_app/pkg/mail"
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

		userBooks, err := service.GetUsersWithDuty()
		if err != nil {
			fmt.Println(err)
		}

		for _, userBook := range userBooks {
			mail.SendEmail(userBook.Email, userBook.DutyCount)
		}
		fmt.Println("Upd")
	}
}

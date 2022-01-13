package autoupdate

import (
	"github.com/Denislite/library_app/pkg/mail"
	"github.com/Denislite/library_app/pkg/service"
	"log"
	"time"
)

func UserChecker() {
	tick := time.Tick(24 * time.Hour)
	for range tick {
		err := service.UpdateDuty()
		if err != nil {
			log.Print(err)
		}

		userBooks, err := service.GetUsersWithDuty()
		if err != nil {
			log.Print(err)
		}

		for _, userBook := range userBooks {
			mail.SendEmail(userBook.Email, userBook.DutyCount)
		}
	}
}

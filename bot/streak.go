package bot

import (
	"familyFormUi/config"
	"time"

	tele "gopkg.in/telebot.v3"
)

func CreateStreak(memberID int, habit string, days int) *Streak {
	return &Streak{
		MemberID:   memberID,
		Habit:      habit,
		StartDate:  time.Now(),
		Days:       make([]int, days),
		TotalDays:  days,
		CurrentDay: 0,
	}
}

func Test() {
	config.B.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("Hello world!")
	})

	config.B.Handle(tele.OnQuery, func(c tele.Context) error {
		return c.Send("Button clicked!")
	})
	config.B.Start()
}

// Make an account

package api

import (
	"context"
	"familyFormUi/config"
	"fmt"
	"log"
	"os"
	"strconv"

	tele "gopkg.in/telebot.v3"
)

func (h *Habit) Create() (err error) {
	key := fmt.Sprintf("habitMember:%d", h.TeleID)
	err = config.Rdb.HSet(context.Background(), key, h).Err()
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("%v, %v", h, *h.TeleID)
	botID, err := strconv.Atoi(os.Getenv("TestingBotID"))
	if err != nil {
		log.Println("err:", err)
		return
	}
	t, err := config.B.Send(tele.ChatID(botID), msg)
	if err != nil {
		log.Println("err:", err)
		return
	}
	log.Println(t)
	return
}

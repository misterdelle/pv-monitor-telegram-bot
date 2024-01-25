package main

import (
	"fmt"
	"pv-monitor-telegram-bot/pkg/data"

	"github.com/yanzay/tbot/v2"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	msg := "Questo è un BOT di misterdelle"
	a.client.SendMessage(m.Chat.ID, msg)
	a.menuHandler(m)
}

// Handle the /menuHandler command here
func (a *application) menuHandler(m *tbot.Message) {
	buttons := makeButtons()
	choosedOption := tbot.OptInlineKeyboardMarkup(buttons)
	a.client.SendMessage(m.Chat.ID, "Scegli un'opzione:", choosedOption)
}

// Handle button presses here
func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	userChoice := cq.Data

	var msg string

	if userChoice == "BatterySOC" {
		b, _ := a.DB.GetBatterySOC()
		batterySOC := b.(float32)

		msg = fmt.Sprintf("Battery SOC: %.2f", batterySOC)
	} else if userChoice == "StationData" {
		lsd, _ := a.DB.GetLastStationData()

		lastStationData := lsd.(data.Station)

		msg = lastStationData.String()
	}

	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, msg)
}

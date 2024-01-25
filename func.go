package main

import (
	"github.com/yanzay/tbot/v2"
)

func makeButtons() *tbot.InlineKeyboardMarkup {
	// Create buttons with visible Text and CallbackData as a value.
	btnPresente := tbot.InlineKeyboardButton{
		Text:         "🔋 BatterySOC",
		CallbackData: "BatterySOC",
	}
	btnMalattia := tbot.InlineKeyboardButton{
		Text:         "⚡ Station Data",
		CallbackData: "StationData",
	}
	buttons := []tbot.InlineKeyboardButton{btnPresente, btnMalattia}

	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			buttons,
		},
	}
}

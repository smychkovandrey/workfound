package main

import (
	"WorkFound/internal/fileworker"
	"WorkFound/internal/models"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	debugmode := false
	var token string
	if len(os.Args) > 0 {
		for i := 1; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "-debug":
				debugmode = true
			case "-t":
				i++
				if i >= len(os.Args) {
					panic("token not found")
				}
				token = os.Args[i]
			}
		}
	}

	se := models.CreateSeacrhEngine()

	if token != "" {
		bot, err := tgbotapi.NewBotAPI(token)
		if err != nil {
			panic(err)
		}
		updateConfig := tgbotapi.NewUpdate(0)
		updateConfig.Timeout = 30

		// Start polling Telegram for updates.
		updates := bot.GetUpdatesChan(updateConfig)

		for update := range updates {
			if update.Message == nil { //смотрим только сообщения
				continue
			}

			if !update.Message.IsCommand() { // ignore any non-command Messages
				continue
			}

			// Extract the command from the Message.
			switch update.Message.Command() {
			case "help":
				sendMessage(bot, update.Message.Chat.ID, "Я знаю только /vacancies  /add_filter  /del_filter /show_filter")
			case "vacancies":
				res := se.GetResults()
				for result := range res {
					if result.Error != nil {
						if debugmode {
							print(result.Name, "->")
							println(result.Error)
						}
						sendMessage(bot, update.Message.Chat.ID, "Ошибка в ", result.Name)
						continue
					}
					sb := make_Result_string(&result)
					sendMessage(bot, update.Message.Chat.ID, sb.String())
				}
			case "add_filter", "del_filter":
				sendMessage(bot, update.Message.Chat.ID, "Пока не работает")
			case "show_filter":
				stopwords, err := fileworker.GetStopWords()
				if err != nil {
					if debugmode {
						println(err)
					}
					sendMessage(bot, update.Message.Chat.ID, "Не могу получить фильтры")
				} else {
					sb := strings.Builder{}
					for _, stopword := range stopwords {
						sb.WriteString(stopword)
						sb.WriteString("\n")
					}
					sendMessage(bot, update.Message.Chat.ID, sb.String())
				}
			default:
				sendMessage(bot, update.Message.Chat.ID, "Неизвестная команда")
			}

		}
	} else {
		res := se.GetResults()
		for result := range res {
			if result.Error != nil {
				print(result.Name, "->")
				println(result.Error)
				continue
			}
			println(make_Result_string(&result).String())
			println()
		}
	}
}

func make_Result_string(result *models.Results) *strings.Builder {
	sb := strings.Builder{}
	sb.WriteString(result.Name)
	sb.WriteString(" (")
	sb.WriteString(strconv.Itoa(len(result.Jobs)))
	sb.WriteString(")")
	sb.WriteString("\n")
	for _, j := range result.Jobs {
		sb.WriteString(j.Name)
		sb.WriteString("->")
		sb.WriteString(j.Url)
		sb.WriteString("\n")
	}
	return &sb
}

func sendMessage(bot *tgbotapi.BotAPI, chatId int64, text ...string) {
	msg := tgbotapi.NewMessage(chatId, strings.Join(text, ""))
	if _, err := bot.Send(msg); err != nil {
		println(err)
	}
}

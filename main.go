package main

import (
	"fmt"
	"log"
	"os"
	"pv-monitor-telegram-bot/pkg/db"
	"pv-monitor-telegram-bot/pkg/repository"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/yanzay/tbot/v2"
)

type score struct {
	wins, draws, losses uint
}

type application struct {
	client *tbot.Client
	score
	DSN         string
	DB          repository.DatabaseRepository
	WebHook     string
	WebHookPort string
}

var (
	app   application
	bot   *tbot.Server
	token string
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TELEGRAM_TOKEN")
	log.Println("Token: ", token)

	app.DSN = os.Getenv("DSN")
	app.WebHook = os.Getenv("WEBHOOK")
	app.WebHookPort = os.Getenv("WEBHOOK_PORT")
}

func main() {
	connRDBMS, err := app.connectToDB()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error %s", err))
	}
	defer connRDBMS.Close()

	app.DB = &db.PostgresDBRepo{DB: connRDBMS}

	if app.WebHook != "" {
		bot = tbot.New(token, tbot.WithWebhook(app.WebHook, app.WebHookPort))
	} else {
		bot = tbot.New(token)
	}

	app.client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("/menu", app.menuHandler)
	bot.HandleCallback(app.callbackHandler)
	log.Fatal(bot.Start())
}

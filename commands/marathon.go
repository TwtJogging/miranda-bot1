package commands

import (
	"log"
	"time"

	tg "gopkg.in/telegram-bot-api.v4"
)

// Rules send rules
func (c Command) Marathon() {
	msg := tg.NewMessage(c.Message.Chat.ID, "<b>Senarai acara terkini larian marathon di Malaysia bagi tahun 2021.</b>-3 minit bacaan\n\nBaca: <a href='https://twitter.com/twt_jogging'>Read here!</a>\n\n1: Kedah Marathon -<b>ditangguhkan</b>\n\n1: Perlis Marathon -ditangguhkan\n\n1: Penang Marathon -<b>ditangguhkan</b>\n\n1: Kuala Kangsar Marathon -<b>ditangguhkan</b>")
	msg.ParseMode = "HTML"
	msg.ReplyToMessageID = c.Message.MessageID

	r, err := c.Bot.Send(msg)

	if err != nil {
		log.Println(err)

		return
	}

	go func() {
		log.Printf("Deleting message %d in 20 seconds...", r.Chat.ID)
		time.Sleep(20 * time.Second)

		// Delete !rules
		twitter := tg.DeleteMessageConfig{
			ChatID:    c.Message.Chat.ID,
			MessageID: c.Message.MessageID,
		}
		c.Bot.DeleteMessage(twitter)

		// Delete Rules after a few second
		reply := tg.DeleteMessageConfig{
			ChatID:    r.Chat.ID,
			MessageID: r.MessageID,
		}
		c.Bot.DeleteMessage(reply)
	}()

}
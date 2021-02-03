package commands

import (
	"log"
	"time"

	tg "gopkg.in/telegram-bot-api.v4"
)

// Rules send rules
func (c Command) Twitter() {
	msg := tg.NewMessage(c.Message.Chat.ID, "<b>The TwtJogging Rules</b>\n\nBaca: <a href='https://twitter.com/twt_jogging'>Read here!</a>")
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

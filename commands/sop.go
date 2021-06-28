package commands

import (
	"log"
	"time"

	tg "gopkg.in/telegram-bot-api.v4"
)

// Rules send rules
func (c Command) Sop() {
	msg := tg.NewMessage(c.Message.Chat.ID, "<b>SOP terbaru dalam Sukan Dan Rekreasi yang telah dikemaskini.</b>-3 minit bacaan\n\nSentiasa ikuti SOP yang telah ditetapkan oleh MKN <a href='https://t.me/MKNRasmi'>Read here!</a>\n\n1: Aktiviti jogging dibenarkan secara individu atau solo run, hanya di dalam kawasan kejiranan sahaja.<b>#KITAJAGAKITA</b>")
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

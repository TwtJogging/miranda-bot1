package commands

import (
	"log"
	"time"

	tg "gopkg.in/telegram-bot-api.v4"
)

// Rules send rules
func (c Command) Order() {
	msg := tg.NewMessage(c.Message.Chat.ID, "<b>Senarai order baju running di TwtJogging.</b>-3 minit bacaan\n\nBaca: <a href='https://t.me/pacrbot'>Read here!</a>\n\n1: @amirasyrafwoi -<b>Size M</b>\n\n1: @afqazt -<b>Size M</b>\n\n2: @KerolWan -<b>Saiz: L</b>\n\n3: Irsyaduddin Ismail -<b>Size M</b>\n\n4: @dzu_im -<b>Update payment</b>\n\n5: @kayyyyyyy -<b>Update payment</b>")
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

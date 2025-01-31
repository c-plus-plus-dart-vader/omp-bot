package card

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func (p *DummyCardCommander) Get(inMsg  *tgbotapi.Message){
	args := inMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if idx < 0 || err != nil {
		fmt.Printf("Invalid idx %s", args)
		return
	}

	card, err := p.cardService.Get(uint64(idx))
	if err != nil {
		fmt.Printf("fail to get card with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inMsg.Chat.ID,
		card.String(),
	)

	p.bot.Send(msg)
}

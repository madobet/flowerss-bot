package handler

import (
	"fmt"

	"go.uber.org/zap"
	tb "gopkg.in/telebot.v3"

	"github.com/indes/flowerss-bot/internal/model"
)

type Start struct {
}

func NewStart() *Start {
	return &Start{}
}

func (s *Start) Command() string {
	return "/start"
}

func (s *Start) Description() string {
	return "初次见面"
}

func (s *Start) Handle(ctx tb.Context) error {
	user, _ := model.FindOrCreateUserByTelegramID(ctx.Chat().ID)
	zap.S().Infof("/start user_id: %d telegram_id: %d", user.ID, user.TelegramID)
	return ctx.Send(fmt.Sprintf("我明白了，外派工作的「契约」已经拟好了，请过目，即日起就可以为你…欸！名字没有签吗？请让我看一下。唔…名字是甘，甘雨…呼，补上了。话，话说回来，工作具体是指哪些呢？"))
}

func (s *Start) Middlewares() []tb.MiddlewareFunc {
	return nil
}

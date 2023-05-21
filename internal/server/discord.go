package server

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kunghim/go-kratos-discord/internal/conf"
	"github.com/kunghim/go-kratos-discord/internal/service"
)

type DiscordServer struct {
	discordService *service.DiscordService
	session        *discordgo.Session
}

func NewDiscordService(c *conf.Server, discordService *service.DiscordService) *DiscordServer {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", c.Discord.BotToken))
	if err != nil {
		panic(err)
	}
	session.AddHandler(discordService.CreateMessage)
	session.AddHandler(discordService.UpdateMessage)
	// 注册新的消息事件
	session.AddHandler(discordService.Something)
	return &DiscordServer{
		discordService: discordService,
		session:        session,
	}
}

func (s *DiscordServer) Start(ctx context.Context) error {
	err := s.session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}
	fmt.Println("Discord session opened")
	return nil
}

func (s *DiscordServer) Stop(ctx context.Context) error {
	err := s.session.Close()
	if err != nil {
		fmt.Println("error closing connection,", err)
	}
	fmt.Println("Discord session closed")
	return nil
}

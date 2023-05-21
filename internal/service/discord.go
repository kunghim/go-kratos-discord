package service

import (
	"github.com/bwmarrin/discordgo"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gookit/goutil/jsonutil"
	"github.com/kunghim/go-kratos-discord/internal/conf"
)

type DiscordService struct {
	log    *log.Helper
	server *conf.Server
}

func NewDiscordService(logger log.Logger, server *conf.Server) *DiscordService {
	return &DiscordService{
		log:    log.NewHelper(log.With(logger, "module", "service/discord")),
		server: server,
	}
}

func (s *DiscordService) CreateMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	s.log.Infof("CreateMessage\n")
	pretty, _ := jsonutil.Pretty(session)
	s.log.Infof("Discord session: %s\n", pretty)
	msg, _ := jsonutil.Pretty(message)
	s.log.Infof("Discord message: %s\n", msg)
}

func (s *DiscordService) UpdateMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	s.log.Infof("CreateMessage\n")
	pretty, _ := jsonutil.Pretty(session)
	s.log.Infof("Discord session: %s\n", pretty)
	msg, _ := jsonutil.Pretty(message)
	s.log.Infof("Discord message: %s\n", msg)
}

func (s *DiscordService) Something(session *discordgo.Session, message *discordgo.MessageDelete) {
	// todo

}

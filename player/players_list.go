package player

import (
	"errors"
	"github.com/Dadard29/go-navis/common"
	"time"
)

type Player struct {
	Username string
	JoinedAsOf string
}

type PlayerList struct {
	accessToken string
	list []Player
}

func NewPlayerList(adminUsername string, token string) PlayerList {
	return PlayerList{
		accessToken: token,
		list: append([]Player{}, Player{
			Username:   adminUsername,
			JoinedAsOf: time.Now().String(),
		}),
	}
}

func (l *PlayerList) AddPlayer(username string, tokenGiven string) error {
	if tokenGiven != l.accessToken {
		return errors.New(common.TOKEN_INVALID)
	}

	l.list = append(l.list, Player{
		Username:   username,
		JoinedAsOf: time.Now().String(),
	})

	return nil
}
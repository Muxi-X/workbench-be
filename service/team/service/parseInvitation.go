package service

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"time"

	"muxi-workbench-team/errno"
	pb "muxi-workbench-team/proto"
	e "muxi-workbench/pkg/err"
)

// ParseInvitation … 解析hash
func (ts *TeamService) ParseInvitation(ctx context.Context, req *pb.ParseInvitationRequest, res *pb.ParseInvitationResponse) error {
	teamID, err := ParseAnInvitation(req.Hash)
	if err != nil {
		return e.ServerErr(errno.ErrLinkExpiration, "链接已过期")
	}
	res.TeamId = teamID
	return nil
}

// ParseAnInvitation parse an invitation
func ParseAnInvitation(hash string) (uint32, error) {
	t := time.Now().Unix()
	sercetKey := viper.GetString("aes.sercet_key")
	words := AesDecrypt(hash, sercetKey)

	teamIDAndDatetime := strings.SplitN(words, " ", 2)

	// string to int64
	datetime, err := strconv.ParseInt(teamIDAndDatetime[1], 10, 64)
	if err != nil {
		return 0, err
	}
	if t >= datetime {
		return 0, errors.New("链接已过期")
	}

	// teamid,_ := strconv.ParseUint(words[:i],10,64)
	teamIDInInt, err := strconv.ParseInt(teamIDAndDatetime[0], 10, 64)
	if err != nil {
		return 0, err
	}

	teamID := uint32(teamIDInInt)
	return teamID, nil
}

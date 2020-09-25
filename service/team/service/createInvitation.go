package service

import (
	"context"
	"github.com/spf13/viper"
	"strconv"
	"time"

	pb "muxi-workbench-team/proto"
)

// CreateInvitation …… 生成邀请码
func (ts *TeamService) CreateInvitation(ctx context.Context, req *pb.CreateInvitationRequest, res *pb.CreateInvitationResponse) error {
	res.Hash = CreateAnInvitation(req.TeamId, req.Expired)
	return nil
}

// CreateAnInvitation create an invitation
func CreateAnInvitation(teamID uint32, expired int64) string {
	t := time.Now().Unix()
	final := t + expired
	words := strconv.FormatUint(uint64(teamID), 10) + " " + strconv.FormatInt(final, 10)
	sercetKey := viper.GetString("aes.sercet_key")
	return AesEncrypt(words, sercetKey)
}

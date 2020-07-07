package model

import (
	"github.com/spf13/viper"
	"strconv"
	"time"
)

func CreateInvitation(team_id uint32, expired int64) string {
	t := time.Now().Unix()
	final := t + expired
	words := strconv.FormatUint(uint64(team_id), 10) + " " + strconv.FormatInt(final, 10)
	sercetKey := viper.GetString("aes.sercet_key")
	return AesEncrypt(words, sercetKey)
}

func ParseInvitation(hash string) (uint32, error) {
	var i int
	t := time.Now().Unix()
	sercetKey := viper.GetString("aes.sercet_key")
	words := AesDecrypt(hash, sercetKey)
	for i = 0; i < len(words); i++ {
		if words[i] == ' ' {
			break
		}
	}
	//string to int64
	datetime, err := strconv.ParseInt(words[i+1:], 10, 64)
	if t >= datetime {
		return 0, err
	}

	//teamid,_ := strconv.ParseUint(words[:i],10,64)
	tmp, _ := strconv.ParseInt(words[:i], 10, 64)
	teamid := uint32(tmp)
	return teamid, nil
}

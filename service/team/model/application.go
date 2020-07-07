package model

import (
	m "github.com/Muxi-X/workbench-be/model"
	"github.com/Muxi-X/workbench-be/pkg/constvar"
	"github.com/spf13/viper"
	"strconv"
	"time"

)

type ApplyUserItem struct {
	ID    uint32
	Name  string
	Eamil string
}

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

func ListApplictions(offset uint32, limit uint32, pagination bool) ([]*ApplyUserItem, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	applicationlist := make([]*ApplyUserItem, 0)

	query := m.DB.Self.Table("applys").Select("id").Order("id desc")

	if pagination {
		query = query.Offset(offset).Limit(limit)
	}

	var count uint64

	if err := query.Scan(&applicationlist).Count(&count).Error; err != nil {
		return applicationlist, count, err
	}

	return applicationlist, count, nil
}
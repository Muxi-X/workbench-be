package handler

import (
	"PROJECT_SCRIPT/model"
)

func ClearUsers() {
	vals := model.GetAllUser2project()
	for _, val := range vals {
		model.ClearUser(val)
	}
}

package main_test

import (
	"muxi-workbench-project/model"
	m "muxi-workbench/model"
	"testing"
)

func TestFib(t *testing.T) {
	// pIDs := []uint32{1, 2}
	var record []*model.SearchResult
	err := m.DB.Self.Error
	if err != nil {
		panic(err)
	}
	// .Where("project_id in ?", pIDs)
	// Raw("select id, filename, last_edit_time, content from docs UNION ALL select id, filename, last_edit_time, content from files")

	// if err := query.Scan(&record).Error; err != nil {
	// 	t.Error(err)
	// }
	// res, count, err := model.SearchTitle(pIDs, "1", 0, 0, 0, false)
	// err2 := errors.New("")

	for _, r := range record {
		t.Logf("%+v\n", r)
	}
}

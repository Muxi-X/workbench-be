package model

import (
	m "muxi-workbench/model"
)

// FolderForDocModel ... 文档文件夹模型
type FolderForDocModel struct {
	FolderModel
}

// TableName ... 物理表名
func (u *FolderForDocModel) TableName() string {
	return "foldersformds"
}

// Create ... 创建文档文件夹
func (u *FolderForDocModel) Create() error {
	return m.DB.Self.Create(&u).Error
}

// Update ... 更新文档文件夹
func (u *FolderForDocModel) Update() error {
	return m.DB.Self.Save(u).Error
}

// GetFolderForDocDetail ... 获取文档文件夹详情
func GetFolderForDocDetail(id uint32) (*FolderDetail, error) {
	s := &FolderDetail{}
	d := m.DB.Self.Table("foldersformds").Select("foldersformds.*, u.name as creator").Joins("left join users u on u.id = foldersformds.create_id").Where("foldersformds.id = ? AND re = 0", id).First(&s)
	return s, d.Error
}

// GetFolderForDocModel ... 获取文档文件夹
func GetFolderForDocModel(id uint32) (*FolderModel, error) {
	s := &FolderModel{}
	d := m.DB.Self.Table("foldersformds").Where("id = ? AND re = 0", id).First(&s)
	return s, d.Error
}

// GetFolderForDocInfoByIds ... 获取文档文件夹信息列表
func GetFolderForDocInfoByIds(ids []uint32) ([]*FolderInfo, error) {
	s := make([]*FolderInfo, 0)
	d := m.DB.Self.Table("foldersformds").Where("id IN (?) AND re = 0", ids).Find(&s)
	return s, d.Error
}

// GetDocChildrenById ... 获取子文档/夹
func GetDocChildrenById(id uint32) (*FolderChildren, error) {
	s := &FolderChildren{}
	d := m.DB.Self.Table("foldersformds").Where("id = ? AND re = 0", id).Find(&s)
	return s, d.Error
}

package service

import (
	"gluten/global"
	"gluten/model"
)

func AddGlutenInfo(info model.GlutenInfo) error {
	err := global.DB.Create(&info).Error
	return err
}

func SelectAllGlutenInfoById(id uint) (err error, info []model.GlutenInfo) {
	err = global.DB.Where(&model.GlutenInfo{UserId: id}).Find(&info).Error
	return
}

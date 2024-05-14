package dao

import (
	"TARG_revenue_report_backend/db"
	"TARG_revenue_report_backend/db/model"
	"gorm.io/gorm"
)

type UserBasicDao struct {
	*gorm.DB
}

func NewUserBasicDao() *UserBasicDao {
	return &UserBasicDao{db.NewDB()}
}

func (dao *UserBasicDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("employee_name=?", userName).Find(&user).Count(&count).Error
	if count == 0 {
		return nil, false, err
	}
	return user, true, nil
}

func (dao *UserBasicDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(&user).Error
	return
}

// 创建用户权限表
func (dao *UserBasicDao) CreateProject(project *model.RolePrivilege) (err error) {
	err = dao.DB.Model(&model.RolePrivilege{}).Create(&project).Error
	return
}

// GetUserById 根据id获取user
func (dao *UserBasicDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id =?", id).First(&user).Error
	return
}

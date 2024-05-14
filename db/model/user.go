package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model          // 生成以ID自增为主的主键以及相关时间记录
	EmployeeId   string `gorm:"type:varchar(20);column:employee_id;unique;not null" json:"employeeId"` // 工号
	EmployeeName string `gorm:"type:varchar(10);column:employee_name" json:"employeeName"`             // 姓名
	Email        string `gorm:"type:varchar(50);column:email" json:"email"`                            // 邮箱
	Phone        string `gorm:"type:varchar(20);column:phone" json:"phone"`                            //用户手机
	Group        string `gorm:"type:varchar(30);column:group" json:"group"`                            //部门从属
	UserExt      string `gorm:"type:varchar(20);column:user_ext" json:"userExt"`                       // 用户分机
	OpenId       string `gorm:"type:varchar(100);column:open_id" json:"openId"`                        // 用户唯一码
}

const (
	PasswordConst        = 12       //密码加密难度
	Active        string = "active" //激活用户
)

//func (user *User) SetPassword(password string) error {
//	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordConst)
//	if err != nil {
//		return err
//	}
//	user.Password = string(bytes)
//	return nil
//}
//
//func (user *User) CheckPassword(password string) bool {
//	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
//	return err == nil
//}

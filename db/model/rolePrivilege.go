package model

import "gorm.io/gorm"

type RolePrivilege struct {
	gorm.Model
	Role         string `gorm:"type:varchar(20);column:role;not null" json:"role"`                  //工号
	EmployeeName string `gorm:"type:varchar(10);column:employee_name;not null" json:"employeeName"` //姓名
	View         int    `gorm:"type:tinyint(1);column:view" json:"view"`                            //是否可观看(0:不可以,1:可以)
	Edit         int    `gorm:"type:tinyint(1);column:edit" json:"edit"`                            //是否可编辑(0:不可以,1：可以)
	ProjectName  string `gorm:"type:varchar(200);column:project_name" json:"projectName"`           // 项目名称
}

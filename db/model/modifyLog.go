package model

import (
	"gorm.io/gorm"
)

type ModifyLog struct {
	gorm.Model
	EmployeeId    string `gorm:"type:varchar(20);column:employ_id;not null" json:"employeeId"`         // 工号
	EmployeeName  string `gorm:"type:varchar(10);column:employee_name;not null" json:"employeeName"`   // 姓名
	ModifyProject string `gorm:"type:varchar(10);column:modify_project;not null" json:"modifyProject"` // 修改专案名
	ModifyColumn  string `gorm:"type:varchar(30);column:modify_column" json:"modifyColumn"`            // 修改列名
	ModifyBefore  string `gorm:"type:text;column:modify_before" json:"modifyBefore"`                   // 修改前的值
	ModifyAfter   string `gorm:"type:text;column:modify_after" json:"modifyAfter"`                     // 修改后的值
}

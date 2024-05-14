package model

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	ProjectName   string `gorm:"type:varchar(200);column:project_name" json:"projectName"`     // 项目名称
	PoNumber      string `gorm:"type:varchar(255);column:po_number" json:"poNumber"`           // PO单号
	PoQuoteprice  int    `gorm:"type:decimal(10,2);column:po_quoteprice" json:"poQuoteprice"`  // PO报价
	MeterialCost  int    `gorm:"type:decimal(10,2);column:meterial_cost" json:"meterialCost"`  // 材料成本
	InvoiceNumber string `gorm:"type:varchar(255);column:invoice_number" json:"invoiceNumber"` // 发票号
}

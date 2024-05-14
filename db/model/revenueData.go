package model

import (
	"gorm.io/gorm"
	"time"
)

type RevenueData struct {
	gorm.Model
	ProjectType               string    `gorm:"type:varchar(50);column:project_type" json:"projectType"`                               // 项目类型
	ProjectName               string    `gorm:"type:varchar(200);column:project_name" json:"ProjectName"`                              // 项目名称
	ProjectStatus             string    `gorm:"type:varchar(20);column:project_status" json:"ProjectStatus"`                           // 项目情况确认
	EstimatedPoTime           time.Time `gorm:"column:estimated_po_time" json:"EstimatedPoTime"`                                       // 预估PO时间
	PoNumber                  string    `gorm:"type:json;column:po_number" json:"PoNumber"`                                            //  PO单号
	ReceivingOrderCorporation string    `gorm:"type:varchar(100);column:receiving_order_corporation" json:"ReceivingOrderCorporation"` // 接单法人
	Owner                     string    `gorm:"type:varchar(30);column:owner" json:"Owner"`                                            // 项目主管
	FactoryArea               string    `gorm:"type:varchar(30);column:factory_area" json:"FactoryArea"`                               //  厂区
	CustomerBase              string    `gorm:"type:varchar(30);column:customer_base" json:"customerBase"`                             // 客户群
	BusinessGroup             string    `gorm:"type:varchar(30);column:business_group" json:"businessGroup"`                           // 事业群
	DemandUnit                string    `gorm:"type:varchar(30);column:demand_unit" json:"demandUnit"`                                 // 需求单位
	DemandOwner               string    `gorm:"type:varchar(30);column:demand_owner" json:"demandOwner"`                               // 需求方owner
	CurrencyType              string    `gorm:"type:varchar(30);column:currency_type" json:"currencyType"`                             // 币别
	EstimatedOrderPrice       int       `gorm:"type:decimal(10,2);column:estimated_order_price" json:"estimatedOrderPrice"`            // 预估订单额
	EstimatedRmb              int       `gorm:"type:decimal(10,2);column:estimated_rmb" json:"estimatedRmb"`                           // 预估本币RMB
	EstimatedRevenue          int       `gorm:"type:decimal(10,2);column:estimated_revenue" json:"estimatedRevenue"`                   // 预估营收
	EstimatedDiff             int       `gorm:"type:decimal(10,2);column:estimated_diff" json:"estimatedDiff"`                         // 金额差异
	PoQuotePrice              int       `gorm:"type:decimal(10,2);column:po_quote_price" json:"poQuotePrice"`                          // PO报价
	PoPriceRmb                int       `gorm:"type:decimal(10,2);column:po_price_rmb" json:"poPriceRmb"`                              // PO本币MRMB
	CustomerDelivery          time.Time `gorm:"column:customer_delivery" json:"customerDelivery"`                                      // 客户交期
	MeterialCost              int       `gorm:"column:meterial_cost" json:"meterialCost"`                                              // 材料成本 出货状态
	ShipmentStatus            string    `gorm:"type:varchar(100);column:shipment_status" json:"shipmentStatus"`                        // 出货状态
	ShipmentDate              time.Time `gorm:"column:shipment_date" json:"shipmentDate"`                                              // 出货日期
	EstimatedRevenueTime      time.Time `gorm:"column:estimated_revenue_time" json:"estimatedRevenueTime"`                             // 预估营收时间
	EstimatedEntryMonth       time.Time `gorm:"column:estimated_entry_month" json:"estimatedEntryMonth"`                               // 预估入账月份
	ActualEntryMonth          time.Time `gorm:"column:actual_entry_month" json:"actualEntryMonth"`                                     // 实际入账月份
	InvoiceNumber             string    `gorm:"type:varchar(50);column:invoice_number" json:"invoiceNumber"`                           // 发票号
	InvoiceDate               time.Time `gorm:"column:invoice_date" json:"InvoiceDate"`                                                // 发票日期
	OverdueDate               time.Time `gorm:"column:overdue_date" json:"overdueDate"`                                                // 逾期日期
	OverdueDateCount          int       `gorm:"type:decimal(10,2);column:overdue_date_count" json:"overdueDateCount"`                  // 逾期天数
	SettlementMethod          string    `gorm:"type:varchar(30);column:settlement_method" json:"settlementMethod"`                     // 结报方式
	SapNumber                 string    `gorm:"type:varchar(50);column:sap_number" json:"sapNumber"`                                   // SAP结报单号or工单号
	ProjectDetail             string    `gorm:"type:text;column:project_detail" json:"projectDetail"`                                  // delay原因/目前进展/状况
	Notes                     string    `gorm:"type:text;column:notes" json:"notes"`                                                   // 备注
	Site                      string    `gorm:"type:varchar(30);column:site" json:"site"`                                              // 厂区代码
	BeforeEstimatedEntryMonth time.Time `gorm:"column:before_estimated_entry_month" json:"beforeEstimatedEntryMonth"`                  // 之前预估入账月份
	UpdateBy                  string    `gorm:"type:varchar(20);column:update_by" json:"updateBy"`                                     // 更新人工号
	IsCancel                  int       `gorm:"type:decimal(10,2);column:is_cancel" json:"isCancel"`                                   // 专案是否取消(0:非取消 1:已取消)
}

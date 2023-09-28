// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameClearingRecord = "clearing_record"

// ClearingRecord mapped from table <clearing_record>
type ClearingRecord struct {
	ClearingRecordID                string     `gorm:"column:clearing_record_id;type:varchar(64);primaryKey" json:"clearing_record_id"`                                 // 主键ID
	ChannelReferenceID              string     `gorm:"column:channel_reference_id;type:varchar(64);not null" json:"channel_reference_id"`                               // 我方流水号（我方送渠道订单号）
	TransactionOrderID              string     `gorm:"column:transaction_order_id;type:varchar(64);not null" json:"transaction_order_id"`                               // 代表我方的支付订单主键
	FundsReconciliationPipeID       *string    `gorm:"column:funds_reconciliation_pipe_id;type:varchar(64)" json:"funds_reconciliation_pipe_id"`                        // 资金对账通道
	ClearingReconciliationPipeID    *string    `gorm:"column:clearing_reconciliation_pipe_id;type:varchar(64)" json:"clearing_reconciliation_pipe_id"`                  // 清算对账通道
	RecordType                      string     `gorm:"column:record_type;type:varchar(32);not null" json:"record_type"`                                                 // 单据类型：收单/外汇
	TransactionType                 string     `gorm:"column:transaction_type;type:varchar(20);not null" json:"transaction_type"`                                       // 支付类型：payment/refund/cb/cb_win/cb_lose
	OrderStatus                     string     `gorm:"column:order_status;type:varchar(12);not null" json:"order_status"`                                               // 订单状态：success/fail
	Amount                          string     `gorm:"column:amount;type:varchar(64);not null" json:"amount"`                                                           // 我方清算金额
	Currency                        string     `gorm:"column:currency;type:varchar(6);not null" json:"currency"`                                                        // 我方清算币种
	OrderTime                       time.Time  `gorm:"column:order_time;type:datetime;not null" json:"order_time"`                                                      // 我方支付时间
	ClearingReconciliationStatus    string     `gorm:"column:clearing_reconciliation_status;type:varchar(8);not null" json:"clearing_reconciliation_status"`            // 清算对账状态: 已对账/未对账
	ClearingReconciliationResult    *string    `gorm:"column:clearing_reconciliation_result;type:varchar(64)" json:"clearing_reconciliation_result"`                    // 清算对账结果
	ClearingTime                    *time.Time `gorm:"column:clearing_time;type:datetime" json:"clearing_time"`                                                         // 清算对账时间
	StatementID                     *string    `gorm:"column:statement_id;type:varchar(64)" json:"statement_id"`                                                        // 我方账单编号
	FundsReconciliationStatus       *string    `gorm:"column:funds_reconciliation_status;type:varchar(10)" json:"funds_reconciliation_status"`                          // 资金对账状态：已对账/未对账
	FundsReconciliationBatch        *string    `gorm:"column:funds_reconciliation_batch;type:varchar(64)" json:"funds_reconciliation_batch"`                            // 资金对账批次号
	Age                             int32      `gorm:"column:age;type:int;not null" json:"age"`                                                                         // 账龄
	AgeLevel                        *string    `gorm:"column:age_level;type:varchar(10)" json:"age_level"`                                                              // 账龄等级
	AgeStatus                       *string    `gorm:"column:age_status;type:varchar(10)" json:"age_status"`                                                            // 账龄处理状态
	ChannelID                       string     `gorm:"column:channel_id;type:varchar(20);not null" json:"channel_id"`                                                   // 渠道
	ChannelClearingBatchNo          *string    `gorm:"column:channel_clearing_batch_no;type:varchar(64)" json:"channel_clearing_batch_no"`                              // 渠道清算批次
	ChannelClearingTime             *string    `gorm:"column:channel_clearing_time;type:varchar(64)" json:"channel_clearing_time"`                                      // 渠道清算时间
	ChannelDealID                   string     `gorm:"column:channel_deal_id;type:varchar(128);not null" json:"channel_deal_id"`                                        // 渠道流水号（渠道订单号）
	ChannelSettlementCurrency       string     `gorm:"column:channel_settlement_currency;type:varchar(6);not null" json:"channel_settlement_currency"`                  // 结算币种
	RecordStatus                    string     `gorm:"column:record_status;type:varchar(32);not null" json:"record_status"`                                             // 凭证处理状态: 正常/异常
	EstimateClearingTime            *string    `gorm:"column:estimate_clearing_time;type:varchar(80)" json:"estimate_clearing_time"`                                    // 预计清算时间
	CreateTime                      time.Time  `gorm:"column:create_time;type:int unsigned;autoCreateTime" json:"create_time"`                                          // 创建时间
	UpdateTime                      *time.Time `gorm:"column:update_time;type:int unsigned;autoUpdateTime" json:"update_time"`                                          // 更新时间
	ClearingCycle                   *string    `gorm:"column:clearing_cycle;type:varchar(30)" json:"clearing_cycle"`                                                    // 清算周期
	ChannelClearingRecordID         string     `gorm:"column:channel_clearing_record_id;type:varchar(64);not null" json:"channel_clearing_record_id"`                   // 标准渠道流水id
	OriginTransactionOrderID        string     `gorm:"column:origin_transaction_order_id;type:varchar(64);not null" json:"origin_transaction_order_id"`                 // 退款原支付单id
	ClearingReconcileDetailResultID string     `gorm:"column:clearing_reconcile_detail_result_id;type:varchar(64);not null" json:"clearing_reconcile_detail_result_id"` // 清算对账结果id
	PayTransactionOrderID           *string    `gorm:"column:pay_transaction_order_id;type:varchar(64)" json:"pay_transaction_order_id"`                                // 交易层支付Id
	RefundTransactionOrderID        *string    `gorm:"column:refund_transaction_order_id;type:varchar(64)" json:"refund_transaction_order_id"`                          // 交易层退款Id
	MerchantID                      *string    `gorm:"column:merchant_id;type:varchar(64)" json:"merchant_id"`                                                          // 商户id
	Region                          *string    `gorm:"column:region;type:varchar(20)" json:"region"`                                                                    // 国家/地区
	PayMethod                       *string    `gorm:"column:pay_method;type:varchar(20)" json:"pay_method"`                                                            // 支付方式
	Extension                       *string    `gorm:"column:extension;type:text" json:"extension"`                                                                     // 拓展字段
	SettleFinishTime                *string    `gorm:"column:settle_finish_time;type:varchar(64)" json:"settle_finish_time"`                                            // 结账时间
	BillID                          *string    `gorm:"column:bill_id;type:varchar(64)" json:"bill_id"`                                                                  // POS结账退款对账id
}

// TableName ClearingRecord's table name
func (*ClearingRecord) TableName() string {
	return TableNameClearingRecord
}

package test

import (
	"fmt"
	"sort"
	"testing"
)

type ReconcileService interface {
	Handle(*interface{}) error
}

type QueryStatementServiceImpl struct {
	ChannelName string
	FundPipeId  string
}

func (q *QueryStatementServiceImpl) Handle(i *interface{}) error {
	fmt.Println("开始查询账单")
	return nil
}

type SummaryAmountServiceImpl struct {
	AmountList     []string
	Identification string
}

func (s *SummaryAmountServiceImpl) Handle(i *interface{}) error {
	fmt.Println("开始汇总金额")
	return nil
}

// 根据类型断言进行排序, 详见 gorm.Open
func TestSortByTypeAssertion(t *testing.T) {
	var items []interface{}

	citiQuery := &QueryStatementServiceImpl{ChannelName: "CKO", FundPipeId: "CITI"}
	hsbcQuery := &QueryStatementServiceImpl{ChannelName: "RevPay", FundPipeId: "HSBC"}
	taishinQuery := &QueryStatementServiceImpl{ChannelName: "Taishin", FundPipeId: "Taishin"}

	citySummary := &SummaryAmountServiceImpl{AmountList: []string{"1.1", "12.2", "109.54"}, Identification: "1"}
	hsbcSummary := &SummaryAmountServiceImpl{AmountList: []string{"10.5"}, Identification: "2"}
	taishinSummary := &SummaryAmountServiceImpl{AmountList: []string{"2", "5.4"}, Identification: "3"}

	items = append(items, citySummary, citiQuery, hsbcSummary, taishinSummary, hsbcQuery, taishinQuery)

	fmt.Printf("第一次执行的结果: %v\n", items)

	sort.Slice(items, func(i, j int) bool {
		_, isConfig := items[i].(*QueryStatementServiceImpl)
		_, isConfig2 := items[j].(*QueryStatementServiceImpl)
		return isConfig && !isConfig2
	})

	fmt.Printf("第二次执行的结果: %v\n", items)

	sort.Slice(items, func(i, j int) bool {
		_, isConfig := items[i].(QueryStatementServiceImpl)
		_, isConfig2 := items[j].(QueryStatementServiceImpl)
		return isConfig && !isConfig2
	})

	fmt.Printf("第三次执行的结果: %v\n", items)

}

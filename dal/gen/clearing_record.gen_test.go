package gen

import (
	"context"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gormCollection/dal/model"
	"reflect"
	"testing"
)

func Test_clearingRecordDo_Assign(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		attrs []field.AssignExpr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Assign(tt.args.attrs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Assign() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Attrs(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		attrs []field.AssignExpr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Attrs(tt.args.attrs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Attrs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Clauses(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		conds []clause.Expression
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Clauses(tt.args.conds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clauses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Create(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		values []*model.ClearingRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if err := c.Create(tt.args.values...); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_clearingRecordDo_CreateInBatches(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		values    []*model.ClearingRecord
		batchSize int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if err := c.CreateInBatches(tt.args.values, tt.args.batchSize); (err != nil) != tt.wantErr {
				t.Errorf("CreateInBatches() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_clearingRecordDo_Debug(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name   string
		fields fields
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Debug(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Debug() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Delete(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		models []*model.ClearingRecord
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult gen.ResultInfo
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			gotResult, err := c.Delete(tt.args.models...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("Delete() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_clearingRecordDo_Distinct(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		cols []field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Distinct(tt.args.cols...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Exists(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		subquery interface{ UnderlyingDB() *gorm.DB }
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Exists(tt.args.subquery); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Find(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*model.ClearingRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			got, err := c.Find()
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_FindByPage(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		offset int
		limit  int
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []*model.ClearingRecord
		wantCount  int64
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			gotResult, gotCount, err := c.FindByPage(tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("FindByPage() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotCount != tt.wantCount {
				t.Errorf("FindByPage() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_clearingRecordDo_FindInBatch(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		batchSize int
		fc        func(tx gen.Dao, batch int) error
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantResults []*model.ClearingRecord
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			gotResults, err := c.FindInBatch(tt.args.batchSize, tt.args.fc)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindInBatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("FindInBatch() gotResults = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}

func Test_clearingRecordDo_FindInBatches(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		result    *[]*model.ClearingRecord
		batchSize int
		fc        func(tx gen.Dao, batch int) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if err := c.FindInBatches(tt.args.result, tt.args.batchSize, tt.args.fc); (err != nil) != tt.wantErr {
				t.Errorf("FindInBatches() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_clearingRecordDo_First(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.ClearingRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			got, err := c.First()
			if (err != nil) != tt.wantErr {
				t.Errorf("First() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_FirstOrCreate(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.ClearingRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			got, err := c.FirstOrCreate()
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstOrCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrCreate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_FirstOrInit(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.ClearingRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			got, err := c.FirstOrInit()
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstOrInit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrInit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_GetByName(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		channelReferenceId string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []*model.ClearingRecord
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			gotResult, err := c.GetByName(tt.args.channelReferenceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("GetByName() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_clearingRecordDo_Group(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		cols []field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Group(tt.args.cols...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Having(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		conds []gen.Condition
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Having(tt.args.conds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Having() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Join(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		table schema.Tabler
		on    []field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Join(tt.args.table, tt.args.on...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Joins(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		fields []field.RelationField
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Joins(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Joins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Last(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.ClearingRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			got, err := c.Last()
			if (err != nil) != tt.wantErr {
				t.Errorf("Last() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_LeftJoin(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		table schema.Tabler
		on    []field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.LeftJoin(tt.args.table, tt.args.on...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LeftJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Limit(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		limit int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Limit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Not(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		conds []gen.Condition
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Not(tt.args.conds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Not() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Offset(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		offset int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Offset(tt.args.offset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Omit(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		cols []field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Omit(tt.args.cols...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Omit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Or(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		conds []gen.Condition
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Or(tt.args.conds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Order(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		conds []field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Order(tt.args.conds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Preload(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		fields []field.RelationField
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Preload(tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Preload() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_ReadDB(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name   string
		fields fields
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.ReadDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Returning(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		value   interface{}
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Returning(tt.args.value, tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Returning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_RightJoin(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		table schema.Tabler
		on    []field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.RightJoin(tt.args.table, tt.args.on...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RightJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Save(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		values []*model.ClearingRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if err := c.Save(tt.args.values...); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_clearingRecordDo_Scan(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		result interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if err := c.Scan(tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_clearingRecordDo_ScanByPage(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		result interface{}
		offset int
		limit  int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantCount int64
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			gotCount, err := c.ScanByPage(tt.args.result, tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScanByPage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotCount != tt.wantCount {
				t.Errorf("ScanByPage() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_clearingRecordDo_Scopes(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		funcs []func(gen.Dao) gen.Dao
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Scopes(tt.args.funcs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scopes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Select(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		conds []field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Select(tt.args.conds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Session(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		config *gorm.Session
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Session(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Take(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name    string
		fields  fields
		want    *model.ClearingRecord
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			got, err := c.Take()
			if (err != nil) != tt.wantErr {
				t.Errorf("Take() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Take() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Unscoped(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name   string
		fields fields
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Unscoped(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unscoped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_Where(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		conds []gen.Condition
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.Where(tt.args.conds...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_WithContext(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.WithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_WriteDB(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	tests := []struct {
		name   string
		fields fields
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.WriteDB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecordDo_withDO(t *testing.T) {
	type fields struct {
		DO gen.DO
	}
	type args struct {
		do gen.Dao
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *clearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &clearingRecordDo{
				DO: tt.fields.DO,
			}
			if got := c.withDO(tt.args.do); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("withDO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecord_Alias(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			if got := c.Alias(); got != tt.want {
				t.Errorf("Alias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecord_As(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	type args struct {
		alias string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *clearingRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			if got := c.As(tt.args.alias); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("As() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecord_GetFieldByName(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	type args struct {
		fieldName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   field.OrderExpr
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			got, got1 := c.GetFieldByName(tt.args.fieldName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFieldByName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetFieldByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_clearingRecord_Table(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	type args struct {
		newTableName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *clearingRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			if got := c.Table(tt.args.newTableName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecord_TableName(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			if got := c.TableName(); got != tt.want {
				t.Errorf("TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecord_WithContext(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   IClearingRecordDo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			if got := c.WithContext(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecord_clone(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   clearingRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			if got := c.clone(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecord_fillFieldMap(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			c.fillFieldMap()
		})
	}
}

func Test_clearingRecord_replaceDB(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   clearingRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			if got := c.replaceDB(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_clearingRecord_updateTableName(t *testing.T) {
	type fields struct {
		clearingRecordDo                clearingRecordDo
		ALL                             field.Asterisk
		ClearingRecordID                field.String
		ChannelReferenceID              field.String
		TransactionOrderID              field.String
		FundsReconciliationPipeID       field.String
		ClearingReconciliationPipeID    field.String
		RecordType                      field.String
		TransactionType                 field.String
		OrderStatus                     field.String
		Amount                          field.String
		Currency                        field.String
		OrderTime                       field.Time
		ClearingReconciliationStatus    field.String
		ClearingReconciliationResult    field.String
		ClearingTime                    field.Time
		StatementID                     field.String
		FundsReconciliationStatus       field.String
		FundsReconciliationBatch        field.String
		Age                             field.Int32
		AgeLevel                        field.String
		AgeStatus                       field.String
		ChannelID                       field.String
		ChannelClearingBatchNo          field.String
		ChannelClearingTime             field.String
		ChannelDealID                   field.String
		ChannelSettlementCurrency       field.String
		RecordStatus                    field.String
		EstimateClearingTime            field.String
		CreateTime                      field.Time
		UpdateTime                      field.Time
		ClearingCycle                   field.String
		ChannelClearingRecordID         field.String
		OriginTransactionOrderID        field.String
		ClearingReconcileDetailResultID field.String
		PayTransactionOrderID           field.String
		RefundTransactionOrderID        field.String
		MerchantID                      field.String
		Region                          field.String
		PayMethod                       field.String
		Extension                       field.String
		SettleFinishTime                field.String
		BillID                          field.String
		fieldMap                        map[string]field.Expr
	}
	type args struct {
		table string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *clearingRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &clearingRecord{
				clearingRecordDo:                tt.fields.clearingRecordDo,
				ALL:                             tt.fields.ALL,
				ClearingRecordID:                tt.fields.ClearingRecordID,
				ChannelReferenceID:              tt.fields.ChannelReferenceID,
				TransactionOrderID:              tt.fields.TransactionOrderID,
				FundsReconciliationPipeID:       tt.fields.FundsReconciliationPipeID,
				ClearingReconciliationPipeID:    tt.fields.ClearingReconciliationPipeID,
				RecordType:                      tt.fields.RecordType,
				TransactionType:                 tt.fields.TransactionType,
				OrderStatus:                     tt.fields.OrderStatus,
				Amount:                          tt.fields.Amount,
				Currency:                        tt.fields.Currency,
				OrderTime:                       tt.fields.OrderTime,
				ClearingReconciliationStatus:    tt.fields.ClearingReconciliationStatus,
				ClearingReconciliationResult:    tt.fields.ClearingReconciliationResult,
				ClearingTime:                    tt.fields.ClearingTime,
				StatementID:                     tt.fields.StatementID,
				FundsReconciliationStatus:       tt.fields.FundsReconciliationStatus,
				FundsReconciliationBatch:        tt.fields.FundsReconciliationBatch,
				Age:                             tt.fields.Age,
				AgeLevel:                        tt.fields.AgeLevel,
				AgeStatus:                       tt.fields.AgeStatus,
				ChannelID:                       tt.fields.ChannelID,
				ChannelClearingBatchNo:          tt.fields.ChannelClearingBatchNo,
				ChannelClearingTime:             tt.fields.ChannelClearingTime,
				ChannelDealID:                   tt.fields.ChannelDealID,
				ChannelSettlementCurrency:       tt.fields.ChannelSettlementCurrency,
				RecordStatus:                    tt.fields.RecordStatus,
				EstimateClearingTime:            tt.fields.EstimateClearingTime,
				CreateTime:                      tt.fields.CreateTime,
				UpdateTime:                      tt.fields.UpdateTime,
				ClearingCycle:                   tt.fields.ClearingCycle,
				ChannelClearingRecordID:         tt.fields.ChannelClearingRecordID,
				OriginTransactionOrderID:        tt.fields.OriginTransactionOrderID,
				ClearingReconcileDetailResultID: tt.fields.ClearingReconcileDetailResultID,
				PayTransactionOrderID:           tt.fields.PayTransactionOrderID,
				RefundTransactionOrderID:        tt.fields.RefundTransactionOrderID,
				MerchantID:                      tt.fields.MerchantID,
				Region:                          tt.fields.Region,
				PayMethod:                       tt.fields.PayMethod,
				Extension:                       tt.fields.Extension,
				SettleFinishTime:                tt.fields.SettleFinishTime,
				BillID:                          tt.fields.BillID,
				fieldMap:                        tt.fields.fieldMap,
			}
			if got := c.updateTableName(tt.args.table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newClearingRecord(t *testing.T) {
	type args struct {
		db   *gorm.DB
		opts []gen.DOOption
	}
	tests := []struct {
		name string
		args args
		want clearingRecord
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newClearingRecord(tt.args.db, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newClearingRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

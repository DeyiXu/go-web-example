package model

const (
	// ExampleTableName 例子表名
	ExampleTableName = "examples"
)

// Example 例子
type Example struct {
	Name string `json:"name"` // 名称
	BaseModel
}

// TableName ...
func (Example) TableName() string {
	return ExampleTableName
}

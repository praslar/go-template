package models

type Demo struct {
	BaseModel
	Name string `json:"name" gorm:"not null"`
	Age  int    `json:"age" gorm:"null"`
}

type CreateDemoRequest struct {
	Name      *string   `json:"name" valid:"Required"`
	Age       *int      `json:"age" valid:"Required"`
}

type UpdateDemoRequest struct {
	Name      *string   `json:"name,omitempty"`
	Age       *int      `json:"age,omitempty"`
}

func (Demo) TableName() string {
	return "demo_table"
}

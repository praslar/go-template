package models

import (
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm/dialects/postgres"
)

const (
	DataTypeList    = "list"
	DataTypeText    = "text"
	DataTypeBoolean = "boolean"
	DataTypeDate    = "date"
)

type Temp struct {
	BaseModel
	AdminName     string         `json:"admin_name" gorm:"column:admin_name;not null"`
	InputType     string         `json:"input_type" gorm:"column:input_type;not null"`
	PlatformKey   string         `json:"platform_key" gorm:"column:platform_key;not null"`
	DisplayNames  postgres.Jsonb `json:"display_names" gorm:"type:jsonb;column:display_names;not null"`
	DisplayLabels postgres.Jsonb `json:"display_labels" gorm:"type:jsonb;column:display_labels;null"`
	ListData      postgres.Jsonb `json:"list_data" gorm:"type:jsonb;column:list_data;null"`
	HelpModalData postgres.Jsonb `json:"help_modal_data" gorm:"type:jsonb;column:help_modal_data;null"`
	IsActive      bool           `json:"is_active" gorm:"column:is_active;not null; default:true"`
}

func (Temp) TableName() string {
	return "temp"
}

type CreateTemp struct {
	//AdminName     *string         `json:"admin_name" valid:"Required"`
	//InputType     *string         `json:"input_type" example:"list" description:"product attribute type, valid values: list, text, boolean, date" `
	//PlatformKey   *string         `json:"platform_key" valid:"Required"`
	//ListData      *postgres.Jsonb `json:"list_data"`
	//HelpModalData *postgres.Jsonb `json:"help_modal_data" valid:"Required"`
	//DisplayNames  *postgres.Jsonb `json:"display_names" valid:"Required"`
	//DisplayLabels *postgres.Jsonb `json:"display_labels" valid:"Required"`
	//IsActive      *bool           `json:"is_active"`
}

type UpdateTemp struct {
	//AdminName     *string         `json:"admin_name" `
	//InputType     *string         `json:"input_type"`
	//PlatformKey   *string         `json:"platform_key"`
	//ListData      *postgres.Jsonb `json:"list_data"`
	//HelpModalData *postgres.Jsonb `json:"help_modal_data"`
	//DisplayNames  *postgres.Jsonb `json:"display_names" `
	//DisplayLabels *postgres.Jsonb `json:"display_labels"`
	//IsActive      *bool           `json:"is_active"`
}

func (p CreateTemp) ValidateNotifyRecursiveType() error {
	//validItems := []string{DataTypeList, DataTypeText, DataTypeBoolean, DataTypeDate}
	//valid := false
	//for _, item := range validItems {
	//	if *p.InputType == item {
	//		valid = true
	//		break
	//	}
	//}
	//
	//if !valid {
	//	return fmt.Errorf("invalid '%[1]s'. Accepted [%[2]s]", *p.InputType, strings.Join(validItems, ", "))
	//}
	return nil
}

func (p CreateTemp) Valid(v *validation.Validation) {
	if err := p.ValidateNotifyRecursiveType(); err != nil {
		_ = v.SetError("notify_recursive_type", err.Error())
	}
}

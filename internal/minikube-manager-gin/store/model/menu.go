package model

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	menuTable = "menu"
)

type Menu struct {
	gorm.Model
	UUID   string   `gorm:"column:uuid; type:string; size:36; not null; unique; <-:create;"`
	Type   string   `gorm:"column:type; type:string; size:5; unique;"` // home、logo、menu
	Title  string   `gorm:"column:title; type:string; size:50; "`
	Image  string   `gorm:"column:image; type:string; omitempty"`
	Href   string   `gorm:"column:href; type:string;"`
	Target string   `gorm:"column:target; type:string;omitempty"`
	Child  []string `gorm:"column:child; type:json;omitempty"`
	Value  string   `gorm:"column:value; type:longtext; "`
}

// TableName returns table name
func (t Menu) TableName() string {
	return tableName(menuTable)
}

// BeforeSave is table operation hooks
func (t *Menu) BeforeSave(tx *gorm.DB) error {
	zap.L().Info("BeforeSave")
	return nil
}

// BeforeCreate is table operation hooks
func (t *Menu) BeforeCreate(tx *gorm.DB) error {
	zap.L().Info("BeforeCreate")
	return nil
}

// BeforeUpdate is table operation hooks
func (t *Menu) BeforeUpdate(tx *gorm.DB) error {
	zap.L().Info("BeforeUpdate")
	return nil
}

// BeforeDelete is table operation hooks
func (t *Menu) BeforeDelete(tx *gorm.DB) error {
	zap.L().Info("BeforeDelete")
	return nil
}

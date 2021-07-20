package repository

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	ActionsTableName = "actions"
)

type ActionType string

const (
	ActionInit ActionType = "init"
)

type ActionSchema struct {
	ID        int `gorm:"primary_key, AUTO_INCREMENT"`
	Type      ActionType
	Params    string
	CreatedAt time.Time
	Error     string
}

func (as *ActionSchema) String() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v",
		as.ID, as.Type, as.Params, as.CreatedAt.Format("2006-01-02 15:04:05"), as.Error)
}

type Actions interface {
	Table
	Size(ctx context.Context) (int, error)
	Add(ctx context.Context, name ActionType, params string, errInfo string) (int, error)
	Items(ctx context.Context, maxCount int) ([]*ActionSchema, error)
}

type actions struct {
	tbl *gorm.DB
}

func NewActions(db *gorm.DB) Actions {
	tbl := db.Table(ActionsTableName)
	tbl.AutoMigrate(&ActionSchema{})
	return &actions{
		tbl: tbl,
	}
}

func (act *actions) Name() string {
	return ActionsTableName
}

func (act *actions) Size(ctx context.Context) (int, error) {
	var count int64
	tx := act.tbl.WithContext(ctx).Count(&count)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return int(count), nil
}

func (act *actions) Add(ctx context.Context, name ActionType, params string, errInfo string) (int, error) {
	action := &ActionSchema{
		Type:      name,
		Params:    params,
		CreatedAt: time.Now(),
		Error:     errInfo,
	}
	result := act.tbl.WithContext(ctx).Create(action)
	if err := result.Error; err != nil {
		return 0, err
	}
	return action.ID, nil
}

func (act *actions) Items(ctx context.Context, maxCount int) ([]*ActionSchema, error) {
	var history []*ActionSchema
	result := act.tbl.WithContext(ctx).Order("ID desc").Limit(maxCount).Find(&history)
	act.Add(ctx, ActionType("history"), fmt.Sprintf("maxCount=%v", maxCount), "")
	return history, result.Error
}

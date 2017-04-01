// Code generated by goagen v1.1.0, command line:
// $ goagen
// --design=github.com/tikasan/eventory-goa/design
// --out=$(GOPATH)
// --version=v1.1.0-dirty
//
// API "eventory": Models
//
// The content of this file is auto-generated, DO NOT MODIFY

package models

import (
	"../app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// イベントジャンル
type EventGenres struct {
	ID        int        `gorm:"primary_key"` // primary key
	EventID   int        // has many EventGenres
	GenreID   int        // has many EventGenres
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m EventGenres) TableName() string {
	return "event_genres"

}

// EventGenresDB is the implementation of the storage interface for
// EventGenres.
type EventGenresDB struct {
	Db *gorm.DB
}

// NewEventGenresDB creates a new storage type.
func NewEventGenresDB(db *gorm.DB) *EventGenresDB {
	return &EventGenresDB{Db: db}
}

// DB returns the underlying database.
func (m *EventGenresDB) DB() interface{} {
	return m.Db
}

// EventGenresStorage represents the storage interface.
type EventGenresStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*EventGenres, error)
	Get(ctx context.Context, id int) (*EventGenres, error)
	Add(ctx context.Context, eventgenres *EventGenres) error
	Update(ctx context.Context, eventgenres *EventGenres) error
	Delete(ctx context.Context, id int) error

	ListEvent(ctx context.Context) []*app.Event
	OneEvent(ctx context.Context, id int) (*app.Event, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *EventGenresDB) TableName() string {
	return "event_genres"

}

// CRUD Functions

// Get returns a single EventGenres as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *EventGenresDB) Get(ctx context.Context, id int) (*EventGenres, error) {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenres", "get"}, time.Now())

	var native EventGenres
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of EventGenres
func (m *EventGenresDB) List(ctx context.Context) ([]*EventGenres, error) {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenres", "list"}, time.Now())

	var objs []*EventGenres
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *EventGenresDB) Add(ctx context.Context, model *EventGenres) error {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenres", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding EventGenres", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *EventGenresDB) Update(ctx context.Context, model *EventGenres) error {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenres", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating EventGenres", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *EventGenresDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "eventGenres", "delete"}, time.Now())

	var obj EventGenres

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting EventGenres", "error", err.Error())
		return err
	}

	return nil
}

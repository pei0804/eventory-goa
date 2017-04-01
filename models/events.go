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

// イベント
type Events struct {
	ID             int `gorm:"primary_key"` // primary key
	APIType        string
	Accepte        int
	Address        string
	EventGenres    []EventGenres // has many EventGenres
	Identifier     string
	Limits         int
	PrefID         int // has many Events
	URL            string
	UserKeepStatus []UserKeepStatus // has many UserKeepStatus
	Wait           int
	CreatedAt      time.Time  // timestamp
	DeletedAt      *time.Time // nullable timestamp (soft delete)
	EndAt          time.Time  // timestamp
	StartAt        time.Time  // timestamp
	UpdatedAt      time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Events) TableName() string {
	return "events"

}

// EventsDB is the implementation of the storage interface for
// Events.
type EventsDB struct {
	Db *gorm.DB
}

// NewEventsDB creates a new storage type.
func NewEventsDB(db *gorm.DB) *EventsDB {
	return &EventsDB{Db: db}
}

// DB returns the underlying database.
func (m *EventsDB) DB() interface{} {
	return m.Db
}

// EventsStorage represents the storage interface.
type EventsStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Events, error)
	Get(ctx context.Context, id int) (*Events, error)
	Add(ctx context.Context, events *Events) error
	Update(ctx context.Context, events *Events) error
	Delete(ctx context.Context, id int) error

	ListEvent(ctx context.Context) []*app.Event
	OneEvent(ctx context.Context, id int) (*app.Event, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *EventsDB) TableName() string {
	return "events"

}

// CRUD Functions

// Get returns a single Events as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *EventsDB) Get(ctx context.Context, id int) (*Events, error) {
	defer goa.MeasureSince([]string{"goa", "db", "events", "get"}, time.Now())

	var native Events
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Events
func (m *EventsDB) List(ctx context.Context) ([]*Events, error) {
	defer goa.MeasureSince([]string{"goa", "db", "events", "list"}, time.Now())

	var objs []*Events
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *EventsDB) Add(ctx context.Context, model *Events) error {
	defer goa.MeasureSince([]string{"goa", "db", "events", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Events", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *EventsDB) Update(ctx context.Context, model *Events) error {
	defer goa.MeasureSince([]string{"goa", "db", "events", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Events", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *EventsDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "events", "delete"}, time.Now())

	var obj Events

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Events", "error", err.Error())
		return err
	}

	return nil
}

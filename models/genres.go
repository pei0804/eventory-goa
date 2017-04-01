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

// ジャンル
type genres struct {
	ID        int `gorm:"primary_key"` // primary key
	Keyword   string
	Name      string
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m genres) TableName() string {
	return "genres"

}

// genresDB is the implementation of the storage interface for
// genres.
type genresDB struct {
	Db *gorm.DB
}

// NewgenresDB creates a new storage type.
func NewgenresDB(db *gorm.DB) *genresDB {
	return &genresDB{Db: db}
}

// DB returns the underlying database.
func (m *genresDB) DB() interface{} {
	return m.Db
}

// genresStorage represents the storage interface.
type genresStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*genres, error)
	Get(ctx context.Context, id int) (*genres, error)
	Add(ctx context.Context, genres *genres) error
	Update(ctx context.Context, genres *genres) error
	Delete(ctx context.Context, id int) error

	ListGenre(ctx context.Context) []*app.Genre
	OneGenre(ctx context.Context, id int) (*app.Genre, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *genresDB) TableName() string {
	return "genres"

}

// CRUD Functions

// Get returns a single genres as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *genresDB) Get(ctx context.Context, id int) (*genres, error) {
	defer goa.MeasureSince([]string{"goa", "db", "genres", "get"}, time.Now())

	var native genres
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of genres
func (m *genresDB) List(ctx context.Context) ([]*genres, error) {
	defer goa.MeasureSince([]string{"goa", "db", "genres", "list"}, time.Now())

	var objs []*genres
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *genresDB) Add(ctx context.Context, model *genres) error {
	defer goa.MeasureSince([]string{"goa", "db", "genres", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding genres", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *genresDB) Update(ctx context.Context, model *genres) error {
	defer goa.MeasureSince([]string{"goa", "db", "genres", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating genres", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *genresDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "genres", "delete"}, time.Now())

	var obj genres

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting genres", "error", err.Error())
		return err
	}

	return nil
}

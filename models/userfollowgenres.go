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
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// ユーザーのフォロージャンル
type UserFollowGenres struct {
	ID        int        `gorm:"primary_key"` // primary key
	GenreID   int        // has many UserFollowGenres
	UserID    int        // has many UserFollowGenres
	CreatedAt time.Time  // timestamp
	DeletedAt *time.Time // nullable timestamp (soft delete)
	UpdatedAt time.Time  // timestamp
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m UserFollowGenres) TableName() string {
	return "user_follow_genres"

}

// UserFollowGenresDB is the implementation of the storage interface for
// UserFollowGenres.
type UserFollowGenresDB struct {
	Db *gorm.DB
}

// NewUserFollowGenresDB creates a new storage type.
func NewUserFollowGenresDB(db *gorm.DB) *UserFollowGenresDB {
	return &UserFollowGenresDB{Db: db}
}

// DB returns the underlying database.
func (m *UserFollowGenresDB) DB() interface{} {
	return m.Db
}

// UserFollowGenresStorage represents the storage interface.
type UserFollowGenresStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*UserFollowGenres, error)
	Get(ctx context.Context, id int) (*UserFollowGenres, error)
	Add(ctx context.Context, userfollowgenres *UserFollowGenres) error
	Update(ctx context.Context, userfollowgenres *UserFollowGenres) error
	Delete(ctx context.Context, id int) error
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *UserFollowGenresDB) TableName() string {
	return "user_follow_genres"

}

// CRUD Functions

// Get returns a single UserFollowGenres as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *UserFollowGenresDB) Get(ctx context.Context, id int) (*UserFollowGenres, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenres", "get"}, time.Now())

	var native UserFollowGenres
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of UserFollowGenres
func (m *UserFollowGenresDB) List(ctx context.Context) ([]*UserFollowGenres, error) {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenres", "list"}, time.Now())

	var objs []*UserFollowGenres
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *UserFollowGenresDB) Add(ctx context.Context, model *UserFollowGenres) error {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenres", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding UserFollowGenres", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *UserFollowGenresDB) Update(ctx context.Context, model *UserFollowGenres) error {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenres", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating UserFollowGenres", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *UserFollowGenresDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "userFollowGenres", "delete"}, time.Now())

	var obj UserFollowGenres

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting UserFollowGenres", "error", err.Error())
		return err
	}

	return nil
}

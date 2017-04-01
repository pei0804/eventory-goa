package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("eventory", func() {
	Description("eventory Model")
	Store("MySQL", gorma.MySQL, func() {
		Description("MySQLのリレーションナルデータベース")
		Model("Events", func() {
			RendersTo(Event)
			Description("イベント")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("api_type", gorma.String, func() {
				Alias("api_type")
			})
			Field("identifier", gorma.String)
			Field("url", gorma.String, func() {
				Alias("url")
			})
			Field("limits", gorma.Integer)
			Field("accepte", gorma.Integer)
			Field("wait", gorma.Integer)
			Field("address", gorma.String)
			Field("start_at", gorma.Timestamp)
			Field("end_at", gorma.Timestamp)
			Field("createdAt", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			HasMany("EventGenres", "EventGenres")
			HasMany("UserKeepStatus", "UserKeepStatus")
		})
		Model("Prefs", func() {
			Description("都道府県")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			HasMany("Events", "Events")
			HasMany("UserFollowPrefs", "UserFollowPrefs")
		})
		Model("Genres", func() {
			RendersTo(Genre)
			Description("ジャンル")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("keyword", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			HasMany("EventGenres", "EventGenres")
			HasMany("UserFollowGenres", "UserFollowGenres")
		})
		Model("Users", func() {
			Description("ユーザー")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("email", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
			HasMany("UserFollowGenres", "UserFollowGenres")
			HasMany("UserFollowPrefs", "UserFollowPrefs")
			HasMany("UserTerminals", "UserTerminals")
			HasMany("UserKeepStatus", "UserKeepStatus")
		})
		Model("EventGenres", func() {
			RendersTo(Event)
			Description("イベントジャンル")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("genre_id", gorma.Integer)
			Field("event_id", gorma.Integer)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
		Model("UserFollowGenres", func() {
			Description("ユーザーのフォロージャンル")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("user_id", gorma.Integer)
			Field("genre_id", gorma.Integer)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
		Model("UserFollowPrefs", func() {
			Description("ユーザーのフォロー都道府県")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("user_id", gorma.Integer)
			Field("pref_id", gorma.Integer)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
		Model("UserKeepStatus", func() {
			Description("ユーザーのキープ状態")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("user_id", gorma.Integer)
			Field("event_id", gorma.Integer)
			Field("status", gorma.String)
			Field("batch_processed ", gorma.Boolean)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
		Model("UserTerminals", func() {
			Description("ユーザーの端末情報")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("user_id", gorma.Integer)
			Field("platform", gorma.String)
			Field("client_version", gorma.String)
			Field("token", gorma.String)
			Field("identifier", gorma.Boolean)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
	})
})

package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("eventory", func() {
	Description("eventory Model")
	Store("MySQL", gorma.MySQL, func() {
		Description("MySQLのリレーションナルデータベース")
		Model("events", func() {
			RendersTo(Event)
			Description("イベント")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("api_type", gorma.String)
			Field("identifier", gorma.String)
			Field("url", gorma.String)
			Field("limits", gorma.Integer)
			Field("accepte", gorma.Integer)
			Field("wait", gorma.Integer)
			Field("address", gorma.String)
			Field("start_at", gorma.Timestamp)
			Field("end_at", gorma.Timestamp)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
		Model("event_genres", func() {
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
		Model("genres", func() {
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
		})
		Model("prefs", func() {
			Description("都道府県")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
		Model("users", func() {
			Description("ユーザー")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
			Field("email", gorma.String)
			Field("created_at", gorma.Timestamp)
			Field("updated_at", gorma.Timestamp)
			Field("deleted_at", gorma.NullableTimestamp)
		})
		Model("user_follow_genres", func() {
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
		Model("user_follow_prefs", func() {
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
		Model("user_keep_status", func() {
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
		Model("user_terminals", func() {
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

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("event", func() {
	BasePath("/event")
	DefaultMedia(Event)
	Action("list", func() {
		Routing(
			// TODO: エンドポイント変える可能生あり
			GET("/"),
		)
		Description("イベント情報取得(クライントはこの処理の実行した時間を保持する)")
		Params(func() {
			Param("updated_at", String, "クライアントが最後にイベントを取得した時間(格納していない場合は空文字で処理します)", func() {
				Default("")
			})
		})
		Response(OK, CollectionOf(Event))
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("genre", func() {
	BasePath("/genre")
	DefaultMedia(Genre)
	Action("create", func() {
		Routing(
			POST("/new"),
		)
		Description("ジャンルの新規作成")
		Params(func() {
			Param("name", String, "ジャンル名(1~30文字)", func() {
				MinLength(1)
				MaxLength(30)
			})
			Required("name")
		})
		Response(Created)
		Response(BadRequest, ErrorMedia)
	})
	Action("list", func() {
		Routing(
			GET("/"),
		)
		Description("all genre get")
		Params(func() {
			Param("keyword", String, "ジャンル名検索に使うキーワード(0~30文字)", func() {
				MinLength(0)
				MaxLength(30)
				Default("")
			})
		})
		Response(OK, CollectionOf(Genre))
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("user", func() {
	BasePath("/user")
	Action("event fav update", func() {
		Routing(
			PUT("/:eventID/keep"),
			PUT("/:eventID/nokeep"),
		)
		Params(func() {
			Param("eventID", Integer, "イベントID（連番）")
			Param("userID", Integer, "ユーザーID（連番）")
			Required("eventID", "userID")
		})
		Description("イベントのお気に入り操作")
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("genre fav update", func() {
		Routing(
			PUT("/:genreID/add"),
			PUT("/:genreID/remove"),
		)
		Params(func() {
			Param("genreID", Integer, "genreID（連番）")
			Param("userID", Integer, "ユーザーID（連番）")
			Required("genreID", "userID")
		})
		Description("ジャンルお気に入り操作")
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger")
})

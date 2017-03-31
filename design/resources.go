package design

import (
	"github.com/goadesign/goa"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("events", func() {
	BasePath("/events")
	Security(UserAuth)
	Action("list", func() {
		Routing(
			GET("/genre/:id"),
			GET("/new"),
			GET("/keep"),
			GET("/nokeep"),
			GET("/popular"),
			GET("/recommend"),
		)
		Description("イベント情報取得")
		Params(func() {
			Param("q", String, "キーワード検索", func() {
				Default("")
			})
			Param("sort", String, "ソート", func() {
				Default("")
			})
			Param("page", Integer, "ページ(1->2->3->4)", func() {
				Minimum(0)
			})
		})
		Response(OK, CollectionOf(Event))
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("keep event", func() {
		Routing(
			PUT("/:eventID/keep"),
		)
		Params(func() {
			Param("eventID", Integer, "イベントID")
			Param("isKeep", Boolean, "キープ操作")
			Required("eventID", "userID", "isKeep")
		})
		Description("イベントのお気に入り操作")
		Response(OK)
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("genres", func() {
	BasePath("/genres")
	Security(UserAuth)
	Action("create", func() {
		Routing(
			POST("/new"),
		)
		Description("ジャンルの新規作成")
		Params(func() {
			Param("name", String, "ジャンル名", func() {
				MinLength(1)
				MaxLength(30)
			})
			Required("name")
		})
		Response(OK, Genre)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("list", func() {
		Routing(
			GET("/"),
		)
		Description("ジャンル取得")
		Params(func() {
			Param("q", String, "ジャンル名検索に使うキーワード", func() {
				MinLength(0)
				MaxLength(30)
				Default("")
			})
		})
		Response(OK, CollectionOf(Genre))
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
	Action("follow genre", func() {
		Routing(
			PUT("/:genreID/follow"),
			DELETE("/:genreID/follow"),
		)
		Params(func() {
			Param("genreID", Integer, "ジャンルID")
			Required("genreID", "userID")
		})
		Description("ジャンルお気に入り操作")
		Response(OK)
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("users", func() {
	BasePath("/users")
	Security(UserAuth)
	Action("tmp account create", func() {
		Routing(
			POST("/tmp"),
		)
		Params(func() {
			Param("client_version", String, "アプリのバージョン", func() {
				MinLength(1)
				MaxLength(10)
			})
			Param("platform", String, "OSとバージョン", func() {
				Enum("ios", "android")
			})
			Param("identifier", String, "識別子(android:Android_ID, ios:IDFV)", func() {
				Pattern("(^[a-z0-9]{16}$|^[a-z0-9\\-]{32}$)")
			})
		})
		Description("一時ユーザーの作成")
		NoSecurity()
		Response(OK, Token)
		Response(BadRequest, ErrorMedia)
	})
	Action("account create", func() {
		Routing(
			POST("/new"),
		)
		Params(func() {
			Param("email", String, "メールアドレス", func() {
				Format(goa.FormatEmail)
			})
			Param("client_version", String, "アプリのバージョン", func() {
				MinLength(1)
				MaxLength(10)
			})
			Param("platform", String, "OSとバージョン", func() {
				Enum("ios", "android")
			})
			Param("identifier", String, "識別子(android:Android_ID, ios:IDFV)", func() {
				Pattern("(^[a-z0-9]{16}$|^[a-z0-9\\-]{32}$)")
			})
		})
		Description("正規ユーザーの作成")
		Response(OK, Token)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("prefs", func() {
	BasePath("/prefs")
	Security(UserAuth)
	Action("pref follow", func() {
		Routing(
			PUT("/:prefID/follow"),
			DELETE("/:prefID/follow"),
		)
		Params(func() {
			Param("prefID", Integer, "都道府県ID")
			Required("prefID", "userID")
		})
		Description("ジャンルお気に入り操作")
		Response(OK)
		Response(NotFound)
		Response(Unauthorized)
		Response(BadRequest, ErrorMedia)
	})
})

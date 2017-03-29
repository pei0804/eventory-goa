package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// レスポンスデータの定義
var Event = MediaType("application/vnd.event+json", func() {
	Description("イベント情報")
	Attributes(func() {
		Attribute("ID", Integer, "ID", func() {
			Example(1)
		})
		Attribute("eventID", String, "イベントID", func() {
			Example("3-12313")
		})
		Attribute("apiID", String, "APIの種類 enum(ATDN=1,CONNPASS=2,DOORKEEPER=3)", func() {
			Example("ATDN")
		})
		Attribute("title", String, "イベント名", func() {
			Example("アジャイル開発勉強会")
		})
		Attribute("url", String, "イベントページURL", func() {
			Example("2016-01-01 10:10:12")
		})
		Attribute("limit", Integer, "参加人数上限", func() {
			Example(10)
		})
		Attribute("accepted", Integer, "参加登録済み人数", func() {
			Example(10)
		})
		Attribute("waitlisted", Integer, "キャンセル待ち人数", func() {
			Example(5)
		})
		Attribute("place", String, "開催地", func() {
			Example("東京都渋谷区3-31-205")
		})
		Attribute("startAt", String, "開催日時", func() {
			Example("2016-01-01 10:10:12")
		})
		Attribute("endAt", String, "終了日時", func() {
			Example("2016-01-01 18:00:00")
		})
	})
	Required("ID", "apiID", "title", "url", "limit", "accepted", "waitlisted", "place", "startAt")
	View("default", func() {
		Attribute("ID")
		Attribute("apiID")
		Attribute("title")
		Attribute("url")
		Attribute("limit")
		Attribute("accepted")
		Attribute("waitlisted")
		Attribute("place")
		Attribute("startAt")
		Attribute("endAt")
	})
})

var Genre = MediaType("application/vnd.genre+json", func() {
	Description("ジャンル")
	Attributes(func() {
		Attribute("ID", Integer, "ジャンルID", func() {
			Example(1)
		})
		Attribute("name", String, "ジャンル名", func() {
			Example("javascript")
		})
	})
	View("default", func() {
		Attribute("ID")
		Attribute("name")
	})
})

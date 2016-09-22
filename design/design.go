package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("study-event", func() {
	Title("The study-event API")
	Description("勉強会情報取得")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("keywords", func() {
	Action("search", func() {
		Routing(GET("search/:text/"))
		Description("複数の勉強会サイトをキーワード検索してレスポンスを返す")
		Params(func() {
			Param("text", String, "キーワード")
		})
		Response(OK, "applicateion/json")
	})

})

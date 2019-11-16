package routers

import (
	"github.com/Ehco1996/oh-my-book/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/render", &controllers.BookRender{})
	beego.Router("/api", &controllers.Api{})
}

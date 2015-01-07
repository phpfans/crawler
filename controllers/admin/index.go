package admin

import (
	"github.com/astaxie/beego"
	"os"
	"runtime"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Main() {
	this.Data["hostname"], _ = os.Hostname()
	this.Data["goversion"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH
	this.Data["beegoversion"] = beego.VERSION
	this.Data["version"] = beego.AppConfig.String("version")
	this.TplNames = "admin/main.tpl"
	this.Layout = "admin/layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Sidebar"] = "admin/layout_sidebar.tpl"
}

func (this *IndexController) Info() {
	this.Redirect("/admin/main", 302)
}

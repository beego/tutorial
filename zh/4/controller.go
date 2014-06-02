type AddController struct {
	beego.Controller
}

func (this *AddController) Get() {
	this.Data["content"] = "value"
	this.Layout = "admin/layout.html"
	this.TplNames = "admin/add.tpl"
}

func (this *AddController) Post() {
	pkgname := this.GetString("pkgname")
	content := this.GetString("content")
	pk := models.GetCruPkg(pkgname)
	if pk.Id == 0 {
		var pp models.PkgEntity
		pp.Pid = 0
		pp.Pathname = pkgname
		pp.Intro = pkgname
		models.InsertPkg(pp)
		pk = models.GetCruPkg(pkgname)
	}
	var at models.Article
	at.Pkgid = pk.Id
	at.Content = content
	models.InsertArticle(at)
	this.Ctx.Redirect(302, "/admin/index")
}

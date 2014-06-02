type BaseAdminController struct {
    baseController
}

func (this *BaseAdminController) NestPrepare() {
    if this.CheckActiveRedirect() {
            return
    }
    // if user isn't admin, then logout user
    if !this.user.IsAdmin {
            models.LogoutUser(&this.Controller)

            // write flash message
            this.FlashWrite("NotPermit", "true")

            this.Redirect("/login", 302)
            return
    }
    // current in admin page
    this.Data["IsAdmin"] = true
}

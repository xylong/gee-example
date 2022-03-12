package v1

import (
	"gee-example/api/model"
	"github.com/gin-gonic/gin"
	"github.com/xylong/gee"
	"net/http"
)

type user struct {
	*gee.Gorm
}

func NewUser() *user {
	return &user{}
}

func (u *user) register(ctx *gin.Context) gee.Model {
	return model.NewUser()
}

func (u *user) login(ctx *gin.Context) string {
	return ""
}

func (u *user) logout(ctx *gin.Context) {

}

func (u *user) delete(ctx *gin.Context) {

}

func (u *user) me(ctx *gin.Context) gee.Model {
	return model.NewUser()
}

func (u *user) profile(ctx *gin.Context) gee.Model {
	return model.NewUser()
}

func (u *user) Route(gee *gee.Gee) {
	gee.Handle(http.MethodPost, "register", u.register)
	gee.Handle(http.MethodPost, "login", u.login)
	gee.Handle(http.MethodGet, "logout", u.logout)
	gee.Handle(http.MethodDelete, "delete", u.delete)
	gee.Handle(http.MethodGet, "me", u.me)
	gee.Handle(http.MethodGet, "users/:id", u.profile)
}

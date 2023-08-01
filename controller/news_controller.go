package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-mini/model"
	"go-web-mini/repository"
	"go-web-mini/vo"
)

type ITestController interface {

	CreateNews(c *gin.Context)           // 创建用户

}


//@middleware auth,transition,transition
//@router /create [post]
func (cc *NewsController) CreateNews(c *gin.Context) {

	   var test vo.CreateNewsRequest
	   c.ShouldBind(&test)

	   var m model.News
	   m.Title= test.Title
	   m.Content= test.Content

		cc.TestRepository.CreateNews(c,&m)
}
//@router /create2 [post]

func (cc *NewsController) CreateNews2(c *gin.Context) {

	var test vo.CreateNewsRequest
	c.ShouldBind(&test)

	var m model.News
	m.Title= test.Title
	m.Content= test.Content

	cc.TestRepository.CreateNews(c,&m)
}

//@middleware auth
//@router /api/news [get]
type NewsController struct {
	TestRepository repository.ITestRepository
}

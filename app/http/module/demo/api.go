package demo

import (
	"log"

	demoService "github.com/developbiao/webcore-demo/app/provider/demo"
	"github.com/developbiao/webcore-demo/framework/contract"
	"github.com/developbiao/webcore-demo/framework/gin"
)

type DemoApi struct {
	service *Service
}

// Register
func Register(r *gin.Engine) error {
	api := NewDemoApi()
	r.Bind(&demoService.DemoProvider{})

	r.GET("/demo/demo", api.Demo)
	r.GET("/demo/demo2", api.Demo2)
	r.GET("/demo/demo_post", api.DemoPost)
	return nil
}

// NewDemoApi
func NewDemoApi() *DemoApi {
	service := NewService()
	return &DemoApi{service: service}
}

// Demo
func (api *DemoApi) Demo(c *gin.Context) {
	appService := c.MustMake(contract.AppKey).(contract.App)
	users := api.service.GetUsers()
	usersDTO := UserModelsToUserDTOs(users)
	// c.JSON(200, usersDTO)
	log.Println(usersDTO)

	baseFolder := appService.BaseFolder()
	c.JSON(200, baseFolder)
}

// Demo2
func (api *DemoApi) Demo2(c *gin.Context) {
	demoProvider := c.MustMake(demoService.DemoKey).(demoService.IService)
	students := demoProvider.GetAllStudent()
	usersDTO := StudentsToUserDTOs(students)
	c.JSON(200, usersDTO)
}

// DemoPost
func (api *DemoApi) DemoPost(c *gin.Context) {
	type Foo struct {
		Name string
	}
	foo := &Foo{}
	err := c.BindJSON(&foo)
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, nil)
}

package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

type user struct {
    Id    int         `form:"-"`
    Name  interface{} `form:"username"`
    Age   int         `form:"age"`
    Email string
}


type Todo struct {
    Id   int
    Title string `orm:"size(100)"`
    Isdone bool
}



func (c *MainController) Get() {
	orm.RegisterModel(new(Todo))
	o := orm.NewOrm()

	todo := Todo{Title:"make a shower",Isdone:false}

	id, _ := o.Insert(&todo)

	c.Data["id"] = id

	v := c.GetSession("admin")
	if v == nil {
		c.SetSession("admin",int(1))
		c.Data["num"]=0
	} else {
		c.SetSession("admin",v.(int)+1)
		c.Data["num"] = v.(int)
	}
	c.Data["Website"] = "beego.me2"

	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}


func (this *MainController) Post() {
	
	jsoninfo := this.GetString("data")
	if jsoninfo == "" {
		this.Ctx.WriteString("data is empty")
		return
	}	
	this.DelSession("admin")
	this.Ctx.WriteString("thanks")
}

func (this *MainController) CreateUser() {

	u := user{}
    this.ParseForm(&u)	
  
   
   this.Ctx.Output.Body([]byte(this.Ctx.Input.Param(":ext")))
}

func (this *MainController) Users() {
	this.Abort("501")
	json := "{name:'John'}"
	this.Data["json"]=json
	this.ServeJson()
}
package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/intogosrc/beego-test/models"
	"math"
	"net/http"
	"strconv"
)

type ApiController struct {
	BaseController
}

func (ctrl *ApiController) CurrentUser(ctx *context.Context) {

	user := &models.User{
		Name:   "Serati Ma",
		Avatar: "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
		UserId: "00000001",
	}
	err := ctx.Output.JSON(user, true, true)

	if err != nil {
		ctx.Output.SetStatus(http.StatusInternalServerError)
	}

	//ctrl.Data["json"] = user
	//ctrl.ServeJSON()
}

func (ctrl *ApiController) UserAdd(ctx *context.Context) {

	resp := make(map[string]interface{})

	resp["status"] = "success"
	resp["msg"] = ""

	defer func(){
		err := ctx.Output.JSON(resp, true, true)

		if err != nil {
			ctx.Output.SetStatus(http.StatusInternalServerError)
		}
	}()


	username := ctx.Input.Query("username")
	age := ctx.Input.Query("age")

	if username == "" || age == "" {
		resp["status"] = "failed"
		resp["msg"] = "parameters not enough"
		return
	}

	model := new(models.UserModel)

	model.Name = username
	var e error
	model.Age,e = strconv.ParseInt(age, 10, 64)

	if e!=nil {
		resp["status"] = "failed"
		resp["msg"] = "wrong format AGE "
		return
	}

	if !model.Save() {
		resp["status"] = "failed"
		resp["msg"] = "save failed"
		return
	}
}

func (ctrl *ApiController) UserList2(ctx *context.Context) {
	pageStr := ctx.Input.Query("page")
	limitStr := ctx.Input.Query("limit")

	logs.Debug("recv page params page %s limit %s", pageStr, limitStr)

	var page int = 1
	var limit int = 10

	p,e := strconv.Atoi(pageStr)

	if e==nil && p>0 {
		page = p
	}

	l,e := strconv.Atoi(limitStr)

	if e==nil && l>0 {
		limit = l
	}

	type L struct {
		Id   int32  `json:"id"`
		Name string `json:"name"`
		Age  int32  `json:"age"`
	}

	model := models.UserModel{}

	//result := make([]*L, 0)
	total := int(model.Count())



	//for i := 0; i < total; i++ {
	//	result = append(result, &L{
	//		Id:   int32(i),
	//		Name: "Leo",
	//		Age:  10,
	//	})
	//}

	logs.Debug("%d/%d", total, limit)

	pageTotal := int(int64(math.Ceil(float64(total)/float64(limit))))
	if page > pageTotal {
		page = pageTotal
	}

	if page<1 {
		page=1
	}

	logs.Debug("recv page params page %v limit %v", page, limit)

	start := (page-1)*limit
	end := start+limit

	logs.Debug("start %v end %v", start, end)

	if end>total {
		end = total
	}

	users := model.GetList(start, limit)

	resp := map[string]interface{} {
		"users": users ,
		"page_info":map[string]int{
			"current": page,
			"limit": limit,
			"amount": total ,
		} ,
	}

	err := ctx.Output.JSON(resp, true, true)

	if err != nil {
		ctx.Output.SetStatus(http.StatusInternalServerError)
	}


}

func (ctrl *ApiController) UserList(ctx *context.Context) {

	pageStr := ctx.Input.Query("page")
	limitStr := ctx.Input.Query("limit")

	logs.Debug("recv page params page %s limit %s", pageStr, limitStr)

	var page int = 1
	var limit int = 10

	p,e := strconv.Atoi(pageStr)

	if e==nil && p>0 {
		page = p
	}

	l,e := strconv.Atoi(limitStr)

	if e==nil && l>0 {
		limit = l
	}

	type L struct {
		Id   int32  `json:"id"`
		Name string `json:"name"`
		Age  int32  `json:"age"`
	}

	result := make([]*L, 0)
	total := 40

	for i := 0; i < total; i++ {
		result = append(result, &L{
			Id:   int32(i),
			Name: "Leo",
			Age:  10,
		})
	}

	logs.Debug("%d/%d", total, limit)

	pageTotal := int(int64(math.Ceil(float64(total)/float64(limit))))
	if page > pageTotal {
		page = pageTotal
	}

	if page<1 {
		page=1
	}

	logs.Debug("recv page params page %v limit %v", page, limit)

	start := (page-1)*limit
	end := start+limit

	logs.Debug("start %v end %v", start, end)

	if end>total {
		end = total
	}

	resp := map[string]interface{} {
		"users": result[start:end] ,
		"page_info":map[string]int{
			"current": page,
			"limit": limit,
			"amount": total ,
		} ,
	}

	err := ctx.Output.JSON(resp, true, true)

	if err != nil {
		ctx.Output.SetStatus(http.StatusInternalServerError)
	}
}

func (ctrl *ApiController) Users(ctx *context.Context) {

	result := []*models.User{
		{
			Key:     "1",
			Name:    "John Brown",
			Age:     32,
			Address: "New York No. 1 Lake Park",
		},
		{
			Key:     "2",
			Name:    "Jim Green",
			Age:     42,
			Address: "London No. 1 Lake Park",
		},
	}

	err := ctx.Output.JSON(result, true, true)

	if err != nil {
		ctx.Output.SetStatus(http.StatusInternalServerError)
	}

	//ctrl.Data["json"] = result
	//ctrl.ServeJSON()
}

func (ctrl *ApiController) Login(ctx *context.Context) {

	userName := ""
	loginType := ""
	password := ""

	for k, v := range ctx.Input.Data() {
		kn, ok := k.(string)
		if !ok {
			continue
		}

		vn, ok := v.(string)
		if !ok {
			continue
		}

		switch kn {
		case "userName":
			userName = vn
		case "password":
			password = vn
		case "type":
			loginType = vn
		}
	} // todo 中间件可以抽象成方法

	result := new(models.UserLogin)

	defer func() {
		err := ctx.Output.JSON(result, true, true)

		if err != nil {
			ctx.Output.SetStatus(http.StatusInternalServerError)
		}
	}()

	if password == "ant.design" && userName == "admin" {
		result.Status = "ok"
		result.Type = loginType
		result.CurrentAuthority = "admin"

		return
	}

	if password == "ant.design" && userName == "user" {
		result.Status = "ok"
		result.Type = loginType
		result.CurrentAuthority = "user"

		return
	}

	if loginType == "mobile" {
		result.Status = "ok"
		result.Type = loginType
		result.CurrentAuthority = "admin"

		return
	}

	result.Status = "error"
	result.Type = loginType
	result.CurrentAuthority = "guest"
}

func (ctrl *ApiController) GoodsList(ctx *context.Context) {
	resultList := []*models.User{
		{
			Key:     "1",
			Name:    "John Lee",
			Age:     32,
			Address: "New York No. 1000 Lake Park",
		},
		{
			Key:     "2",
			Name:    "Jim Blue",
			Age:     42,
			Address: "London No. 2011 Lake Park",
		},
	}

	result := map[string]interface{}{
		"list": resultList,
	}

	err := ctx.Output.JSON(result, true, true)

	if err != nil {
		ctx.Output.SetStatus(http.StatusInternalServerError)
	}
}

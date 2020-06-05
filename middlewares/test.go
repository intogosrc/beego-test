package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"strings"
)

func Test (ctx *context.Context) {
	fmt.Println("test middleware is invoked")

	username := ctx.Input.Query("username") // 从 Request 获取值

	ctx.Input.SetData("test_param", "tttttt "+username) // 可以设置值

	// todo 检验用户权限
}

func ParseJsonBody (ctx *context.Context) {
	fmt.Println("ParseJsonBody middleware is invoked")

	contentType := ctx.Request.Header.Get("Content-Type")
	fmt.Println("content-type", contentType)

	if strings.HasPrefix(contentType, "application/json") {
		content,_ := ioutil.ReadAll(ctx.Request.Body)

		container := make(map[string]interface{})

		err := json.Unmarshal(content, &container)

		if err!=nil {
			logs.Error("ParseJsonBody middleware exec error %s", err.Error())
			return
		}

		fmt.Println(string(content), container)

		for k,v := range container {
			ctx.Input.SetData(k,v)
		}
	}

}
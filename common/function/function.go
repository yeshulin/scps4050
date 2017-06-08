package function

import (
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego/context"
)

func ConvertT(in int64) (out string) {
	tm := time.Unix(in, 0)
	out = tm.Format("2006-01-02 15:04:05")
	return out
}
func TransparentStatic(ctx *context.Context) {
	if strings.Index(ctx.Request.URL.Path, "api/") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "path/"+ctx.Request.URL.Path)
}

package routers

import (
	"webproject/4050/controllers"
	"webproject/4050/controllers/admin"
	"webproject/4050/controllers/api"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &admin.UserController{}, "get:Login")
	beego.Router("/admin/logout", &admin.UserController{}, "get:LogOut")
	beego.Router("/user/login", &admin.UserController{}, "post:Post")
	beego.Router("/user/reg", &admin.UserController{}, "post:Reg")
	beego.Router("/user/regweixin", &admin.UserController{}, "post:RegWeixin")
	beego.Router("/user/apply", &admin.UserController{}, "post:Apply")
	beego.Router("/user/verify", &admin.UserController{}, "post:Verify")
	beego.Router("/user/getapply", &admin.UserController{}, "get:GetApply")
	beego.Router("/register", &admin.UserController{}, "get:Register")
	beego.Router("/admin", &admin.IndexController{})
	beego.Router("/admin/user/find", &admin.UserController{}, "get:Get")
	beego.Router("/admin/userlist", &admin.UserController{}, "get:UserList")
	beego.Router("/admin/user/add", &admin.UserController{}, "get:Add;post:Add")
	beego.Router("/admin/user/addpost", &admin.UserController{}, "post:AddPost")
	beego.Router("/admin/user/delete", &admin.UserController{}, "post:Delete")
	beego.Router("/admin/user/view", &admin.UserController{}, "get:View")
	beego.Router("/admin/user/edit", &admin.UserController{}, "get:Edit")
	beego.Router("/admin/user/editpost", &admin.UserController{}, "post:EditPost")
	beego.Router("/admin/role", &admin.RoleController{}, "get:RoleList")
	beego.Router("/admin/role/find", &admin.RoleController{}, "get:Get")
	beego.Router("/admin/role/add", &admin.RoleController{}, "get:Add")
	beego.Router("/admin/role/addpost", &admin.RoleController{}, "post:AddPost")
	beego.Router("/admin/role/delete", &admin.RoleController{}, "post:Delete")
	beego.Router("/admin/role/view", &admin.RoleController{}, "get:View")
	beego.Router("/admin/role/edit", &admin.RoleController{}, "get:Edit")
	beego.Router("/admin/role/editpost", &admin.RoleController{}, "post:EditPost")
	beego.Router("/admin/role/user", &admin.RoleController{}, "get:User")
	beego.Router("/admin/role/userfind", &admin.RoleController{}, "get:UserFind")
	beego.Router("/admin/role/deluser", &admin.RoleController{}, "post:DelUser")
	beego.Router("/admin/role/adduser", &admin.RoleController{}, "post:AddUser")
	beego.Router("/admin/role/node", &admin.RoleController{}, "get:Node")
	beego.Router("/admin/role/nodefind", &admin.RoleController{}, "get:NodeFind")
	beego.Router("/admin/role/delnode", &admin.RoleController{}, "post:DelNode")
	beego.Router("/admin/role/addnode", &admin.RoleController{}, "post:AddNode")
	beego.Router("/admin/node", &admin.NodeController{}, "get:NodeList")
	beego.Router("/admin/node/find", &admin.NodeController{}, "get:Get")
	beego.Router("/admin/node/add", &admin.NodeController{}, "get:Add")
	beego.Router("/admin/node/addpost", &admin.NodeController{}, "post:AddPost")
	beego.Router("/admin/node/delete", &admin.NodeController{}, "post:Delete")
	beego.Router("/admin/node/view", &admin.NodeController{}, "get:View")
	beego.Router("/admin/node/edit", &admin.NodeController{}, "get:Edit")
	beego.Router("/admin/node/editpost", &admin.NodeController{}, "post:EditPost")
	beego.Router("/admin/newstype", &admin.NewsTypeController{}, "get:NewsTypeList")
	beego.Router("/admin/newstype/find", &admin.NewsTypeController{}, "get:Get")
	beego.Router("/admin/newstype/add", &admin.NewsTypeController{}, "get:Add")
	beego.Router("/admin/newstype/addpost", &admin.NewsTypeController{}, "post:AddPost")
	beego.Router("/admin/newstype/delete", &admin.NewsTypeController{}, "post:Delete")
	beego.Router("/admin/newstype/view", &admin.NewsTypeController{}, "get:View")
	beego.Router("/admin/newstype/edit", &admin.NewsTypeController{}, "get:Edit")
	beego.Router("/admin/newstype/editpost", &admin.NewsTypeController{}, "post:EditPost")
	beego.Router("/admin/news", &admin.NewsController{}, "get:NewsList")
	beego.Router("/admin/news/find", &admin.NewsController{}, "get:Get")
	beego.Router("/admin/news/add", &admin.NewsController{}, "get:Add")
	beego.Router("/admin/news/addpost", &admin.NewsController{}, "post:AddPost")
	beego.Router("/admin/news/delete", &admin.NewsController{}, "post:Delete")
	beego.Router("/admin/news/view", &admin.NewsController{}, "get:View")
	beego.Router("/admin/news/edit", &admin.NewsController{}, "get:Edit")
	beego.Router("/admin/news/editpost", &admin.NewsController{}, "post:EditPost")
	beego.Router("/admin/menu", &admin.MenuController{})
	beego.Router("/admin/news", &admin.NewsController{})
	beego.Router("/admin/newstype", &admin.NewsTypeController{})
	beego.Router("/admin/course", &admin.CourseController{})
	beego.Router("/admin/coursetype", &admin.CourseTypeController{})
	beego.Router("/admin/verify/find", &admin.VerifyController{}, "get:Get")
	beego.Router("/admin/verify", &admin.VerifyController{}, "get:VerifyList")
	beego.Router("/admin/verify/pass", &admin.VerifyController{}, "Post:Pass")
	beego.Router("/admin/verify/reject", &admin.VerifyController{}, "Post:Reject")
	beego.Router("/admin/verify/setremark", &admin.VerifyController{}, "Post:Setremark")
	beego.Router("/admin/verify/tongji", &admin.VerifyController{}, "get:Tongji")
	beego.Router("/admin/verify/view", &admin.VerifyController{}, "get:View")
	beego.Router("/admin/verify/applys", &admin.VerifyController{}, "get:Applys")
	beego.Router("/admin/verify/passapplys", &admin.VerifyController{}, "Post:Passapplys")
	beego.Router("/admin/verify/rejectapplys", &admin.VerifyController{}, "Post:Rejectapplys")
	beego.Router("/admin/verify/import", &admin.VerifyController{}, "Get:Import")
	beego.Router("/admin/verify/importpost", &admin.VerifyController{}, "Post:Importpost")
	beego.Router("/admin/signs/find", &admin.SignsController{}, "get:Get")
	beego.Router("/admin/signs", &admin.SignsController{}, "get:SignsList")
	beego.Router("/admin/signs/pass", &admin.SignsController{}, "Post:Pass")
	beego.Router("/admin/signs/reject", &admin.SignsController{}, "Post:Reject")
	beego.Router("/admin/signs/view", &admin.SignsController{}, "get:View")
	beego.Router("/admin/index/welcome", &admin.IndexController{}, "get:Welcome")

	/*接口路由*/
	beego.Router("/file/upload", &controllers.FileController{}, "post:Upload")
	beego.Router("/file/uploadpage", &controllers.FileController{}, "get:UploadPage")
	beego.Router("/api/signs", &api.SignsController{})
	beego.Router("/api/signs/list", &api.SignsController{}, "get:List")
	beego.Router("/api/zones", &api.ZonesController{})
	beego.Router("/api/applys/list", &api.ApplysController{}, "get:List")
	beego.Router("/api/users", &api.UsersController{})
	/*oss路由*/
	beego.Router("/oss/webupload", &controllers.OssController{}, "get:WebUpload")
}

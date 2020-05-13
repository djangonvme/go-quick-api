package params

type Login struct {
	Mobile string `form:"mobile" binding:"required"`
	Pwd    string `form:"pwd" binding:"required"`
}

// form, binding 是gin的验证规则
// more validate rules refer: https://godoc.org/gopkg.in/go-playground/validator.v8
type SearchUserList struct {
	//Mobile string `form:"mobile" binding:"required"`
	Mobile string `form:"mobile"`
	//Page   int64  `form:"page" binding:"required,gt=0"`
	Page int64 `form:"page"`
}

//
type AddUser struct {
	Mobile string `form:"mobile" binding:"required,len=11"`
	Pwd    string `form:"pwd" binding:"required,min=6,max=32"` //6<= len(pwd) <=32
	Name   string `form:"name" binding:""`
}

type UserItem struct {
	Id     int64  `json:"id"`
	Mobile string `json:"mobile"`
	Name   string `json:"name"`
}
type UserList struct {
	Total    int64      `json:"total"`
	PageSize int64      `json:"page_size"`
	List     []UserItem `json:"list"`
}

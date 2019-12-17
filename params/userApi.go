package params

type UserListSearch struct {
	Mobile string `json:"mobile"`
	Page   int64  `json:"page"`
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

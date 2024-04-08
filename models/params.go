package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

//用来定义请求的参数结构体

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票参数
type ParamVoteData struct {
	//UserId  从请求中获取当前用户
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //投票类型，赞成(1)，反对(-1),取消投票(0)
}

// ParamPostList 获取帖子列表请求参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"` // 可以为空
	Page        int64  `json:"page" form:"page"`
	Size        int64  `json:"size" form:"size"`
	Order       string `json:"order" form:"order"`
}

//// ParamCommunityPostList 按照社区获取帖子列表请求参数
//type ParamCommunityPostList struct {
//	*ParamPostList
//	CommunityID int64 `json:"community_id" form:"community_id"`
//}

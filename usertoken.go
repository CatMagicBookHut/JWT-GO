package jwt

import (
	"NyaLog/gin-blog-server/utils/errmsg"
)

var UserTokenMap map[string]string

// Init Hash table
// 初始化哈希表
func init() {
	UserTokenMap = make(map[string]string)
}

// Save uid and token to Hash table
// 登陆时存入uid和对应的token
func UserLogin(uid string, token string) int {
	UserTokenMap[uid] = token
	return errmsg.SUCCESS
}

// Get token by uid
// 提取uid对应的token
func GetToken(uid string) string {
	return UserTokenMap[uid]
}

// Delete token by uid
// 删除uid对应的token
func DeleteToken(uid string) int {
	delete(UserTokenMap, uid)
	return errmsg.SUCCESS
}

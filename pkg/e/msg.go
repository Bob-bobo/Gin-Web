package e

/**
此处将标签改为——————活动

*/
var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_USERNAME:                  "用户名错误",
	ERROR_EXIST_CATEGO:              "已存在该活动名称",
	ERROR_EXIST_CATEGO_FAIL:         "获取已存在活动失败",
	ERROR_NOT_EXIST_CATEGO:          "该活动不存在",
	ERROR_GET_CATEGOS_FAIL:          "获取所有活动失败",
	ERROR_COUNT_CATEGO_FAIL:         "统计活动失败",
	ERROR_ADD_CATEGO_FAIL:           "新增活动失败",
	ERROR_EDIT_CATEGO_FAIL:          "修改活动失败",
	ERROR_DELETE_CATEGO_FAIL:        "删除活动失败",
	ERROR_CHECK_EXIST_CATEGORY_FAIL: "检查分类失败",
	ERROR_IMPORT_CATEGO_FAIL:        "导入活动失败",
	ERROR_NOT_EXIST_ACTIVITY:        "该活动不存在",
	ERROR_ADD_ACTIVITY_FAIL:         "新增活动失败",
	ERROR_DELETE_ACTIVITY_FAIL:      "删除活动失败",
	ERROR_CHECK_EXIST_ACTIVITY_FAIL: "检查活动是否存在失败",
	ERROR_EDIT_ACTIVITY_FAIL:        "修改活动失败",
	ERROR_COUNT_ACTIVITY_FAIL:       "统计活动失败",
	ERROR_GET_ACTIVITIES_FAIL:       "获取多个活动失败",
	ERROR_GET_ACTIVITY_FAIL:         "获取单个活动失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_ADD_COMMENT_FAIL:          "添加评论失败",
	ERROR_JOIN_ACTIVITY_FAIL:        "参加活动失败",
	ERROR_NOT_EXIST_ACTUSER:         "用户未参加该活动",
	ERROR_DELETE_ACTUSER_FAIL:       "退出活动失败",
	ERROR_CHECK_INFO_FALI:           "查看个人信息失败",
	ERROR_GET_USER_FALI:             "查看所有用户失败",
	ERROR_GET_CALLIGRAPHY_FAIL:      "查询所有书法失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

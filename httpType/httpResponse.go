package httpType

type Response struct {
	Code int64       `json:"code"` // 错误码，非 0 表示失败
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"` // 错误描述
}

type (
	AccessTokenResp struct {
		AccessToken      string `json:"access_token"`       // user_access_token，用于获取用户资源
		AvatarBig        string `json:"avatar_big"`         // 用户头像 640x640
		AvatarMiddle     string `json:"avatar_middle"`      // 用户头像 240x240
		AvatarThumb      string `json:"avatar_thumb"`       // 用户头像 72x72
		AvatarURL        string `json:"avatar_url"`         // 用户头像
		Email            string `json:"email"`              // 用户邮箱，**字段权限要求**：获取用户邮箱信息
		EnName           string `json:"en_name"`            // 用户英文名称
		EnterpriseEmail  string `json:"enterprise_email"`   // 企业邮箱，请先确保已在管理后台启用飞书邮箱服务。**字段权限要求**：获取用户受雇信息
		ExpiresIn        int64  `json:"expires_in"`         // access_token 的有效期，单位: 秒
		Mobile           string `json:"mobile"`             // 用户手机号。**字段权限要求**：获取用户手机号
		Name             string `json:"name"`               // 用户姓名
		OpenID           string `json:"open_id"`            // 用户在应用内的唯一标识
		RefreshExpiresIn int64  `json:"refresh_expires_in"` // refresh_token 的有效期，单位: 秒
		RefreshToken     string `json:"refresh_token"`      // 刷新用户 access_token 时使用的 token
		Sid              string `json:"sid"`                // 用户当前登录态session的唯一标识，为空则不返回
		TenantKey        string `json:"tenant_key"`         // 当前企业标识
		TokenType        string `json:"token_type"`         // token 类型
		UnionID          string `json:"union_id"`           // 用户统一ID
		UserID           string `json:"user_id"`            // 用户 user_id。**字段权限要求**：获取用户 user ID
	}
)

type (
	FreshAccessTokenResp struct {
		AccessToken      string `json:"access_token"`       // user_access_token，用于获取用户资源
		AvatarBig        string `json:"avatar_big"`         // 用户头像 640x640
		AvatarMiddle     string `json:"avatar_middle"`      // 用户头像 240x240
		AvatarThumb      string `json:"avatar_thumb"`       // 用户头像 72x72
		AvatarURL        string `json:"avatar_url"`         // 用户头像
		Email            string `json:"email"`              // 用户邮箱。**字段权限要求**：获取用户邮箱信息
		EnName           string `json:"en_name"`            // 用户英文名称
		EnterpriseEmail  string `json:"enterprise_email"`   // 企业邮箱，请先确保已在管理后台启用飞书邮箱服务。**字段权限要求**：获取用户受雇信息
		ExpiresIn        int64  `json:"expires_in"`         // access_token 的有效期，单位: 秒
		Mobile           string `json:"mobile"`             // 用户手机号。**字段权限要求**：获取用户手机号
		Name             string `json:"name"`               // 用户姓名
		OpenID           string `json:"open_id"`            // 用户在应用内的唯一标识
		RefreshExpiresIn int64  `json:"refresh_expires_in"` // refresh_token 的有效期，单位: 秒
		RefreshToken     string `json:"refresh_token"`      // 刷新用户 access_token 时使用的 token
		Sid              string `json:"sid"`                // 用户当前登录态session的唯一标识，为空则不返回
		TenantKey        string `json:"tenant_key"`         // 当前企业标识
		TokenType        string `json:"token_type"`         // token 类型
		UnionID          string `json:"union_id"`           // 用户统一ID
		UserID           string `json:"user_id"`            // 用户 user_id。**字段权限要求**：获取用户 user ID
	}
	FreshAccessTokenRequest struct {
		RefreshToken string `json:"refresh_token,omitempty"` // 刷新 user_access_token 需要的凭证 获取user_access_token接口和本接口均返回; refresh_token，每次请求，请注意使用最新获取到的refresh_token; 示例值："ur-oQ0mMq6MCcueAv0pwx2fQQhxqv__CbLu6G8ySFwafeKww2Def2BJdOkW3.9gCFM.LBQgFri901QaqeuL"
	}
)

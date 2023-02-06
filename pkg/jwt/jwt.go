package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const TokenExpireDuration = time.Hour * 24 // 过期时间

// CustomSecret 用于加盐的字符串
var CustomSecret = []byte("嘿咻嘿咻")

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个UserID字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	UserID               int64  `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 内嵌标准的声明
}

// GenToken 生成JWT
func GenToken(username string, userId int64) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		userId, // 自定义字段
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)), // 设置过期时间
			Issuer:    "dogTailBlog",                                           // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(CustomSecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	cc := new(CustomClaims)
	// / 解析token 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法, 这里就直接解析到变量 cc 里面了
	token, err := jwt.ParseWithClaims(tokenString, cc, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	//if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
	//	return claims, nil
	//}
	if token.Valid {
		return cc, nil
	}
	return nil, errors.New("invalid token")
}

/*
上面是获取Access token, 访问资源时候的token，此外我们还可以设置另外一个种token, refresh token
refresh token过期时间较长，refreshToken就是用来在accessToken过期以后来重新获取accessToken的
例如:
accessToken 1天过期
refreshToken 30天过期

使用流程
	客户端 --> 登录 --> 服务端
	客户端 <-- 返回accessToken和refreshToken --< 服务端
	客户端 --> 携带accessToken 请求业务接口使用 --> 服务端
	客户端 <-- 返回数据 --< 服务端
	客户端 --> accessToken 过期, 携带 refreshToken 访问刷新接口 --> 服务端
	客户端 <-- 返回新的 accessToken --< 服务端
	客户端 --> 携带新的 accessToken 请求业务接口使用 --> 服务端
	客户端 <-- 返回数据 --< 服务端
	每次访问刷新接口的时候,同时刷新 refreshToken, 则用户会一直处于登录状态

-> 1.登录成功获得 refresh token 并持久化
-> 2.通过 refresh token 请求刷新得到 access token 并临时储存
-> 3.请求业务接口使用 access token
-> 4.access token 过期或者快过期再次回到「 2 」
-> 5.refresh token 也过期则生命周期结束，需重新登录

为啥用 refreshToken
	1.因为安全问题，防止token盗用, accessToken有效时间段，使用refreshToken可以长时间保持登录状态
	2.accessToken 过期后，使用 refreshToken 需要读取额外的状态来确认是否继续发token (额外的状态可以存数据库或缓存)
	3.refreshToken 只有在第一次获取和刷新access token时才会在网络中传输, 被盗风险小一点
*/

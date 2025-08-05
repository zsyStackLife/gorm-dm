package dameng

import (
	"net/url"
	"strings"
)

// DriverName 数据库驱动、连接字符串协议名称
const DriverName = "dm"

// BuildUrl 构建达梦数据库连接字符串
//
//	 url:[scheme:][//[userinfo@]host][/]path[?query][#fragment]
//	- 如： dm://user:password@host:port?schema=SYSDBA[&propName2=propValue2]…
//	- 若要指定用户登录后的当前模式，请在 options 中设置 schema，缺省为用户的默认模式，如 SYSDBA
//	- 参考链接： https://eco.dameng.com/document/dm/zh-cn/pm/go-rogramming-guide.html#11.3%20%E8%BF%9E%E6%8E%A5%E4%B8%B2%E5%B1%9E%E6%80%A7%E8%AF%B4%E6%98%8E
func BuildDMUrl(user, password, dns string, urlOptions ...map[string]string) string {
	var b = strings.Builder{}
	b.WriteString(DriverName)
	b.WriteByte(':')
	if dns != "" || user != "" {
		b.WriteString("//")
	}
	if user != "" {
		b.WriteString(user)
	}
	if user != "" && password != "" {
		b.WriteByte(':')
		b.WriteString(password)
		b.WriteByte('@')
	}
	if dns != "" {
		b.WriteString(dns)
	}
	var propQuery = url.Values{}
	for _, option := range urlOptions {
		for key, value := range option {
			propQuery.Add(key, value)
		}
	}
	if propQuery.Encode() != "" {
		b.WriteByte('?')
		b.WriteString(propQuery.Encode())
	}

	return b.String()
}

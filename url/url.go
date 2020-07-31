package url

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// 拼接url路径
func UrlJoin(parts ...string) string {
	l := len(parts)
	if l == 1 {
		return parts[0]
	}
	ps := make([]string, l)
	for i, part := range parts {
		if part == "" {
			continue
		}
		if i == 0 {
			ps[i] = strings.TrimRight(part, "/")
		} else {
			ps[i] = strings.TrimLeft(part, "/")
		}
	}
	return strings.Join(ps, "/")
}

// 将传入的url补充成完整的url, urlObj必须是完整的url obj, 类似: http://baidu.com
func ParseUrl(urlObj *url.URL, u string) (*url.URL, error) {
	if u == "" {
		return nil, errors.New("param 'u' is empty")
	}
	if !strings.HasPrefix(u, "http") && !strings.HasPrefix(u, "/") {
		// like <link rel="apple-touch-icon-precomposed" sizes="152x152" href="gitbook/images/apple-touch-icon-precomposed-152.png">
		if urlObj.Path == "" {
			u = UrlJoin(fmt.Sprintf("%s://%s", urlObj.Scheme, urlObj.Host), u)
		} else {
			u = UrlJoin(fmt.Sprintf("%s://%s", urlObj.Scheme, urlObj.Host), urlObj.Path, u)
		}
	}
	if strings.HasPrefix(u, "//") { // 缺少了https
		// like //github.com/image/xxxxxxxxx
		u = fmt.Sprintf("%s:%s", urlObj.Scheme, u)
	}

	if !strings.HasPrefix(u, urlObj.Scheme) && strings.HasPrefix(u, "/") {
		// like /image/xxxxx
		u = fmt.Sprintf("%s://%s%s", urlObj.Scheme, urlObj.Host, u)
	}
	return url.Parse(u)
}

// 获取二级域名, 输入www.baidu.com,返回baidu.com
func GetDomain(domain string) string {
	if len(domain) == 0 {
		return domain
	}
	// port
	name := strings.Split(domain, ":")[0]
	domains := strings.Split(name, ".")
	if len(domains) <= 1 {
		return domain
	}
	if len(domains) == 2 {
		return strings.Join(domains, ".")
	}
	return strings.Join(domains[len(domains)-2:], ".")
}

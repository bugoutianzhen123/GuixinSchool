package service

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const (
	CookieKey1 = "ASP.NET_SessionId"
	CookieKey2 = "JSESSIONID"
	GinStuIDKey = "stuID"
)

type AuthSvc struct{}

func (a *AuthSvc) Login(ctx context.Context,stuID,password string) error {
	client, infos, err := a.getNecessaryInfo(ctx)
	if err != nil {
		return fmt.Errorf("get necessary info failed: %w", err)
	}

	// 登录
	err = a.login(ctx, client, stuID, password, infos["lt"], infos["execution"])
	if err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	log.Println("Login successful")
	return nil
}

var jwtSecret = []byte("GuiBao") 

func (a *AuthSvc) GetToken(ctx context.Context, stuID string) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		GinStuIDKey: stuID,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(), // Token expires in 30 days
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func (a *AuthSvc) verifyToken(ctx context.Context, tokenString string) (string, error) {
	// Parse and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC and matches
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token: %w", err)
	}

	// Extract claims and return stuID
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if stuID, ok := claims["stuID"].(string); ok {
			return stuID, nil
		}
		return "", fmt.Errorf("stuID not found in token claims")
	}

	return "", fmt.Errorf("invalid token claims")
}




func (a *AuthSvc) JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			return
		}

		// 检查是否为 Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization format"})
			return
		}

		tokenStr := parts[1]

		// 验证 token
		stuID, err := a.verifyToken(c.Request.Context(), tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// 将 stuID 保存到 gin.Context
		c.Set(GinStuIDKey, stuID)

		c.Next() // 继续处理后续请求
	}
}



func (a *AuthSvc) getNecessaryInfo(ctx context.Context) (*http.Client, map[string]string, error) {
	infos := make(map[string]string)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Println("Redirected to:", req.URL)
			return nil // 允许重定向，模拟浏览器自动跳转
		},
		Transport: tr,
	}

	req, err := http.NewRequestWithContext(ctx, "GET", "http://kjyy.ccnu.edu.cn/clientweb/xcus/ic2/Default.aspx?version=3.00.20181109", nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()
	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	lt, _ := doc.Find("input[name='lt']").Attr("value")
	execution, _ := doc.Find("input[name='execution']").Attr("value")
	if lt == "" || execution == "" {
		return nil, nil, fmt.Errorf("failed to find lt or execution in the response")
	}

	infos["lt"] = lt
	infos["execution"] = execution

	domains := []string{
		"http://kjyy.ccnu.edu.cn",
		"https://account.ccnu.edu.cn/cas",
	}

	var getCookieKey1, getCookieKey2 bool

	for _, domain := range domains {
		rootURL, _ := url.Parse(domain)
		for _, cookie := range jar.Cookies(rootURL) {
			if cookie.Name == CookieKey1 {
				getCookieKey1 = true
				infos[cookie.Name] = cookie.Value
			}
			if cookie.Name == CookieKey2 {
				getCookieKey2 = true
				infos[cookie.Name] = cookie.Value
			}
			fmt.Println("Cookie:", cookie.Name, "Value:", cookie.Value, "Domain:", cookie.Domain)
		}
	}

	if !getCookieKey1 || !getCookieKey2 {
		return nil, nil, fmt.Errorf("failed to get cookies, expected 2 cookies")
	}

	log.Printf("necessary info: %+v\n", infos)

	return client, infos, nil
}

func (a *AuthSvc) login(ctx context.Context, client *http.Client, stuID, pwd, lt, execution string) error {
	var redirected bool

	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirected = true
		return nil
	}

	//登录
	form := url.Values{
		"username":  {stuID},
		"password":  {pwd},
		"lt":        {lt},
		"execution": {execution},
		"_eventId":  {"submit"},
		"submit":    {"登录"},
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=", strings.NewReader(form.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create login request: %w", err)
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://account.ccnu.edu.cn")
	req.Header.Set("Referer", "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="137", "Chromium";v="137", "Not/A)Brand";v="24"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)

	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("send request failed: %w", err)
	}

	// 检查是否重定向
	// 如果没有，则代表失败
	if !redirected {
		return fmt.Errorf("login did not redirect, check your stuID and password")
	}

	return nil
}

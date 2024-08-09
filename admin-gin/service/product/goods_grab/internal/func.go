package internal

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

func UrlInitByCategoryID(page int, categoryID uint) (string, error) {
	// 获取当前时间
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	// by=bottom_price  sort=ASC
	urlStr := "https://e.reduxingxuan.com/open/offical/product/index-search-v2?" +
		"kw=&cm=&pp0=&pp1=&f0=&f1=&pf=&" +
		"by=id&" +
		"xxbp=0&xxgy=0&jrsx=0&xsjl=0&" +
		"sort=&" +
		"platform=&" +
		"type=&" +
		"page=" + strconv.Itoa(page) + "&" +
		"member_id=&" +
		"product_category_id=&" +
		"product_category_two_id=&" +
		"product_category_three_id=" + strconv.Itoa(int(categoryID)) + "&" +
		"is_guarantee=0&" +
		"product_tag_ids=&" +
		"is_displayreward=&" +
		"is_newyeartd=&" +
		"timestamp=" + strconv.FormatInt(timestamp, 10) + "&sign_type=md5"

	// 解析URL
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	// 获取查询参数部分
	queryParams := u.Query()
	// 删除名为"sign"的参数
	delete(queryParams, "sign")
	// 将参数按照字母顺序排序
	keys := make([]string, 0, len(queryParams))
	for key := range queryParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	// 构建排序后的查询参数字符串
	var sortedParams []string
	for _, key := range keys {
		value := queryParams[key]
		for _, v := range value {
			sortedParams = append(sortedParams, fmt.Sprintf("%s=%s", key, v))
		}
	}
	a := strings.Join(sortedParams, "&")
	// 构建新的查询参数字符串
	newQueryString := strings.Join(sortedParams, "&") + "&secret=EC7259DFCAD1715C23D400B7A88407A3"
	hash := md5.Sum([]byte(newQueryString))
	// 将哈希值转换为十六进制格式的字符串，并转换为大写形式
	hashStr := strings.ToUpper(hex.EncodeToString(hash[:]))
	u.RawQuery = a + "&sign=" + hashStr
	return u.String(), nil
}

func SendHttp(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	// 设置请求头
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36")
	// 发送 HTTP 请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	reader := resp.Body
	// 读取响应体数据
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

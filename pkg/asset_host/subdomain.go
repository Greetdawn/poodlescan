// 该功能主要实现子域探测功能
// 功能一：基于fofa的子域收集
// 功能二：基于subfinder的子域接口扫描

package asset_host

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/projectdiscovery/subfinder/v2/pkg/passive"
	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

// 定义所需参数
var (
	FOFA_EMAIL      string = "cannopznjooss@gmail.com"
	FOFA_KEY        string = "e0a8689fd9ee6f8249a57ad7230a7065"
	Page            string = "1"
	Size            string
	Full            string       = "true"
	respjson        ResponseJson //定义json串结构体解析
	subdomianResult []string     //定义结果存放切片
)

// 定义json返回结构体
type ResponseJson struct {
	Error   bool       `json:"error"`
	Size    int        `json:"size"`
	Page    int        `json:"page"`
	Mode    string     `json:"mode"`
	Query   string     `json:"query"`
	Results [][]string `json:"results"`
}

// 调用子域收集功能
func ScanSubDomain(domain string) []string {
	return _subfinder(domain)
}

// 基于fofa接口实现子域数据获取
// FOFA_EMAIL=xxx@gmail.com
// FOFA_KEY=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
// https://fofa.info/api/v1/search/all?email={}&key={}&page={}&qbase64={}
// "email", FOFA_EMAIL 用户的fofa账号
// "key", FOFA_KEY 用户的fofa key
// "page", page 查询返回结果的页面
// "qbase64", domain 传入需要查询的域名
func _getSubdomainByFofa(domain string) []string {
	// 生成fofa接口固定域名搜过格式 domain="baidu.com" 进过base64编码操作
	domain = base64.StdEncoding.EncodeToString([]byte("domain=" + "\"" + domain + "\""))

	// 定义get传参变量
	params := url.Values{}
	params.Set("email", FOFA_EMAIL)
	params.Set("key", FOFA_KEY)
	params.Set("page", Page)
	params.Set("size", Size)
	params.Set("full", Full)
	params.Set("qbase64", domain)

	for {
		// 定义初始化参数

		// 进行url参数解析
		Url, err := url.Parse("https://fofa.info/api/v1/search/all")
		if err != nil {
			return nil
		}

		// 如果url中存在中文，会对其进行url编码操作
		Url.RawQuery = params.Encode()

		// 构造完整请求目标的URl
		fofa_url := Url.String()
		fmt.Printf("fofa_url: %v\n", fofa_url)

		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(fofa_url)
		if err != nil {
			return nil
		}
		defer resp.Body.Close()

		respContent, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil
		}

		// 解析网页返回结果的json串，保存至结构体中
		json.Unmarshal(respContent, &respjson)

	}

	// for _, v := range respjson.Results {
	// 	fmt.Printf("v[0]: %v\n", v[0])
	// }

	return nil
}

// 集成subfinder的子域扫描接口
func _subfinder(domain string) []string {
	runnerInstance, err := runner.NewRunner(&runner.Options{
		Threads:            100,                             // Thread controls the number of threads to use for active enumerations
		Timeout:            30,                              // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 10,                              // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		Resolvers:          resolve.DefaultResolvers,        // Use the default list of resolvers by marshaling it to the config
		Sources:            passive.DefaultSources,          // Use the default list of passive sources
		AllSources:         passive.DefaultAllSources,       // Use the default list of all passive sources
		Recursive:          passive.DefaultRecursiveSources, // Use the default list of recursive sources
		Providers:          &runner.Providers{},             // Use empty api keys for all providers
	})

	buf := bytes.Buffer{}
	err = runnerInstance.EnumerateSingleDomain(context.Background(), domain, []io.Writer{&buf})
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)

	return nil
}

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
	"net/http"
	"poodle/pkg/logger"
	"strings"
	"time"

	"github.com/projectdiscovery/subfinder/v2/pkg/passive"
	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

type GetSubdomainMethod int

const (
	GSM_FOFA_API  GetSubdomainMethod = 100
	GSM_SUBFINDER GetSubdomainMethod = 200
)

// 定义所需参数
var (
	FOFA_EMAIL string = "cannopznjooss@gmail.com"
	FOFA_KEY   string = "a63e6f1bd58eba5877dbb0420addb6e9"
	// 子域信息收集保存结果
	Results []string
	// 子域收集方法
	// 默认使用fofa暴露接口进行子域收集
	Get_Subdomain_Method GetSubdomainMethod = GSM_FOFA_API
)

// 定义json返回结构体
type ResponseJson struct {
	Error   bool     `json:"error"`
	ErrMsg  string   `json:"errmsg"`
	Size    int      `json:"size"`
	Results []string `json:"results"`
}

// 调用子域收集功能
func ScanSubDomain(domain string) []string {
	switch Get_Subdomain_Method {
	case GSM_FOFA_API:
		return _getSubdomainByFofa(domain)
	case GSM_SUBFINDER:
		return _subfinder(domain)
	default:
		return _getSubdomainByFofa(domain)
	}
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
	qbase64 := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("domain=\"%s\"", domain)))

	// 构造完整请求目标的fofa API
	// 详情参阅 https://fofa.info/static_pages/api_help
	fofaURL := fmt.Sprintf("https://fofa.info/api/v1/search/all?full=true&fields=host&page=1&size=10000&email=%s&key=%s&qbase64=%s", FOFA_EMAIL, FOFA_KEY, qbase64)

	// fmt.Printf("fofaURL: %v\n", fofaURL)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(fofaURL)
	if err != nil {
		logger.LogError(err.Error(), logger.LOG_TERMINAL)
		return nil
	}
	defer resp.Body.Close()

	respContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogError(err.Error(), logger.LOG_TERMINAL)
		return nil
	}

	// 初始化结构体
	var respjson ResponseJson

	// 解析网页返回结果的json串，保存至结构体中
	err = json.Unmarshal(respContent, &respjson)
	if err != nil {
		logger.LogError(err.Error(), logger.LOG_TERMINAL)
		return nil
	}

	if respjson.Error {
		logger.LogError(respjson.ErrMsg, logger.LOG_TERMINAL)
		return nil
	}

	if respjson.Size > 0 {
		for _, subdomain := range respjson.Results {
			if strings.HasPrefix(strings.ToLower(subdomain), "http://") || strings.HasPrefix(strings.ToLower(subdomain), "https://") {
				// fmt.Printf("subdomain[strings.Index(subdomain, \"//\")+2:]: %v\n", subdomain[strings.Index(subdomain, "//")+2:])
				// 按照格式取出主机,可能是ip也可能子域名
				subdomain = strings.Split(subdomain[strings.Index(subdomain, "//")+2:], ":")[0]
				// fmt.Printf("subdomain: %v\n", subdomain)
				Results = append(Results, subdomain)
			}
		}
		// fmt.Printf("Results: %v\n", Results)
	}

	return Results
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
		logger.LogError(err.Error(), logger.LOG_TERMINAL)
	}

	data, err := io.ReadAll(&buf)
	if err != nil {
		logger.LogError(err.Error(), logger.LOG_TERMINAL)
	}

	fmt.Printf("%s", data)

	return nil
}

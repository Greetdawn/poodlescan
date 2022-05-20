package test

import (
	"fmt"
	"poodle/pkg/asset_host"
	"testing"
)

func TestMain(m *testing.M) {
	res := asset_host.ScanSubDomain("baidu.com")
	fmt.Println()
	fmt.Printf("res: %v\n", res)
	return
}

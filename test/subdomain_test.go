package test

import (
	"fmt"
	"poodle/pkg/asset_host"
	"testing"
)

func TestFofa(m *testing.M) {
	got := asset_host.ScanSubDomain("baidu.com")
	fmt.Println()
	fmt.Printf("got: %v\n", got)
}

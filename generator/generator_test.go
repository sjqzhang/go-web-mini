package generator

import "testing"

func TestGenerate(t *testing.T) {

	// 代码生成 // 123:456@tcp(127.0.0.1)/databaseName?tls=true
	Generate("root:root@tcp(127.0.0.1)/go_web_mini?tls=false",
		"go_web_mini",
		[]string{"news"},
		"go-web-mini")
}

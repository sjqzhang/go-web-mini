package generator

import "testing"

func TestGenerate(t *testing.T) {

	conf := Config{
		DSN:         "root:root@tcp(127.0.0.1)/go_web_mini?tls=false",
		ModuleName:  "go-web-mini",
		Tables:      []string{ "news"},
		WebRoot:     "/Users/junqiang.zhang/repo/js/go-web-mini-ui",
		ServerRoot:  "/Users/junqiang.zhang/repo/go/go-web-mini",
		TablePrefix: "",
		TableSuffix: "_tab",
	}
	InitConfig(conf)

	DoGenerate()
}

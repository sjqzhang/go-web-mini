package generator

import "testing"

func TestGenerate(t *testing.T) {

	conf := Config{
		DSN:        "root:root@tcp(127.0.0.1)/go_web_mini?tls=false",
		ModuleName: "go-web-mini",
		Tables:     []string{"branch_tab"},
		WebRoot:    "/Users/junqiang.zhang/repo/js/go-web-mini-ui",


	}
	InitConfig(conf)

	DoGenerate()
}

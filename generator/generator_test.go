package generator

import "testing"

func TestGenerate(t *testing.T) {

	conf := Config{
		DSN:         "root:root@tcp(127.0.0.1)/dms?tls=false",
		ModuleName:  "go-web-mini",
		Tables:      []string{"news"},
		WebRoot:     "/Users/junqiang.zhang/repo/js/go-web-mini-ui",
		ServerRoot:  "/Users/junqiang.zhang/repo/go/go-web-mini",
		TablePrefix: "",
		TableSuffix: "_tab",
	}
	InitConfig(conf)

	var tmpTables []string
	tables := getAllTableNames(nil, "dms")
	for _, table := range tables {
		if table == "users" || table == "roles" || table == "user_roles" || table == "role_permissions" || table == "user" ||
			table == "user_tab" {
			continue
		}
		tmpTables = append(tmpTables, table)
	}

	conf.Tables = tmpTables
	conf.Tables= []string{"ci_tab"}

	DoGenerate(&conf)
}

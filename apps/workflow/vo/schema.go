package vo

type GetSchemaInfosResp struct {
	SchemaInfos []SchemaInfo `json:"schema_infos"`
}

type SchemaInfo struct {
	ServiceTasks []ServiceTasks `json:"service_tasks"`
	Commands     []Command      `json:"commands"`
}

type SchemaFileResp struct {
	SchemaCode string `json:"schema_code"`
}


type GetSchemaReq struct {
	SchemaCodes []string `json:"schema_codes"`
}

type SchemaFileReq struct {
	SchemaCode string `json:"schema_code"`
	Schema     string `json:"schema"`
}

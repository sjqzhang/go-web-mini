{{define "api"}}
import request from '@/utils/request'

// 获取{{.Table.TableComment}}列表
export function get{{.Table.TableName}}(params) {
  return request({
    url: '/api/{{.Table.Uri}}',
    method: 'get',
    params
  })
}

// 创建{{.Table.TableComment}}
export function create{{.Table.TableName}}(data) {
  return request({
    url: '/api/{{.Table.Uri}}',
    method: 'post',
    data
  })
}

// 更新{{.Table.TableComment}}
export function update{{.Table.TableName}}(Id, data) {
  return request({
    url: '/api/{{.Table.Uri}}/' + Id,
    method: 'put',
    data
  })
}

// 删除{{.Table.TableComment}}
export function delete{{.Table.TableName}}(data) {
  return request({
    url: '/api/{{.Table.Uri}}',
    method: 'delete',
    data
  })
}

{{end}}

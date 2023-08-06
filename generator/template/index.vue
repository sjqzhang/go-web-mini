{{define "index.vue"}}

<template>
  <div>
    <el-card class="container-card" shadow="always">
      <el-form size="mini" :inline="true" :model="params" class="demo-form-inline">
        {{range .Fields}}{{if  checkField .ColumnName}}
        <el-form-item label="{{.ColumnCommentForView}}"><el-input v-model.trim="params.{{.ColumnName}}" clearable placeholder="{{.ColumnCommentForView}}" @clear="search" />
        </el-form-item>{{end}}{{end}}

        <el-form-item>
          <el-button :loading="loading" icon="el-icon-search" type="primary" @click="search">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button :loading="loading" icon="el-icon-plus" type="warning" @click="create">新增</el-button>
        </el-form-item>
        <el-form-item>
          <el-button :disabled="multipleSelection.length === 0" :loading="loading" icon="el-icon-delete" type="danger" @click="batchDelete">批量删除</el-button>
        </el-form-item>
      </el-form>

      <el-table v-loading="loading" :data="tableData" border stripe style="width: 100%" @selection-change="handleSelectionChange">
        <el-table-column type="selection" width="55" align="center" />
        {{range .Fields}}{{if  checkField .ColumnName}}
        <el-table-column show-overflow-tooltip sortable prop="{{.ColumnName}}" label="{{.ColumnCommentForView}}" />{{end}}{{end}}
        <el-table-column fixed="right" label="操作" align="center" width="120">
          <template slot-scope="scope">
            <el-tooltip content="编辑" effect="dark" placement="top">
              <el-button size="mini" icon="el-icon-edit" circle type="primary" @click="update(scope.row)" />
            </el-tooltip>
            <el-tooltip class="delete-popover" content="删除" effect="dark" placement="top">
              <el-popconfirm title="确定删除吗？" @onConfirm="singleDelete(scope.row.id)">
                <el-button slot="reference" size="mini" icon="el-icon-delete" circle type="danger" />
              </el-popconfirm>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        :current-page="params.pageNum"
        :page-size="params.pageSize"
        :total="total"
        :page-sizes="[1, 5, 10, 30]"
        layout="total, prev, pager, next, sizes"
        background
        style="margin-top: 10px;float:right;margin-bottom: 10px;"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />

      <el-dialog :title="dialogFormTitle" :visible.sync="dialogFormVisible">
        <el-form ref="dialogForm" size="small" :model="dialogFormData" :rules="dialogFormRules" label-width="120px">
          {{range .Fields}}{{if  checkField .ColumnName}}
          <el-form-item label="{{.ColumnCommentForView}}" prop="{{.ColumnName}}">
            <el-input v-model.trim="dialogFormData.{{.ColumnName}}" placeholder="{{.ColumnCommentForView}}" />
          </el-form-item>{{end}}{{end}}
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button size="mini" @click="cancelForm()">取 消</el-button>
          <el-button size="mini" :loading="submitLoading" type="primary" @click="submitForm()">确 定</el-button>
        </div>
      </el-dialog>

    </el-card>
  </div>
</template>

<script>
import { get{{.Table.TableName}}, create{{.Table.TableName}}, update{{.Table.TableName}}, delete{{.Table.TableName}} } from '@/api/business/{{.Table.Uri}}'

export default {
  name: 'Api',
  filters: {
    methodTagFilter(val) {
      if (val === 'GET') {
        return ''
      } else if (val === 'POST') {
        return 'success'
      } else if (val === 'PUT') {
        return 'info'
      } else if (val === 'PATCH') {
        return 'warning'
      } else if (val === 'DELETE') {
        return 'danger'
      } else {
        return 'info'
      }
    }
  },
  data() {
    return {
      // 查询参数
      params: {
      {{range .Fields}}{{if  checkField .ColumnName}}
        // content: '',
            {{.ColumnName}}:'',
            {{end}}{{end}}

        pageNum: 1,
        pageSize: 10
      },
      // 表格数据
      tableData: [],
      total: 0,
      loading: false,

      // dialog对话框
      submitLoading: false,
      dialogFormTitle: '',
      dialogType: '',
      dialogFormVisible: false,
      dialogFormData: {
        id:'',
            {{range .Fields}}{{if  checkField .ColumnName}}
      // content: '',
      {{.ColumnName}}:'',
          {{end}}{{end}}
      },
      // dialogFormRules: {
      //   path: [
      //     { required: true, message: '请输入访问路径', trigger: 'blur' },
      //     { min: 1, max: 100, message: '长度在 1 到 100 个字符', trigger: 'blur' }
      //   ],
      //   category: [
      //     { required: true, message: '请输入所属类别', trigger: 'blur' },
      //     { min: 1, max: 50, message: '长度在 1 到 50 个字符', trigger: 'blur' }
      //   ],
      //   method: [
      //     { required: true, message: '请选择请求方式', trigger: 'change' }
      //   ],
      //   desc: [
      //     { required: false, message: '说明', trigger: 'blur' },
      //     { min: 0, max: 100, message: '长度在 0 到 100 个字符', trigger: 'blur' }
      //   ]
      // },

      // 删除按钮弹出框
      popoverVisible: false,
      // 表格多选
      multipleSelection: []
    }
  },
  created() {
    this.getTableData()
  },
  methods: {
    // 查询
    search() {
      this.params.pageNum = 1
      this.getTableData()
    },

    // 获取表格数据
    async getTableData() {
      this.loading = true
      try {
        const { data } = await get{{.Table.TableName}}(this.params)
        this.tableData = data.list
        this.total = data.total
      } finally {
        this.loading = false
      }
    },

    // 新增
    create() {
      this.dialogFormTitle = '新增接口'
      this.dialogType = 'create'
      this.dialogFormVisible = true
    },

    // 修改
    update(row) {
      this.dialogFormData.id = row.id

      {{range .Fields}}{{if  checkField .ColumnName}}
      // content: '',
      this.dialogFormData.{{.ColumnName}}=row.{{.ColumnName}}
          {{end}}{{end}}
      // this.dialogFormData.title = row.title
      // this.dialogFormData.content = row.content
      // this.dialogFormData.method = row.method
      // this.dialogFormData.desc = row.desc

      this.dialogFormTitle = '修改接口'
      this.dialogType = 'update'
      this.dialogFormVisible = true
    },

    // 提交表单
    submitForm() {
      this.$refs['dialogForm'].validate(async valid => {
        if (valid) {
          let msg = ''
          this.submitLoading = true
          try {
            if (this.dialogType === 'create') {
              const { message } = await create{{.Table.TableName}}(this.dialogFormData)
              msg = message
            } else {
              const { message } = await update{{.Table.TableName}}(this.dialogFormData.id, this.dialogFormData)
              msg = message
            }
          } finally {
            this.submitLoading = false
          }

          this.resetForm()
          this.getTableData()
          this.$message({
            showClose: true,
            message: msg,
            type: 'success'
          })
        } else {
          debugger
          this.$message({
            showClose: true,
            message: '表单校验失败',
            type: 'error'
          })
          return false
        }
      })
    },

    // 提交表单
    cancelForm() {
      this.resetForm()
    },

    resetForm() {
      this.dialogFormVisible = false
      this.$refs['dialogForm'].resetFields()
      this.dialogFormData = {
      {{range .Fields}}{{if  checkField .ColumnName}}
      // content: '',
          {{.ColumnName}}:'',
      {{end}}{{end}}

      }
    },

    // 批量删除
    batchDelete() {
      this.$confirm('此操作将永久删除, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(async res => {
        this.loading = true
        const apiIds = []
        this.multipleSelection.forEach(x => {
          apiIds.push(x.ID)
        })
        let msg = ''
        try {
          const { message } = await delete{{.Table.TableName}}({ apiIds: apiIds })
          msg = message
        } finally {
          this.loading = false
        }

        this.getTableData()
        this.$message({
          showClose: true,
          message: msg,
          type: 'success'
        })
      }).catch(() => {
        this.$message({
          showClose: true,
          type: 'info',
          message: '已取消删除'
        })
      })
    },

    // 表格多选
    handleSelectionChange(val) {
      this.multipleSelection = val
    },

    // 单个删除
    async singleDelete(Id) {
      this.loading = true
      let msg = ''
      try {
        const { message } = await delete{{.Table.TableName}}({ id: Id })
        msg = message
      } finally {
        this.loading = false
      }

      this.getTableData()
      this.$message({
        showClose: true,
        message: msg,
        type: 'success'
      })
    },

    // 分页
    handleSizeChange(val) {
      this.params.pageSize = val
      this.getTableData()
    },
    handleCurrentChange(val) {
      this.params.pageNum = val
      this.getTableData()
    }
  }
}
</script>

<style scoped>
  .container-card{
    margin: 10px;
  }

  .delete-popover{
    margin-left: 10px;
  }
</style>

{{end}}

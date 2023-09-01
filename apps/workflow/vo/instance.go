package vo

type InstanceTransitions struct {
	Transitions []InstanceTransition `json:"transitions"`
}

type InstanceTransition struct {
	InstanceId     int         `json:"instance_id"`
	IdentityId     string      `json:"identity_id"`
	FromStateName  string      `json:"from_state_name"`
	ToStateName    string      `json:"to_state_name"`
	CommandName    string      `json:"command_name"`
	Variables      []Variables `json:"variables"`
	TransitionTime int         `json:"transition_time"`
}
type CreateInstanceResp struct {
	InstanceID int `json:"instance_id"`
}

type ServiceTasks struct {
	TaskId    string      `json:"task_id"`
	TaskName  string      `json:"task_name"`
	Variables []Variables `json:"variables"`
}

type Variables struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetInstanceResp struct {
	InstanceID      int            `json:"instance_id"`
	SchemeCode      string         `json:"scheme_code"`
	State           string         `json:"state"`
	CurrentTaskID   string         `json:"current_task_Id"`
	CurrentTaskName string         `json:"current_task_name"`
	ServiceTasks    []ServiceTasks `json:"service_tasks"`
	Commands        []Command      `json:"commands"`
	Variables       []Variables    `json:"variables"`
}

type InstanceTasks struct {
	InstanceID      int            `json:"instance_id"`
	CurrentTaskName string         `json:"current_task_name"`
	Tasks           []ServiceTasks `json:"tasks"`
}

type GetInstanceTasksResp struct {
	InstanceTasksList []InstanceTasks `json:"instance_tasks_list"`
}

type InstanceSchemaCode struct {
	InstanceID        int    `json:"instance_id"`
	SchemaCode        string `json:"schema_code"`
	InstanceCurTaskId string `json:"instance_cur_task_id"`
}

type GetInstanceSchemaCodeResp struct {
	InstanceSchemaCodeList []InstanceSchemaCode `json:"instance_schema_code_list"`
}

type InstanceCommand struct {
	InstanceID int     `json:"instance_id"`
	Command    Command `json:"command"`
}

type GetInstanceCommandResp struct {
	InstanceCommandList []InstanceCommand `json:"instance_command_list"`
}





type CreateInstance struct {
	SchemeCode string      `json:"scheme_code"`
	Variables  []Variables `json:"variables"`
}



type SetInstanceVariablesReq struct {
	InstanceId int    `uri:"instance_id" binding:"required"`
	Name       string `json:"name"`
	Value      string `json:"value"`
}

type GetInstanceReq struct {
	InstanceId int `uri:"instance_id" binding:"required"`
}

type GetInstanceVariablesReq struct {
	InstanceId int    `uri:"instance_id" binding:"required"`
	Name       string `form:"name"`
}

type GetInstanceTasksReq struct {
	InstanceIdList []int `json:"instance_id_list"`
}

type InstanceSchemaCodeReq struct {
	InstanceIdList []int  `json:"instance_id_list"`
	SchemaCode     string `json:"schema_code"`
}

type GetInstanceCommandReq struct {
	InstanceIdList []int  `json:"instance_id_list"`
	CommandKey     string `json:"command_key"`
}

type GetInstanceAvailableCommandsReq struct {
	InstanceId int    `uri:"instance_id" binding:"required"`
	IdentityId string `form:"identity_id"`
}

type GetInsListAvailableCommandsReq struct {
	InstanceIds []int  `json:"instance_ids"`
	IdentityId  string `form:"identity_id"`
}

type CommandExecution struct {
	InstanceId int         `uri:"instance_id" binding:"required"`
	IdentityId string      `json:"identity_id"`
	Command    string      `json:"command"`
	Variables  []Variables `json:"variables"`
}

type CommandExecutionBatch struct {
	InstanceId int         `json:"instance_id"`
	IdentityId string      `json:"identity_id"`
	Command    string      `json:"command"`
	Variables  []Variables `json:"variables"`
}

type CommandExecutionBatchReq struct {
	CommandExecutionBatchList []CommandExecutionBatch
}

type CommandOptimize struct {
	Command   string      `json:"command"`
	Variables []Variables `json:"variables"`
}

type CommandExecutionOptimizeReq struct {
	InstanceId  int               `json:"instance_id"`
	IdentityId  string            `json:"identity_id"`
	CommandList []CommandOptimize `json:"command_list"`
}

type GetInstanceTransitions struct {
	InstanceId int `uri:"instance_id" binding:"required"`
}

type GetInstanceTransitionsByIds struct {
	Ids []int `json:"ids"`
}

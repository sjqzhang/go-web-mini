package vo


type GetInstanceAvailableCommandsReq struct {
	InstanceId int    `uri:"instance_id" binding:"required"`
	IdentityId string `form:"identity_id"`
}

type GetInstanceAvailableCommands struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	Key       string   `json:"key"`
	Variables []string `json:"variables"`
}

type InstanceTransitions struct {
	Transitions []InstanceTransition `json:"transitions"`
}

type InstanceTransition struct {
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
	TaskId   string `json:"task_id"`
	TaskName string `json:"task_name"`
}

type Variables struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetInstanceResp struct {
	InstanceID    int            `json:"instance_id"`
	SchemeCode    string         `json:"scheme_code"`
	State         string         `json:"state"`
	CurrentTaskID string         `json:"current_task_Id"`
	ServiceTasks  []ServiceTasks `json:"service_tasks"`
	Variables     []Variables    `json:"variables"`
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



type CommandExecution struct {
	InstanceId int         `uri:"instance_id" binding:"required"`
	IdentityId string      `json:"identity_id"`
	Command    string      `json:"command"`
	Variables  []Variables `json:"variables"`
}

type GetInstanceTransitions struct {
	InstanceId int `uri:"instance_id" binding:"required"`
}

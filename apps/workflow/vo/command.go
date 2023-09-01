package vo

type GetInstanceAvailableCommands struct {
	Commands []Command `json:"commands"`
}

type Command struct {
	Id        string      `json:"id"`
	Key       string      `json:"key"`
	Variables []Variables `json:"variables"`
}

type CommandExecResp struct {
	TransitionIds []int `json:"transition_ids"`
}

type InsIdTransitionIds struct {
	InstanceId    int   `json:"instance_id"`
	TransitionIds []int `json:"transition_ids"`
}

type CommandExecBatchResp struct {
	InsIdTransitionIdsList []InsIdTransitionIds `json:"ins_id_transition_ids_list"`
}

type InsAvailableCommands struct {
	InstanceId  int       `json:"instance_id"`
	CommandList []Command `json:"command_list"`
}

type GetInsListAvailableCommandsResp struct {
	InsAvailableCommandsList []InsAvailableCommands `json:"ins_available_commands_list"`
}

package models

type RunState struct {
	LifeCycleState RunLifeCycleState `json:"life_cycle_state,omitempty" url:"life_cycle_state,omitempty"`
	ResultState    RunResultState    `json:"result_state,omitempty" url:"result_state,omitempty"`
	StateMessage   string            `json:"state_message,omitempty" url:"state_message,omitempty"`
}

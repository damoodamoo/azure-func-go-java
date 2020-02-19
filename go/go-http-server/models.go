package main

type ReturnValue struct {
	Data string
}
type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

type InvokeResponseStringReturnValue struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue string
}

type InvokeRequest struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}

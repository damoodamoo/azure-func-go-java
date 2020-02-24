package main

type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

type InvokeRequest struct {
	Data     map[string]interface{}
	Metadata map[string]interface{}
}

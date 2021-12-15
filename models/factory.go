package models

type Deployment struct {
	Name              string       `json:"name"`
	DeployName        string       `json:"deploy_name"`
	PodName           string       `json:"pod_name"`
	NodeName          string       `json:"node_name"`
	Image             string       `json:"image" bson:"image"`
	Tag               string       `json:"tag" bson:"tag"`
	ServiceName       string       `json:"service_name"`
	InternalAddress   string       `json:"internal_address"`
	Port              int32        `json:"port"`
	TotalReplicas     int32        `json:"total_replicas"`
	AvailableReplicas int32        `json:"available_replicas"`
	Size              int          `json:"full_size" bson:"full_size"`
	LastUpdated       string       `json:"last_updated" bson:"last_updated"`
	MemoryLimit       string       `json:"memory_limit" bson:"memory_limit"`
}
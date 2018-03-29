package entity

//heapster 获取的事件类型
type HEvent struct {
	Metadata       Metadata       `json:"metadata"`
	InvolvedObject InvolvedObject `json:"involvedObject"`
	Reason         string         `json:"reason"`
	Message        string         `json:"message"`
	Source         source         `json:"source"`
	FirstTimestamp string         `json:"firstTimestamp"`
	LastTimestamp  string         `json:"LastTimestamp"`
	Count          int            `json:"count"`
	Type           string         `json:"type"`
}

type Metadata struct {
	Name              string `json:"name"`
	NameSpace         string `json:"namespace"`
	SelfLink          string `json:"selfLink"`
	Uid               string `json:"uid"`
	ResourceVersion   string `json:"resourceVersion"`
	CreationTimestamp string `json:"creationTimestamp"`
}

type InvolvedObject struct {
	Kind            string `json:"kind"`
	Namespace       string `json:"namespace"`
	Name            string `json:"name"`
	Uid             string `json:"uid"`
	ApiVersion      string `json:"apiVersion"`
	ResourceVersion string `json:"resourceVersion"`
}
type source struct {
	Component string `json:"component"`
}

package api

// Tag ...
type Tag struct {
	Key   string `json:"key" tf:"force_new"`
	Value string `json:"value" tf:"force_new"`
}

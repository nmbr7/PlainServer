package types

type Compute_getInstanceParam struct {
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags. Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored (both PUT & PATCH) when empty.
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty" yaml:"resourceManagerTags,omitempty"`
}

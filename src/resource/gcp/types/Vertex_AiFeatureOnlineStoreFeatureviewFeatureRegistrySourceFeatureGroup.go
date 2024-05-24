package types

type Vertex_AiFeatureOnlineStoreFeatureviewFeatureRegistrySourceFeatureGroup struct {
	// Identifier of the feature group.
	FeatureGroupId string `json:"featureGroupId,omitempty" yaml:"featureGroupId,omitempty"`

	// Identifiers of features under the feature group.
	FeatureIds []string `json:"featureIds,omitempty" yaml:"featureIds,omitempty"`
}
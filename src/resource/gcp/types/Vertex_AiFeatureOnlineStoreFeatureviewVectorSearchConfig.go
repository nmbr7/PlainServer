package types

type Vertex_AiFeatureOnlineStoreFeatureviewVectorSearchConfig struct {
	/*
	   Configuration options for the tree-AH algorithm (Shallow tree + Asymmetric Hashing). Please refer to this paper for more details: https://arxiv.org/abs/1908.10396
	   Structure is documented below.
	*/
	TreeAhConfig Vertex_AiFeatureOnlineStoreFeatureviewVectorSearchConfigTreeAhConfig `json:"treeAhConfig,omitempty" yaml:"treeAhConfig,omitempty"`

	// Configuration options for using brute force search, which simply implements the standard linear search in the database for each query. It is primarily meant for benchmarking and to generate the ground truth for approximate search.
	BruteForceConfig Vertex_AiFeatureOnlineStoreFeatureviewVectorSearchConfigBruteForceConfig `json:"bruteForceConfig,omitempty" yaml:"bruteForceConfig,omitempty"`

	// Column of crowding. This column contains crowding attribute which is a constraint on a neighbor list produced by nearest neighbor search requiring that no more than some value k' of the k neighbors returned have the same value of crowdingAttribute.
	CrowdingColumn string `json:"crowdingColumn,omitempty" yaml:"crowdingColumn,omitempty"`

	/*
	   The distance measure used in nearest neighbor search.
	   For details on allowed values, see the [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1beta1/projects.locations.featureOnlineStores.featureViews#DistanceMeasureType).
	   Possible values are: `SQUARED_L2_DISTANCE`, `COSINE_DISTANCE`, `DOT_PRODUCT_DISTANCE`.
	*/
	DistanceMeasureType string `json:"distanceMeasureType,omitempty" yaml:"distanceMeasureType,omitempty"`

	// Column of embedding. This column contains the source data to create index for vector search.
	EmbeddingColumn string `json:"embeddingColumn,omitempty" yaml:"embeddingColumn,omitempty"`

	// The number of dimensions of the input embedding.
	EmbeddingDimension int `json:"embeddingDimension,omitempty" yaml:"embeddingDimension,omitempty"`

	// Columns of features that are used to filter vector search results.
	FilterColumns []string `json:"filterColumns,omitempty" yaml:"filterColumns,omitempty"`
}

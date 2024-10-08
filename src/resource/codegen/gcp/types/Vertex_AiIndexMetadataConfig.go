package types

type Vertex_AiIndexMetadataConfig struct {
	/*
	   The configuration with regard to the algorithms used for efficient search.
	   Structure is documented below.
	*/
	AlgorithmConfig Vertex_AiIndexMetadataConfigAlgorithmConfig `json:"algorithmConfig,omitempty" yaml:"algorithmConfig,omitempty"`

	/*
	   The default number of neighbors to find via approximate search before exact reordering is
	   performed. Exact reordering is a procedure where results returned by an
	   approximate search algorithm are reordered via a more expensive distance computation.
	   Required if tree-AH algorithm is used.
	*/
	ApproximateNeighborsCount int `json:"approximateNeighborsCount,omitempty" yaml:"approximateNeighborsCount,omitempty"`

	// The number of dimensions of the input vectors.
	Dimensions int `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	/*
	   The distance measure used in nearest neighbor search. The value must be one of the followings:
	   - SQUARED_L2_DISTANCE: Euclidean (L_2) Distance
	   - L1_DISTANCE: Manhattan (L_1) Distance
	   - COSINE_DISTANCE: Cosine Distance. Defined as 1 - cosine similarity.
	   - DOT_PRODUCT_DISTANCE: Dot Product Distance. Defined as a negative of the dot product
	*/
	DistanceMeasureType string `json:"distanceMeasureType,omitempty" yaml:"distanceMeasureType,omitempty"`

	/*
	   Type of normalization to be carried out on each vector. The value must be one of the followings:
	   - UNIT_L2_NORM: Unit L2 normalization type
	   - NONE: No normalization type is specified.
	*/
	FeatureNormType string `json:"featureNormType,omitempty" yaml:"featureNormType,omitempty"`

	/*
	   Index data is split into equal parts to be processed. These are called "shards".
	   The shard size must be specified when creating an index. The value must be one of the followings:
	   - SHARD_SIZE_SMALL: Small (2GB)
	   - SHARD_SIZE_MEDIUM: Medium (20GB)
	   - SHARD_SIZE_LARGE: Large (50GB)
	*/
	ShardSize string `json:"shardSize,omitempty" yaml:"shardSize,omitempty"`
}

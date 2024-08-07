package types

type Blockchainnodeengine_BlockchainNodesEthereumDetailsGethDetails struct {
	/*
	   Blockchain garbage collection modes. Only applicable when NodeType is FULL or ARCHIVE.
	   Possible values are: `FULL`, `ARCHIVE`.

	   <a name="nested_additional_endpoints"></a>The `additional_endpoints` block contains:
	*/
	GarbageCollectionMode string `json:"garbageCollectionMode,omitempty" yaml:"garbageCollectionMode,omitempty"`
}

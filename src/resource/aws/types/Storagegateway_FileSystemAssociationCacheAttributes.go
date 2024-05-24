package types

type Storagegateway_FileSystemAssociationCacheAttributes struct {
	/*
	   Refreshes a file share's cache by using Time To Live (TTL).
	   TTL is the length of time since the last refresh after which access to the directory would cause the file gateway
	   to first refresh that directory's contents from the Amazon S3 bucket. Valid Values: `0` or `300` to `2592000` seconds (5 minutes to 30 days). Defaults to `0`
	*/
	CacheStaleTimeoutInSeconds int `json:"cacheStaleTimeoutInSeconds,omitempty" yaml:"cacheStaleTimeoutInSeconds,omitempty"`
}

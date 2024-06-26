package types

type Glue_CrawlerRecrawlPolicy struct {
	// Specifies whether to crawl the entire dataset again, crawl only folders that were added since the last crawler run, or crawl what S3 notifies the crawler of via SQS. Valid Values are: `CRAWL_EVENT_MODE`, `CRAWL_EVERYTHING` and `CRAWL_NEW_FOLDERS_ONLY`. Default value is `CRAWL_EVERYTHING`.
	RecrawlBehavior string `json:"recrawlBehavior,omitempty" yaml:"recrawlBehavior,omitempty"`
}

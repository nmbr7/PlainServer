package types

type Secretmanager_getSecretTopic struct {
	/*
	   The resource name of the Pub/Sub topic that will be published to, in the following format: projects/-/topics/-.
	   For publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

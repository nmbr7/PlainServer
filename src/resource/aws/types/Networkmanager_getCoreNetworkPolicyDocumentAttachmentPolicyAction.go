package types

type Networkmanager_getCoreNetworkPolicyDocumentAttachmentPolicyAction struct {
	// Defines how a segment is mapped. Values can be `constant` or `tag`. `constant` statically defines the segment to associate the attachment to. `tag` uses the value of a tag to dynamically try to map to a segment.reference_policies_elements_condition_operators.html) to evaluate.
	AssociationMethod string `json:"associationMethod,omitempty" yaml:"associationMethod,omitempty"`

	// Determines if this mapping should override the segment value for `require_attachment_acceptance`. You can only set this to `true`, indicating that this setting applies only to segments that have `require_attachment_acceptance` set to `false`. If the segment already has the default `require_attachment_acceptance`, you can set this to inherit segment’s acceptance value.
	RequireAcceptance bool `json:"requireAcceptance,omitempty" yaml:"requireAcceptance,omitempty"`

	// Name of the `segment` to share as defined in the `segments` section. This is used only when the `association_method` is `constant`.
	Segment string `json:"segment,omitempty" yaml:"segment,omitempty"`

	// Maps the attachment to the value of a known key. This is used with the `association_method` is `tag`. For example a `tag` of `stage = “test”`, will map to a segment named `test`. The value must exactly match the name of a segment. This allows you to have many segments, but use only a single rule without having to define multiple nearly identical conditions. This prevents creating many similar conditions that all use the same keys to map to segments.
	TagValueOfKey string `json:"tagValueOfKey,omitempty" yaml:"tagValueOfKey,omitempty"`
}

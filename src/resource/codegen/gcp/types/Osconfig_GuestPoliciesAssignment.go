package types

type Osconfig_GuestPoliciesAssignment struct {
	/*
	   Targets instances in any of these zones. Leave empty to target instances in any zone.
	   Zonal targeting is uncommon and is supported to facilitate the management of changes by zone.
	*/
	Zones []string `json:"zones,omitempty" yaml:"zones,omitempty"`

	/*
	   Targets instances matching at least one of these label sets. This allows an assignment to target disparate groups,
	   for example "env=prod or env=staging".
	   Structure is documented below.
	*/
	GroupLabels []Osconfig_GuestPoliciesAssignmentGroupLabel `json:"groupLabels,omitempty" yaml:"groupLabels,omitempty"`

	/*
	   Targets VM instances whose name starts with one of these prefixes.
	   Like labels, this is another way to group VM instances when targeting configs,
	   for example prefix="prod-".
	   Only supported for project-level policies.
	*/
	InstanceNamePrefixes []string `json:"instanceNamePrefixes,omitempty" yaml:"instanceNamePrefixes,omitempty"`

	/*
	   Targets any of the instances specified. Instances are specified by their URI in the form
	   zones/[ZONE]/instances/[INSTANCE_NAME].
	   Instance targeting is uncommon and is supported to facilitate the management of changes
	   by the instance or to target specific VM instances for development and testing.
	   Only supported for project-level policies and must reference instances within this project.
	*/
	Instances []string `json:"instances,omitempty" yaml:"instances,omitempty"`

	/*
	   Targets VM instances matching at least one of the following OS types.
	   VM instances must match all supplied criteria for a given OsType to be included.
	   Structure is documented below.
	*/
	OsTypes []Osconfig_GuestPoliciesAssignmentOsType `json:"osTypes,omitempty" yaml:"osTypes,omitempty"`
}

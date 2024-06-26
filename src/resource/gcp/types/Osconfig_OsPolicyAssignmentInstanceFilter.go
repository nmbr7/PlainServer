package types

type Osconfig_OsPolicyAssignmentInstanceFilter struct {
	/*
	   Target all VMs in the project. If true, no other criteria
	   is permitted.
	*/
	All bool `json:"all,omitempty" yaml:"all,omitempty"`

	/*
	   List of label sets used for VM exclusion. If
	   the list has more than one label set, the VM is excluded if any of the label
	   sets are applicable for the VM. Structure is
	   documented below.
	*/
	ExclusionLabels []Osconfig_OsPolicyAssignmentInstanceFilterExclusionLabel `json:"exclusionLabels,omitempty" yaml:"exclusionLabels,omitempty"`

	/*
	   List of label sets used for VM inclusion. If
	   the list has more than one `LabelSet`, the VM is included if any of the
	   label sets are applicable for the VM. Structure is
	   documented below.
	*/
	InclusionLabels []Osconfig_OsPolicyAssignmentInstanceFilterInclusionLabel `json:"inclusionLabels,omitempty" yaml:"inclusionLabels,omitempty"`

	/*
	   List of inventories to select VMs. A VM is
	   selected if its inventory data matches at least one of the following
	   inventories. Structure is documented below.
	*/
	Inventories []Osconfig_OsPolicyAssignmentInstanceFilterInventory `json:"inventories,omitempty" yaml:"inventories,omitempty"`
}

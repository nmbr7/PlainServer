package types

type Compute_getRegionInstanceTemplateReservationAffinity struct {
	// Specifies the label selector for the reservation to use.
	SpecificReservations []Compute_getRegionInstanceTemplateReservationAffinitySpecificReservation `json:"specificReservations,omitempty" yaml:"specificReservations,omitempty"`

	// The accelerator type resource to expose to this instance. E.g. `nvidia-tesla-k80`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

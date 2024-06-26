package types

type Compute_InstanceFromTemplateReservationAffinity struct {
	// The type of reservation from which this instance can consume resources.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Specifies the label selector for the reservation to use.
	SpecificReservation Compute_InstanceFromTemplateReservationAffinitySpecificReservation `json:"specificReservation,omitempty" yaml:"specificReservation,omitempty"`
}

package types

type Diagflow_CxEntityTypeEntity struct {
	/*
	   The primary value associated with this entity entry. For example, if the entity type is vegetable, the value could be scallions.
	   For KIND_MAP entity types: A canonical value to be used in place of synonyms.
	   For KIND_LIST entity types: A string that can contain references to other entity types (with or without aliases).
	*/
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	/*
	   A collection of value synonyms. For example, if the entity type is vegetable, and value is scallions, a synonym could be green onions.
	   For KIND_LIST entity types: This collection must contain exactly one synonym equal to value.

	   - - -
	*/
	Synonyms []string `json:"synonyms,omitempty" yaml:"synonyms,omitempty"`
}

package types

type Connect_getUserHierarchyGroupHierarchyPath struct {
	// Details of level five. See below.
	LevelFives []Connect_getUserHierarchyGroupHierarchyPathLevelFife `json:"levelFives,omitempty" yaml:"levelFives,omitempty"`

	// Details of level four. See below.
	LevelFours []Connect_getUserHierarchyGroupHierarchyPathLevelFour `json:"levelFours,omitempty" yaml:"levelFours,omitempty"`

	// Details of level one. See below.
	LevelOnes []Connect_getUserHierarchyGroupHierarchyPathLevelOne `json:"levelOnes,omitempty" yaml:"levelOnes,omitempty"`

	// Details of level three. See below.
	LevelThrees []Connect_getUserHierarchyGroupHierarchyPathLevelThree `json:"levelThrees,omitempty" yaml:"levelThrees,omitempty"`

	// Details of level two. See below.
	LevelTwos []Connect_getUserHierarchyGroupHierarchyPathLevelTwo `json:"levelTwos,omitempty" yaml:"levelTwos,omitempty"`
}

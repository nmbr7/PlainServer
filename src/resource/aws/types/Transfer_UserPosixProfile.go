package types

type Transfer_UserPosixProfile struct {
	// The POSIX group ID used for all EFS operations by this user.
	Gid int `json:"gid,omitempty" yaml:"gid,omitempty"`

	// The secondary POSIX group IDs used for all EFS operations by this user.
	SecondaryGids []int `json:"secondaryGids,omitempty" yaml:"secondaryGids,omitempty"`

	// The POSIX user ID used for all EFS operations by this user.
	Uid int `json:"uid,omitempty" yaml:"uid,omitempty"`
}

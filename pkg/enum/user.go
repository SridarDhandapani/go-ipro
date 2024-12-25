package enum

import "fmt"

type ACL int

const (
	ACLAdministrator ACL = iota + 1
	ACLCameraControl
	ACLLiveOnly
)

func (a ACL) String() string {
	return fmt.Sprintf("%d", a)
}

// generated by stringer -type=DevState dev_state.go; DO NOT EDIT

package directory

import "fmt"

const _DevState_name = "DevStateInvalidDevStateNewDevStateReady"

var _DevState_index = [...]uint8{0, 15, 26, 39}

func (i DevState) String() string {
	if i >= DevState(len(_DevState_index)-1) {
		return fmt.Sprintf("DevState(%d)", i)
	}
	return _DevState_name[_DevState_index[i]:_DevState_index[i+1]]
}

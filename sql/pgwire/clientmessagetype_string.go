// Code generated by "stringer -type=clientMessageType"; DO NOT EDIT

package pgwire

import "fmt"

const (
	_clientMessageType_name_0 = "clientMsgBindclientMsgCloseclientMsgDescribeclientMsgExecute"
	_clientMessageType_name_1 = "clientMsgParseclientMsgSimpleQuery"
	_clientMessageType_name_2 = "clientMsgSync"
	_clientMessageType_name_3 = "clientMsgTerminate"
)

var (
	_clientMessageType_index_0 = [...]uint8{0, 13, 27, 44, 60}
	_clientMessageType_index_1 = [...]uint8{0, 14, 34}
	_clientMessageType_index_2 = [...]uint8{0, 13}
	_clientMessageType_index_3 = [...]uint8{0, 18}
)

func (i clientMessageType) String() string {
	switch {
	case 66 <= i && i <= 69:
		i -= 66
		return _clientMessageType_name_0[_clientMessageType_index_0[i]:_clientMessageType_index_0[i+1]]
	case 80 <= i && i <= 81:
		i -= 80
		return _clientMessageType_name_1[_clientMessageType_index_1[i]:_clientMessageType_index_1[i+1]]
	case i == 83:
		return _clientMessageType_name_2
	case i == 88:
		return _clientMessageType_name_3
	default:
		return fmt.Sprintf("clientMessageType(%d)", i)
	}
}

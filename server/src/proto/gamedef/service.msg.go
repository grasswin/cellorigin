// This code is generated by github.com/davyxu/cellnet/protoc-gen-msg, DO NOT EDIT
// Generated from: service.proto
package gamedef

import (
	"github.com/davyxu/cellnet"
)

func init() {
	cellnet.RegisterMessage("gamedef.ServiceDefine", (*ServiceDefine)(nil))
	cellnet.RegisterMessage("gamedef.ServiceFile", (*ServiceFile)(nil))
	cellnet.RegisterMessage("gamedef.ServiceMap", (*ServiceMap)(nil))
	cellnet.RegisterMessage("gamedef.LocalConfig", (*LocalConfig)(nil))
}

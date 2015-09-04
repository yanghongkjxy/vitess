package callinfo

import (
	"fmt"
	"html/template"

	"github.com/youtube/vitess/go/rpcwrap/proto"
	"golang.org/x/net/context"
)

// RPCWrapCallInfo takes a context generated by rpcwrap, and
// returns one that has CallInfo filled in.
func RPCWrapCallInfo(ctx context.Context) context.Context {
	remoteAddr, _ := proto.RemoteAddr(ctx)
	return NewContext(ctx, &rpcWrapCallInfoImpl{
		remoteAddr: remoteAddr,
	})
}

type rpcWrapCallInfoImpl struct {
	remoteAddr string
}

func (rwci *rpcWrapCallInfoImpl) RemoteAddr() string {
	return rwci.remoteAddr
}

func (rwci *rpcWrapCallInfoImpl) Username() string {
	return ""
}

func (rwci *rpcWrapCallInfoImpl) Text() string {
	return fmt.Sprintf("%s", rwci.remoteAddr)
}

func (rwci *rpcWrapCallInfoImpl) HTML() template.HTML {
	return template.HTML("<b>RemoteAddr:</b> " + rwci.remoteAddr + "</br>\n")
}
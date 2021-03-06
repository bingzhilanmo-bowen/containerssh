package signal

import (
	"github.com/janoszen/containerssh/backend"
	"github.com/janoszen/containerssh/log"
	channelRequest "github.com/janoszen/containerssh/ssh/channel/request"

	"golang.org/x/crypto/ssh"
)

type requestMsg struct {
	signal string
}


type ChannelRequestHandler struct {
	logger log.Logger
}

func New(logger log.Logger) channelRequest.TypeHandler {
	return &ChannelRequestHandler{
		logger: logger,
	}
}


func (e ChannelRequestHandler) GetRequestObject() interface{} {
	return &requestMsg{}
}

func (e ChannelRequestHandler) HandleRequest(request interface{}, reply channelRequest.Reply, channel ssh.Channel, session backend.Session) {
	e.logger.DebugF("Signal request: %s", request.(*requestMsg).signal)
	//todo should the list of signals allowed be filtered?
	err := session.SendSignal("SIG" + request.(*requestMsg).signal)
	if err != nil {
		e.logger.DebugF("Failed signal request (%s)", err)
		reply(false, nil)
	} else {
		reply(true, nil)
	}
}

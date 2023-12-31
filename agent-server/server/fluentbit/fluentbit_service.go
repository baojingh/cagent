package fluentbit

import (
	"agent-server/command"
	"agent-server/constant"
	logger "agent-server/log"
	"agent-server/pb"
	"agent-server/utils"
	"context"
	"time"
)

var log = logger.New()

type AgentServcie struct {
	pb.UnimplementedAgentActionServer
}

func (agentService *AgentServcie) UpdateFluentbitHost(ctx context.Context,
	request *pb.AgentServiceRequest) (*pb.AgentResponse, error) {
	log.Info("Server receive Request, %v", request)
	res := &pb.AgentResponse{
		Timestamp: time.Now().Format(constant.DATE_FORMAT),
		Status:    pb.Status_OK,
	}

	cmdName := "fluentbit_host"
	paramMap := request.ParamMap
	param := utils.Map2StrWithEqu(paramMap)

	out, err := command.ExecuteShellCmd(cmdName, param)
	if err != nil {
		log.Errorf("Sth wrong executing cmd: %s, err: %v, out: %v", cmdName, err, out)
		res.Status = pb.Status_ERR
		res.Data = err.Error()
		return res, err
	}
	log.Infof("CMD %s is executed success.", cmdName)
	res.Data = out
	return res, nil
}

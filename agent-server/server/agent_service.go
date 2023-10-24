package server

import (
	"agent-server/constant"
	"agent-server/pb"
	"context"
	"time"
)

type AgentServcie struct {
	pb.UnimplementedAgentActionServer
}

func (agentService *AgentServcie) UpdateServiceFile(ctx context.Context,
	request *pb.AgentServiceRequest) (*pb.AgentResponse, error) {

	ip := request.HopstIP
	component := request.Component
	paramMap := request.ParamMap

	log.Infof("request.%s, %s, %v", ip, component, paramMap)

	res := &pb.AgentResponse{
		Data:      "Hello, cobra",
		Timestamp: time.Now().Format(constant.DATE_FORMAT),
		Status:    pb.Status_OK,
	}
	return res, nil
}

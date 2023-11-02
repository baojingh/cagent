package suricata

import (
	"agent-server/constant"
	logger "agent-server/log"
	"agent-server/pb"
	"io"
	"path/filepath"
	"time"
)

var log = logger.New()

type AgentFileServiceServer struct {
	pb.UnimplementedAgentFileServiceServer
}

func (service *AgentFileServiceServer) UploadBigFile(
	stream pb.AgentFileService_UploadBigFileServer) error {
	file := NewFile()
	var fileSize uint32
	fileName := ""
	fileSize = 0
	defer func() {
		if err := file.Close(); err != nil {
			log.Errorf("Fail to close file %s, %v", file.FilePath, err)
			return
		}
	}()
	for {
		req, err := stream.Recv()
		fileName = req.GetFileName()
		if file.FilePath == "" {
			file.SetFile(fileName, constant.DEFAULT_RULE_PATH)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		chunk := req.GetChunk()
		fileSize += uint32(len(chunk))
		if err := file.Write(chunk); err != nil {
			log.Fatal(err)
		}
	}
	name := filepath.Base(file.FilePath)
	err := stream.SendAndClose(
		&pb.AgentResponse{
			Timestamp: time.Now().Format(constant.DATE_FORMAT),
			Status:    pb.Status_OK,
			FileName:  name,
		})
	log.Infof("Receive file %s, file size: %d.", file.FilePath, fileSize)
	return err
}

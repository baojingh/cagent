package suricata

import (
	"agent-server/constant"
	logger "agent-server/log"
	"agent-server/pb"
	"io"
	"time"
)

var log = logger.New()

type FileServiceServer struct {
	pb.UnimplementedAgentFileServiceServer
}

func NewServer() *FileServiceServer {
	s := &FileServiceServer{}
	return s
}

func (g *FileServiceServer) ReceiveBigFile(stream pb.AgentFileService_UploadBigFileServer) error {
	file := NewFile()
	var fileSize uint32
	fileSize = 0
	defer func() {
		if err := file.OutputFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	for {
		req, err := stream.Recv()
		if file.FilePath == "" {
			file.SetFile("rules", "./")
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
	res := &pb.AgentResponse{
		Timestamp: time.Now().Format(constant.DATE_FORMAT),
		Status:    pb.Status_OK,
	}
	err := stream.SendAndClose(res)
	return err

}

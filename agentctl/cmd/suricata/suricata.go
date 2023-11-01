package suricata

import (
	"agentctl/pb"
	"agentctl/utils"
	"context"
	"fmt"
	"io"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AgentFileServiceServer struct {
	pb.UnimplementedAgentFileServiceServer
}

type ClientService struct {
	ip        string
	port      string
	filePath  string
	batchSize int
	client    pb.AgentFileServiceClient
}

func NewService(ip string, port string, filepath string, batchSize int) *ClientService {
	cli := &ClientService{
		batchSize: batchSize,
	}
	return cli
}

func (cli *ClientService) UploadSuricataFile() {
	ip := cli.ip
	port := cli.port
	file := cli.filePath
	isValid := utils.ValidateIP(ip)
	isExist := utils.IsFileExist(file)
	if !isValid {
		log.Errorf("Param ip %s is not validated", ip)
		return
	}
	if !isExist {
		log.Errorf("File %s not exist.", file)
		return
	}
	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Info("Connect to server success.")
	defer conn.Close()
	cli.client = pb.NewAgentFileServiceClient(conn)
	err = cli.doUpload()
	if err != nil {
		log.Error(err)
	}
}

func (cli *ClientService) doUpload() error {
	filePath := cli.filePath
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Fail to open file %s, err: %v", filePath, err)
		return err
	}
	defer file.Close()

	ctx, cancel := context.WithCancel(context.Background())
	stream, err := cli.client.UploadBigFile(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	buf := make([]byte, 1024*1024)
	for {
		br, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error(err)
			return err
		}
		req := &pb.AgentFileRequest{
			Chunk: buf[:br],
		}
		if err := stream.Send(req); err != nil {
			log.Error("Client fail to send data.", err)
			return err
		}
	}
	res, _ := stream.CloseAndRecv()
	cancel()
	log.Infof("Client get res from server, %v", res)
	return nil

}

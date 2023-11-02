package suricata

import (
	logger "agentctl/log"
	"agentctl/pb"
	"agentctl/utils"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	path "path/filepath"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = logger.New()

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
		ip:        ip,
		port:      port,
		filePath:  filepath,
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
	log.Infof("Connect to server %s %s success.", ip, port)
	cli.client = pb.NewAgentFileServiceClient(conn)
	err = cli.doUpload(conn)
	if err != nil {
		log.Error(err)
	}
}

func (cli *ClientService) doUpload(conn *grpc.ClientConn) error {
	filePath := cli.filePath
	file, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Fail to open file %s, err: %v", filePath, err)
		return err
	}
	defer file.Close()
	defer conn.Close()

	filename := path.Base(file.Name())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := cli.client.UploadBigFile(ctx)
	if err != nil {
		log.Error(err)
		return err
	}
	buf := make([]byte, cli.batchSize)
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
			Chunk:    buf[:br],
			FileName: filename,
		}
		if err := stream.Send(req); err != nil {
			log.Error("Client fail to send data.", err)
			return err
		}
	}
	res, _ := stream.CloseAndRecv()
	if strings.TrimSpace(res.GetStatus().String()) != "OK" {
		msg := fmt.Sprintf("Server said that upload file %s failure.", res.GetFileName())
		log.Errorf(msg)
		return errors.New(msg)
	} else {
		log.Infof("Server said that upload file %s success.", res.GetFileName())
	}
	return nil

}

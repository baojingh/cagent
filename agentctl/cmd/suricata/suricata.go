package suricata

import (
	logger "agentctl/log"
	"agentctl/pb"
	"agentctl/utils"
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = logger.New()

var SuricataCmd = &cobra.Command{
	Use:   "suricata",
	Short: "Manage suricata service",
	Long:  ``,
	Run:   UpdateSuricataFile,
}

func UpdateSuricataFile(cmd *cobra.Command, args []string) {
	ip, _ := cmd.Flags().GetString("ip")
	port, _ := cmd.Flags().GetString("port")
	config, _ := cmd.Flags().GetString("config")
	isValid := utils.ValidateIP(ip)
	isEmp := utils.IsEmpty(config)

	if !isValid {
		log.Errorf("Param ip %s is not validated", ip)
		return
	}
	if isEmp {
		log.Error("Param config is empty")
		return
	}
	paramMap := make(map[string]string)

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", ip, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Info("Connect to server success.")
	defer conn.Close()
	c := pb.NewAgentActionClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.AgentServiceRequest{
		HopstIP:   ip,
		Component: "fluentbit",
		ParamMap:  paramMap,
	}
	log.Infof("Client compose cmd params, %v", req)
	r, err := c.UpdateFluentbitHost(ctx, req)
	if err != nil {
		log.Fatalf("%v, %v", err, r)
	}
	log.Printf("Update fluentbit service success, %v", r)
}

func init() {
	SuricataCmd.AddCommand(SuricataUploadCmd)
}

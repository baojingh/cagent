package fluentbit

import (
	"agentctl/constant"
	logger "agentctl/log"
	"agentctl/pb"
	"agentctl/utils"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var log = logger.New()

var FluentbitCmd = &cobra.Command{
	Use:   "fluentbit",
	Short: "Manage fluentbit service",
	Long:  ``,
	Run:   UpdateFluentbitHost,
}

func UpdateFluentbitHost(cmd *cobra.Command, args []string) {
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
	paramMap := splitConfig2Map(config)

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
	r, err := c.UpdateFluentbitHost(ctx, req)
	if err != nil {
		log.Fatalf("%v, %v", err, r)
	}
	log.Printf("Update fluentbit service success, info %v", req)
}

func init() {
	FluentbitCmd.Flags().String(constant.FLUENTBIT_PARAM_CONFIG, "", "fluentbit config")
}

func splitConfig2Map(str string) map[string]string {
	paris := utils.SplitStr(str, ",")
	configMap := make(map[string]string)
	for _, val := range paris {
		equ := utils.SplitStr(val, "=")
		k := strings.TrimSpace(equ[0])
		v := strings.TrimSpace(equ[1])
		configMap[k] = v
	}
	return configMap
}

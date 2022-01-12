package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/Jamshid-Ismoilov/mytaxi-api-gateway/config"
	pb "github.com/Jamshid-Ismoilov/mytaxi-api-gateway/genproto"
)

type IServiceManager interface {
	TaxiService() pb.TaxiServiceClient
}

type serviceManager struct {
	taxiService pb.TaxiServiceClient
}

func (s *serviceManager) TaxiService() pb.TaxiServiceClient {
	return s.taxiService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connTaxi, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.TaxiServiceHost, conf.TaxiServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		taxiService: pb.NewTaxiServiceClient(connTaxi),
	}

	return serviceManager, nil
}

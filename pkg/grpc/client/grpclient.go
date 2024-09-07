package grpclient

import (
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	errMissingGrpcIp          = errors.New("missing grpc ip")
	errMissingGrpcPort        = errors.New("missing grpc port")
	errMakingConnectionToGrpc = errors.New("error while establishing grpc connection")
)

func NewGrpcClient(grpcIp, grpcPort string) (*grpc.ClientConn, error) {
	err := CheckConnectionCredentials(grpcIp, grpcPort)
	if err != nil {
		return nil, err
	}

	conn, err := grpc.NewClient(grpcIp+":"+grpcPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errMakingConnectionToGrpc
	}
	return conn, nil
}

func CheckConnectionCredentials(grpcIP, grpcPort string) error {
	if grpcIP == "" {
		return errMissingGrpcIp
	}
	if grpcPort == "" {
		return errMissingGrpcPort
	}
	return nil
}

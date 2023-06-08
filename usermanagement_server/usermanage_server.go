package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/Amizhthanmd/Golang_Multipleclient_Grpc/usermanagement"
	"google.golang.org/grpc"
)

const (
	Port = ":8081"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received : %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	if in.GetName() == "Amizhthan" {
		return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: 5005}, nil
	} else {
		return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
	}
}

func (s *UserManagementServer) CreatePcSpecs(ctx context.Context, specin *pb.PcSpecs) (*pb.PcSpecsResponse, error) {
	
	log.Printf("Hostname : %v", specin.GetHostname())
	log.Printf("Os Name : %v", specin.GetOsName())
	log.Printf("Os Version : %v", specin.GetOsVersion())
	log.Printf("Os Arch : %v", specin.GetOsArch())
	log.Printf("Number of Cpu Cores : %v", specin.GetNumberCpuCores())
	log.Printf("Total Disk Space : %.2f GB", specin.GetTotalSpace())
	log.Printf("Used Disk Space : %.2f GB", specin.GetUsedSpace())
	log.Printf("Free Disk Space : %.2f GB", specin.GetFreeSpace())
	var returnmsg = "PC Specifications received from Client..."
	return &pb.PcSpecsResponse{Message: returnmsg}, nil
}

func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Printf("Cannot start a server : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})

	log.Printf("Server listening @ : %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to listen : %v", err)
	}
}

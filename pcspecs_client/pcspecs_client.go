package main

import (
	"context"
	_"fmt"
	"log"
	"runtime"

	pb "github.com/Amizhthanmd/Golang_Multipleclient_Grpc/usermanagement"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

const (
	address = "localhost:8081"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to server: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//------------ retrieve host ----------------

	hostInfo, err := host.Info()
	if err != nil {
		log.Printf("Did not retrieve host information")
	}
	/* fmt.Printf("Hostname: %v\n", hostInfo.Hostname)
	fmt.Printf("OS Name: %v\n", hostInfo.Platform)
	fmt.Printf("OS Version: %v %v\n", hostInfo.PlatformVersion)
	fmt.Printf("OS Kernel Version: %v\n", hostInfo.KernelVersion)
	fmt.Printf("OS Arch: %v\n", runtime.GOARCH) */
	numCPU := runtime.NumCPU()
	numCPUCores := int32(numCPU)
	/* fmt.Printf("Number of CPU Cores: %d\n", numCPUCores) */

	//-------------- Retrieve disk information ---------------
	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Printf("Did not retrieve disk information")
	}
	var totalSpaceGB float64
	var usedSpaceGB float64
	var freeSpaceGB float64
	for _, partition := range partitions {
		if partition.Mountpoint == "/System/Volumes/Data" {
			usage, err := disk.Usage(partition.Mountpoint)
			if err == nil {
				totalSpaceGB = float64(usage.Total) / (1024 * 1024 * 1024)
				usedSpaceGB = float64(usage.Used) / (1024 * 1024 * 1024)
				freeSpaceGB = float64(usage.Free) / (1024 * 1024 * 1024)
				/* fmt.Printf("Mount Point: %v\n", partition.Mountpoint)
				fmt.Printf("Total Space: %.2f GB\n", totalSpaceGB)
				fmt.Printf("Used Space: %.2f GB\n", usedSpaceGB)
				fmt.Printf("Free Space: %.2f GB\n", freeSpaceGB) */
			}
		}
	}

	r, err := c.CreatePcSpecs(ctx, &pb.PcSpecs{Hostname: hostInfo.Hostname, OsName: hostInfo.Platform,
		OsVersion: hostInfo.PlatformVersion, OsArch: runtime.GOARCH, NumberCpuCores: numCPUCores,
		TotalSpace: totalSpaceGB, UsedSpace: usedSpaceGB, FreeSpace: freeSpaceGB})

	if err != nil {
		log.Fatalf("could not create user: %v", err)
	}

	log.Printf("Message from server : %v", r.GetMessage())

}

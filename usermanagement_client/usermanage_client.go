package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Amizhthanmd/Golang_Multipleclient_Grpc/usermanagement"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	var newUsers = make(map[string]int32)
	newUsers["Amizhthan"] = 22
	newUsers["Mugesh"] = 23

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				for name, age := range newUsers {
					r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
					if err != nil {
						log.Fatalf("could not create user: %v", err)
					}
					log.Printf("\n New User:\nName: %s\nAge: %d\nID: %d", r.GetName(), r.GetAge(), r.GetId())
				}
			}
			time.Sleep(10 * time.Second)
		}
	}()

	select {}
}

/* package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Amizhthanmd/Golang_grpc/usermanagement"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8081"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did'nt connect to server : %v", err)
	}
	defer conn.Close()

	c := pb.NewUserManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_users = make(map[string]int32)

	new_users["Amizhthan"] = 22
	new_users["Mugesh"] = 23

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("Could not create user : %v", err)
		}
		log.Printf("NewUsers : \n Name : %s \n Age : %d \n Id : %d ", r.GetName(), r.GetAge(), r.GetId())
	}

}

*/

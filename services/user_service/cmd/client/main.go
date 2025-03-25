/*
	This is a simple client to test the user service.
*/

package main

import (
	"context"
	"fmt"
	pb "user_service/pb"

	"github.com/flashhhhh/pkg/env"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost" + env.GetEnv("USER_PORT", ":50051"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx := context.Background()

	/*
		Test GetUserByID
	*/
	// req := &pb.IDRequest{
	// 	Id: 4,
	// }

	// res, err := client.GetUserByID(ctx, req)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("User ID:", res.Id)
	// fmt.Println("Username:", res.Username)
	// fmt.Println("Name:", res.Name)






	/*
		------------------------------------------------------------------------
	*

	/*
		Test Login
	*/

	res, err := client.Login(ctx, &pb.LoginRequest{
		Username: "admin",
		Password: "admin",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Token:", res.Token)







	/*
		------------------------------------------------------------------------
	*/

	/*
		Test GetAllUsers
	*/
	// req := &pb.EmptyRequest{}

	// res, err := client.GetAllUsers(ctx, req)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, user := range res.Users {
	// 	fmt.Println("User ID:", user.Id)
	// 	fmt.Println("Username:", user.Username)
	// 	fmt.Println("Name:", user.Name)
	// 	fmt.Println()
	// }
}
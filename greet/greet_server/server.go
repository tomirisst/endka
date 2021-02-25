package main

import (
	"com.grpc.tleu/greet/greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type Server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v \n", req)

	number := int(req.GetGreeting().GetNumber())
	var str string
	for i := 2;number > i; i++  {
		for number%i == 0 {
			a := strconv.Itoa(int(i))
			str = str + a + ", "
			number = number/i;
		}
	}
	if number >2 {
		num := strconv.Itoa(number)
		str = str + num
	}

	res := &greetpb.GreetResponse{
		Result: str,
	}

	return res, nil
}
func main() {
	l, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})
	log.Println("Server is running on port:50051")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}

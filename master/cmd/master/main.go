package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	apiv1 "aide/proto/pkg/apiv1"
	jobv1 "aide/proto/pkg/jobv1"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

type aideServer struct {
	apiv1.UnimplementedAIDEServer

	// A map from job name -> job
	jobs map[string]*jobv1.Job
}

func newServer() *aideServer {
	s := &aideServer{jobs: make(map[string]*jobv1.Job)}
	return s
}

func (s *aideServer) CreateJob(ctx context.Context, createJobRequest *apiv1.CreateJobRequest) (*apiv1.CreateJobResponse, error) {
	_, ok := s.jobs[createJobRequest.Job.Name]
	if ok {
		return nil, status.Error(codes.AlreadyExists, "Already exists.")
	}
	job := &jobv1.Job{Name: createJobRequest.Job.Name, Type: createJobRequest.Job.Type, Size: createJobRequest.Job.Size}
	s.jobs[job.Name] = job
	return &apiv1.CreateJobResponse{Job: job}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	apiv1.RegisterAIDEServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

package rpc

import (
	pb "github.com/rlvoyer/simple_bazel_project/protos"
)

type Server struct {
	pb.UnimplementedFooServiceServer
}

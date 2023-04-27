package redisrpc

import (
	"redisrpc/rpc"

	"google.golang.org/grpc/metadata"
)

func MakeMetadata(md metadata.MD) *rpc.Metadata {
	if md == nil || md.Len() == 0 {
		return nil
	}
	result := make(map[string]*rpc.Strings, md.Len())
	for key, values := range md {
		result[key] = &rpc.Strings{
			Values: values,
		}
	}
	return &rpc.Metadata{
		Md: result,
	}
}

package igrpc

import (
	"errors"

	"github.com/coldwind/artist/pkg/iutils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 根据addr配置获取grpc的conn 使用后记得关闭
// 关闭代码 defer conn.Close()
func GetGrpcClientHandle(grpcList []string) (*grpc.ClientConn, error) {
	if addr, err := getGrpcAddr(grpcList); err == nil {
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}

		return conn, nil
	} else {
		return nil, err
	}
}

func getGrpcAddr(grpcList []string) (string, error) {
	listLen := len(grpcList)
	if listLen == 0 {
		return "", errors.New("empty grpc")
	}

	iutils.Seed()
	index := iutils.Rand(0, int32(listLen)-1)
	return grpcList[index], nil
}

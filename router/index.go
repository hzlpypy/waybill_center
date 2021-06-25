package router

import (
	"context"
	"fmt"
	"github.com/hzlpypy/waybill_center/waybill_server"
	"github.com/hzlpypy/waybill_center/init_service"
	"google.golang.org/grpc"
	protos "github.com/hzlpypy/waybill_center/proto_info/protos"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"strconv"
)

type initServiceConfig struct {
	grpcPort int
	InitService init_service.InitService

}
func NewServiceConfig()  {

}
//RunGrpcServer creates a gRPC server which has no service registered and has not
func RunGrpcServer(isc *initServiceConfig) error {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(filter))


	protos.RegisterWaybillCenterServer(grpcServer, &waybill_server.WaybillServer{})

	isc.InitService.InitVCTable()
	listen, err := net.Listen("tcp", ":"+strconv.Itoa(isc.grpcPort))
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("User center service has started=> grpc:127.0.0.1:%d", isc.grpcPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Printf("failed to grpc-serve:%v", err)
		return err
	}
	return nil
}

func filter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(500, fmt.Sprintf("panic: %v", r))
			log.Printf("%v", err)
		}
	}()
	logReqUri(ctx, info.FullMethod)
	return handler(ctx, req)
}

func logReqUri(ctx context.Context, reqUri string) {
	var ip string
	md, _ := metadata.FromIncomingContext(ctx)
	contentType := md.Get("content-type")[0]
	if val, ok := md[":authority"]; ok {
		ip = val[0]
		p, ok := peer.FromContext(ctx)
		if ok {
			log.Printf("%v -> %v%v [%v]", p.Addr.String(), ip, reqUri, contentType)
		} else {
			log.Printf("ip is none: %v", ctx)
		}
	} else {
		log.Printf("ip is none: %v", ctx)
	}

}
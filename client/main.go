package main

import (
	"flag"
	"io"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	pb "github.com/ozym/slink-proto/protobuf"
)

var serverAddr = flag.String("server_addr", "127.0.0.1:19000", "The server address in the format of host:port")
var selectors = flag.String("selectors", "???", "The default selectors to request")
var streams = flag.String("streams", "NZ_CAW", "The default streams to request")

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewSeedLinkClient(conn)

	stream, err := client.Stream(context.Background(), &pb.Selection{
		Streams:   *streams,
		Selectors: *selectors,
	})

	if err != nil {
		grpclog.Fatalf("%v.RouteChat(_) = _, %v", client, err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				grpclog.Fatalf("Failed to receive a note : %v", err)
			}
			var v0, v1 int32
			t := time.Unix(int64(in.Start.Nanoseconds)/int64(time.Second), int64(in.Start.Nanoseconds)%int64(time.Second))
			if len(in.Samples) > 0 {
				v0, v1 = in.Samples[0].Value, in.Samples[len(in.Samples)-1].Value
			}
			grpclog.Printf("Got packet: %s_%s_%s_%s %s sps=%g n=%d %d ... %d", in.Network, in.Station, in.Location, in.Channel, t.String(), in.Sps, in.Length, v0, v1)
		}
	}()

	stream.CloseSend()
	<-waitc
}

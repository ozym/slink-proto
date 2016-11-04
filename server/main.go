package main

import (
	"github.com/ozym/mseed"
	"github.com/ozym/slink"

	"flag"
	"fmt"
	pb "github.com/ozym/slink-proto/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"strings"
	"time"
)

var port = flag.Int("port", 19000, "The server port")
var server = flag.String("server", "link.geonet.org.nz:18000", "The Seedlink Server")
var netdly = flag.Int("netdly", 0, "provide network delay")
var netto = flag.Int("netto", 300, "provide network timeout")
var keepalive = flag.Int("keepalive", 0, "provide keep-alive")

type seedlinkServer struct {
}

func (sl *seedlinkServer) Stream(sel *pb.Selection, stream pb.SeedLink_StreamServer) error {

	// initial seedlink handle
	slconn := slink.NewSLCD()
	defer slink.FreeSLCD(slconn)

	// seedlink settings
	slconn.SetNetDly(*netdly)
	slconn.SetNetTo(*netto)
	slconn.SetKeepAlive(*keepalive)

	// conection
	slconn.SetSLAddr(*server)
	defer slconn.Disconnect()

	var streams, selectors string
	if streams = sel.Streams; streams == "" {
		streams = "*_*"
	}
	if selectors = sel.Selectors; selectors == "" {
		selectors = "???"
	}

	// configure streams selectors to recover
	slconn.ParseStreamList(streams, selectors)

	// make space for miniseed blocks
	msr := mseed.NewMSRecord()
	defer mseed.FreeMSRecord(msr)

loop:
	for {
		switch p, rc := slconn.Collect(); rc {
		case slink.SLTERMINATE:
			break loop
		case slink.SLNOPACKET:
			continue loop
		case slink.SLPACKET:
			// just in case we're shutting down
			if p != nil && p.PacketType() == slink.SLDATA {
				msr.Unpack(p.GetMSRecord(), 512, 1, 0)
				sps, cnt := msr.Samprate(), msr.Samplecnt()

				var samples []*pb.Sample
				values, err := msr.DataSamples()
				if err != nil {
					return err
				}
				for i, v := range values {
					dt := time.Duration(float64(time.Duration(i)*time.Second) / float64(sps))
					samples = append(samples, &pb.Sample{
						Value: v,
						Epoch: &pb.Timestamp{
							Nanoseconds: msr.Starttime().Add(dt).UnixNano(),
						},
					})

				}

				if err := stream.Send(&pb.Packet{
					Network:  strings.TrimRight(msr.Network(), "\u0000"),
					Station:  strings.TrimRight(msr.Station(), "\u0000"),
					Location: strings.TrimRight(msr.Location(), "\u0000"),
					Channel:  strings.TrimRight(msr.Channel(), "\u0000"),
					Start: &pb.Timestamp{
						Nanoseconds: msr.Starttime().UnixNano(),
					},
					End: &pb.Timestamp{
						Nanoseconds: msr.Endtime().UnixNano(),
					},
					Length:  cnt,
					Sps:     sps,
					Samples: samples,
				}); err != nil {
					return err
				}
			}
		default:
			break loop
		}

	}

	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterSeedLinkServer(grpcServer, &seedlinkServer{})
	grpcServer.Serve(lis)
}

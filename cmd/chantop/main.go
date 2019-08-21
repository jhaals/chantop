package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/gosuri/uilive"
	"github.com/olekukonko/tablewriter"
	"google.golang.org/grpc"
	z "google.golang.org/grpc/channelz/grpc_channelz_v1"
)

var (
	address = flag.String("address", "localhost:8080", "gRPC server address")
	timeout = flag.Int("timeout", 1, "Timeout in seconds to complete channel requests")
	watch   = flag.Bool("watch", false, "Watch for changes")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	client := z.NewChannelzClient(conn)

	writer := uilive.New()
	writer.Start()

	for {
		ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*time.Duration(*timeout)))

		resp, err := client.GetTopChannels(ctx, &z.GetTopChannelsRequest{})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		tw := tablewriter.NewWriter(writer)
		tw.SetHeader([]string{"Target", "State", "Started", "Succeeded", "Failed", "Last Call"})

		for _, c := range resp.GetChannel() {
			appendTable(c.GetData(), tw)

			for _, sr := range c.GetSubchannelRef() {
				resp, err := client.GetSubchannel(ctx, &z.GetSubchannelRequest{SubchannelId: sr.GetSubchannelId()})
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				appendTable(resp.GetSubchannel().GetData(), tw)
			}
		}

		tw.Render()
		if !*watch {
			break
		}
		time.Sleep(1000 * time.Millisecond)
	}
	writer.Stop()
}

func appendTable(cd *z.ChannelData, tw *tablewriter.Table) {
	tw.Append([]string{
		cd.GetTarget(),
		cd.GetState().GetState().String(),
		fmt.Sprintf("%d", cd.GetCallsStarted()),
		fmt.Sprintf("%d", cd.GetCallsSucceeded()),
		fmt.Sprintf("%d", cd.GetCallsFailed()),
		fmt.Sprintf("%s", ptypes.TimestampString(cd.GetLastCallStartedTimestamp())),
	})
}

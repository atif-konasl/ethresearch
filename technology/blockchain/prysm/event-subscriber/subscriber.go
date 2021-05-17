package main

import (
	"context"
	ptypes "github.com/gogo/protobuf/types"
	"sync"
	"time"
)

func (c *GRPCClient) chainHeadSubscription(wg *sync.WaitGroup) {
	stream, err := c.beaconClient.StreamChainHead(
		c.ctx,
		&ptypes.Empty{},
		)
	if err != nil {
		log.WithError(err).Fatal("Failed to subscribe to StreamChainHead")
	}

	log.Info("Successfully subscribed to chain header event")
	wg.Add(1)
	go func() {
		for {
			chainHeader, err := stream.Recv()
			if err != nil {
				log.WithError(err).Error("Failed to receive chain header")
				return
				wg.Done()
			}
			//chainHeader := data.(*types.ChainHeader)
			log.WithField("headerSlot", chainHeader.HeadSlot).WithField("headBlockRoot", chainHeader.HeadBlockRoot).Info("Got chain header")
		}
	}()
}

func (c *GRPCClient) newPendingBlocksSubscription(wg *sync.WaitGroup) {
	stream, err := c.beaconClient.StreamNewPendingBlocks(
		c.ctx,
		&ptypes.Empty{},
	)
	if err != nil { log.WithError(err).Fatal("Failed to subscribe to StreamPendingBlocks") }

	log.Info("Successfully subscribed to StreamPendingBlocks event")
	wg.Add(1)
	go func() {
		for {
			block, err := stream.Recv()
			if err != nil {
				log.WithError(err).Error("Failed to receive chain header")
				return
				wg.Done()
			}
			blockRoot, err := block.HashTreeRoot()
			if err != nil {
				log.Error("Test")
			}
			log.WithField("proposerIndex", block.ProposerIndex).WithField(
				"blockRoot", blockRoot).WithField("slot", block.Slot).Info("Got block from vanguard")
		}
	}()
}

func main ()  {
	ctx := context.Background()
	wg := new(sync.WaitGroup)

	grpcClient, err := Dial(ctx, "34.90.144.142:4000", 5 * time.Second, 5, 100000)
	if err != nil {
		log.Fatal("failed to initiate grpc client")
	}

	//grpcClient.chainHeadSubscription(wg)
	grpcClient.newPendingBlocksSubscription(wg)
	wg.Wait()
}
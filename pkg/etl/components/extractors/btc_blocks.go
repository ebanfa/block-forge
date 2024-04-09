package extractors

import (
	"fmt"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/edward1christian/block-forge/pkg/application/common/context"
	"github.com/edward1christian/block-forge/pkg/application/common/event"
	"github.com/edward1christian/block-forge/pkg/application/system"
	"github.com/edward1christian/block-forge/pkg/etl/process"
)

// BtcBlockStreamingExtractorETLComponent represents an ETL component for streaming Bitcoin block data.
type BtcBlockStreamingExtractorETLComponent struct {
	id          string
	name        string
	description string
	rpcClient   *rpcclient.Client
	system      system.SystemInterface
}

// NewBtcBlockStreamingExtractorETLComponent creates a new instance of BtcBlockStreamingExtractorETLComponent.
func NewBtcBlockStreamingExtractorETLComponent(rpcClient *rpcclient.Client) *BtcBlockStreamingExtractorETLComponent {
	return &BtcBlockStreamingExtractorETLComponent{
		id:          "btc-block-streaming-extractor",
		name:        "Bitcoin Block Streaming Extractor",
		description: "ETL component for streaming Bitcoin block data",
		rpcClient:   rpcClient,
	}
}

// Initialize initializes the module.
// Returns an error if the initialization fails.
func (c *BtcBlockStreamingExtractorETLComponent) Initialize(ctx *context.Context, sys system.SystemInterface) error {
	// No additional initialization required, the RPC client and event bus are already set up
	return nil
}

// Start starts the component.
// Returns an error if the start operation fails.
func (c *BtcBlockStreamingExtractorETLComponent) Start(ctx *context.Context) error {
	recordChannel, err := c.OpenStream(ctx, source)
	if err != nil {
		return err
	}
	go func() {
		for record := range recordChannel {
			c.system.EventBus().Publish(event.Event{
				Type: event.EventTypeDataExtracted,
				Data: record,
			})
		}
	}()

	return nil
}

// Stop stops the component.
// Returns an error if the stop operation fails.
func (c *BtcBlockStreamingExtractorETLComponent) Stop(ctx *context.Context) error {
	// No additional actions required to stop the component
	return nil
}

// OpenStream establishes a streaming connection with the Bitcoin source and returns a channel for receiving records.
func (c *BtcBlockStreamingExtractorETLComponent) OpenStream(ctx *context.Context, source process.SourceInterface) (<-chan process.RecordInterface, error) {
	// Create a channel for receiving records
	recordChannel := make(chan process.RecordInterface)

	go func() {
		defer close(recordChannel)

		for {
			// Get the hash of the latest block
			bestBlockHash, err := c.rpcClient.GetBestBlockHash()
			if err != nil {
				fmt.Printf("Error getting best block hash: %v\n", err)
				return
			}

			// Fetch the latest block from the Bitcoin RPC client asynchronously
			future := c.rpcClient.GetBlockVerboseAsync(bestBlockHash)
			block, err := future.Receive()
			if err != nil {
				fmt.Printf("Error fetching block: %v\n", err)
				return
			}

			// Create a record with the block data
			record := &RecordData{
				Data:     block,
				Metadata: map[string]interface{}{"source": source.GetID()},
			}

			select {
			case <-ctx.Done():
				return
			case recordChannel <- record:
			}
		}
	}()

	return recordChannel, nil
}

// RecordData is an implementation of the RecordInterface.
type RecordData struct {
	process.RecordInterface
	Data     interface{}
	Metadata map[string]interface{}
}

func (r *RecordData) GetData() interface{} {
	return r.Data
}

func (r *RecordData) GetMetadata() map[string]interface{} {
	return r.Metadata
}

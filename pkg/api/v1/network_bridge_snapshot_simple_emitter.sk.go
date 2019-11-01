// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	fmt "fmt"
	"time"

	"go.opencensus.io/stats"

	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

type NetworkBridgeSimpleEmitter interface {
	Snapshots(ctx context.Context) (<-chan *NetworkBridgeSnapshot, <-chan error, error)
}

func NewNetworkBridgeSimpleEmitter(aggregatedWatch clients.ResourceWatch) NetworkBridgeSimpleEmitter {
	return NewNetworkBridgeSimpleEmitterWithEmit(aggregatedWatch, make(chan struct{}))
}

func NewNetworkBridgeSimpleEmitterWithEmit(aggregatedWatch clients.ResourceWatch, emit <-chan struct{}) NetworkBridgeSimpleEmitter {
	return &networkBridgeSimpleEmitter{
		aggregatedWatch: aggregatedWatch,
		forceEmit:       emit,
	}
}

type networkBridgeSimpleEmitter struct {
	forceEmit       <-chan struct{}
	aggregatedWatch clients.ResourceWatch
}

func (c *networkBridgeSimpleEmitter) Snapshots(ctx context.Context) (<-chan *NetworkBridgeSnapshot, <-chan error, error) {
	snapshots := make(chan *NetworkBridgeSnapshot)
	errs := make(chan error)

	untyped, watchErrs, err := c.aggregatedWatch(ctx)
	if err != nil {
		return nil, nil, err
	}

	go errutils.AggregateErrs(ctx, errs, watchErrs, "networkBridge-emitter")

	go func() {
		currentSnapshot := NetworkBridgeSnapshot{}
		timer := time.NewTicker(time.Second * 1)
		var previousHash uint64
		sync := func() {
			currentHash := currentSnapshot.Hash()
			if previousHash == currentHash {
				return
			}

			previousHash = currentHash

			stats.Record(ctx, mNetworkBridgeSnapshotOut.M(1))
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		defer func() {
			close(snapshots)
			close(errs)
		}()

		for {
			record := func() { stats.Record(ctx, mNetworkBridgeSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case untypedList := <-untyped:
				record()

				currentSnapshot = NetworkBridgeSnapshot{}
				for _, res := range untypedList {
					switch typed := res.(type) {
					case *MeshBridge:
						currentSnapshot.MeshBridges = append(currentSnapshot.MeshBridges, typed)
					default:
						select {
						case errs <- fmt.Errorf("NetworkBridgeSnapshotEmitter "+
							"cannot process resource %v of type %T", res.GetMetadata().Ref(), res):
						case <-ctx.Done():
							return
						}
					}
				}

			}
		}
	}()
	return snapshots, errs, nil
}

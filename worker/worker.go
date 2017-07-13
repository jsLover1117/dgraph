/*
 * Copyright (C) 2017 Dgraph Labs, Inc. and Contributors
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// Package worker contains code for internal worker communication to perform
// queries and mutations.
package worker

import (
	"fmt"
	"log"
	"math"
	"net"
	"sync"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/group"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	pstore           *badger.KV
	workerServer     *grpc.Server
	leaseGid         uint32
	pendingProposals chan struct{}
	// In case of flaky network connectivity we would try to keep upto maxPendingEntries in wal
	// so that the nodes which have lagged behind leader can just replay entries instead of
	// fetching snapshot if network disconnectivity is greater than the interval at which snapshots
	// are taken

	emptyMembershipUpdate protos.MembershipUpdate
)

func workerPort() int {
	return x.Config.PortOffset + Config.BaseWorkerPort
}

func Init(ps *badger.KV) {
	pstore = ps
	// needs to be initialized after group config
	leaseGid = group.BelongsTo("_lease_")
	pendingProposals = make(chan struct{}, Config.NumPendingProposals)
	workerServer = grpc.NewServer(
		grpc.MaxRecvMsgSize(x.GrpcMaxSize),
		grpc.MaxSendMsgSize(x.GrpcMaxSize),
		grpc.MaxConcurrentStreams(math.MaxInt32))
}

// grpcWorker struct implements the gRPC server interface.
type grpcWorker struct {
	sync.Mutex
	reqids map[uint64]bool
}

// addIfNotPresent returns false if it finds the reqid already present.
// Otherwise, adds the reqid in the list, and returns true.
func (w *grpcWorker) addIfNotPresent(reqid uint64) bool {
	w.Lock()
	defer w.Unlock()
	if w.reqids == nil {
		w.reqids = make(map[uint64]bool)
	} else if _, has := w.reqids[reqid]; has {
		return false
	}
	w.reqids[reqid] = true
	return true
}

// Hello rpc call is used to check connection with other workers after worker
// tcp server for this instance starts.
func (w *grpcWorker) Echo(ctx context.Context, in *protos.Payload) (*protos.Payload, error) {
	return &protos.Payload{Data: in.Data}, nil
}

// RunServer initializes a tcp server on port which listens to requests from
// other workers for internal communication.
func RunServer(bindall bool) {
	laddr := "localhost"
	if bindall {
		laddr = "0.0.0.0"
	} else if len(Config.MyAddr) > 0 {
		fmt.Printf("--my flag is provided without bindall, Did you forget to specify bindall?\n")
	}
	var err error
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", laddr, workerPort()))
	if err != nil {
		log.Fatalf("While running server: %v", err)
		return
	}
	log.Printf("Worker listening at address: %v", ln.Addr())

	protos.RegisterWorkerServer(workerServer, &grpcWorker{})
	workerServer.Serve(ln)
}

// StoreStats returns stats for data store.
func StoreStats() string {
	return "Currently no stats for badger"
}

// BlockingStop stops all the nodes, server between other workers and syncs all marks.
func BlockingStop() {
	stopAllNodes()              // blocking stop all nodes
	workerServer.GracefulStop() // blocking stop server
	// blocking sync all marks
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	if err := syncAllMarks(ctx); err != nil {
		x.Printf("Error in sync watermarks : %s", err.Error())
	}
	snapshotAll()
}

package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	cbor "github.com/ipfs/go-ipld-cbor"
)

func main() {
	var api lotusapi.FullNodeStruct
	head := http.Header{}

	closer, err := jsonrpc.NewMergeClient(
		context.Background(),
		"https://api.calibration.node.glif.io/rpc/v1",
		"Filecoin",
		lotusapi.GetInternalStructs(&api),
		head,
	)

	if err != nil {
		panic(err)
	}
	defer closer()

	minerAddr, err := address.NewFromString("f01001")

	err = build.UseNetworkBundle("calibrationnet")
	if err != nil {
		panic(err)
	}

	actor, err := api.StateGetActor(context.Background(), minerAddr, types.TipSetKey{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Miner actor code from StateGetActor: %s\n", actor.Code.String())

	actorCode, success := actors.GetActorCodeID(actorstypes.Version(actors.LatestVersion), manifest.MinerKey)
	if !success {
		panic("failed to get actor codes")
	}
	fmt.Printf("Miner actor code from GetActorCodeID: %s\n", actorCode)

	tbs := blockstore.NewTieredBstore(blockstore.NewAPIBlockstore(&api), blockstore.NewMemory())

	state, err := miner.Load(adt.WrapStore(context.Background(), cbor.NewCborStore(tbs)), actor)
	if err != nil {
		panic(err)
	}

	fmt.Println("State", state)
}

package main

import (
	"context"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/venus/venus-shared/actors/builtin"
	sharedTypes "github.com/filecoin-project/venus/venus-shared/types"

	types "github.com/filecoin-project/venus/venus-shared/types/gateway"
	"github.com/ipfs-force-community/venus-gateway/marketevent"
	"github.com/ipfs-force-community/venus-gateway/proofevent"
	"github.com/ipfs-force-community/venus-gateway/walletevent"
)

type IGatewayPushAPI interface {
	proofevent.IProofEvent
	walletevent.IWalletEvent
}

type IGatewayAPI interface {
	proofevent.IProofEventAPI
	walletevent.IWalletEventAPI
	IGatewayPushAPI
}

var _ IGatewayAPI = (*GatewayAPI)(nil)

type GatewayAPI struct {
	proofevent.IProofEventAPI
	pe *proofevent.ProofEventStream

	walletevent.IWalletEventAPI
	we *walletevent.WalletEventStream

	marketevent.IMarketEventAPI
	me *marketevent.MarketEventStream
}

func NewGatewayAPI(pe *proofevent.ProofEventStream, we *walletevent.WalletEventStream, me *marketevent.MarketEventStream) *GatewayAPI {
	return &GatewayAPI{
		IProofEventAPI:  proofevent.NewProofEventAPI(pe),
		IWalletEventAPI: walletevent.NewWalletEventAPI(we),
		IMarketEventAPI: marketevent.NewMarketEventAPI(me),
		pe:              pe,
		we:              we,
		me:              me,
	}
}

func (g *GatewayAPI) ComputeProof(ctx context.Context, miner address.Address, sectorInfos []builtin.ExtendedSectorInfo, rand abi.PoStRandomness, height abi.ChainEpoch, nwVersion network.Version) ([]builtin.PoStProof, error) {
	return g.pe.ComputeProof(ctx, miner, sectorInfos, rand, height, nwVersion)
}

func (g *GatewayAPI) ListConnectedMiners(ctx context.Context) ([]address.Address, error) {
	return g.pe.ListConnectedMiners(ctx)
}

func (g *GatewayAPI) ListMinerConnection(ctx context.Context, addr address.Address) (*types.MinerState, error) {
	return g.pe.ListMinerConnection(ctx, addr)
}

func (g *GatewayAPI) WalletHas(ctx context.Context, supportAccount string, addr address.Address) (bool, error) {
	return g.we.WalletHas(ctx, supportAccount, addr)
}

func (g *GatewayAPI) WalletSign(ctx context.Context, account string, addr address.Address, toSign []byte, meta sharedTypes.MsgMeta) (*crypto.Signature, error) {
	return g.we.WalletSign(ctx, account, addr, toSign, meta)
}

func (g *GatewayAPI) ListWalletInfo(ctx context.Context) ([]*types.WalletDetail, error) {
	return g.we.ListWalletInfo(ctx)
}

func (g *GatewayAPI) ListWalletInfoByWallet(ctx context.Context, wallet string) (*types.WalletDetail, error) {
	return g.we.ListWalletInfoByWallet(ctx, wallet)
}

func (g *GatewayAPI) IsUnsealed(ctx context.Context, miner address.Address, pieceCid cid.Cid, sector storage.SectorRef, offset sharedTypes.PaddedByteIndex, size abi.PaddedPieceSize) (bool, error) {
	return g.me.IsUnsealed(ctx, miner, pieceCid, sector, offset, size)
}

func (g *GatewayAPI) SectorsUnsealPiece(ctx context.Context, miner address.Address, pieceCid cid.Cid, sector storage.SectorRef, offset sharedTypes.PaddedByteIndex, size abi.PaddedPieceSize, dest string) error {
	return g.me.SectorsUnsealPiece(ctx, miner, pieceCid, sector, offset, size, dest)
}

func (g *GatewayAPI) ListMarketConnectionsState(ctx context.Context) ([]types.MarketConnectionState, error) {
	return g.me.ListMarketConnectionsState(ctx)
}

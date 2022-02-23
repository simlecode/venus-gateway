package walletevent

import (
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	sharedTypes "github.com/filecoin-project/venus/venus-shared/types"
	types "github.com/filecoin-project/venus/venus-shared/types/gateway"
)

type IWalletEvent interface {
	ListWalletInfo(ctx context.Context) ([]*types.WalletDetail, error)
	ListWalletInfoByWallet(ctx context.Context, wallet string) (*types.WalletDetail, error)

	WalletHas(ctx context.Context, supportAccount string, addr address.Address) (bool, error)
	WalletSign(ctx context.Context, account string, addr address.Address, toSign []byte, meta sharedTypes.MsgMeta) (*crypto.Signature, error)
}

type IWalletEventAPI interface {
	ResponseWalletEvent(ctx context.Context, resp *types.ResponseEvent) error
	ListenWalletEvent(ctx context.Context, policy *types.WalletRegisterPolicy) (<-chan *types.RequestEvent, error)
	SupportNewAccount(ctx context.Context, channelId sharedTypes.UUID, account string) error
	AddNewAddress(ctx context.Context, channelId sharedTypes.UUID, newAddrs []address.Address) error
	RemoveAddress(ctx context.Context, channelId sharedTypes.UUID, newAddrs []address.Address) error
}

var _ IWalletEventAPI = (*WalletEventAPI)(nil)

type WalletEventAPI struct {
	walletEvent *WalletEventStream
}

func NewWalletEventAPI(walletEvent *WalletEventStream) *WalletEventAPI {
	return &WalletEventAPI{walletEvent: walletEvent}
}

func (w *WalletEventAPI) ResponseWalletEvent(ctx context.Context, resp *types.ResponseEvent) error {
	return w.walletEvent.ResponseEvent(ctx, resp)
}

func (w *WalletEventAPI) ListenWalletEvent(ctx context.Context, policy *types.WalletRegisterPolicy) (<-chan *types.RequestEvent, error) {
	return w.walletEvent.ListenWalletEvent(ctx, policy)
}

func (w *WalletEventAPI) SupportNewAccount(ctx context.Context, channelId sharedTypes.UUID, account string) error {
	return w.walletEvent.SupportNewAccount(ctx, channelId, account)
}

func (w *WalletEventAPI) AddNewAddress(ctx context.Context, channelId sharedTypes.UUID, newAddrs []address.Address) error {
	return w.walletEvent.AddNewAddress(ctx, channelId, newAddrs)
}

func (w *WalletEventAPI) RemoveAddress(ctx context.Context, channelId sharedTypes.UUID, newAddrs []address.Address) error {
	return w.walletEvent.RemoveAddress(ctx, channelId, newAddrs)
}

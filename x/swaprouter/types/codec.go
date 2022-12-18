package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

// RegisterLegacyAminoCodec registers the necessary x/swaprouter interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*PoolI)(nil), nil)
	cdc.RegisterConcrete(&MsgSwapExactAmountIn{}, "osmosis/swaprouter/swap-exact-amount-in", nil)
	cdc.RegisterConcrete(&MsgSwapExactAmountOut{}, "osmosis/swaprouter/swap-exact-amount-out", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterInterface(
		"osmosis.swaprouter.v1beta1.PoolI",
		(*PoolI)(nil),
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgSwapExactAmountIn{},
		&MsgSwapExactAmountOut{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	sdk.RegisterLegacyAminoCodec(amino)
	RegisterLegacyAminoCodec(authzcodec.Amino)

	amino.Seal()
}
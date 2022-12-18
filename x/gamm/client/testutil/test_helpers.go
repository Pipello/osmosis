package testutil

import (
	"fmt"

	gammcli "github.com/osmosis-labs/osmosis/v13/x/gamm/client/cli"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// commonArgs is args for CLI test commands.
var commonArgs = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
}

// MsgJoinPool broadcast pool join message.
func MsgJoinPool(clientCtx client.Context, owner fmt.Stringer, poolID uint64, shareAmtOut string, maxAmountsIn []string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%d", gammcli.FlagPoolId, poolID),
		fmt.Sprintf("--%s=%s", gammcli.FlagShareAmountOut, shareAmtOut),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, owner.String()),
	}

	for _, maxAmt := range maxAmountsIn {
		args = append(args, fmt.Sprintf("--%s=%s", gammcli.FlagMaxAmountsIn, maxAmt))
	}
	args = append(args, commonArgs...)
	return clitestutil.ExecTestCLICmd(clientCtx, gammcli.NewJoinPoolCmd(), args)
}

// MsgExitPool broadcast a pool exit message.
func MsgExitPool(clientCtx client.Context, owner fmt.Stringer, poolID uint64, shareAmtIn string, minAmountsOut []string, extraArgs ...string) (testutil.BufferWriter, error) {
	args := []string{
		fmt.Sprintf("--%s=%d", gammcli.FlagPoolId, poolID),
		fmt.Sprintf("--%s=%s", gammcli.FlagShareAmountIn, shareAmtIn),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, owner.String()),
	}

	for _, maxAmt := range minAmountsOut {
		args = append(args, fmt.Sprintf("--%s=%s", gammcli.FlagMinAmountsOut, maxAmt))
	}

	args = append(args, commonArgs...)
	return clitestutil.ExecTestCLICmd(clientCtx, gammcli.NewExitPoolCmd(), args)
}
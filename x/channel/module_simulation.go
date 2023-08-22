package channel

import (
	"math/rand"

	"channel/testutil/sample"
	channelsimulation "channel/x/channel/simulation"
	"channel/x/channel/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = channelsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgOpenchannel = "op_weight_msg_openchannel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgOpenchannel int = 100

	opWeightMsgClosechannel = "op_weight_msg_closechannel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClosechannel int = 100

	opWeightMsgCommitment = "op_weight_msg_commitment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCommitment int = 100

	opWeightMsgWithdrawTimelock = "op_weight_msg_withdraw_timelock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawTimelock int = 100

	opWeightMsgWithdrawHashlock = "op_weight_msg_withdraw_hashlock"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawHashlock int = 100

	opWeightMsgFund = "op_weight_msg_fund"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFund int = 100

	opWeightMsgAcceptfund = "op_weight_msg_acceptfund"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAcceptfund int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	channelGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&channelGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgOpenchannel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgOpenchannel, &weightMsgOpenchannel, nil,
		func(_ *rand.Rand) {
			weightMsgOpenchannel = defaultWeightMsgOpenchannel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgOpenchannel,
		channelsimulation.SimulateMsgOpenchannel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgClosechannel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClosechannel, &weightMsgClosechannel, nil,
		func(_ *rand.Rand) {
			weightMsgClosechannel = defaultWeightMsgClosechannel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClosechannel,
		channelsimulation.SimulateMsgClosechannel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCommitment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCommitment, &weightMsgCommitment, nil,
		func(_ *rand.Rand) {
			weightMsgCommitment = defaultWeightMsgCommitment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCommitment,
		channelsimulation.SimulateMsgCommitment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawTimelock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawTimelock, &weightMsgWithdrawTimelock, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawTimelock = defaultWeightMsgWithdrawTimelock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawTimelock,
		channelsimulation.SimulateMsgWithdrawTimelock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawHashlock int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawHashlock, &weightMsgWithdrawHashlock, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawHashlock = defaultWeightMsgWithdrawHashlock
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawHashlock,
		channelsimulation.SimulateMsgWithdrawHashlock(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFund int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFund, &weightMsgFund, nil,
		func(_ *rand.Rand) {
			weightMsgFund = defaultWeightMsgFund
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFund,
		channelsimulation.SimulateMsgFund(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAcceptfund int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAcceptfund, &weightMsgAcceptfund, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptfund = defaultWeightMsgAcceptfund
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAcceptfund,
		channelsimulation.SimulateMsgAcceptfund(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

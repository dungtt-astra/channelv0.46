package types

import (
	"testing"

	"channel/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgSenderwithdrawtimelock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSenderwithdrawtimelock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgSenderwithdrawtimelock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgSenderwithdrawtimelock{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

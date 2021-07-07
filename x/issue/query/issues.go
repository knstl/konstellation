package query

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/konstellation/konstellation/x/issue/keeper"
	"github.com/konstellation/konstellation/x/issue/types"
)

func Issues(ctx sdk.Context, k keeper.Keeper, data []byte) ([]byte, error) {
	var params types.IssuesParams
	if err := k.GetCodec().UnmarshalBinaryBare(data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, sdkerrors.ErrJSONUnmarshal.Error())
	}

	issueList := types.CoinIssueList{CoinIssues: []*types.CoinIssue{}}
	issues := k.List(ctx, params)
	for _, issue := range issues {
		issueList.CoinIssues = append(issueList.CoinIssues, issue)
	}
	bz, err := k.GetCodec().MarshalBinaryBare(&issueList)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, sdkerrors.ErrJSONMarshal.Error())
	}

	return bz, nil
}

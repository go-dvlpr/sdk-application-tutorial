package nameservice

import (
	"github.com/BigCodilo/sdk-application-tutorial/x/nameservice/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewMsgBuyName = types.NewMsgBuyName
	NewMsgSetName = types.NewMsgSetName
	NewMsgCreateUser = types.NewMsgCreateUser
	NewMsgSend		= types.NewMsgSend
	NewWhois      = types.NewWhois
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
	//NewXtsDump	  = types.NewTxsDump
)

type (
	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	MsgCreateUser   = types.MsgCreateUser
	MsgSend			= types.MsgSend
	//Account 		= types.Account
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois
	TxsDump			= types.TxsDump
)

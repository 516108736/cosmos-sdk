package auth

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
	"github.com/cosmos/cosmos-sdk/store/iavl"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
)

// AccountKeeper encodes/decodes accounts using the go-amino (binary)
// encoding/decoding library.
type AccountKeeper struct {
	// The (unexposed) key used to access the store from the Context.
	key sdk.StoreKey

	// The prototypical Account constructor.
	proto func() exported.Account

	// The codec codec for binary encoding/decoding of accounts.
	cdc *codec.Codec

	paramSubspace subspace.Subspace
}

// NewAccountKeeper returns a new sdk.AccountKeeper that uses go-amino to
// (binary) encode and decode concrete sdk.Accounts.
// nolint
func NewAccountKeeper(
	cdc *codec.Codec, key sdk.StoreKey, paramstore subspace.Subspace, proto func() exported.Account,
) AccountKeeper {

	return AccountKeeper{
		key:           key,
		proto:         proto,
		cdc:           cdc,
		paramSubspace: paramstore.WithKeyTable(types.ParamKeyTable()),
	}
}

// Logger returns a module-specific logger.
func (ak AccountKeeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// NewAccountWithAddress implements sdk.AccountKeeper.
func (ak AccountKeeper) NewAccountWithAddress(ctx sdk.Context, addr sdk.AccAddress) exported.Account {
	acc := ak.proto()
	err := acc.SetAddress(addr)
	if err != nil {
		// Handle w/ #870
		panic(err)
	}
	err = acc.SetAccountNumber(ak.GetNextAccountNumber(ctx))
	if err != nil {
		// Handle w/ #870
		panic(err)
	}
	return acc
}

// NewAccount creates a new account
func (ak AccountKeeper) NewAccount(ctx sdk.Context, acc exported.Account) exported.Account {
	if err := acc.SetAccountNumber(ak.GetNextAccountNumber(ctx)); err != nil {
		panic(err)
	}
	return acc
}

// GetAccount implements sdk.AccountKeeper.
func (ak AccountKeeper) GetAccount(ctx sdk.Context, addr sdk.AccAddress) exported.Account {

	store := ctx.KVStore(ak.key)
	bz := store.Get(types.AddressStoreKey(addr))
	if bz == nil {
		return nil
	}

	acc := ak.DecodeAccount(bz)
	return acc
}

func (ak AccountKeeper) GetAccounts(ctx sdk.Context, from int,to int, length int,random bool,shouldCoins sdk.Coins) []exported.Account {
	fmt.Println("IAVL GET Start  from",from,"to",to,"length",length,"random",random)
	ts:=time.Now()
	store := ctx.KVStore(ak.key)
	data:=sdk.AccAddress{}
	for index := 0; index < length; index++ {
		ii:=index+from
		if random{
			ii=rand.Intn(to-from)+from
		}

		data=sdk.IntToAccAddress(ii)
		bz := store.Get(types.AddressStoreKey(sdk.IntToAccAddress(ii)))
		if len(bz) == 0 {
			fmt.Println("Idex", hex.EncodeToString(sdk.IntToAccAddress(ii)), ii)
			panic("sb-GetAccounts")
		}
	}

	if !ak.GetAccount(ctx,data).GetCoins().IsEqual(shouldCoins){
		fmt.Print("???????????")
		panic("GetAccount check: should equal here")
	}
	fmt.Println("IAVL GET End ",time.Now().Sub(ts).Seconds())
	return nil
}

// GetAllAccounts returns all accounts in the accountKeeper.
func (ak AccountKeeper) GetAllAccounts(ctx sdk.Context) []exported.Account {
	accounts := []exported.Account{}
	appendAccount := func(acc exported.Account) (stop bool) {
		accounts = append(accounts, acc)
		return false
	}
	ak.IterateAccounts(ctx, appendAccount)
	return accounts
}

// SetAccount implements sdk.AccountKeeper.
func (ak AccountKeeper) SetAccount(ctx sdk.Context, acc exported.Account) {
	addr := acc.GetAddress()
	store := ctx.KVStore(ak.key)
	bz, err := ak.cdc.MarshalBinaryBare(acc)
	if err != nil {
		panic(err)
	}
	store.Set(types.AddressStoreKey(addr), bz)
}


func (ak AccountKeeper)SetAccounts(ctx sdk.Context,from int, to int,length int, newCoin sdk.Coins,random bool)  {
	fmt.Println("IAVL UPDATE Start from",from,"to",to,"length",length,"random",random)
	store := ctx.KVStore(ak.key)
	ts := time.Now()
	data:=sdk.AccAddress{}
	for index:=0;index<length;index++{
		ii:=index+from
		if random{
			ii=rand.Intn(to-from)+from
		}
		addr := sdk.AccAddress(sdk.IntToAccAddress(ii))
		data=addr
		acc := BaseAccount{
			Address: addr,
			Coins:   newCoin,
		}
		bz, err := ak.cdc.MarshalBinaryBare(acc)
		if err != nil {
			panic(err)
		}
		store.Set(types.AddressStoreKey(addr), bz)

		if index != 0 && index%100000 == 0 {
			fmt.Println("SetAccounts handle index",index,time.Now().Sub(ts).Seconds())
		}
		if index != 0 && index%10000000 == 0 {
			commitID := store.(*iavl.Store).Commit()
			fmt.Println("SetAccounts commit index", index, hex.EncodeToString(commitID.Hash), time.Now().Sub(ts).Seconds())
		}

	}
	store.(*iavl.Store).Commit()
	fmt.Println("IAVL UPDATE End",time.Now().Sub(ts).Seconds())
	if ak.GetAccount(ctx,data)==nil{
		panic("SetAccounts check : should !=nil")
	}
}


// RemoveAccount removes an account for the account mapper store.
// NOTE: this will cause supply invariant violation if called
func (ak AccountKeeper) RemoveAccount(ctx sdk.Context, acc exported.Account) {
	addr := acc.GetAddress()
	store := ctx.KVStore(ak.key)
	store.Delete(types.AddressStoreKey(addr))
}


// RemoveAccount removes an account for the account mapper store.
// NOTE: this will cause supply invariant violation if called
func (ak AccountKeeper) RemoveAccounts(ctx sdk.Context,from int,to int,length int ,random bool) {
	ts:=time.Now()
	fmt.Println("IAVL Delete  Start from",from,"to",to,"length",length,"random",random)


	data:=sdk.AccAddress{}

	store := ctx.KVStore(ak.key)
	for index := 0; index < length; index++ {
		ii:=index+from
		if random{

			ii=rand.Intn(to-from)+from
		}
		addr := sdk.AccAddress(sdk.IntToAccAddress(ii))
		data=addr
		store.Delete(types.AddressStoreKey(addr))



	}
	store.(*iavl.Store).Commit()
	fmt.Println("IAVL Delete End",time.Now().Sub(ts).Seconds())

	if ak.GetAccount(ctx,data)!=nil{
		panic("RemoveAccounts check: should nil here")
	}

}


// IterateAccounts implements sdk.AccountKeeper.
func (ak AccountKeeper) IterateAccounts(ctx sdk.Context, process func(exported.Account) (stop bool)) {
	store := ctx.KVStore(ak.key)
	iter := sdk.KVStorePrefixIterator(store, types.AddressStoreKeyPrefix)
	defer iter.Close()
	for {
		if !iter.Valid() {
			return
		}
		val := iter.Value()
		acc := ak.DecodeAccount(val)
		if process(acc) {
			return
		}
		iter.Next()
	}
}

// GetPubKey Returns the PubKey of the account at address
func (ak AccountKeeper) GetPubKey(ctx sdk.Context, addr sdk.AccAddress) (crypto.PubKey, sdk.Error) {
	acc := ak.GetAccount(ctx, addr)
	if acc == nil {
		return nil, sdk.ErrUnknownAddress(fmt.Sprintf("account %s does not exist", addr))
	}
	return acc.GetPubKey(), nil
}

// GetSequence Returns the Sequence of the account at address
func (ak AccountKeeper) GetSequence(ctx sdk.Context, addr sdk.AccAddress) (uint64, sdk.Error) {
	acc := ak.GetAccount(ctx, addr)
	if acc == nil {
		return 0, sdk.ErrUnknownAddress(fmt.Sprintf("account %s does not exist", addr))
	}
	return acc.GetSequence(), nil
}

// GetNextAccountNumber Returns and increments the global account number counter
func (ak AccountKeeper) GetNextAccountNumber(ctx sdk.Context) uint64 {
	var accNumber uint64
	store := ctx.KVStore(ak.key)
	bz := store.Get(types.GlobalAccountNumberKey)
	if bz == nil {
		accNumber = 0
	} else {
		err := ak.cdc.UnmarshalBinaryLengthPrefixed(bz, &accNumber)
		if err != nil {
			panic(err)
		}
	}

	bz = ak.cdc.MustMarshalBinaryLengthPrefixed(accNumber + 1)
	store.Set(types.GlobalAccountNumberKey, bz)

	return accNumber
}

// -----------------------------------------------------------------------------
// Params

// SetParams sets the auth module's parameters.
func (ak AccountKeeper) SetParams(ctx sdk.Context, params types.Params) {
	ak.paramSubspace.SetParamSet(ctx, &params)
}

// GetParams gets the auth module's parameters.
func (ak AccountKeeper) GetParams(ctx sdk.Context) (params types.Params) {
	ak.paramSubspace.GetParamSet(ctx, &params)
	return
}

// -----------------------------------------------------------------------------
// Misc.

func (ak AccountKeeper) DecodeAccount(bz []byte) (acc exported.Account) {
	err := ak.cdc.UnmarshalBinaryBare(bz, &acc)
	if err != nil {
		panic(err)
	}
	return
}

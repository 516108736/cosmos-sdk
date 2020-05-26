package solomachine

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clientexported "github.com/cosmos/cosmos-sdk/x/ibc/02-client/exported"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/02-client/types"
)

// CheckMisbehaviourAndUpdateState determines whether or not the currently registered
// public key signed over two different messages with the same sequence. If this is true
// the client state is updated to a frozen status.
func CheckMisbehaviourAndUpdateState(
	clientState clientexported.ClientState,
	consensusState clientexported.ConsensusState,
	misbehaviour clientexported.Misbehaviour,
) (clientexported.ClientState, error) {

	// cast the interface to specific types before checking for misbehaviour
	smClientState, ok := clientState.(ClientState)
	if !ok {
		return nil, sdkerrors.Wrapf(clienttypes.ErrInvalidClientType, "client state type %T is not solo machine", clientState)
	}

	smConsensusState, ok := consensusState.(ConsensusState)
	if !ok {
		return nil, sdkerrors.Wrapf(clienttypes.ErrInvalidConsensus, "consensus state type %T, expected %T", consensusState, ConsensusState{})
	}

	evidence, ok := misbehaviour.(Evidence)
	if !ok {
		return nil, sdkerrors.Wrapf(clienttypes.ErrInvalidClientType, "evidence type %T is not solo machine", misbehaviour)
	}

	if smClientState.IsFrozen() {
		return nil, sdkerrors.Wrapf(clienttypes.ErrClientFrozen, "client is already frozen")
	}

	// verify evidence
	if err := checkMisbehaviour(smClientState, consensusState, evidence); err != nil {
		return nil, err
	}

	smClientState.Frozen = true
	return smClientState, nil
}

// checkMisbehaviour checks if the currently registered public key has signed
// over two different messages at the same sequence.
func checkMisbehaviour(clientState ClientState, consensusState ConsensusState, evidence Evidence) error {
	pubKey := consensusState.PubKey

	// assert that provided signature data are different
	if bytes.Equal(evidence.SignatureOne.Data, evidence.SignatureTwo.Data) {
		return sdkerrors.Wrap(clienttypes.ErrInvalidEvidence, "evidence signatures have identical data messages")
	}

	data := EvidenceSignBytes(sequence, evidence.SignatureOne.Data)

	// check first signature
	if err := CheckSignature(pubKey, data, evidence.SignatureOne.Signature); err != nil {
		return sdkerrors.Wrap(err, "evidence signature one failed to be verified")
	}

	data = EvidenceSignBytes(sequence, evidence.SignatureTwo.Data)

	// check second signature
	if err := CheckSignature(pubKey, data, evidence.SignatureTwo.Signature); err != nil {
		return sdkerrors.Wrap(err, "evidence signature two failed to be verified")
	}

	return nil
}

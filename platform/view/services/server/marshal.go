/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package server

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"

	view "github.com/hyperledger-labs/fabric-smart-client/platform/view"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/api"
	hash2 "github.com/hyperledger-labs/fabric-smart-client/platform/view/services/hash"
	"github.com/hyperledger-labs/fabric-smart-client/platform/view/services/server/protos"
)

// UnmarshalCommand unmarshal Command messages
func UnmarshalCommand(raw []byte) (*protos.Command, error) {
	command := &protos.Command{}
	err := proto.Unmarshal(raw, command)
	if err != nil {
		return nil, err
	}

	return command, nil
}

type TimeFunc func() time.Time

// ResponseMarshaler produces SignedCommandResponse
type ResponseMarshaler struct {
	sp   api.ServiceProvider
	time TimeFunc
}

func NewResponseMarshaler(sp api.ServiceProvider) (*ResponseMarshaler, error) {
	return &ResponseMarshaler{
		sp:   sp,
		time: time.Now,
	}, nil
}

func (s *ResponseMarshaler) MarshalCommandResponse(command []byte, responsePayload interface{}) (*protos.SignedCommandResponse, error) {
	cr, err := commandResponseFromPayload(responsePayload)
	if err != nil {
		return nil, err
	}

	ts, err := ptypes.TimestampProto(s.time())
	if err != nil {
		return nil, err
	}

	did := api.GetIdentityProvider(s.sp).DefaultIdentity()
	cr.Header = &protos.CommandResponseHeader{
		Creator:     did,
		CommandHash: s.computeHash(command),
		Timestamp:   ts,
	}

	return s.createSignedCommandResponse(cr)
}

func (s *ResponseMarshaler) createSignedCommandResponse(cr *protos.CommandResponse) (*protos.SignedCommandResponse, error) {
	raw, err := proto.Marshal(cr)
	if err != nil {
		return nil, err
	}

	did := api.GetIdentityProvider(s.sp).DefaultIdentity()
	dSigner, err := view.GetSigService(s.sp).GetSigner(did)
	if err != nil {
		return nil, err
	}
	signature, err := dSigner.Sign(raw)
	if err != nil {
		return nil, err
	}

	return &protos.SignedCommandResponse{
		Response:  raw,
		Signature: signature,
	}, nil
}

func (s *ResponseMarshaler) computeHash(data []byte) (hash []byte) {
	hash, err := hash2.GetHasher(s.sp).Hash(data)
	if err != nil {
		panic(fmt.Errorf("failed computing hash on [% x]", data))
	}
	return
}

func commandResponseFromPayload(payload interface{}) (*protos.CommandResponse, error) {
	switch t := payload.(type) {
	case *protos.CommandResponse_Err:
		return &protos.CommandResponse{Payload: t}, nil
	case *protos.CommandResponse_InitiateViewResponse:
		return &protos.CommandResponse{Payload: t}, nil
	case *protos.CommandResponse_TrackViewResponse:
		return &protos.CommandResponse{Payload: t}, nil
	case *protos.CommandResponse_CallViewResponse:
		return &protos.CommandResponse{Payload: t}, nil
	case *protos.CommandResponse_IsTxFinalResponse:
		return &protos.CommandResponse{Payload: t}, nil
	case *protos.CommandResponse_IsHashFinalResponse:
		return &protos.CommandResponse{Payload: t}, nil
	default:
		return nil, errors.Errorf("command type not recognized: %T", t)
	}
}

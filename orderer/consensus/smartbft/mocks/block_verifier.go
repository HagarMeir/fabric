// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	common "github.com/hyperledger/fabric/protos/common"
	mock "github.com/stretchr/testify/mock"

	protoutil "github.com/hyperledger/fabric/protoutil"
)

// BlockVerifier is an autogenerated mock type for the BlockVerifier type
type BlockVerifier struct {
	mock.Mock
}

// VerifyBlockSignature provides a mock function with given fields: _a0, _a1
func (_m *BlockVerifier) VerifyBlockSignature(_a0 []*protoutil.SignedData, _a1 *common.ConfigEnvelope) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*protoutil.SignedData, *common.ConfigEnvelope) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

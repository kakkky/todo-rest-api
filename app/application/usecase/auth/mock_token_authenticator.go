// Code generated by MockGen. DO NOT EDIT.
// Source: ./interface_token_authenticator.go
//
// Generated by this command:
//
//	mockgen -package=auth -source=./interface_token_authenticator.go -destination=./mock_token_authenticator.go
//

// Package auth is a generated GoMock package.
package auth

import (
	reflect "reflect"

	jwt "github.com/golang-jwt/jwt"
	gomock "go.uber.org/mock/gomock"
)

// MockTokenAuthenticator is a mock of TokenAuthenticator interface.
type MockTokenAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenAuthenticatorMockRecorder
	isgomock struct{}
}

// MockTokenAuthenticatorMockRecorder is the mock recorder for MockTokenAuthenticator.
type MockTokenAuthenticatorMockRecorder struct {
	mock *MockTokenAuthenticator
}

// NewMockTokenAuthenticator creates a new mock instance.
func NewMockTokenAuthenticator(ctrl *gomock.Controller) *MockTokenAuthenticator {
	mock := &MockTokenAuthenticator{ctrl: ctrl}
	mock.recorder = &MockTokenAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenAuthenticator) EXPECT() *MockTokenAuthenticatorMockRecorder {
	return m.recorder
}

// GenerateToken mocks base method.
func (m *MockTokenAuthenticator) GenerateToken(sub, jwtID string) *jwt.Token {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", sub, jwtID)
	ret0, _ := ret[0].(*jwt.Token)
	return ret0
}

// GenerateToken indicates an expected call of GenerateToken.
func (mr *MockTokenAuthenticatorMockRecorder) GenerateToken(sub, jwtID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockTokenAuthenticator)(nil).GenerateToken), sub, jwtID)
}

// GetJWTIDFromClaim mocks base method.
func (m *MockTokenAuthenticator) GetJWTIDFromClaim(token *jwt.Token) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJWTIDFromClaim", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJWTIDFromClaim indicates an expected call of GetJWTIDFromClaim.
func (mr *MockTokenAuthenticatorMockRecorder) GetJWTIDFromClaim(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJWTIDFromClaim", reflect.TypeOf((*MockTokenAuthenticator)(nil).GetJWTIDFromClaim), token)
}

// GetSubFromClaim mocks base method.
func (m *MockTokenAuthenticator) GetSubFromClaim(token *jwt.Token) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubFromClaim", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubFromClaim indicates an expected call of GetSubFromClaim.
func (mr *MockTokenAuthenticatorMockRecorder) GetSubFromClaim(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubFromClaim", reflect.TypeOf((*MockTokenAuthenticator)(nil).GetSubFromClaim), token)
}

// SignToken mocks base method.
func (m *MockTokenAuthenticator) SignToken(token *jwt.Token) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignToken", token)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignToken indicates an expected call of SignToken.
func (mr *MockTokenAuthenticatorMockRecorder) SignToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignToken", reflect.TypeOf((*MockTokenAuthenticator)(nil).SignToken), token)
}

// VerifyExpiresAt mocks base method.
func (m *MockTokenAuthenticator) VerifyExpiresAt(token *jwt.Token) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyExpiresAt", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyExpiresAt indicates an expected call of VerifyExpiresAt.
func (mr *MockTokenAuthenticatorMockRecorder) VerifyExpiresAt(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyExpiresAt", reflect.TypeOf((*MockTokenAuthenticator)(nil).VerifyExpiresAt), token)
}

// VerifyToken mocks base method.
func (m *MockTokenAuthenticator) VerifyToken(signedToken string) (*jwt.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", signedToken)
	ret0, _ := ret[0].(*jwt.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockTokenAuthenticatorMockRecorder) VerifyToken(signedToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockTokenAuthenticator)(nil).VerifyToken), signedToken)
}
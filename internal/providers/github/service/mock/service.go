// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go
//
// Generated by this command:
//
//	mockgen -package mock_service -destination=./mock/service.go -source=./service.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	json "encoding/json"
	http "net/http"
	reflect "reflect"

	uuid "github.com/google/uuid"
	db "github.com/stacklok/minder/internal/db"
	gomock "go.uber.org/mock/gomock"
	oauth2 "golang.org/x/oauth2"
)

// MockGitHubProviderService is a mock of GitHubProviderService interface.
type MockGitHubProviderService struct {
	ctrl     *gomock.Controller
	recorder *MockGitHubProviderServiceMockRecorder
}

// MockGitHubProviderServiceMockRecorder is the mock recorder for MockGitHubProviderService.
type MockGitHubProviderServiceMockRecorder struct {
	mock *MockGitHubProviderService
}

// NewMockGitHubProviderService creates a new mock instance.
func NewMockGitHubProviderService(ctrl *gomock.Controller) *MockGitHubProviderService {
	mock := &MockGitHubProviderService{ctrl: ctrl}
	mock.recorder = &MockGitHubProviderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitHubProviderService) EXPECT() *MockGitHubProviderServiceMockRecorder {
	return m.recorder
}

// CreateGitHubAppProvider mocks base method.
func (m *MockGitHubProviderService) CreateGitHubAppProvider(ctx context.Context, token oauth2.Token, stateData db.GetProjectIDBySessionStateRow, installationID int64, state string) (*db.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGitHubAppProvider", ctx, token, stateData, installationID, state)
	ret0, _ := ret[0].(*db.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGitHubAppProvider indicates an expected call of CreateGitHubAppProvider.
func (mr *MockGitHubProviderServiceMockRecorder) CreateGitHubAppProvider(ctx, token, stateData, installationID, state any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGitHubAppProvider", reflect.TypeOf((*MockGitHubProviderService)(nil).CreateGitHubAppProvider), ctx, token, stateData, installationID, state)
}

// CreateGitHubAppWithoutInvitation mocks base method.
func (m *MockGitHubProviderService) CreateGitHubAppWithoutInvitation(ctx context.Context, qtx db.Querier, userID, installationID int64) (*db.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGitHubAppWithoutInvitation", ctx, qtx, userID, installationID)
	ret0, _ := ret[0].(*db.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGitHubAppWithoutInvitation indicates an expected call of CreateGitHubAppWithoutInvitation.
func (mr *MockGitHubProviderServiceMockRecorder) CreateGitHubAppWithoutInvitation(ctx, qtx, userID, installationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGitHubAppWithoutInvitation", reflect.TypeOf((*MockGitHubProviderService)(nil).CreateGitHubAppWithoutInvitation), ctx, qtx, userID, installationID)
}

// DeleteGitHubAppInstallation mocks base method.
func (m *MockGitHubProviderService) DeleteGitHubAppInstallation(ctx context.Context, installationID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGitHubAppInstallation", ctx, installationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteGitHubAppInstallation indicates an expected call of DeleteGitHubAppInstallation.
func (mr *MockGitHubProviderServiceMockRecorder) DeleteGitHubAppInstallation(ctx, installationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGitHubAppInstallation", reflect.TypeOf((*MockGitHubProviderService)(nil).DeleteGitHubAppInstallation), ctx, installationID)
}

// DeleteInstallation mocks base method.
func (m *MockGitHubProviderService) DeleteInstallation(ctx context.Context, providerID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInstallation", ctx, providerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInstallation indicates an expected call of DeleteInstallation.
func (mr *MockGitHubProviderServiceMockRecorder) DeleteInstallation(ctx, providerID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInstallation", reflect.TypeOf((*MockGitHubProviderService)(nil).DeleteInstallation), ctx, providerID)
}

// GetConfig mocks base method.
func (m *MockGitHubProviderService) GetConfig(ctx context.Context, class db.ProviderClass, userConfig json.RawMessage) (json.RawMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", ctx, class, userConfig)
	ret0, _ := ret[0].(json.RawMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockGitHubProviderServiceMockRecorder) GetConfig(ctx, class, userConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockGitHubProviderService)(nil).GetConfig), ctx, class, userConfig)
}

// ValidateGitHubAppWebhookPayload mocks base method.
func (m *MockGitHubProviderService) ValidateGitHubAppWebhookPayload(r *http.Request) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGitHubAppWebhookPayload", r)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateGitHubAppWebhookPayload indicates an expected call of ValidateGitHubAppWebhookPayload.
func (mr *MockGitHubProviderServiceMockRecorder) ValidateGitHubAppWebhookPayload(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGitHubAppWebhookPayload", reflect.TypeOf((*MockGitHubProviderService)(nil).ValidateGitHubAppWebhookPayload), r)
}

// ValidateGitHubInstallationId mocks base method.
func (m *MockGitHubProviderService) ValidateGitHubInstallationId(ctx context.Context, token *oauth2.Token, installationID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGitHubInstallationId", ctx, token, installationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateGitHubInstallationId indicates an expected call of ValidateGitHubInstallationId.
func (mr *MockGitHubProviderServiceMockRecorder) ValidateGitHubInstallationId(ctx, token, installationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGitHubInstallationId", reflect.TypeOf((*MockGitHubProviderService)(nil).ValidateGitHubInstallationId), ctx, token, installationID)
}

// VerifyProviderTokenIdentity mocks base method.
func (m *MockGitHubProviderService) VerifyProviderTokenIdentity(ctx context.Context, remoteUser, accessToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyProviderTokenIdentity", ctx, remoteUser, accessToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyProviderTokenIdentity indicates an expected call of VerifyProviderTokenIdentity.
func (mr *MockGitHubProviderServiceMockRecorder) VerifyProviderTokenIdentity(ctx, remoteUser, accessToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyProviderTokenIdentity", reflect.TypeOf((*MockGitHubProviderService)(nil).VerifyProviderTokenIdentity), ctx, remoteUser, accessToken)
}

package mocks

import (
    "github.com/stretchr/testify/mock"
    "github.com/amir/raidman"
)

type MockRiemannClient struct {
    mock.Mock
}

func (m *MockRiemannClient) Close() {
    m.Called()
}

func (m *MockRiemannClient) Send(event *raidman.Event) error {
    return m.Called(event).Error(0)
}

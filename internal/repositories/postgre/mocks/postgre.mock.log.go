package mocks

import (
	"avito-backend-assignment/internal/models"
	"github.com/stretchr/testify/mock"
)

type LogRepository struct {
	mock.Mock
}

func (_m *LogRepository) Save(log models.Log) (string, error) {
	ret := _m.Called(log)

	var r0 string
	if rf, ok := ret.Get(0).(func(models.Log) string); ok {
		r0 = rf(log)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *LogRepository) FindById(id string) (models.Log, error) {
	//TODO implement me
	panic("implement me")
}

func (_m *LogRepository) FindByRequestId(request string) (models.Log, error) {
	//TODO implement me
	panic("implement me")
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLogRepository(t mockConstructorTestingTNewRepository) *LogRepository {
	mock := &LogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

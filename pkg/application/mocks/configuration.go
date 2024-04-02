package mocks

type MockConfiguration struct{}

func (m *MockConfiguration) Get(key string) interface{} {
	return nil
}

func (m *MockConfiguration) Set(key string, value interface{}) {
}

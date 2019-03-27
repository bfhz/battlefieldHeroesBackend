package kvstore

import (
	"fmt"
	"sync"
)

type HashMap struct {
	Data map[string]string
	mu   sync.Mutex
}

func (m *HashMap) Get(ident string, key string) string {
	val, ok := m.Data[fmt.Sprintf("%s/%s", ident, key)]
	if !ok {
		return ""
	}
	return string(val)
}

func (m *HashMap) Set(ident string, key, value string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Data[fmt.Sprintf("%s/%s", ident, key)] = value
	return nil
}

func (m *HashMap) SetM(ident string, hash map[string]string) error {
	for k, v := range hash {
		m.Set(ident, k, v)
	}
	return nil
}

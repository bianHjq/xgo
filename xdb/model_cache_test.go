package xdb

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestCache struct {
	cache map[string]string
}

func (t *TestCache) Get(ctx context.Context, key string) (string, error) {
	c, _ := t.cache[key]
	return c, nil
}

func (t *TestCache) Del(ctx context.Context, key string) error {
	delete(t.cache, key)
	return nil
}

func (t *TestCache) Set(ctx context.Context, key string, data string) error {
	t.cache[key] = data
	return nil
}

func (t *TestCache) SetWithTTL(ctx context.Context, key string, data string, ttl time.Duration) error {
	t.cache[key] = data
	return nil
}

func init() {
	cache = &TestCache{
		cache: map[string]string{},
	}
}

func TestModel_FindBy(t *testing.T) {
	var u User
	err := m.FindBy("1").Binding(&u)
	assert.Equal(t, nil, err)

	err = m.FindBy("1").Binding(&u)
	assert.Equal(t, nil, err)
	fmt.Println("row", fmt.Sprintf("%+v", u))
}

func TestModel_UpdateBy(t *testing.T) {
	_, err := m.UpdateBy("1", map[string]any{
		"name": "圣斗士-星矢",
		"profile": map[string]any{
			"hobby": "技能-天马流行拳",
		},
	})
	assert.Equal(t, nil, err)
}

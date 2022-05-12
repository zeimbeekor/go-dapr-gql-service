package utils

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/zeimbeekor/go-dapr-gql-service/graph/model"
)

func TestGenericRequest(t *testing.T) {
	t.Parallel()
	t.Run("Get users", func(t *testing.T) {
		t.Parallel()
		var err error
		mUsers := []*model.User{}
		if body, _, err := GenericRequest("users"); err == nil {
			json.Unmarshal(body, &mUsers)
		}
		assert.NotNil(t, mUsers)
		assert.Nil(t, err)
		assert.NotEqual(t, 0, len(mUsers))
	})
}

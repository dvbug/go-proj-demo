package scripts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserJson(t *testing.T) {
	user := &User{FirstName: "Lei", LastName: "Li", Location: Location{City: "Beijing", Postcode: 100010}}
	Info(user)
	jsonStr := ToJson(user)
	assert.NotEmpty(t, jsonStr)

	user2 := &User{}
	err := FromJson(jsonStr, user2)
	assert.Nil(t, err)
	assert.NotNil(t, user2)
	assert.Equal(t, user2, user)
}

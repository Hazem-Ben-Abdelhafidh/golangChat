package controllers

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	args := RequestBody{
		Name:     "pain",
		Password: "test123",
	}

	user, err := client.User.Create().SetName(args.Name).SetPassword(args.Password).Save(ctx)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, args.Name, user.Name)
	require.Equal(t, args.Password, user.Password)
	require.NotZero(t, user.ID)
}

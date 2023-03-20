package controllers

import (
	"context"
	"os"
	"testing"

	"github.com/Hazem-Ben-Abdelhafidh/golangChat/config"
	"github.com/Hazem-Ben-Abdelhafidh/golangChat/ent"
)

var (
	client *ent.Client
	ctx    context.Context
)

func TestMain(m *testing.M) {
	client = config.ConnectDB()
	ctx = context.Background()
	os.Exit(m.Run())
}

package db

import (
	"context"
	"github.com/TungstenRust/simplebank/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	connPool, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("can't connect to the database;", err)
	}
	testStore = NewStore(connPool)
	os.Exit(m.Run())
}

package repos

import (
	"testing"

	"github.com/yasseldg/go-simple/repos/rMongo"
)

func TestDrop(t *testing.T) {
	// Initialize the MongoDB admin client
	mongoAdmin := rMongo.NewAdmin()
	mongoAdmin.SetDebug(true)

	// Get the client
	client, err := mongoAdmin.GetClient("", "WRITE")
	if err != nil {
		t.Fatalf("GetClient() error = %v", err)
	}

	// List database names
	database_names, err := client.ListDatabaseNames(context.Background(), primitive.D{})
	if err != nil {
		t.Fatalf("ListDatabaseNames() error = %v", err)
	}

	// Filter database names
	db_list := []string{}
	for _, database_name := range database_names {
		if strings.HasPrefix(database_name, "BYBIT") {
			if database_name == "BYBIT_BTCUSDT" || database_name == "BYBIT_WLDUSDT" {
				continue
			}
			db_list = append(db_list, database_name)
		}
	}

	// Drop databases
	drop_dbs(client, db_list)
}

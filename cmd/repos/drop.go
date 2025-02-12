package repos

import (
	"context"
	"strings"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/repos/rMongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Drop() {

	mongoAdmin := rMongo.NewAdmin()
	// mongoAdmin.SetDebug(true)

	conn_name := "PP_Historic_Trades_R"

	sLog.Warn("conn_name: %s", conn_name)

	client, err := mongoAdmin.GetClient("", conn_name)
	if err != nil {
		sLog.Panic("GetClient( ): %s", err.Error())
	}

	database_names, err := client.ListDatabaseNames(context.Background(), primitive.D{})
	if err != nil {
		sLog.Panic("ListDatabaseNames( ): %s", err.Error())
	}

	db_list := []string{}

	for _, database_name := range database_names {

		if database_name == "admin" || database_name == "config" || database_name == "local" {
			continue
		}

		if strings.HasPrefix(database_name, "BITMEX") {
			// if database_name == "BYBIT_BTCUSDT" || database_name == "BYBIT_WLDUSDT" {
			// }

			continue
		}

		db_list = append(db_list, database_name)
	}

	drop_dbs(client, db_list)
}

func drop_dbs(client rMongo.InterAdminClient, db_list []string) {

	for _, database_name := range db_list {
		sLog.Info(database_name)

		database, err := client.GetDatabase(database_name)
		if err != nil {
			sLog.Error("GetDatabase( %s ): %s", database_name, err.Error())
			continue
		}

		coll_list, err := database.ListCollectionNames(context.Background(), primitive.D{})
		if err != nil {
			sLog.Error("ListCollectionNames( %s ): %s", database_name, err.Error())
			continue
		}

		drop_colls(database, coll_list)
	}
}

func drop_colls(database rMongo.InterAdminDatabase, coll_list []string) {

	for _, coll_name := range coll_list {
		// sLog.Info(coll_name)

		if !strings.HasPrefix(coll_name, "trend_") {
			// if coll_name != "historic_trades" {
			continue
		}

		coll := database.NewCollection(coll_name)
		if coll == nil {
			sLog.Error("NewCollection( %s ) is nil", coll_name)
			continue
		}

		sLog.Warn("Dropping %s", coll_name)

		if err := coll.Drop(context.Background()); err != nil {
			sLog.Error("Drop( %s ): %s", coll_name, err.Error())
		}
	}
}

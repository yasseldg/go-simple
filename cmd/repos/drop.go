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
	mongoAdmin.SetDebug(true)

	client, err := mongoAdmin.GetClient("", "WRITE")
	if err != nil {
		sLog.Panic("GetClient( ): %s", err.Error())
	}

	sLog.Warn("client: %s", client.Env())

	database_names, err := client.ListDatabaseNames(context.Background(), primitive.D{})
	if err != nil {
		sLog.Panic("ListDatabaseNames( ): %s", err.Error())
	}

	db_list := []string{}

	for _, database_name := range database_names {

		if strings.HasPrefix(database_name, "BYBIT") {
			if database_name == "BYBIT_BTCUSDT" || database_name == "BYBIT_WLDUSDT" {
				continue
			}

			db_list = append(db_list, database_name)
		}
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
		sLog.Info(coll_name)

		coll := database.NewCollection(coll_name)
		if coll == nil {
			sLog.Error("NewCollection( %s ) is nil", coll_name)
			continue
		}

		if err := coll.Drop(context.Background()); err != nil {
			sLog.Error("Drop( %s ): %s", coll_name, err.Error())
		}
	}
}

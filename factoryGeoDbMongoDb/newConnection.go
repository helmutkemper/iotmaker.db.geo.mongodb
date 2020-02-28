package factoryGeoDbMongoDb

import iotmaker_db_geo_mongodb "github.com/helmutkemper/iotmaker.db.geo.mongodb"

func NewConnection(connString, dbName string) (error, *iotmaker_db_geo_mongodb.DbFunctionsFromMap) {
	conn := &iotmaker_db_geo_mongodb.DbFunctionsFromMap{
		CollectionWay:              "way",
		CollectionTmpWay:           "wayTmp",
		CollectionWayToPopulate:    "wayToPopulate",
		CollectionSurrounding:      "surrounding",
		CollectionSurroundingRight: "surroundingRight",
		CollectionSurroundingLeft:  "surroundingLeft",
	}

	err := conn.Connect(connString, dbName)

	return err, conn
}

package iotmaker_db_geo_mongodb

import (
	iotmaker_geo_osm "github.com/helmutkemper/iotmaker.geo.osm"
	"github.com/helmutkemper/osmpbf"
)

type DbFunctionsFromMapInterface interface {
	Connect(connection ...interface{}) error
	Disconnect() error
	WayTmpInsert(data interface{}) error
	WayTmpFind(query interface{}, pointerToResult *[]osmpbf.Way) error
	WayTmpDeleteByOsmId(id int64) error
	WayToPopulateFind(query interface{}, pointerToResult *[]iotmaker_geo_osm.WayStt) error
	WayToPopulateUpdateLocations(id int64, loc, rad [][2]float64) error
	WayToPopulateDeleteByOsmId(id int64) error
	WayToPopulateInsert(data interface{}) error
	WayCount(query interface{}) (error, int64)
	WayInsert(data interface{}) error
	WayFind(query interface{}, pointerToResult *[]iotmaker_geo_osm.WayStt) error
	SurroundingInsert(data interface{}) error
	SurroundingLeftInsert(data interface{}) error
	SurroundingRightInsert(data interface{}) error
}

package iotmaker_db_geo_mongodb

import (
	"context"
	"errors"
	iotmaker_geo_osm "github.com/helmutkemper/iotmaker.geo.osm"
	"github.com/helmutkemper/osmpbf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbFunctionsFromMap struct {
	Client   interface{}
	dbString string

	CollectionWay              string
	CollectionTmpWay           string
	CollectionWayToPopulate    string
	CollectionSurrounding      string
	CollectionSurroundingRight string
	CollectionSurroundingLeft  string
}

//"mongodb://0.0.0.0:27017"
func (el *DbFunctionsFromMap) Connect(connection ...interface{}) error {
	var err error
	var connString string
	var dbString string

	if len(connection) != 2 {
		return errors.New("connection must be a string like 'mongodb://0.0.0.0:27017', 'server_name'")
	}

	switch connection[0].(type) {
	case string:
		connString = connection[0].(string)

		if connString == "" {
			return errors.New("connection must be a string like 'mongodb://0.0.0.0:27017', 'server_name'")
		}

	default:
		return errors.New("connection must be a string like 'mongodb://0.0.0.0:27017', 'server_name'")
	}

	switch connection[1].(type) {
	case string:
		dbString = connection[1].(string)

		if dbString == "" {
			return errors.New("connection must be a string like 'mongodb://0.0.0.0:27017', 'server_name'")
		}

		el.dbString = dbString

	default:
		return errors.New("connection must be a string like 'mongodb://0.0.0.0:27017', 'server_name'")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(connString)

	// Connect to MongoDB
	el.Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = el.Client.(*mongo.Client).Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	return nil
}

func (el *DbFunctionsFromMap) Disconnect() error {
	return el.Client.(*mongo.Client).Disconnect(context.TODO())
}

func (el *DbFunctionsFromMap) WayTmpInsert(data interface{}) error {
	var err error
	_, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionTmpWay).InsertOne(context.TODO(), data)
	return err
}

func (el *DbFunctionsFromMap) WayTmpFind(query interface{}, pointerToResult *[]osmpbf.Way) error {
	var err error
	var cursor *mongo.Cursor
	var toDecode osmpbf.Way

	*pointerToResult = make([]osmpbf.Way, 0)
	cursor, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionTmpWay).Find(context.TODO(), query)

	if cursor == nil {
		return errors.New("mongodb.find().error: cursor is nil")
	}
	defer cursor.Close(context.TODO())

	if err = cursor.Err(); err != nil {
		return err
	}

	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&toDecode)
		if err != nil {
			return err
		}

		*pointerToResult = append(*pointerToResult, toDecode)
	}

	return nil
}

func (el *DbFunctionsFromMap) WayTmpDeleteByOsmId(id int64) error {
	var err error
	_, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionTmpWay).DeleteOne(context.TODO(), bson.M{"id": id})
	return err
}

func (el *DbFunctionsFromMap) WayToPopulateUpdateLocations(id int64, loc, rad [][2]float64) error {
	var err error
	_, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionWayToPopulate).UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": bson.M{"rad": rad, "loc": loc}})
	return err
}

func (el *DbFunctionsFromMap) WayToPopulateDeleteByOsmId(id int64) error {
	var err error
	_, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionWayToPopulate).DeleteOne(context.TODO(), bson.M{"id": id})
	return err
}

func (el *DbFunctionsFromMap) WayCount(query interface{}) (error, int64) {
	count, err := el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionWay).CountDocuments(context.TODO(), query)
	return err, count
}

func (el *DbFunctionsFromMap) WayInsert(data interface{}) error {
	var err error
	_, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionWay).InsertOne(context.TODO(), data)
	return err
}

func (el *DbFunctionsFromMap) WayFind(query interface{}, pointerToResult *[]iotmaker_geo_osm.WayStt) error {
	var err error
	var cursor *mongo.Cursor
	var toDecode iotmaker_geo_osm.WayStt

	*pointerToResult = make([]iotmaker_geo_osm.WayStt, 0)
	cursor, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionWay).Find(context.TODO(), query)

	if cursor == nil {
		return errors.New("mongodb.find().error: cursor is nil")
	}
	defer cursor.Close(context.TODO())

	if err = cursor.Err(); err != nil {
		return err
	}

	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&toDecode)
		if err != nil {
			return err
		}

		*pointerToResult = append(*pointerToResult, toDecode)
	}

	return nil
}

func (el *DbFunctionsFromMap) SurroundingInsert(data interface{}) error {
	var err error
	_, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionSurrounding).InsertOne(context.TODO(), data)
	return err
}

func (el *DbFunctionsFromMap) SurroundingLeftInsert(data interface{}) error {
	var err error
	_, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionSurroundingLeft).InsertOne(context.TODO(), data)
	return err
}

func (el *DbFunctionsFromMap) SurroundingRightInsert(data interface{}) error {
	var err error
	_, err = el.Client.(*mongo.Client).Database(el.dbString).Collection(el.CollectionSurroundingRight).InsertOne(context.TODO(), data)
	return err
}

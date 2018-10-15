package db

import mgo "gopkg.in/mgo.v2"

var mgoSession *mgo.Session

var mongo_conn_str = "mongodb:27017"
//var mongo_conn_str = "localhost:27017"

func GetMongoSession() (*mgo.Session, error) {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongo_conn_str)
		if err != nil {
			return nil, err
		}
	}
	return mgoSession.Clone(), nil
}

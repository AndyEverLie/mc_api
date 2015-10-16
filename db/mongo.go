package db
import "gopkg.in/mgo.v2"

const DB_URL = "mongodb://localhost:27017/mc_api"

var (
	mgoSession *mgo.Session
)

func getSession() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(DB_URL)
		if err != nil {
			panic(err)
		}
		mgoSession.SetPoolLimit(2048)	// TODO param
	}
	return mgoSession.Clone()
}

func Query(collection string, s func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB("").C(collection)
	return s(c)
}

package store

/////////

import (
	"log"

	"github.com/fast-user/fast/pkg/fast-server/models"

	"github.com/globalsign/mgo/bson"
"github.com/globalsign/mgo"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
)

//// mongo session
type FastStore struct {
Store
}
func NewFastStore(database *mgo.Session,dbName string) *FastStore {
	fastUser := new(models.FastUser)

	fastStore := new(FastStore)
  fastStore.Store = Store {
	DBName : dbName,
	CollectionName :fastUser.CollectionName(),
	Session:database,
}
	return fastStore
}

func (fs *FastStore) CreateAccount(ctx *gin.Context, fastUserAccount *models.FastUser) (string, error) {

	log.Println("dbname ", fs.Store.DBName)
	cursor := fs.Store.Session.DB(fs.Store.DBName).C(fs.Store.CollectionName)
	log.Println("creating Account")
log.Println("fast user account ",fastUserAccount)
	err := cursor.Insert(&fastUserAccount)

	if err != nil {
		log.Fatal("problem saving Account: ", errors.Details(err))
		return "", errors.Trace(err)
	}
	log.Println("Account created")

	return fastUserAccount.ID.Hex(), nil
}

// etAccountByID returns an Account in the DB by ID.
func (fs *FastStore) GetFastAccountByID(ctx *gin.Context, id string) (*models.FastUser, error) {

	if !bson.IsObjectIdHex(id) {
		log.Println("Account ID is invalid")
		return nil, errors.NotValidf("Account ID: '%s'", id)
	}
	//session, dbName := GetSessionForMongo()
	cursor := fs.Store.Session.DB(fs.Store.DBName).C(fs.Store.CollectionName)
	var fastUser models.FastUser

	err := cursor.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&fastUser)

	if err != nil {
		log.Println("problem getting account by id: ", errors.Details(err))
		return nil, errors.Wrap(err, errors.NotFoundf("Account with ID: '%s'", id))
	}
	return &fastUser, nil
}

//// get by stripe id
func (fs *FastStore) GetFastAccountByStripe(ctx *gin.Context, stripeID string) (*models.FastUser, error) {
cursor := fs.Store.Session.DB(fs.Store.DBName).C(fs.Store.CollectionName)
	var fastUser models.FastUser

	err := cursor.Find(bson.M{"stripe_id": stripeID}).One(&fastUser)

	if err != nil {
		log.Println("problem getting account by stripeid: ", errors.Details(err))
		return nil, errors.Wrap(err, errors.NotFoundf("Account with stripeID: '%s'",stripeID ))
	}
	return &fastUser, nil
}
/// get by brainwireID
func (fs *FastStore) GetFastAccountByBrainTreeID(ctx *gin.Context, brainTreeID string) (*models.FastUser, error) {

	cursor := fs.Store.Session.DB(fs.Store.DBName).C(fs.Store.CollectionName)
	var fastUser models.FastUser

	err := cursor.Find(bson.M{"brain_tree_id": brainTreeID}).One(&fastUser)

	if err != nil {
		log.Println("problem getting account by brainTreeID: ", errors.Details(err))
		return nil, errors.Wrap(err, errors.NotFoundf("Account with brainTreeID: '%s'",brainTreeID))
	}
	return &fastUser, nil
}

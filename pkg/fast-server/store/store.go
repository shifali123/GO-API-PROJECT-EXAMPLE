package store

import (
	"log"
	"os"
"github.com/joho/godotenv"
	"github.com/globalsign/mgo"
)
type Store struct {
	Session *mgo.Session
	DBName string
	CollectionName string
}
func  GetSessionForMongo() (*mgo.Session, string) {
	url, dbName := GetEnv()
	session, err := SetupDatabase(url, dbName)
	if err != nil {
		log.Panic("error in creating session ", err)
	}
	log.Println("dbname is ", dbName)
	return session, dbName

}
func SetupDatabase(databaseURL string, databaseName string) (*mgo.Session, error) {

	return Connect(databaseURL, databaseName)
}
func Connect(url string, databaseName string) (*mgo.Session, error) {

	database, err := mgo.Dial(url)
	if err != nil {
		log.Panic("Encountered error while connecting to database")

	}
	database.SetMode(mgo.Monotonic, true)
	session := database

	return session, err
}

func GetEnv() (string, string) {
	err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file ",err)
}


	for _, env := range os.Environ() {
		log.Println("env is ", env)
	}
	databaseURL, dbURLSet := os.LookupEnv("MONGO_HOST")
	if !dbURLSet {

		return "", ""
	}
	databaseName, dbNameSet := os.LookupEnv("MONGO_NAME")
	if !dbNameSet {

		return "", ""
	}

	return databaseURL, databaseName
}

// CollectionName collection name for MongoDB.
func CollectionName() string {
	return "fastuseraccount"
}

//// get env for stripe key
func GetEnvStripe() (string) {
	err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file ",err)
}
	stripeKey, _ := os.LookupEnv("STRIPE_KEY")
	log.Println("stripe key is ",stripeKey)
return stripeKey
}

///// get envs for braintreekey
func GetEnvBrainTree() (string,string,string) {
	err := godotenv.Load()
if err != nil {
	log.Fatal("Error loading .env file ",err)
}
brainTreeMerchID, _ := os.LookupEnv("BRAINTREE_MERCH_ID")
brainTreePubKey, _ := os.LookupEnv("BRAINTREE_PUB_KEY")
brainTreePrivateKey, _ := os.LookupEnv("BRAINTREE_PRIV_KEY")
return brainTreeMerchID,brainTreePubKey,brainTreePrivateKey
}

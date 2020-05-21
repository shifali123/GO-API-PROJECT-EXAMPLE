package handler

import (
	"net/http"
"log"
	"github.com/fast-user/fast/pkg/fast-server/models"
"context"
	"github.com/fast-user/fast/pkg/fast-server/store"
"github.com/globalsign/mgo"
	"github.com/gin-gonic/gin"
	 "github.com/stripe/stripe-go/v71"
	 "github.com/stripe/stripe-go/v71/customer"
"strings"
"github.com/braintree-go/braintree-go"
"time"



)

type CreateFastUserHandler struct {

	FastStore         *store.FastStore

}

func NewCreateFastUserHandler(db *mgo.Session,dbName string) *CreateFastUserHandler {
	return &CreateFastUserHandler{
		FastStore:         store.NewFastStore(db,dbName),

	}
}

//// create fastuseraccount
func (fast *CreateFastUserHandler) CreateFastUser(ctx *gin.Context) {
	// Get the AddressBook model from the request
	var fastUserRequest models.FastUser


	err := ctx.ShouldBindJSON(&fastUserRequest)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, "invalid create user request")
		return
	}
stripeKey :=	store.GetEnvStripe()
brainTreeMerchID,brainTreePubKey,brainTreePrivateKey:= store.GetEnvBrainTree()


if brainTreeMerchID != "" && brainTreePubKey != "" && brainTreePrivateKey != "" {
bt := braintree.New(
	braintree.Sandbox,
	brainTreeMerchID,
	brainTreePubKey,
	brainTreePrivateKey,
//	"2mg4xm26m8s49cg7",
	//"s9jdf43t9vd28cdn",
    //"8e6863af06ec4444fc0e5d1cb14bfe99",
  )

	log.Println("br is ",bt)
customerRequest :=	&braintree.CustomerRequest{
	FirstName :fastUserRequest.FirstName,
	LastName:fastUserRequest.LastName ,
	}

	ctxBt := context.Background()

customerRequestCreated,err :=	bt.Customer().Create(ctxBt,customerRequest)
if err != nil {
	fastUserRequest.BrainTreeID = "brain tree user is not created due to other reason such as json not correct"
		 ctx.Error(err)
}
fastUserRequest.BrainTreeID = customerRequestCreated.Id
} else {
	fastUserRequest.BrainTreeID = "brain tree id is not set because brain tree keys are not set in the env "
}
if stripeKey != "" {
	stripe.Key = stripeKey
//stripe.Key = "sk_test_bgKphlwP7uDHxAagQkD1t3sm00nSxu4FAH"

var name string
if fastUserRequest.FirstName != "" {
	name +=fastUserRequest.FirstName + " "
}
if fastUserRequest.LastName != "" {
	name +=fastUserRequest.LastName + " "
}
if fastUserRequest.MiddleName != "" {
	name +=fastUserRequest.MiddleName + " "
}
nameStripe := &name
if strings.TrimSpace(fastUserRequest.Address.Country) == "" {
	fastUserRequest.Address.Country = "USA"
}
address := &stripe.AddressParams {
	City: &fastUserRequest.Address.City,
	Line1:&fastUserRequest.Address.Line1,
	Line2:&fastUserRequest.Address.Line2,
	Country:&fastUserRequest.Address.Country,
	PostalCode:&fastUserRequest.Address.PostalCode,
	State:&fastUserRequest.Address.SubDivsion,}
params := &stripe.CustomerParams{
  Description: stripe.String("Customer Creation for Fast"),
Name : nameStripe,
Address: address,
}
stripeAccountCreate, err := customer.New(params)
if err != nil {
		 ctx.Error(err)
}

fastUserRequest.StripeID = stripeAccountCreate.ID
} else {
	fastUserRequest.StripeID = "stripe account cannot be created because env for stripe is not set "
}

fastUserRequest.CreatedAt = time.Now().Unix()
	_, err = fast.FastStore.CreateAccount(ctx, &fastUserRequest)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, "unable to create record in database")
		return
	}
	log.Println("fastuser ",fastUserRequest)
	ctx.JSON(http.StatusOK, gin.H{"stripe_id":fastUserRequest.StripeID ,"braintree_id":fastUserRequest.BrainTreeID,"user_created":true,"timestamp":fastUserRequest.CreatedAt})

}

//// getfastuseraccount
func (fast *CreateFastUserHandler) GetFastUserByID(ctx *gin.Context) {
	userID := ctx.Param("id")

 stripID := ctx.DefaultQuery("stripID", "")
 brainWireID := ctx.DefaultQuery("brainwireID", "")



	fastUser, err := fast.FastStore.GetFastAccountByID(ctx, userID)
	if err != nil {
		if stripID != "" {
		 fastUser, err := 	fast.FastStore.GetFastAccountByStripe(ctx, stripID)
		 if err != nil {
		 ctx.Error(err)
		 }
		 ctx.JSON(http.StatusOK, fastUser)
		 return
		}
		if brainWireID != "" {
			 fastUser, err :=  fast.FastStore.GetFastAccountByBrainTreeID(ctx, brainWireID)
			 if err != nil {
				 ctx.Error(err)
			 }
				ctx.JSON(http.StatusOK, fastUser)
				return
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, "unable to find user by id")
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, fastUser)
}

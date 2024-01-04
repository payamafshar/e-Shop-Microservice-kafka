package authenticationservice

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(group *gin.RouterGroup) {



}
func getMongoCollection(mongoURL, dbName, collectionName string) *mongo.Collection {
	// client, err := mongo.Connect(context.Background() )
	clientOptions := ClientOptions()
	var cred options.Credential

	cred.Username = os.Getenv("MongoUsername")
	cred.Password = os.Getenv("MongoPassword")
	mongoUrl := os.Getenv("MongoUrl")
	clientOptions = clientOptions.ApplyURI(mongoUrl).SetAuth(cred)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}


func Testhandler(ctx *gin.Context) {

	var createAccountDto CreateAccountDto
	if err := ctx.ShouldBindJSON(&createAccountDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	writer, closeWriter := NewWriter[CreateAccountDto]("kafka:9092", "test-topic")
	err := writer.WriteBatch(ctx, createAccountDto)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
	}
	ctx.JSON(http.StatusOK, &createAccountDto)
	defer closeWriter()
}

type CreateAccountDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}


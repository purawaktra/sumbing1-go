package entities

type ProfilePicture struct {
	AccountId string `bson:"account_id"`
	Data      []byte `bson:"data"`
}

type BsonData struct {
	RequestId string `bson:"request_id"`
	Data      any    `bson:"data"`
}

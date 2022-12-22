package model

type Business struct {
	Name  string `bson:"name"`
	Admin User   `bson:"admin"`
	Menu  []Menu `bson:"menu"`
}

// menu status
const (
	Ready    = 1
	NotReady = 2
)

type Menu struct {
	Name   string   `bson:"name"`
	Status int      `bson:"status"`
	Price  int      `bson:"price"`
	Origin string   `bson:"origin"`
	Review []Review `bson:"review"`
	Score  float32  `bson:"score"`
}

type Review struct {
	Orderer User   `bson:"orderer"`
	Content string `bson:"content"`
	Score   int    `bson:"score"`
}

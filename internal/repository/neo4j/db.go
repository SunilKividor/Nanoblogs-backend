package neo4j

import (
	"context"
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var (
	neo4jDB neo4j.DriverWithContext
)

func ConnectNeo4j() {
	host := os.Getenv("NEO4JHOST")
	username := os.Getenv("NEO4JUSERNAME")
	password := os.Getenv("NEO4JPASSWORD")
	driver, err := neo4j.NewDriverWithContext(host, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic(err)
	}
	neo4jDB = driver
	ctx := context.Background()
	info, err := (neo4jDB.GetServerInfo(ctx))
	if err != nil {
		panic(err)
	}
	fmt.Println(info.Address())
}

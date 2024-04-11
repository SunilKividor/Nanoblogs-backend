package neo4j

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/net/context"
)

func RegisterNewUserNeo4j(username string, categories []string) error {
	ctx := context.Background()
	session := neo4jDB.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	res, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		res, err := tx.Run(
			ctx,
			"MERGE (p:Person {name : $username}) RETURN id(p)",
			map[string]interface{}{"username": username},
		)
		if err != nil {
			return nil, err
		}
		for _, category := range categories {
			_, err := tx.Run(
				ctx,
				"MATCH(p:Person {name : $username})"+
					"MATCH(c:Category {name : $category})"+
					"MERGE (p)-[r:SELECTED_CATEGORY]->(c)"+
					"ON CREATE SET r.frequency = 1, c.frequency = c.frequency + 1",
				map[string]interface{}{"username": username, "category": category},
			)
			if err != nil {
				return nil, err
			}
		}
		return nil, res.Err()
	})
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

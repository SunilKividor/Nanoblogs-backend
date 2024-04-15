package neo4j

import (
	"context"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func IncreaseCategoryFreq(cat string) error {
	//searches category . If found increases its frequency
	ctx := context.Background()
	sesion := neo4jDB.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer sesion.Close(ctx)
	smt := "MATCH (c:Category {name : $category}) SET c.frequency = c.frequency + 1"
	_, err := sesion.Run(ctx, smt, map[string]interface{}{
		"category": cat,
	})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func CreateNewCategory(cat string) error {
	//creates a new category in the db

	ctx := context.Background()
	sesion := neo4jDB.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer sesion.Close(ctx)
	smt := "MERGE (c:Category {name : $category}) ON MATCH SET c.frequency = c.frequency + 1 ON CREATE SET c.frequency = $frequency"
	_, err := sesion.Run(ctx, smt, map[string]interface{}{
		"category":  cat,
		"frequency": 1,
	})
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil
}

func GetTopCategories() (interface{}, error) {
	//get top categories with most frequencies
	ctx := context.Background()
	session := neo4jDB.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)
	res, err := session.ExecuteRead(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		stm := "MATCH (n: Category) RETURN n.name ORDER BY n.frequency DESC LIMIT 25"
		res, err := tx.Run(ctx, stm, nil)
		if err != nil {
			return nil, err
		}
		var category []string
		for res.Next(ctx) {
			record := res.Record()
			nameValue, found := record.Get("n.name")
			if !found {
				log.Println("No name found")
			}
			name, ok := nameValue.(string)
			if ok {
				category = append(category, name)
			} else {
				log.Println("name property is not string")
			}
		}
		log.Println(category)
		return category, nil
	})
	if err != nil {

		return nil, err
	}
	return res, nil
}

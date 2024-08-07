package middleware

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func MonitorFunctionPerformance(client *mongo.Client, collection string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Path() == "/monitoring" {
			return c.Next()
		}

		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		humanReadableDuration := formatDuration(duration)

		coll := client.Database("Monitoring").Collection(collection)
		document := map[string]interface{}{
			"path":     c.Path(),
			"method":   c.Method(),
			"duration": humanReadableDuration,
			"time":     time.Now(),
		}

		result, err := coll.InsertOne(context.Background(), document)
		if err != nil {
			log.Println("Failed to insert performance data:", err)
		} else {
			log.Println("Inserted performance data:", result.InsertedID)
			log.Printf("Document: %+v\n", document)
		}

		return err
	}
}

func formatDuration(d time.Duration) string {
	if d < time.Millisecond {
		return d.String() // in microseconds
	} else if d < time.Second {
		return d.Truncate(time.Millisecond).String() // in milliseconds
	} else {
		return d.Truncate(time.Second).String() // in seconds
	}
}

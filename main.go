package zuoye4

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

var redisDB *redis.Client

func initClient() (err error) {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisDB.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func main() {
	router := gin.Default()
	router.POST("/love", func(c *gin.Context) {
		name := c.PostForm("name")
		user, err := c.Cookie("name")
		if err != nil {
			return
		}
		err = redisDB.Set(user, name, -1).Err()
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
}

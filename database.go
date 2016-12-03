
import "gopkg.in/redis.v5"


type DB struct{
    client redis.Client
}

func NewDatabase(address string,port int,  password string) *DB {
    client := redis.NewClient(&redis.Options{
        Addr:     address +":"+port,
        Password: password, // no password set
        DB:       0,  // use default DB
    })

    return new &DB{client}
}

func (*DB) isAccessible() boolean {
    _, err := client.Ping().Result()
    return err == nil
}
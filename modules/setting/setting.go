package setting

//import "github.com/go-ini/ini"

var (
	Connection *Connection
)

type Connection struct {
	Host string
	Username string
	Password string
	Database string
	Schema string
}

func (c *Connection) GetConnectionSettings() Connection {
	return &Connection{
		Host: "localhost",
		Schema: "public",
		Database: "goclerk",
	}
}
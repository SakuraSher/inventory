package inventory

const DB_CONNECTIONSTRING = "postgresql://postgres:1234qwer@localhost:5432/inventory"
const DB_SERVERPORT = "8080"

type Config struct {
	DBConnectionString string
	ServerPort         string
}

func LoadConfig() *Config {
	return &Config{
		DBConnectionString: DB_CONNECTIONSTRING, // os.Getenv("DB_CONNECTIONSTRING"),
		ServerPort:         DB_SERVERPORT,       //      os.Getenv("DB_SERVERPORT"),
	}
}

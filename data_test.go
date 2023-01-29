package goconfigure

const CONFIG_COMPLEX_FILE = "./test/testdata/config-complex.json"

const CONFIG_EMPTY_FILE = "./test/testdata/config-empty.json"

const CONFIG_COMPLEX_STRING = `
{
  "server": {
    "port": 8081
  },
  "database": {
    "employees": {
      "type": "postgres",
      "postgres": {
        "url": "postgres://localhost:12345/employees_db"
      }
    }
  },
  "messaging": {
    "kafka": {
      "type": "kafka",
      "kafka": {
        "bootstrapServers": "localhost:12345"
      }
    }
  },
  "emails": [
    { "name":  "test1", "email": "test1@test.local"},
    { "name":  "test2", "email": "test2@test.local"},
    { "name":  "test3", "email": "test3@test.local"}
  ]
}
`

const CONFIG_EMPTY_STRING = "{}"

var CONFIG_COMPLEX_PROPERTIES = ConfigProperties{
	"server": map[string]any{
		"port": 8081,
	},
	"database": map[string]any{
		"employees": map[string]any{
			"type": "postgres",
			"postgres": map[string]any{
				"url": "postgres://localhost:12345/employees_db",
			},
		},
	},
	"messaging": map[string]any{
		"kafka": map[string]any{
			"type": "kafka",
			"kafka": map[string]any{
				"bootstrapServers": "localhost:12345",
			},
		},
	},
	"emails": []struct {
		name  string
		email string
	}{
		{"test1", "test1@test.local"},
		{"test2", "test2@test.local"},
		{"test3", "test3@test.local"},
	},
}

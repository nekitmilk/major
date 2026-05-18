package parser

import (
	"major/internal/models"
	"testing"
)

func FuzzWithDefaulthData(f *testing.F) {
	configReq := &models.CheckRequest{ // пример yml файла для дальнейшей мутации
		Config: `{
  "version": "1.0",
  "debug": true,
  "log": {
    "level": "debug",
    "output": "stdout"
  },
  "database": {
    "host": "localhost",
    "user": "admin",
    "password": "superSecret123",
    "name": "appdb"
  },
  "storage": {
    "digest_algorithm": "MD5"
  },
  "server": {
    "bind": "127.0.0.1",
    "port": 8080
  }
}`,
	}
	f.Add([]byte(configReq.Config))

	f.Fuzz(func(t *testing.T, data []byte) {
		_, err := Parse(data)
		if err != nil {
			t.Logf("got expected error for input %q: %v", data, err)
		}
	})
}

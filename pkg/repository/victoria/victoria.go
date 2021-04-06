package victoria

import (
	"database/sql"

	"github.com/luckyshmo/gateway/models"
	"github.com/sirupsen/logrus"
)

type Victoria struct {
	SqlDB *sql.DB
}

func (vi *Victoria) WriteData(data ...models.Data) error {
	logrus.Info("write to victoria")
	return nil
}

package victoria

import (
	"database/sql"

	"github.com/luckyshmo/gateway/models"
)

type Victoria struct {
	SqlDB *sql.DB
}

func (vi *Victoria) WriteData(...models.RawData) error {
	return nil
}

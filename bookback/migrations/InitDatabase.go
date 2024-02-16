package migrations

import "fmt"

type Configuration struct {
	Database struct {
		Migration bool
	}
}

func InitMasterData(cfg Configuration) error {
	const op = "migration.initMasterData"

	if cfg.Database.Migration {
		fmt.Println(op, "Migration is enabled")
	}

	return nil
}

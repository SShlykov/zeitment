package app

import "github.com/SShlykov/zeitment/bookback/internal/infrastructure/minio"

func (app *App) initMinio() error {
	client, err := minio.NewMinioClient()
	if err != nil {
		return err
	}
	app.minio = client
	return nil
}

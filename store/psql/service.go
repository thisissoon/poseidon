package psql

import (
	"fmt"

	"compass/store"

	"github.com/jmoiron/sqlx"
)

// Service table name
const ServiceTableName = "service"

// Service upsert query
var UpsertServiceQry = fmt.Sprintf(`
	INSERT INTO "%s" (
		"logical_name",
		"namespace",
		"description")
	VALUES ($1,$2,$3)
	ON CONFLICT ON CONSTRAINT service_tbl_pk DO
	UPDATE SET
		update_date=timezone('UTC'::text, now()),
		namespace=excluded.namespace,
		description=excluded.description
	RETURNING *;`, ServiceTableName)

// ServiceStore handles CRUD opperations for services in psql
type ServiceStore struct {
	db *sqlx.DB
}

// PutService implements the store.ServicePutter and executes
// an Upsert query to create or update a service
func (store *ServiceStore) PutService(service *store.Service) (*store.Service, error) {
	return nil, nil
}

// upsertService executes an Upsert query to create or update a service
func upsertService(db sqlx.Queryer, service *store.Service) (*store.Service, error) {
	var svc *store.Service
	err := QueryRowx(
		db,
		RowHandlerFunc(func(row *sqlx.Row) error {
			return row.StructScan(svc)
		}),
		UpsertServiceQry,
		service.LogicalName,
		service.Namesapace,
		service.Description)
	return svc, err
}

// NewServiceStore returns a new ServiceStore
func NewServiceStore(db *sqlx.DB) *ServiceStore {
	return &ServiceStore{
		db: db,
	}
}
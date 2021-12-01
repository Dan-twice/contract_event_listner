package db

import (
	"database/sql"
	"eth_service/internal/config"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"eth_service/internal/transaction"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const eventsTable = "events"

var eventsColumns = []string{"id", "works"}

// squirrel uses for more complicated queries with some additional
// conditions, like WHERE

type Eventer struct {
	db *pgdb.DB
	selector squirrel.SelectBuilder
}

func NewEventer(db *pgdb.DB) config.EventerInterface {
	return &Eventer{
		db: db,
		selector: squirrel.Select(eventsColumns...).From(eventsTable),
	}
}

func (ev Eventer) Add(event *transaction.TransactionTestEvent) error {
	if event == nil{
		return nil
	}

	query := squirrel.Insert(eventsTable).
		Columns(eventsColumns[1:]...).
		Values(event.Success)

	err := ev.db.Exec(query)
	if err != nil {
		return errors.Wrap(err, "failed to execute insert query for events table")
	}

	return nil
}

func (ev Eventer) GetById(id int64) config.EventerInterface {
	// Will it change the selector field
	ev.selector = ev.selector.Where(squirrel.Eq{"id": id})  // ev.selector
	return ev
}

func (ev Eventer) Select() ([]config.Event, error) {
	result := make([]config.Event, 0, 2)
	err := ev.db.Select(&result, ev.selector)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to select issuances")
	}

	return result, nil
}

package migrations

import (
	"fmt"
	"gopkg.in/go-pg/migrations.v4"
	"gopkg.in/pg.v4"
	"strconv"
	"strings"
)

var (
	Migrations []migrations.Migration
)

func Register(migration migrations.Migration) {

	Migrations = append(Migrations, migration)
}

func Run(db migrations.DB, a ...string) (oldVersion, newVersion int64, err error) {
	err = createTables(db)
	if err != nil {
		return
	}

	oldVersion, err = migrations.Version(db)
	if err != nil {
		return
	}
	newVersion = oldVersion

	var cmd string
	if len(a) > 0 {
		cmd = a[0]
	}

	switch cmd {
	case "version":
		return
	case "up", "":
		for i := range Migrations {
			m := &Migrations[i]
			if m.Version <= oldVersion {
				continue
			}
			err = m.Up(db)
			if err != nil {
				return
			}
			newVersion = m.Version
			err = migrations.SetVersion(db, newVersion)
			if err != nil {
				return
			}
		}
		return
	case "down":
		if oldVersion == 0 {
			return
		}

		var m *migrations.Migration
		for i := len(Migrations) - 1; i >= 0; i-- {
			mm := &Migrations[i]
			if mm.Version <= oldVersion {
				m = mm
				break
			}
		}
		if m == nil {
			err = fmt.Errorf("migration %d not found\n", oldVersion)
			return
		}

		if m.Down != nil {
			err = m.Down(db)
			if err != nil {
				return
			}
		}

		newVersion = m.Version - 1
		err = migrations.SetVersion(db, newVersion)
		if err != nil {
			return
		}
		return
	case "set_version":
		if len(a) < 2 {
			err = fmt.Errorf("set_version requires version as 2nd arg, e.g. set_version 42")
			return
		}

		newVersion, err = strconv.ParseInt(a[1], 10, 64)
		if err != nil {
			return
		}
		err = migrations.SetVersion(db, newVersion)
		return
	default:
		err = fmt.Errorf("unsupported command: %q", cmd)
		return
	}
}

func createTables(db migrations.DB) error {
	if ind := strings.Index(migrations.TableName, "."); ind >= 0 {
		_, err := db.Exec(`CREATE SCHEMA IF NOT EXISTS ?`, pg.Q(migrations.TableName[:ind]))
		if err != nil {
			return err
		}
	}

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS ? (
			id serial,
			version bigint,
			created_at timestamptz
		)
	`, pg.Q(migrations.TableName))
	return err
}

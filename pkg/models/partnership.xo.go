// Package models contains the database interaction model code
//
// GENERATED BY GOSCHEMA. DO NOT EDIT.
package models

import (
	"fmt"

	"github.com/Jacobbrewer1/patcher/inserter"
	"github.com/prometheus/client_golang/prometheus"
)

// Partnership represents a row from 'partnership'.
type Partnership struct {
	Id        int `db:"id,autoinc,pk"`
	PlayerAId int `db:"player_a_id"`
	PlayerBId int `db:"player_b_id"`
	TeamId    int `db:"team_id"`
}

// Insert inserts the Partnership to the database.
func (m *Partnership) Insert(db DB) error {
	t := prometheus.NewTimer(DatabaseLatency.WithLabelValues("insert_Partnership"))
	defer t.ObserveDuration()

	const sqlstr = "INSERT INTO partnership (" +
		"`player_a_id`, `player_b_id`, `team_id`" +
		") VALUES (" +
		"?, ?, ?" +
		")"

	DBLog(sqlstr, m.PlayerAId, m.PlayerBId, m.TeamId)
	res, err := db.Exec(sqlstr, m.PlayerAId, m.PlayerBId, m.TeamId)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	m.Id = int(id)
	return nil
}

func InsertManyPartnerships(db DB, ms ...*Partnership) error {
	if len(ms) == 0 {
		return nil
	}

	t := prometheus.NewTimer(DatabaseLatency.WithLabelValues("insert_many_Partnership"))
	defer t.ObserveDuration()

	vals := make([]any, 0, len(ms))
	for _, m := range ms {
		// Dereference the pointer to get the struct value.
		vals = append(vals, []any{*m})
	}

	sqlstr, args, err := inserter.NewBatch(vals, inserter.WithTable("partnership")).GenerateSQL()
	if err != nil {
		return fmt.Errorf("failed to create batch insert: %w", err)
	}

	DBLog(sqlstr, args...)
	res, err := db.Exec(sqlstr, args...)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	for i, m := range ms {
		m.Id = int(id + int64(i))
	}

	return nil
}

// IsPrimaryKeySet returns true if all primary key fields are set to none zero values
func (m *Partnership) IsPrimaryKeySet() bool {
	return IsKeySet(m.Id)
}

// Update updates the Partnership in the database.
func (m *Partnership) Update(db DB) error {
	t := prometheus.NewTimer(DatabaseLatency.WithLabelValues("update_Partnership"))
	defer t.ObserveDuration()

	const sqlstr = "UPDATE partnership " +
		"SET `player_a_id` = ?, `player_b_id` = ?, `team_id` = ? " +
		"WHERE `id` = ?"

	DBLog(sqlstr, m.PlayerAId, m.PlayerBId, m.TeamId, m.Id)
	res, err := db.Exec(sqlstr, m.PlayerAId, m.PlayerBId, m.TeamId, m.Id)
	if err != nil {
		return err
	}

	// Requires clientFoundRows=true
	if i, err := res.RowsAffected(); err != nil {
		return err
	} else if i <= 0 {
		return ErrNoAffectedRows
	}

	return nil
}

// InsertWithUpdate inserts the Partnership to the database, and tries to update
// on unique constraint violations.
func (m *Partnership) InsertWithUpdate(db DB) error {
	t := prometheus.NewTimer(DatabaseLatency.WithLabelValues("insert_update_Partnership"))
	defer t.ObserveDuration()

	const sqlstr = "INSERT INTO partnership (" +
		"`player_a_id`, `player_b_id`, `team_id`" +
		") VALUES (" +
		"?, ?, ?" +
		") ON DUPLICATE KEY UPDATE " +
		"`player_a_id` = VALUES(`player_a_id`), `player_b_id` = VALUES(`player_b_id`), `team_id` = VALUES(`team_id`)"

	DBLog(sqlstr, m.PlayerAId, m.PlayerBId, m.TeamId)
	res, err := db.Exec(sqlstr, m.PlayerAId, m.PlayerBId, m.TeamId)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	m.Id = int(id)
	return nil
}

// Save saves the Partnership to the database.
func (m *Partnership) Save(db DB) error {
	if m.IsPrimaryKeySet() {
		return m.Update(db)
	}
	return m.Insert(db)
}

// SaveOrUpdate saves the Partnership to the database, but tries to update
// on unique constraint violations.
func (m *Partnership) SaveOrUpdate(db DB) error {
	if m.IsPrimaryKeySet() {
		return m.Update(db)
	}
	return m.InsertWithUpdate(db)
}

// Delete deletes the Partnership from the database.
func (m *Partnership) Delete(db DB) error {
	t := prometheus.NewTimer(DatabaseLatency.WithLabelValues("delete_Partnership"))
	defer t.ObserveDuration()

	const sqlstr = "DELETE FROM partnership WHERE `id` = ?"

	DBLog(sqlstr, m.Id)
	_, err := db.Exec(sqlstr, m.Id)

	return err
}

// PartnershipById retrieves a row from 'partnership' as a Partnership.
//
// Generated from primary key.
func PartnershipById(db DB, id int) (*Partnership, error) {
	t := prometheus.NewTimer(DatabaseLatency.WithLabelValues("insert_Partnership"))
	defer t.ObserveDuration()

	const sqlstr = "SELECT `id`, `player_a_id`, `player_b_id`, `team_id` " +
		"FROM partnership " +
		"WHERE `id` = ?"

	DBLog(sqlstr, id)
	var m Partnership
	if err := db.Get(&m, sqlstr, id); err != nil {
		return nil, err
	}

	return &m, nil
}

// GetPlayerBIdPlayer Gets an instance of Player
//
// Generated from constraint partnership_player_id_fk
func (m *Partnership) GetPlayerBIdPlayer(db DB) (*Player, error) {
	return PlayerById(db, m.PlayerBId)
}

// GetPlayerAIdPlayer Gets an instance of Player
//
// Generated from constraint partnership_player_id_fk2
func (m *Partnership) GetPlayerAIdPlayer(db DB) (*Player, error) {
	return PlayerById(db, m.PlayerAId)
}

// GetTeamIdTeam Gets an instance of Team
//
// Generated from constraint partnership_team_id_fk
func (m *Partnership) GetTeamIdTeam(db DB) (*Team, error) {
	return TeamById(db, m.TeamId)
}

// PartnershipByPlayerAIdPlayerBIdTeamId retrieves a row from 'partnership' as a *Partnership.
//
// Generated from index 'partnership_player_a_id_player_b_id_uindex' of type 'unique'.
func PartnershipByPlayerAIdPlayerBIdTeamId(db DB, playerAId int, playerBId int, teamId int) (*Partnership, error) {
	t := prometheus.NewTimer(DatabaseLatency.WithLabelValues("insert_Partnership"))
	defer t.ObserveDuration()

	const sqlstr = "SELECT `id`, `player_a_id`, `player_b_id`, `team_id` " +
		"FROM partnership " +
		"WHERE `player_a_id` = ? AND `player_b_id` = ? AND `team_id` = ?"

	DBLog(sqlstr, playerAId, playerBId, teamId)
	var m Partnership
	if err := db.Get(&m, sqlstr, playerAId, playerBId, teamId); err != nil {
		return nil, err
	}

	return &m, nil
}

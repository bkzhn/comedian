package storage

import (
	"time"

	// This line is must for working MySQL database
	_ "github.com/go-sql-driver/mysql"
	"gitlab.com/team-monitoring/comedian/model"
)

// CreateStanduper creates comedian entry in database
func (m *DB) CreateStanduper(s model.Standuper) (model.Standuper, error) {
	err := s.Validate()
	if err != nil {
		return s, err
	}
	res, err := m.DB.Exec(
		"INSERT INTO `channel_members` (team_id, user_id, channel_id, submitted_standup_today, role_in_channel, created) VALUES (?,?,?,?,?,?)",
		s.TeamID, s.UserID, s.ChannelID, true, s.RoleInChannel, time.Now())
	if err != nil {
		return s, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return s, err
	}
	s.ID = id

	return s, nil
}

// UpdateStanduper updates Standuper entry in database
func (m *DB) UpdateStanduper(st model.Standuper) (model.Standuper, error) {
	_, err := m.DB.Exec(
		"UPDATE `channel_members` SET submitted_standup_today=?, role_in_channel=? WHERE id=?",
		st.SubmittedStandupToday, st.RoleInChannel, st.ID,
	)
	if err != nil {
		return st, err
	}
	var i model.Standuper
	err = m.DB.Get(&i, "SELECT * FROM `channel_members` WHERE id=?", st.ID)
	return i, err
}

//FindStansuperByUserID finds user in channel
func (m *DB) FindStansuperByUserID(userID, channelID string) (model.Standuper, error) {
	var u model.Standuper
	err := m.DB.Get(&u, "SELECT * FROM `channel_members` WHERE user_id=? AND channel_id=?", userID, channelID)
	return u, err
}

// ListStandupers returns array of standup entries from database
func (m *DB) ListStandupers() ([]model.Standuper, error) {
	items := []model.Standuper{}
	err := m.DB.Select(&items, "SELECT * FROM `channel_members`")
	return items, err
}

//GetStanduper returns a standuper
func (m *DB) GetStanduper(id int64) (model.Standuper, error) {
	standuper := model.Standuper{}
	err := m.DB.Get(&standuper, "SELECT * FROM channel_members where id=?", id)
	return standuper, err
}

// ListChannelStandupers returns array of standup entries from database
func (m *DB) ListChannelStandupers(channelID string) ([]model.Standuper, error) {
	items := []model.Standuper{}
	err := m.DB.Select(&items, "SELECT * FROM `channel_members` WHERE channel_id=?", channelID)
	return items, err
}

// DeleteStanduper deletes channel_members entry from database
func (m *DB) DeleteStanduper(id int64) error {
	_, err := m.DB.Exec("DELETE FROM `channel_members` WHERE id=?", id)
	return err
}
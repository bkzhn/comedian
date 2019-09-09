package storage

import (
	"testing"
	"time"

	"github.com/maddevsio/comedian/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotification(t *testing.T) {
	tt := time.Now().Unix() + 30*60
	n := model.NotificationThread{
		ChannelID:        "1",
		UserID:           "1",
		NotificationTime: tt,
		ReminderCounter:  0,
	}

	notification, err := db.CreateNotificationThread(n)
	require.NoError(t, err)
	assert.Equal(t, "1", notification.ChannelID)
	assert.Equal(t, "1", notification.UserID)
	assert.Equal(t, tt, notification.NotificationTime)
	assert.Equal(t, 0, notification.ReminderCounter)

	thread, err := db.SelectNotificationsThread(notification.ChannelID)
	require.NoError(t, err)
	assert.Equal(t, thread.ChannelID, notification.ChannelID)

	err = db.DeleteNotificationThread(notification.ID)
	require.NoError(t, err)

	thread, err = db.SelectNotificationsThread(notification.ChannelID)
	assert.Equal(t, 0, thread.ReminderCounter)
	assert.Equal(t, "", thread.UserID)
	assert.Equal(t, int64(0), thread.NotificationTime)
	assert.Equal(t, "", thread.ChannelID)

	n = model.NotificationThread{
		ChannelID:        "1",
		UserID:           "User1",
		NotificationTime: tt,
		ReminderCounter:  0,
	}

	nt, err := db.CreateNotificationThread(n)
	require.NoError(t, err)

	err = db.UpdateNotificationThread(nt.ID, nt.ChannelID, tt)
	require.NoError(t, err)

	thread, err = db.SelectNotificationsThread(nt.ChannelID)
	require.NoError(t, err)
	assert.Equal(t, 1, thread.ReminderCounter)

	err = db.DeleteNotificationThread(nt.ID)
	require.NoError(t, err)
}

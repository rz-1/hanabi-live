package table

type chatNormalData struct {
	userID   int
	username string
	msg      string
}

func (m *Manager) ChatNormal(userID int, username string, msg string) {
	m.newRequest(requestTypeChatNormal, &chatNormalData{ // nolint: errcheck
		userID:   userID,
		username: username,
		msg:      msg,
	})
}

func (m *Manager) chatNormal(data interface{}) {
	var d *chatNormalData
	if v, ok := data.(*chatNormalData); !ok {
		m.logger.Errorf("Failed type assertion for data of type: %T", d)
		return
	} else {
		d = v
	}

	m.chat(d.userID, d.username, d.msg, false)
}
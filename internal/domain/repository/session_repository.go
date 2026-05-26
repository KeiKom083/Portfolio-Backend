package repository

// SessionRepository はセッションの永続化に関するインターフェース。
type SessionRepository interface {
	Create(userID string) (sessionID string, err error)
	Get(sessionID string) (userID string, ok bool)
	Delete(sessionID string)
}

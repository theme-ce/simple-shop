package token

import "time"

type Maker interface {
	CreateToken(username string, is_admin bool, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

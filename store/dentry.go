package store

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type DentryPutter interface {
	PutDentry(*Dentry) (*Dentry, error)
}

type DentryStore interface {
	DentryPutter
}

type Dentry struct {
	Id          uuid.UUID
	CreateDate  time.Time
	UpdateDate  time.Time
	Prefix      string
	Destination string
	Namespace   string
	Priority    int32
	ServiceId   string
}
package facade

import (
	"facade/notify"
	"facade/storage"
	"facade/validate"
)

type Facade struct {
	to      string
	comment string

	tkn  validate.Token
	perm validate.Permission

	store    storage.Storage
	notifier notify.Email
}

func New(to, comment, token, user string) Facade {
	return Facade{
		to:       to,
		comment:  comment,
		tkn:      validate.NewToken(token),
		perm:     validate.NewPermission(user),
		store:    storage.New("postgres"),
		notifier: notify.New(),
	}
}

func (f *Facade) Comment() error {
	if err := f.tkn.Validate(); err != nil {
		return err
	}

	if err := f.perm.Validate(); err != nil {
		return err
	}

	f.store.Save(f.comment)
	f.notifier.Send(f.to, f.comment)

	return nil
}

func (f *Facade) Notify() {
	f.notifier.Send(f.to, f.comment)
}

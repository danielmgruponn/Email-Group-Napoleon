package mine

import (
	"context"
	"napoleon-email/src/app/domain/application/mine"
	contactmine "napoleon-email/src/app/domain/model/contact_mine"
	contactminerepositoryinterface "napoleon-email/src/app/domain/repository/contact_mine_repository_interface"
	"time"
)

type NapoMineEmailApplicationImp struct {
	contactNapoMineRepository contactminerepositoryinterface.ContactNapoleonGoldMineRepositoryInterface
}

var _ mine.ContactMineEmailApplicationInterface = &NapoMineEmailApplicationImp{}

func NewNapoMineEmailApplicationImp(contactMineRepository contactminerepositoryinterface.ContactNapoleonGoldMineRepositoryInterface) *NapoMineEmailApplicationImp {
	return &NapoMineEmailApplicationImp{
		contactNapoMineRepository: contactMineRepository,
	}
}

func (app *NapoMineEmailApplicationImp) CreateNapoMineEmail(ctx context.Context, email, name, to, subject, text, html string) (*contactmine.ContactNapoleonMine, error) {
	sending := true
	model := &contactmine.ContactNapoleonMine{
		ClientEmail: email,
		ClientName:  name,
		CreatedAt:   time.Now(),
		To:          to,
		Message: contactmine.MessageMineContact{
			HTML:    html,
			Subject: subject,
			Text:    text,
		},
		Sending: &sending,
	}
	return app.contactNapoMineRepository.Create(ctx, model)
}

func (app *NapoMineEmailApplicationImp) UpdateNapoMineEmail(ctx context.Context, contact *contactmine.ContactNapoleonMine) (*contactmine.ContactNapoleonMine, error) {
	return app.contactNapoMineRepository.Update(ctx, contact)
}

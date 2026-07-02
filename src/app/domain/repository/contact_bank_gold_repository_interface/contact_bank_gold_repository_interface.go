package contactbankgoldrepositoryinterface

import (
	"context"
	contactbankgold "napoleon-email/src/app/domain/model/contact_bank_gold"
)

type ContactBankGoldRepositoryInterface interface {
	Create(ctx context.Context, contact *contactbankgold.ContactBankGold) (*contactbankgold.ContactBankGold, error)
	FindAll(ctx context.Context) ([]contactbankgold.ContactBankGold, error)
	FindById(ctx context.Context, id string) (*contactbankgold.ContactBankGold, error)
	Update(ctx context.Context, contact *contactbankgold.ContactBankGold) (*contactbankgold.ContactBankGold, error)
	Remove(ctx context.Context, id string) error
}

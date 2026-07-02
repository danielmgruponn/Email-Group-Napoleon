package contactbankgoldrepositoryimp

import (
	"context"
	contactbankgold "napoleon-email/src/app/domain/model/contact_bank_gold"
	contactbankgoldrepositoryinterface "napoleon-email/src/app/domain/repository/contact_bank_gold_repository_interface"
	"napoleon-email/src/pkg/logger"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ContactBankGoldRepositoryImp struct {
	fstore *firestore.Client
}

var _ contactbankgoldrepositoryinterface.ContactBankGoldRepositoryInterface = &ContactBankGoldRepositoryImp{}

const (
	collectionName = "contacto_bank_gold"
)

func NewContactBankGoldRepositoryImp(fstore *firestore.Client) *ContactBankGoldRepositoryImp {
	return &ContactBankGoldRepositoryImp{fstore: fstore}
}

func (r *ContactBankGoldRepositoryImp) Create(ctx context.Context, contact *contactbankgold.ContactBankGold) (*contactbankgold.ContactBankGold, error) {
	docRef, _, err := r.fstore.Collection(collectionName).Add(ctx, contact)
	if err != nil {
		logger.LogError("Error in create contact bank gold", err, logger.LogStruct{Action: "Error in create contact bank gold"})
		return nil, err
	}
	contact.ID = docRef.ID
	return contact, nil 
}

func (r *ContactBankGoldRepositoryImp) FindAll(ctx context.Context) ([]contactbankgold.ContactBankGold, error) {
	var contacts []contactbankgold.ContactBankGold

	iter := r.fstore.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.LogError("Error in iteration contacts of bank gold", err, logger.LogStruct{Action: "Error in iteration contacts"})
			return nil, err
		}
		var contact contactbankgold.ContactBankGold
		if err := doc.DataTo(&contact); err != nil {
			logger.LogError("Error in unmarshal data of contact bank gold", err, logger.LogStruct{Action: "Error in unmarshal data of contact bank gold"})
			return nil, err
		}
		contact.ID = doc.Ref.ID
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactBankGoldRepositoryImp) FindById(ctx context.Context, id string) (*contactbankgold.ContactBankGold, error) {
	doc, err := r.fstore.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		logger.LogError("Error in find contact bank gold by id", err, logger.LogStruct{Action: "Error in find contact bank gold by id"})
		return nil, err
	}
	var contact contactbankgold.ContactBankGold
	if err := doc.DataTo(&contact); err != nil {
		logger.LogError("Error in unmarshal data of contact bank gold", err, logger.LogStruct{Action: "Error in unmarshal data of contact bank gold"})
		return nil, err
	}
	contact.ID = doc.Ref.ID
	return &contact, nil
}

func (r *ContactBankGoldRepositoryImp) Update(ctx context.Context, contact *contactbankgold.ContactBankGold) (*contactbankgold.ContactBankGold, error) {
	_, err := r.fstore.Collection(collectionName).Doc(contact.ID).Set(ctx, contact)
	if err != nil {
		logger.LogError("Error in update contact bank gold", err, logger.LogStruct{Action: "Error in update contact bank gold"})
		return nil, err
	}
	return contact, nil
}

func (r *ContactBankGoldRepositoryImp) Remove(ctx context.Context, id string) error {
	_, err := r.fstore.Collection(collectionName).Doc(id).Delete(ctx)
	if err != nil {
		logger.LogError("Error in remove contact bank gold", err, logger.LogStruct{Action: "Error in remove contact bank gold"})
		return err
	}
	return nil
}

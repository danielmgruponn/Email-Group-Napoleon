package contactminerepositoryimp

import (
	"context"
	contactmine "napoleon-email/src/app/domain/model/contact_mine"
	contactminerepositoryinterface "napoleon-email/src/app/domain/repository/contact_mine_repository_interface"
	"napoleon-email/src/pkg/logger"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type ContactNapoleonGoldMineRepositoryImp struct {
	fstore *firestore.Client
}

var _ contactminerepositoryinterface.ContactNapoleonGoldMineRepositoryInterface = &ContactNapoleonGoldMineRepositoryImp{}

const (
	collectionName = "contacto_mine"
)

func NewContactNapoleonGoldMineRepositoryImp(fstore *firestore.Client) *ContactNapoleonGoldMineRepositoryImp {
	return &ContactNapoleonGoldMineRepositoryImp{fstore: fstore}
}

func (r *ContactNapoleonGoldMineRepositoryImp) Create(ctx context.Context, contact *contactmine.ContactNapoleonMine) (*contactmine.ContactNapoleonMine, error) {
	docRef, _, err := r.fstore.Collection(collectionName).Add(ctx, contact)
	if err != nil {
		logger.LogError("Error in create contact napoleon gold mine", err, logger.LogStruct{Action: "Error in create contact napoleon gold mine"})
		return nil, err
	}
	contact.ID = docRef.ID
	return contact, nil 
}

func (r *ContactNapoleonGoldMineRepositoryImp) FindAll(ctx context.Context) ([]contactmine.ContactNapoleonMine, error) {
	var contacts []contactmine.ContactNapoleonMine

	iter := r.fstore.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.LogError("Error in iteration contacts of napoleon gold mine", err, logger.LogStruct{Action: "Error in iteration contacts"})
			return nil, err
		}
		var contact contactmine.ContactNapoleonMine
		if err := doc.DataTo(&contact); err != nil {
			logger.LogError("Error in unmarshal data of contact napoleon gold mine", err, logger.LogStruct{Action: "Error in unmarshal data of contact napoleon gold mine"})
			return nil, err
		}
		contact.ID = doc.Ref.ID
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ContactNapoleonGoldMineRepositoryImp) FindById(ctx context.Context, id string) (*contactmine.ContactNapoleonMine, error) {
	doc, err := r.fstore.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		logger.LogError("Error in find contact napoleon gold mine by id", err, logger.LogStruct{Action: "Error in find contact napoleon gold mine by id"})
		return nil, err
	}
	var contact contactmine.ContactNapoleonMine
	if err := doc.DataTo(&contact); err != nil {
		logger.LogError("Error in unmarshal data of contact napoleon gold mine", err, logger.LogStruct{Action: "Error in unmarshal data of contact napoleon gold mine"})
		return nil, err
	}
	contact.ID = doc.Ref.ID
	return &contact, nil
}

func (r *ContactNapoleonGoldMineRepositoryImp) Update(ctx context.Context, contact *contactmine.ContactNapoleonMine) (*contactmine.ContactNapoleonMine, error) {
	_, err := r.fstore.Collection(collectionName).Doc(contact.ID).Set(ctx, contact)
	if err != nil {
		logger.LogError("Error in update contact napoleon gold mine", err, logger.LogStruct{Action: "Error in update contact napoleon gold mine"})
		return nil, err
	}
	return contact, nil
}

func (r *ContactNapoleonGoldMineRepositoryImp) Remove(ctx context.Context, id string) error {
	_, err := r.fstore.Collection(collectionName).Doc(id).Delete(ctx)
	if err != nil {
		logger.LogError("Error in remove contact napoleon gold mine", err, logger.LogStruct{Action: "Error in remove contact napoleon gold mine"})
		return err
	}
	return nil
}

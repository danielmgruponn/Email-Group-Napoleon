package infrastructure

import (
	contactbankgoldrepositoryinterface "napoleon-email/src/app/domain/repository/contact_bank_gold_repository_interface"
	contactgroupnapoleonrepositoryinterface "napoleon-email/src/app/domain/repository/contact_group_napoleon_repository_interface"
	contactminerepositoryinterface "napoleon-email/src/app/domain/repository/contact_mine_repository_interface"
	contactnapoleonrepositoryinterface "napoleon-email/src/app/domain/repository/contact_napoleon_repository_interface"
	contactbankgoldrepositoryimp "napoleon-email/src/app/infrastructure/contact_bank_gold_repository_imp"
	contactgroupnapoleonrepositoryimp "napoleon-email/src/app/infrastructure/contact_group_napoleon_repository_imp"
	contactminerepositoryimp "napoleon-email/src/app/infrastructure/contact_mine_repository_imp"
	contactnapoleonrepositoryimp "napoleon-email/src/app/infrastructure/contact_napoleon_repository_imp"
	"napoleon-email/src/pkg/logger"
	"sync"

	"cloud.google.com/go/firestore"
)

type Kernel struct {
	fstore *firestore.Client
	mu sync.Mutex

	contactNapoRepository contactnapoleonrepositoryinterface.ContactNapoleonRepositoryInterface
	contactNapoMineRepository contactminerepositoryinterface.ContactNapoleonGoldMineRepositoryInterface
	contactGroupNapoRepository contactgroupnapoleonrepositoryinterface.ContactGroupNapoleonRepositoryInterface
	contactBankGoldRepository contactbankgoldrepositoryinterface.ContactBankGoldRepositoryInterface
}

var (
	instance *Kernel
	once sync.Once
)

func InitKernel(fstore *firestore.Client) {
	once.Do(func() {
		instance = &Kernel{
			fstore: fstore,
		}
	})
}

func GetKernel() *Kernel {
	if instance == nil {
		logger.LogError("Kernel is not initialized", nil, logger.LogStruct{Action: "GetKernel"})
		panic("Kernel is not initialized")
	}
	return instance
}

func (k *Kernel) GetContactNapoleonRepository() contactnapoleonrepositoryinterface.ContactNapoleonRepositoryInterface {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.contactNapoRepository == nil {
		k.contactNapoRepository = contactnapoleonrepositoryimp.NewContactNapoleonRepositoryImp(k.fstore)
	}
	return k.contactNapoRepository
}

func (k *Kernel) GetContactNapoleonMineRepository() contactminerepositoryinterface.ContactNapoleonGoldMineRepositoryInterface {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.contactNapoMineRepository == nil {
		k.contactNapoMineRepository = contactminerepositoryimp.NewContactNapoleonGoldMineRepositoryImp(k.fstore)
	}
	return k.contactNapoMineRepository
}

func (k *Kernel) GetContactGroupNapoleonRepository() contactgroupnapoleonrepositoryinterface.ContactGroupNapoleonRepositoryInterface {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.contactGroupNapoRepository == nil {
		k.contactGroupNapoRepository = contactgroupnapoleonrepositoryimp.NewContactGroupNapoleonRepositoryImp(k.fstore)
	}
	return k.contactGroupNapoRepository
}

func (k *Kernel) GetContactBankGoldRepository() contactbankgoldrepositoryinterface.ContactBankGoldRepositoryInterface {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.contactBankGoldRepository == nil {
		k.contactBankGoldRepository = contactbankgoldrepositoryimp.NewContactBankGoldRepositoryImp(k.fstore)
	}
	return k.contactBankGoldRepository
}

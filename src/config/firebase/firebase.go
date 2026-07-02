package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"napoleon-email/src/pkg/logger"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type FirebaseConfig struct {
	ProjectID string `json:"project_id"`
}

func ConnectionFirestore(ctx context.Context, credentials string) (*firestore.Client, error) {
	fileBytes, err := os.ReadFile(credentials)
	if err != nil {
		logger.LogError("Error reading credentials file", err, logger.LogStruct{Action: "Read File Credentials"} )
		return nil, err
	} 
	var config FirebaseConfig
	err = json.Unmarshal(fileBytes, &config)
	if err != nil {
		logger.LogError("Error unmarshalling credentials", err, logger.LogStruct{Action: "Unmarshal Credentials"})
		return nil, err
	}
	if config.ProjectID == "" {
		logger.LogError("Project ID is empty", err, logger.LogStruct{Action: "Empty Project ID"})
		return nil, fmt.Errorf("project_id is empty in credentials file")
	}
	opt := option.WithAuthCredentialsFile(option.ServiceAccount, credentials)

	client, err := firestore.NewClient(ctx, config.ProjectID, opt)
	if err != nil {
		logger.LogError("Error creating Firestore client", err, logger.LogStruct{Action: "Create Firestore Client"})
		return nil, err
	}
	return client, nil
}
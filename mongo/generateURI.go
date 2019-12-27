package mongo

import (
	"fmt"
	"github.com/codedbypm/gcloud-secret-manager/secretmanager"
)

func GenerateURI(projectID string) (string, error) {
	const mongoUserSecretName = "agora-secret-mongo-user"
	const mongoPassSecretName = "agora-secret-mongo-pass"

	mongoUserData, err := secretmanager.GetSecretData(mongoUserSecretName, projectID)
	if err != nil {
		return "", fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoUserSecretName, err)
	}

	mongoPassData, err := secretmanager.GetSecretData(mongoPassSecretName, projectID)
	if err != nil {
		return "", fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoPassSecretName, err)
	}

	var uri = fmt.Sprintf("mongodb+srv://%s:%s@agorapolis-001-ymzlz.gcp.mongodb.net", string(mongoUserData), string(mongoPassData))
	return uri, nil
}

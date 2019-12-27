package mongo

import (
	"fmt"
	secretmanager "github.com/codedbypm/gcloud-secret-manager/secretmanager"
)

func generateURI(projectID string, keyRingName string, keyName string) (string, error) {
	const mongoUserSecretName = "agora-secret-mongo-user"
	const mongoPassSecretName = "agora-secret-mongo-pass"
	const keyID = fmt.Sprintf("projects/%s/locations/europe-west1/keyRings/%s/cryptoKeys/%s/cryptoKeyVersions/latest", projectID, keyRingName, keyName)

	mongoUser, err := secretmanager.DecryptSecretSymmetric(mongoUserSecretName, projectID, keyID)
	if err != nil {
		return "", fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoUserSecretName, err)
	}

	mongoPass, err := secretmanager.DecryptSecretSymmetric(mongoPassSecretName, projectID, keyID)
	if err != nil {
		return "", fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoPassSecretName, err)
	}

	var uri = fmt.Sprint("mongodb+srv://%s:%s@agorapolis-001-ymzlz.gcp.mongodb.net", mongoUser, mongoPass)
	return uri, nil
}

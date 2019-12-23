package mongo

import (
	"fmt"
	secretmanager "github.com/codedbypm/gcloud-secret-manager/secretmanager"
)

func generateURI() (string, error) {
	const mongoUserSecretName = "agora-secret-mongo-user"
	const mongoPassSecretName = "agora-secret-mongo-pass"
	const keyID = "projects/agora-262523/locations/europe-west1/keyRings/agora-key-ring/cryptoKeys/agora-key/cryptoKeyVersions/latest"

	mongoUser, err := secretmanager.DecryptSecretSymmetric(mongoUserSecretName, "agora-262523", keyID)
	if err != nil {
		return "", fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoUserSecretName, err)
	}

	mongoPass, err := secretmanager.DecryptSecretSymmetric(mongoPassSecretName, "agora-262523", keyID)
	if err != nil {
		return "", fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoPassSecretName, err)
	}

	var uri = fmt.Sprint("mongodb+srv://%s:%s@agorapolis-001-ymzlz.gcp.mongodb.net", mongoUser, mongoPass)
	return uri, nil
}

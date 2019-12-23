package connection

import (
	"fmt"

	"github.com/codedbypm/gcloud-secret-manager/access"
	"github.com/codedbypm/gcloud-secret-manager/decrypt"
)

func GenerateURI() (string, error) {
	const mongoUserSecretName = "agora-secret-mongo-user"
	const mongoPassSecretName = "agora-secret-mongo-pass"
	const keyID = "projects/agora-262523/locations/europe-west1/keyRings/agora-key-ring/cryptoKeys/agora-key/cryptoKeyVersions/latest"

	mongoUserSecret, secretError := access.GetSecret("agora-262523", mongoUserSecretName)
	if secretError != nil {
		return "", fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoUserSecretName, secretError)
	}

	mongoPassSecret, secretError := access.GetSecret("agora-262523", mongoPassSecretName)
	if secretError != nil {
		return "", fmt.Errorf("Error: could not retrieve secret %s (%s)", mongoPassSecretName, secretError)
	}

	mongoUser, decryptError := decrypt.DecryptSymmetric(keyID, mongoUserSecret.Payload.Data)
	if decryptError != nil {
		return "", fmt.Errorf("Error: could not decrypt secret %s (%s)", mongoUserSecretName, decryptError)
	}

	mongoPass, decryptError := decrypt.DecryptSymmetric(keyID, mongoPassSecret.Payload.Data)
	if decryptError != nil {
		return "", fmt.Errorf("Error: could not decrypt secret %s (%s)", mongoPassSecretName, decryptError)
	}

	var uri = fmt.Sprint("mongodb+srv://%s:%s@agorapolis-001-ymzlz.gcp.mongodb.net", mongoUser, mongoPass)
	return uri, nil

}

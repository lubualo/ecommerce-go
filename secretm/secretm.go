package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/ddessilvestri/ecommerce-go/awsgo"
	"github.com/ddessilvestri/ecommerce-go/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println(" > Secret Request " + secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return secretData, err

	}

	json.Unmarshal([]byte(*key.SecretString), &secretData)
	fmt.Println(" > Successfully Secret Retrieve " + secretName)
	return secretData, nil
}

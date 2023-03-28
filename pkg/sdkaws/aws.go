package sdkaws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lexruntimev2"
)

type AwsSDK struct {
	props          AwsProperties
	botAliasId     string
	botId          string
	localeId       string
	sessionId      string
	accesKeyId     string
	secretAccesKey string
	sessionToken   string
	region         string
}

type AwsProperties struct {
	*lexruntimev2.RecognizeTextInput
}

func NewAwsClient(properties AwsProperties, botaliasid string, botid string, localeid string, sessionid string, acceskeyid string, secretacceskey string, sessiontoken string, region string) AwsSDK {
	awsclient := AwsSDK{
		props:          properties,
		botAliasId:     botaliasid,
		botId:          botid,
		localeId:       localeid,
		sessionId:      sessionid,
		accesKeyId:     acceskeyid,
		secretAccesKey: secretacceskey,
		sessionToken:   sessiontoken,
		region:         region,
	}
	return awsclient
}

func (a *AwsSDK) SendMessage(client AwsProperties) (*lexruntimev2.RecognizeTextOutput, error) {
	mySession, err := session.NewSession(&aws.Config{
		Region:      aws.String(a.region),
		Credentials: credentials.NewStaticCredentials(a.accesKeyId, a.secretAccesKey, ""),
	})

	if err != nil {
		return nil, err
	}
	svc := lexruntimev2.New(mySession)

	req, resp := svc.RecognizeTextRequest(client.RecognizeTextInput)
	err = req.Send()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

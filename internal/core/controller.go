package core

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/service/lexruntimev2"
	"github.com/gin-gonic/gin"
	"github.com/iriscompanyio/awslex-bot/pkg/config"
	"github.com/iriscompanyio/awslex-bot/pkg/helpers"
	"github.com/iriscompanyio/awslex-bot/pkg/sdkaws"
)

var message string = "Hola tienes flores"

func WebhookHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		params := sdkaws.AwsProperties{
			RecognizeTextInput: &lexruntimev2.RecognizeTextInput{
				BotAliasId: &config.Cfg.BotAliasId,
				BotId:      &config.Cfg.BotId,
				LocaleId:   &config.Cfg.LocaleId,
				SessionId:  &config.Cfg.SessionId,
				Text:       &message,
			},
		}

		client := sdkaws.NewAwsClient(params, config.Cfg.BotAliasId, config.Cfg.BotId, config.Cfg.LocaleId, config.Cfg.SessionId, config.Cfg.AccessKeyId, config.Cfg.SecretAccessKey, config.Cfg.SessionToken, config.Cfg.Region)
		fmt.Println("Parámetros enviados: ", params)
		resp, err := client.SendMessage(params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helpers.MessageError(0, err))
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println("Respuesta recibida: ", resp)

		ctx.JSON(http.StatusAccepted, helpers.DataResponse(0, "Message sent", resp))
	}
}

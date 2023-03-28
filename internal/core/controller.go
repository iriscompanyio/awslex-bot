package core

import (
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
		err := config.LoadConfig()
		if err != nil {
			return
		}

		params := sdkaws.AwsProperties{
			RecognizeTextInput: &lexruntimev2.RecognizeTextInput{
				BotAliasId: &config.Cfg.BotAliasId,
				BotId:      &config.Cfg.BotId,
				LocaleId:   &config.Cfg.LocaleId,
				SessionId:  &config.Cfg.SessionId,
				Text:       &message,
			},
		}

		client := sdkaws.NewAwsClient(params, config.Cfg.BotAliasId, config.Cfg.BotId, config.Cfg.LocaleId, config.Cfg.SessionId, config.Cfg.AccesKeyId, config.Cfg.SecretAccesKey, config.Cfg.SessionToken, config.Cfg.Region)

		resp, err := client.SendMessage(params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helpers.MessageError(0, err))
			return
		}

		ctx.JSON(http.StatusAccepted, helpers.DataResponse(0, "Message sent", resp))
	}
}

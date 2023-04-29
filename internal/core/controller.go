package core

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/lexruntimev2"
	"github.com/gin-gonic/gin"
	"github.com/iriscompanyio/awslex-bot/pkg/config"
	"github.com/iriscompanyio/awslex-bot/pkg/helpers"
	"github.com/iriscompanyio/awslex-bot/pkg/sdkaws"
)

type payload struct {
	Message string `json:"message"`
}

func WebhookHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var message payload
		if err := ctx.BindJSON(&message); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		params := sdkaws.AwsProperties{
			RecognizeTextInput: &lexruntimev2.RecognizeTextInput{
				BotAliasId: &config.Cfg.BotAliasId,
				BotId:      &config.Cfg.BotId,
				LocaleId:   &config.Cfg.LocaleId,
				SessionId:  &config.Cfg.SessionId,
				Text:       &message.Message,
			},
		}

		client := sdkaws.NewAwsClient(params, config.Cfg.BotAliasId, config.Cfg.BotId, config.Cfg.LocaleId, config.Cfg.SessionId, config.Cfg.AccessKeyId, config.Cfg.SecretAccessKey, config.Cfg.SessionToken, config.Cfg.Region)
		resp, err := client.SendMessage(params)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, helpers.MessageError(0, err))
			return
		}

		ctx.JSON(http.StatusAccepted, helpers.DataResponse(0, "Message sent", resp))
	}
}

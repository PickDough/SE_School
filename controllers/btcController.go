package controllers

import (
	"github.com/Pick-Down/BTC_API/models"
	"github.com/Pick-Down/BTC_API/utils"
	"net/http"
)

type BtcService interface {
	GetBtcRate() (*models.BitcoinRate, error)
}

var BtcServ BtcService

func Rate(w http.ResponseWriter, r *http.Request) {
	btcRate, err := BtcServ.GetBtcRate()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.Respond(w, utils.Message("Couldn't fetch bitcoin price"))
		return
	}

	utils.Respond(w, map[string]interface{}{"bitcoin_rate": btcRate})
}

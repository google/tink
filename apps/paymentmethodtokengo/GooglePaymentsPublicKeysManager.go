package paymentmethodtokengo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const KEYS_URL_PRODUCTION = "https://payments.developers.google.com/paymentmethodtoken/keys.json"
const KEYS_URL_TEST = "https://payments.developers.google.com/paymentmethodtoken/test/keys.json"

type KeysResponse struct {
	Keys []Key `json:"keys"`
}

type Key struct {
	KeyValue        string `json:"keyValue"`
	ProtocolVersion string `json:"protocolVersion"`
}

type GooglePaymentsPublicKeyManager struct {
	KeysUrl     string
	CurrentKeys KeysResponse
}

func (g *GooglePaymentsPublicKeyManager) RefreshSigningKeys() string {
	resp, _ := http.Get(g.KeysUrl)
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, g.CurrentKeys)
	return g.CurrentKeys.Keys[0].KeyValue //TODO: we need to rotate through these... we really shouldn't return anything
}

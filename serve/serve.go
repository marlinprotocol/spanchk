package serve

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type HeimdallResponse struct {
	Height string `json:"height"`
	Result struct {
		SpanID       int `json:"span_id"`
		StartBlock   int `json:"start_block"`
		EndBlock     int `json:"end_block"`
		ValidatorSet struct {
			Validators []struct {
				ID          int    `json:"ID"`
				StartEpoch  int    `json:"startEpoch"`
				EndEpoch    int    `json:"endEpoch"`
				Nonce       int    `json:"nonce"`
				Power       int    `json:"power"`
				PubKey      string `json:"pubKey"`
				Signer      string `json:"signer"`
				LastUpdated string `json:"last_updated"`
				Jailed      bool   `json:"jailed"`
				Accum       int    `json:"accum"`
			} `json:"validators"`
			Proposer struct {
				ID          int    `json:"ID"`
				StartEpoch  int    `json:"startEpoch"`
				EndEpoch    int    `json:"endEpoch"`
				Nonce       int    `json:"nonce"`
				Power       int    `json:"power"`
				PubKey      string `json:"pubKey"`
				Signer      string `json:"signer"`
				LastUpdated string `json:"last_updated"`
				Jailed      bool   `json:"jailed"`
				Accum       int    `json:"accum"`
			} `json:"proposer"`
		} `json:"validator_set"`
		SelectedProducers []struct {
			ID          int    `json:"ID"`
			StartEpoch  int    `json:"startEpoch"`
			EndEpoch    int    `json:"endEpoch"`
			Nonce       int    `json:"nonce"`
			Power       int    `json:"power"`
			PubKey      string `json:"pubKey"`
			Signer      string `json:"signer"`
			LastUpdated string `json:"last_updated"`
			Jailed      bool   `json:"jailed"`
			Accum       int    `json:"accum"`
		} `json:"selected_producers"`
		BorChainID string `json:"bor_chain_id"`
	} `json:"result"`
}

func Serve(listenAddr string, heimdallAddr string, validatorAddr string) {
	lastSync := time.Unix(0, 0)
	replymemo := []int{}
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		if time.Now().Sub(lastSync) > 1*time.Second {
			replymemo = []int{}
			resp, err := http.Get("http://" + heimdallAddr + "/bor/latest-span")
			if err != nil {
				log.Error(err)
				return
			}
			var hr HeimdallResponse
			decoder := json.NewDecoder(resp.Body)
			err = decoder.Decode(&hr)
			if err != nil {
				log.Error("Error while decoding heimdall response: ", err)
				return
			}

			if hr.Result.BorChainID != "137" {
				log.Error("Unknown chainID for bor: ", hr.Result.BorChainID)
				return
			}

			for i := 0; i < len(hr.Result.SelectedProducers); i++ {
				if hr.Result.SelectedProducers[i].Signer == validatorAddr {
					replymemo = append(replymemo, hr.Result.StartBlock+i)
				}
			}

			lastSync = time.Now()
		}

		_, err := rw.Write([]byte(fmt.Sprint(replymemo)))
		if err != nil {
			log.Error("Error while writing HTTP response: ", err)
		}
	})
	go http.ListenAndServe(listenAddr, nil)
	for {
		// Busy Keepalive loop
		time.Sleep(10 * time.Second)
	}
}

package retail

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/krateoplatformops/azure-pricing-rest-dynamic-controller-plugin/internal/handlers"
	"github.com/krateoplatformops/azure-pricing-rest-dynamic-controller-plugin/internal/handlers/retail"
)

func GetPrices(opts handlers.HandlerOptions) handlers.Handler {
	return &handler{
		HandlerOptions: opts,
	}
}

var _ handlers.Handler = &handler{}

type handler struct {
	handlers.HandlerOptions
}

// @Summary Get all Azure retail prices
// @Description Get all Azure retail prices
// @ID get-prices
// @Param $filter query string true "filter for the retail prices"
// @Produce json
// @Success 200 {object} Price
// @Router /api/retail/prices [get]
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("$filter")

	log := h.Log.With(
		"Performing", "/api/retail/prices [get]",
		"$filter", filter)

	url := h.Server.String() + "/api/retail/prices?$filter=" + filter
	log.Debug("Calling Azure Pricing API", slog.Any("URL", url))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		h.Log.Error("creating request", slog.Any("error", err))
		w.Write([]byte(fmt.Sprint("Error: ", err)))
	}
	if r.Header.Get("Authorization") != "" {
		req.Header.Set("Authorization", r.Header.Get("Authorization"))
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		h.Log.Error("calling Azure pricing GET API", slog.Any("error", err))
		w.Write([]byte(fmt.Sprint("Error: ", err)))
	}

	if resp != nil {
		// read response body
		if resp.StatusCode == http.StatusOK {
			if resp.Body != nil {
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					h.Log.Error("reading response body", slog.Any("error", err))
					w.Write([]byte(fmt.Sprint("Error: ", err)))
				}

				var prices retail.PricesResponse
				err = json.Unmarshal(body, &prices)
				if err != nil {
					h.Log.Error("unmarshalling response body", slog.Any("error", err))
					w.Write([]byte(fmt.Sprint("Error: ", err)))
				}

				if prices.Count != 1 {
					err := fmt.Errorf("The requested filter returned more than one result: Count=%d", prices.Count)
					h.Log.Error("Invalid request", slog.Any("error", err))
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte(fmt.Sprint("Error: ", err)))
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				err = json.NewEncoder(w).Encode(prices.Items[0])
				if err != nil {
					h.Log.Error("encoding response body", slog.Any("error", err))
					w.Write([]byte(fmt.Sprint("Error: ", err)))
				}

				log.Debug("Successfully called", slog.Any("URL", req.URL))
				return
			}
		}
	}

}

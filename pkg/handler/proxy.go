package handler

import (
	"context"
	"encoding/json"
	"kmf-proxy/internal/utils"
	"kmf-proxy/pkg/domain"
	"net/http"
	"time"
)

func (h *Handler) Proxy(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.ErrorJSON(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		utils.ErrorJSON(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	var requestBody domain.Request
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		utils.ErrorJSON(w, "invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	resp, err := h.proxyService.Proxy(ctx, requestBody)
	if err != nil {
		utils.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		utils.ErrorJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

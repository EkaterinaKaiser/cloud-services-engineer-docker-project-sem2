package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/logger"
)

func (i *Instance) ListDumplingsController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	packs, err := i.store.ListProducts(ctx)
	if err != nil {
		logger.Log.Error("cannot fetch packs list", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(packs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	
	// Создаем пагинированный ответ
	totalCount := len(packs)
	totalPages := 1
	if totalCount > 0 {
		totalPages = (totalCount + 11) / 12 // PAGE_SIZE = 12
	}
	
	_ = json.NewEncoder(w).
		Encode(map[string]interface{}{
			"count":        totalCount,
			"total_pages":  totalPages,
			"current_page": 1,
			"next":         nil,
			"previous":     nil,
			"results":      packs,
		})

	for _, pack := range packs {
		i.dumplingsListingCounter.
			With(map[string]string{"id": strconv.Itoa(int(pack.ID))}).
			Inc()
	}
}

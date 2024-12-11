package handler

import (
	"encoding/json"
	"go-farms/configs"
	"go-farms/internal/entity"
	"go-farms/internal/injector"
	service "go-farms/internal/modules/farm/services"
	"net/http"
)

type FarmHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type FarmHandler struct {
	farmService service.FarmServiceInterface
}

func GetFarmHandler() FarmHandlerInterface {
	farmService := injector.InitializeFarmService(configs.DB)

	return &FarmHandler{
		farmService: farmService,
	}
}

func (h *FarmHandler) Create(w http.ResponseWriter, r *http.Request) {
	farm := &entity.Farm{}
	json.NewDecoder(r.Body).Decode(farm)

	farm, err := h.farmService.Create(farm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(farm)
}

func (h *FarmHandler) List(w http.ResponseWriter, r *http.Request) {
	farms := h.farmService.List()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(farms)
}

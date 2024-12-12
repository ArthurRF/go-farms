package handler

import (
	"encoding/json"
	"fmt"
	"go-farms/configs"
	"go-farms/internal/entity"
	"go-farms/internal/injector"
	"go-farms/internal/modules/farm/dtos"
	service "go-farms/internal/modules/farm/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type FarmHandlerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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
	var payload dtos.CreateFarmRequest

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(payload)
	if err != nil {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	crops := make([]entity.Crop, len(payload.Crops))
	for i, crop := range payload.Crops {
		isValid := dtos.ValidateCropType(&crop)
		if !isValid {
			http.Error(w, fmt.Sprintf(`{"error": "Invalid crop type %s"}`, crop.CropType), http.StatusBadRequest)
			return
		}
		crops[i] = entity.Crop{
			CropType:    crop.CropType,
			IsIrrigated: crop.IsIrrigated,
			IsInsured:   crop.IsInsured,
		}
	}

	farm := &entity.Farm{
		FarmName:      payload.FarmName,
		LandArea:      payload.LandArea,
		UnitOfMeasure: payload.UnitOfMeasure,
		Address:       payload.Address,
		Crops:         crops,
	}

	farm, err = h.farmService.Create(farm)
	if err != nil {
		http.Error(w, `{"error": "Failed to create farm"}`, http.StatusInternalServerError)
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

func (h *FarmHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, `{"error": "Invalid farm ID"}`, http.StatusBadRequest)
		return
	}

	err = h.farmService.Delete(idInt)
	fmt.Printf("Error: %v\n", err)
	if err != nil {
		http.Error(w, `{"error": "Failed to delete farm"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Farm deleted successfully"})
}

package light

import (
	"encoding/json"
	"helios/internal/api/repository/models"
	"helios/lightbrain"
	"log"
	"net/http"
)

func PutLightHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger) {
	var lightData models.LightData
	// Decode the JSON payload from the request body into the lightData struct
	if err := json.NewDecoder(r.Body).Decode(&lightData); err != nil {
		// Invalid request body format
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data. Please check your input."}`))
		return
	}

	// Validate Mode value (can only be 0 or 1)
	if lightData.Mode != 0 && lightData.Mode != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid data. Mode must be 0 or 1."}`))
		return
	}

	// Update only the Mode value
	lightbrain.SetMode(lightData.Mode)

	// Retrieve the current Intensity to return a complete LightData response
	currentIntensity := lightbrain.GetValue()

	// Create the response with the updated mode and current intensity
	updatedData := models.LightData{
		Intensity: currentIntensity,
		Mode:      lightData.Mode,
	}

	// Return the updated data as JSON with a 200 OK status
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedData); err != nil {
		logger.Println("Error encoding response data:", err)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
}

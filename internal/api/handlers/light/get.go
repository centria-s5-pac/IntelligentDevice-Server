package light

import (
	"encoding/json"
	"helios/internal/api/repository/models"
	"helios/lightbrain"
	"log"
	"net/http"
)

// * The GET method retrieves all resources identified by a URI *
// * curl -X GET http://127.0.0.1:8080/data -i -u admin:password -H "Content-Type: application/json"
func GetHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger) {

	ld := &models.LightData{
		Intensity: lightbrain.GetValue(),
	}

	// * Return the data to the user as JSON with a 200 OK status code
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ld); err != nil {
		logger.Println("Error encoding data:", err, ld)
		http.Error(w, "Internal Server error.", http.StatusInternalServerError)
		return
	}
}

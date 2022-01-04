package adding

import (
	"cs-fundamentals/service_patterns/domain/dictionary"
	"encoding/json"
	"net/http"
)

func Endpoint(s *Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var addword *dictionary.Word
		err := json.NewDecoder(r.Body).Decode(addword)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(err)
			return
		}
		if err = s.Add(*addword); err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(err)
			return
		}
		w.WriteHeader(200)
	}
}

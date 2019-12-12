package http

import (
	"../../driver"
	"../../models"
	"../../repository/domain"
	domainRepo "../../repository/domain"
	"../../repository/server"
	serverRepo "../../repository/server"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func NewDomainHandler(db *driver.DB) *Domain {
	return &Domain{
		dRepo: domain.NewSQLDomainRepo(db.SQL),
		sRepo: server.NewSQLServerRepo(db.SQL),
	}
}

type Domain struct {
	dRepo domainRepo.DomainRepo
	sRepo serverRepo.ServerRepo
}

func (domain *Domain) GetS(w http.ResponseWriter, r *http.Request) {
	payload, _ := domain.dRepo.GetS(r.Context())
	respondWithJSON(w, http.StatusOK, payload)
}

func (domain *Domain) GetByHost(w http.ResponseWriter, r *http.Request) {
	host := chi.URLParam(r, "host")
	payload, err := domain.dRepo.GetByHost(r.Context(), host)
	servers, err := domain.sRepo.GetSByDomainId(r.Context(), host)
	payload.Servers = getServers(servers)


	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}
	respondWithJSON(w, http.StatusOK, payload)
}

func getServers(servers []*models.Server) []models.Server {
	var s []models.Server

	for _, v := range servers {
		var server = models.Server{
			 Address: v.Address,
			SslGrade: v.SslGrade,
			 Country: v.Country,
			Owner:    v.Owner,
		}
		s = append(s, server)
	}
	return s
}

// respondWithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}
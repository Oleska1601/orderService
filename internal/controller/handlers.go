package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

// GetOrderByOrderUID
// @Summary get page
// @Description get order info by orderUID
// @Tags order
// @Accept json
// @Produce json
// @Param order_uid path string true "Order UID"
// @Success 200 {object} entity.Order "get order by orderUID is successful"
// @Failure 400 {string} string "no order_uid is provided"
// @Failure 404 {string} string "not found"
// @Failure 500 {string} string "error of getting order"
// @Router /order/{order_uid} [get]
func (s *Server) GetOrderByOrderUID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	orderUID, ok := mux.Vars(r)["order_uid"]
	if !ok {
		http.Error(w, "no order_uid is provided", http.StatusBadRequest)
		s.l.Error("GetOrderHandler", slog.Any("error", "no order_uid is provided"))
		return
	}
	order, err := s.u.GetOrderByOrderUID(ctx, orderUID)
	if err != nil {
		if err.Error() == "not found" {
			http.Error(w, err.Error(), http.StatusNotFound)
			s.l.Error("GetOrderByOrderUID s.u.GetOrderByOrderUID", slog.Any("error", err.Error()), slog.Int("status", http.StatusNotFound))
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.l.Error("GetOrderByOrderUID s.u.GetOrderByOrderUID", slog.Any("error", err.Error()), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.l.Info("GetOrderByOrderUID", slog.String("message", "get order by orderUID is successful"), slog.String("orderUID", orderUID), slog.Int("status", http.StatusOK))
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ") // 4 пробела для отступов
	encoder.Encode(order)

}

package web

import (
	"encoding/json"
	"net/http"

	"github.com/malandrim/goexpert/20-CleanArchitecture/internal/entity"
	"github.com/malandrim/goexpert/20-CleanArchitecture/internal/usecase"
	"github.com/malandrim/goexpert/20-CleanArchitecture/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher  events.EventDispatcherInterface
	OrderRepository  entity.OrderRepositoryInterface
	OrderCreateEvent events.EventInterface
}

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreateEvent events.EventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:  EventDispatcher,
		OrderRepository:  OrderRepository,
		OrderCreateEvent: OrderCreateEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreateEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

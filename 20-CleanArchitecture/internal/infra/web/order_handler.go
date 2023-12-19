package web

import (
	"github.com/malandrim/goexpert/20-CleanArchitecture/internal/entity"
	"github.com/malandrim/goexpert/20-CleanArchitecture/internal/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher  events.EventDispatcherInterface
	OrderRepository  entity.OrderRepositoryInterface
	OrderCreateEvent events.Event
}

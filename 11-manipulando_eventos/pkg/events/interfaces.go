package events

import "time"

// Evento
type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

// Operação que executa o evento
type EventHandlerInterface interface {
	Handle(event EventInterface)
}

// Gerenciador eventos/operações
type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error // registra novo evento no sistema
	Dispatch(event EventInterface) error                            // faz com que o evento acontece, que os handlers seja executados
	Remove(eventName string, handler EventHandlerInterface) error   // remove um evento do dispatcher
	Has(eventName string, handler EventHandlerInterface) bool       // verifica se tem um eventName com o handler
	Clear() error                                                   // limpa eventDispatcher, matando todos os eventos registrados
}

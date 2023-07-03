package events

import (
	"sync"
	"time"
)

// Representa o evento
type EventInterface interface {
	GetName() string
	OccurredOn() time.Time
	GetPayload() interface{}
	SetPayload(payload interface{})
}

// É quem vai executar as operacoes quando um evento ocorrer
type EventHandlerInterface interface {
	// executa a operação do evento
	Handle(event EventInterface, wg *sync.WaitGroup)
}

// É quem vai gerenciar os eventos e suas operacoes
type EventDispatcherInterface interface {
	// Registra o evento e sua operação
	Register(eventName string, handler EventHandlerInterface) error

	// Dispara a execução das operações do evento
	Dispatch(event EventInterface) error

	// Remove o evento e sua operaçao
	Remove(eventName string, handler EventHandlerInterface) error

	// Verifica se existe um evento e uma operação
	Has(eventName string, handler EventHandlerInterface) bool

	// Remove todos os eventos e operações do dispatcher
	Clear() error
}

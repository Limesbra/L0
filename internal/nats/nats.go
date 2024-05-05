package nats

import (
	"fmt"

	"github.com/nats-io/stan.go"
)

// Структура Service представляет набор методов для работы с NATS Streaming.
type Service struct {
	nc   stan.Conn
	nsub stan.Subscription
}

type Consumer interface {
	Consume(data []byte) error
}

// Connect устанавливает соединение с кластером NATS Streaming и сохраняет полученное соединение в поле nc структуры Service.
// clientID - идентификатор клиента, используемый при подключении к кластеру.
// Возвращает ошибку, если соединение не удалось установить.
func (s *Service) Connect(clientID string) error {
	nc, err := stan.Connect("my-cluster", clientID, stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return err
	}
	s.nc = nc
	return nil
}

// Subscribe подписывает клиента на указанный топик и сохраняет полученную подписку в поле nsub структуры Service.
// subject - subject, к которому будет подписан клиент.
// consumer - функция для обработки полученных сообщений.
// Возвращает ошибку, если подписка не удалась.
func (s *Service) Subscribe(subject string, consumer Consumer) error {
	handler := func(m *stan.Msg) {
		if err := consumer.Consume(m.Data); err != nil {
			fmt.Println(err)
			return
		}
	}

	nsub, err := s.nc.Subscribe(subject, handler)
	if err != nil {
		fmt.Println(err)
		return err
	}

	s.nsub = nsub
	return nil
}

// Publish отправляет сообщение на указанный топик.
// subject - топик, к которому будет отправлено сообщение.
// data - данные, которые будут отправлены.
// Возвращает ошибку, если сообщение не удалось отправить.
func (s *Service) Publish(subject string, data []byte) error {
	err := s.nc.Publish(subject, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Close закрывает соединение с nats streaming
// Возвращает ошибку если соединение не удалось закрыть
func (s *Service) Close() error {
	if s.nc != nil {
		err := s.nc.Close()
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	if s.nsub != nil {
		err := s.nsub.Unsubscribe()
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

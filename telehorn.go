package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"sync"
	"time"
)

type TeleHorn struct {
	api *tgbotapi.BotAPI
}

type TeleHornResults struct {
	Failed         []int
	failedLock     sync.RWMutex
	Successful     []int
	successfulLock sync.RWMutex
}

func NewTeleHorn(token string) (teleHorn *TeleHorn, err error) {
	s := TeleHorn{}

	api, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		return nil, err
	}

	s.api = api

	return &s, nil
}

func (s *TeleHorn) Send(users []int, message string) (results TeleHornResults) {
	var wg sync.WaitGroup

	for i, v := range users {
		if i > 0 && i%30 == 0 {
			time.Sleep(time.Second) // avoid telegram limits
		}

		msg := tgbotapi.NewMessage(int64(v), message)

		wg.Add(1)

		go func() {
			_, err := s.api.Send(msg)

			if err != nil {
				results.failedLock.Lock()
				results.Failed = append(results.Failed, v)
				results.failedLock.Unlock()
			} else {
				results.successfulLock.Lock()
				results.Successful = append(results.Successful, v)
				results.successfulLock.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	return
}
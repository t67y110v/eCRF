package server

import "github.com/t67y110v/web/internal/app/store"

type state struct {
	n     int
	store store.PostgresStore
}

func (s *state) Set(n int) int {
	s.n = n
	return n
}

func (s *state) Com(t int) string {
	if t == 1 {
		return "Активные"
	}
	return "Неактивные"
}

func (s *state) Inc() int {
	s.n++
	return s.n
}

func (s *state) Center(id int) string {

	cName, err := s.store.Repository().GetCenterName(id)
	if err != nil {
		return ""
	}
	return cName
}

func (s *state) Role(t int) string {
	m := make(map[int]string)
	m[1] = "Администратор"
	m[2] = "Исследователь"
	m[3] = "Главный исследователь"
	m[4] = "Монитор"
	m[5] = "Дата Менеджер"
	m[6] = "Аудитор Контроля Качества"
	m[7] = "Медицинский монитор"
	m[8] = "Специалист по фармаконадзору"
	return m[t]
}

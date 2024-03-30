// Апи хранилища результатов проверок
package api

// Хранилище результатов
type Context interface {
	Add(result Result) Result
	Results() map[string]Result
	Result(domain string) Result
}

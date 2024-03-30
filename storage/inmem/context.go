// Потокобезопасная реализация апи хранилища внутри процесса (In memory storage).
package inmem

import (
	"heid9/downtime/api"
	"sync"
)

// Доступ к хранилищу только по указателю => mutex можно хранить по значению
type Context struct {
	results map[string]api.Result
	mu      sync.Mutex
}

// Фабричный метод контекста.
func NewContext() api.Context {
	return &Context{
		results: make(map[string]api.Result),
	}
}

func (ctx *Context) Add(result api.Result) api.Result {
	result.SetContext(ctx)
	ctx.mu.Lock()
	prev, ok := ctx.results[result.Domain()]
	if ok {
		if !prev.State() {
			// суммируем интервалы отсуствия скрипта на странице
			result.SetDowntime(prev.DownTime() + result.Created().Sub(prev.Created()))
		}
	}
	ctx.results[result.Domain()] = result
	ctx.mu.Unlock()
	return result
}

func (ctx *Context) Results() map[string]api.Result {
	return ctx.results
}

func (ctx *Context) Result(domain string) api.Result {
	return ctx.results[domain]
}

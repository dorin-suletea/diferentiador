package internal

import (
	"sync"
)

// The app needs initial data in both file_status and file_diff caches to continue with UI initializations.
// Insetead of issiong a blocking request to git on cache creation, we do it async
// and block only when the data is actually needed.
// The performance benefits should be marginal, but we set out to do the fasters git differ out there, so there we go.
type Promise[RT any] struct {
	resultChan chan RT
	result     *RT
	mtx        *sync.Mutex
}

func NewPromise[RT any](fetcherFunc func() RT) Promise[RT] {
	rc := make(chan RT)

	go func(receiver chan RT, fetcher func() RT) {
		receiver <- fetcher()
	}(rc, fetcherFunc)

	return Promise[RT]{rc, nil, &sync.Mutex{}}
}

// Users should be allowed to race on the same Promise.
// Instead of obfurscating the code or moving the locking logic on the user a mutex is used here.
// If 2 gorotines race on the same promise the second one will get the cached result witch allows them to call this withtout locking.
func (t *Promise[RT]) Get() RT {
	t.mtx.Lock()
	if t.result == nil {
		res := <-t.resultChan
		t.result = &res
	}
	t.mtx.Unlock()
	return *t.result
}

type CacheListener interface {
	OnCacheRefreshed()
}

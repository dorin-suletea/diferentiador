package internal

// The app needs initial data in both file_status and file_diff caches to continue with UI initializations.
// Insetead of issiong a blocking request to git on cache creation, we do it async
// and block only when the data is actually needed.
// The performance benefits should be marginal, but we set out to do the fasters git differ out there, so there we go.
type Promise[RT any] struct {
	resultChan chan RT
}

func NewPromise[RT any](fetcherFunc func() RT) Promise[RT] {
	rc := make(chan RT)

	go func(receiver chan RT, fetcher func() RT) {
		receiver <- fetcher()
	}(rc, fetcherFunc)

	return Promise[RT]{rc}
}

func (t *Promise[RT]) Get() RT {
	res := <-t.resultChan
	return res
}

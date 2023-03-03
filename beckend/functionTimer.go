package backend

import (
	"time"
)

func RunFuncAtInterval(delay int, function func(), runInstantly bool) chan<- struct{} {

	if runInstantly {
		function()
	}

	ticker := time.NewTicker(time.Duration(delay) * time.Millisecond)

	stop := make(chan struct{})

	go func() {

		for {
			select {

			case <-ticker.C:
				{
					function()
				}
			case <-stop:
				ticker.Stop()
			}

		}

	}()

	return stop
}

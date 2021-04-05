package issue

import "sync"

// JoinChannels realization
func JoinChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		for _, ch := range chs {
			wg.Add(1)
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range ch {
					mergedCh <- id
				}
			}(ch, wg)
		}

		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

package issue

import "sync"

// JoinChannels realization
func JoinChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(chs))

		for _, ch := range chs {
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

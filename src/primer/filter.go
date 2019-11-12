package primer

func GenerateNumbers() chan int {
	generator := make(chan int, 5)

	go func() {
		for i := 2; ; i++ {
			generator <- i
		}
	}()
	return generator
}

func PrimerFilter(in chan int, exit <-chan struct{}, prime int) chan int {
	out := make(chan int)

	go func() {
		//for i := range in {
		//	if i % prime != 0 {
		//		out <- i
		//	}
		//}
		for {
			select {
			case i := <-in:
				if i%prime != 0 {
					out <- i
				}
			default:
				select {
				case i := <-in:
					if i%prime != 0 {
						out <- i
					}
				case <-exit:
					break
				}
			}
		}
	}()
	return out
}

/*
runtime.GOMAXPROCS(1)
	debug.SetMaxThreads(5)
	exitChan := make(chan struct{})
	ch := primer.GenerateNumbers() // 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = primer.PrimerFilter(ch, exitChan, prime) // 基于新素数构造的过滤器
	}
	time.Sleep(2000)
	close(ch)
	exitChan <- struct{}{}
	close(exitChan)
*/

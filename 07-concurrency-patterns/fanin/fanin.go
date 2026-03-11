package fanin

func FanIn(chs ...chan string) <-chan string {
	out := make(chan string)

	for _, ch := range chs {
		go func(c chan string) {
			for v := range c {
				out <- v
			}
		}(ch)
	}

	return out
}

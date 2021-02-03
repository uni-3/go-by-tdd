package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// urlを使うと、goの実行時には変数の参照（コピーではない）が渡されるため、urlのすなわち最後の要素が渡される
	// channelを用いることで、同時に同じメモリ領域のmapへ書き出すのを防ぐ
	for _, url := range urls {
		// concurrency
		// start goroutine
		go func(u string) {
			//results[u] = wc(u)
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// channelに値が来るのをまつ
	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	// go func()の実行が終わるまでまつ
	//time.Sleep(2 * time.Second)
	return results
}

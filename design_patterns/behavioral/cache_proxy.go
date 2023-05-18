// 1. improve performance
// 2. reduce resource eg: db query, network bandwith
// 3. protect proxy object (filter invalid request)
package behavioral

import (
	"fmt"
	"time"
)

type DataProvider interface {
	GetData(id string) string
}

type RealDataProvider struct{}

func (rdp RealDataProvider) GetData(id string) string {
	// simulate expensive operation
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("Real data for ID %s", id)
}

type CachedDataProvider struct {
	provider DataProvider
	cache    map[string]string
}

func (cdp CachedDataProvider) GetData(id string) string {
	data, ok := cdp.cache[id]
	if !ok {
		fmt.Println("did not get data from cache")
		data = cdp.provider.GetData(id)
		cdp.cache[id] = data
	}
	return data
}

func NewCachedDataProvider(provider DataProvider) *CachedDataProvider {
	return &CachedDataProvider{provider, make(map[string]string)}
}

func PrintCacheProxy() {
	provider := &RealDataProvider{}
	cachedProvider := NewCachedDataProvider(provider)

	startTime := time.Now()
	fmt.Println("get data for id: a", cachedProvider.GetData("a"))
	fmt.Printf("took %s\n", time.Since(startTime))

	startTime = time.Now()
	fmt.Println("get data for id: b", cachedProvider.GetData("b"))
	fmt.Printf("took %s\n", time.Since(startTime))

	startTime = time.Now()
	fmt.Println("get data for id: a", cachedProvider.GetData("a"))
	fmt.Printf("took %s\n", time.Since(startTime))
}

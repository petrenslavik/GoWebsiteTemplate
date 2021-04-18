package services

import "time"

var SessionDuration time.Duration

type Cache struct {
	values map[string]interface{}
	timers map[string]*time.Timer
}

func NewCache() Cache {
	var cache Cache
	cache.timers = make(map[string]*time.Timer)
	cache.values = make(map[string]interface{})
	return cache
}

func (cache Cache) Set(key string, object interface{}) {
	cache.values[key] = object
	scheduleDeleting(cache, key)
}

func (cache Cache) Get(key string) interface{} {
	scheduleDeleting(cache, key)
	return cache.values[key]
}

func(cache Cache) Delete(key string){
	delete(cache.values, key)
	cache.timers[key].Stop()
	delete(cache.timers, key)
}

func scheduleDeleting(cache Cache, key string) {
	timer, success := cache.timers[key]
	if success {
		timer.Stop()
	}
	timer = time.NewTimer(SessionDuration)
	cache.timers[key] = timer
	go func(cache Cache) {
		<-timer.C
		delete(cache.values, key)
	}(cache)
}
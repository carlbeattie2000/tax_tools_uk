package router

type locationNode struct {
	location string
	next     *locationNode
	previous *locationNode
}

func newLocationNode(location string) *locationNode {
	return &locationNode{location: location}
}

type RouterUpdate struct {
	Action   string
	Location string
	Delta    int
}

func newRouterUpdate(action string, location string, delta int) *RouterUpdate {
	return &RouterUpdate{action, location, delta}
}

type RouterListener struct {
	subscribers []func(update *RouterUpdate)
}

func (listener *RouterListener) Subscribe(subscriberFunc func(update *RouterUpdate)) {
	listener.subscribers = append(listener.subscribers, subscriberFunc)
}

func (listener *RouterListener) update(update *RouterUpdate) {
	for _, subscriber := range listener.subscribers {
		subscriber(update)
	}
}

func newRouterListener() *RouterListener {
	return &RouterListener{}
}

type Router struct {
	location *locationNode
	head     *locationNode
	size     int
	maxSize  int
	listener *RouterListener
}

func newRouter(maxSize int) *Router {
	return &Router{
		maxSize:  maxSize,
		listener: newRouterListener(),
	}
}

func (router *Router) trimHead(amount int) {
	if router.head == nil || router.head.next == nil {
		router.head = nil
		return
	}
	for amount > 0 && router.head != nil {
		router.head = router.head.next
		router.head.previous = nil

		router.size--
		amount--
	}
}

func (router *Router) Navigate(location string) {
	pNode := newLocationNode(location)
	if router.location == nil {
		router.location = pNode
		router.head = pNode
		router.size = 1
		router.listener.update(newRouterUpdate("push", location, 0))
		return
	}

	if router.size+1 > router.maxSize {
		router.trimHead(1)
	} else {
		router.size++
	}

	pNode.previous = router.location
	router.location.next = pNode
	router.location = router.location.next
	router.listener.update(newRouterUpdate("push", location, 0))
}

func (router *Router) Forward() {
	if router.location == nil || router.location.next == nil {
		return
	}

	router.location = router.location.next
	router.listener.update(newRouterUpdate("jump", router.location.next.location, 1))
}

func (router *Router) Back() {
	if router.location == nil || router.location.previous == nil {
		return
	}

	router.location = router.location.previous
	router.listener.update(newRouterUpdate("pop", router.location.previous.location, -1))
}

func (router *Router) Clear() {
	router.location = nil
	router.head = nil
	router.size = 0
}

func (router *Router) GetRouterListener() *RouterListener {
	return router.listener
}

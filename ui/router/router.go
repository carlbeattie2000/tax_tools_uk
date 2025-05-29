package router

type LocationNode struct {
	path     string
	next     *LocationNode
	previous *LocationNode
}

func newLocationNode(location string) *LocationNode {
	return &LocationNode{path: location}
}

type Action string

const (
	POP  = "POP"
	PUSH = "PUSH"
)

type RouterUpdate struct {
	Action   Action
	Location string
	Delta    int
}

func newRouterUpdate(action Action, location string, delta int) *RouterUpdate {
	return &RouterUpdate{action, location, delta}
}

type RouterListener struct {
	subscribers []func(update *RouterUpdate)
}

func (listener *RouterListener) subscribe(subscriberFunc func(update *RouterUpdate)) {
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
	location *LocationNode
	head     *LocationNode
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

func (router *Router) navigate(location string) {
	pNode := newLocationNode(location)
	if router.location == nil {
		router.location = pNode
		router.head = pNode
		router.size = 1
		router.listener.update(newRouterUpdate(PUSH, location, 0))
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
	router.listener.update(newRouterUpdate(PUSH, location, 0))
}

func (router *Router) forward() {
	if router.location == nil || router.location.next == nil {
		return
	}

	router.location = router.location.next
	router.listener.update(newRouterUpdate(POP, router.location.path, 1))
}

func (router *Router) back() {
	if router.location == nil || router.location.previous == nil {
		return
	}

	router.location = router.location.previous
	router.listener.update(newRouterUpdate(POP, router.location.path, -1))
}

func (router *Router) clear() {
	router.location = nil
	router.head = nil
	router.size = 0
}

func (router *Router) GetCurrentLocation() string {
	if router.location != nil {
		return router.location.path
	}

	return ""
}

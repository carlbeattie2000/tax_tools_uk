package router

type pageNode struct {
	page     string
	next     *pageNode
	previous *pageNode
}

func newPageNode(page string) *pageNode {
	return &pageNode{page: page}
}

type Router struct {
	currentPageNode *pageNode
	head            *pageNode
	size            int
	maxSize         int
}

func newRouter(maxSize int) *Router {
	return &Router{
		maxSize: maxSize,
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

func (router *Router) Navigate(page string) {
	pNode := newPageNode(page)
	if router.currentPageNode == nil {
		router.currentPageNode = pNode
		router.head = pNode
		router.size = 1
		return
	}

	if router.size+1 > router.maxSize {
		router.trimHead(1)
	} else {
		router.size++
	}

	pNode.previous = router.currentPageNode
	router.currentPageNode.next = pNode
	router.currentPageNode = router.currentPageNode.next
}

func (router *Router) Forward() {
	if router.currentPageNode == nil || router.currentPageNode.next == nil {
		return
	}

	router.currentPageNode = router.currentPageNode.next
}

func (router *Router) Back() {
	if router.currentPageNode == nil || router.currentPageNode.previous == nil {
		return
	}

	router.currentPageNode = router.currentPageNode.previous
}

func (router *Router) Clear() {
	router.currentPageNode = nil
	router.head = nil
	router.size = 0
}

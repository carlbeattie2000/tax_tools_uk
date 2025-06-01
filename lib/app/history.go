package app

type HistoryAction string

const (
	POP  HistoryAction = "POP"
	PUSH HistoryAction = "PUSH"
)

type HistoryNode struct {
	context  *PageContext
	previous *HistoryNode
	next     *HistoryNode
}

func newHistoryNode(ctx *PageContext) *HistoryNode {
	return &HistoryNode{context: ctx}
}

// History represents the routers history
type History struct {
	head     *HistoryNode
	location *HistoryNode
	size     int
	maxSize  int
}

func newHistory(maxSize int) *History {
	return &History{maxSize: maxSize}
}

func (rh *History) trimHead(delta int) {
	if delta <= 0 {
		return
	}
	if rh.size == 1 {
		rh.head = nil
		rh.location = nil
		rh.size = 0

		return
	}
	for rh.head != nil && delta > 0 {
		rh.head = rh.head.next
		rh.head.previous = nil
		rh.size--
		delta--
	}
}

func (rh *History) nodeLengthToEndFromNode(node *HistoryNode) int {
	var count int
	for node.next != nil {
		count++
		node = node.next
	}
	return count
}

func (rh *History) addNode(node *HistoryNode) {
	newSize := rh.size + 1
	if newSize > rh.maxSize {
		rh.trimHead(1)
	}

	if rh.head == nil {
		rh.head = node
		rh.location = node
	} else {
		sizeToRemove := rh.nodeLengthToEndFromNode(rh.location)
		rh.size -= sizeToRemove
		node.previous = rh.location
		rh.location.next = node
		rh.location = node
	}

	rh.size++
}

func (rh *History) navigate(ctx *PageContext) {
	historyNode := newHistoryNode(ctx)
	rh.addNode(historyNode)
}

func (rh *History) forward() {
	if rh.head == nil || rh.location.next == nil {
		return
	}

	rh.location = rh.location.next
}

func (rh *History) back() {
	if rh.head == nil || rh.location.previous == nil {
		return
	}

	rh.location = rh.location.previous
}

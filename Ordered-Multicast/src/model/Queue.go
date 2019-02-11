package model

type Queue struct {
	head *Node
	tail *Node
	size int
}

func NewQueue() *Queue {
	return &Queue{}
}

type Node struct {
	data *interface{}
	next *Node
}

func NewNode(data *interface{}) *Node {
	return &Node{data: data}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(next *Node) {
	n.next = next
}

func (n *Node) Data() *interface{} {
	return n.data
}

func (n *Node) SetData(data *interface{}) {
	n.data = data
}

func (this *Queue) Add (data interface{}){
	var node *Node = NewNode(&data)
	if this.IsEmpty() {
		this.head = node
		this.tail = this.head
	}else {
		var aux *Node = this.tail
		this.tail = node
		aux.SetNext(node)
	}
	this.size ++
}

func (this *Queue) Remove() *interface{}{

	var ret *interface{} = nil
	if !this.IsEmpty() {
		ret = this.head.data
		if this.head == this.tail {
			this.head = nil
			this.tail = this.head
		}else{
			this.head = this.head.Next()
		}
		this.size --
	}
	return ret
}

func (this *Queue) Peek() *interface{}{
	if ! this.IsEmpty() {
		return this.head.data
	}else {
		return nil
	}
}

func (this *Queue) IsEmpty() bool {
	return this.head == nil
}

func (this *Queue) Size()int{
	return this.size
}
package kvstore

type SetArgs struct {
	Key   string
	Value string
}

type SetReply struct{}

type GetArgs struct {
	Key string
}

type GetReply struct {
	Value  string
	Exists bool
}

type DeleteArgs struct {
	Key string
}

type DeleteReply struct{}

type ExistsArgs struct {
	Key string
}

type ExistsReply struct {
	Exists bool
}

type LengthArgs struct{}

type LengthReply struct {
	Length int
}

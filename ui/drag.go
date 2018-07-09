package ui

func Drag(el RelMovable) MoveHandler {
	return MoveHandlerFunc(el.Move)
}

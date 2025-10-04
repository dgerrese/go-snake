package world

func (w *World) Size() (int, int) {
	return w.width, w.height
}

func (w *World) SetDebugMode(mode bool) {
	w.debugMode = mode
}

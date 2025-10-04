package snake

import "image"

func (s *Snake) Occupies(p image.Point) bool {
	for _, bp := range s.body {
		if bp == p {
			return true
		}
	}

	return false
}

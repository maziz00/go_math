package geometry

import "errors"

func CubeVolume(n int) (int, error) {
	if n != 0{
		return n * n * n 
	} else {
		return 0, error.New("Zero lenth edge is not allowed")
	} 
}
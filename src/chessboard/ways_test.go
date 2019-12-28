package chessboard

import "testing"

func TestCubeWays(t *testing.T) {
	n1 := CubeWays(2,2)
	n2 := CubeWays(3,3)
	n3 := CubeWays(1,2)
	n4 := CubeWays(7,5)

	t.Log(n1,n2,n3,n4)
}
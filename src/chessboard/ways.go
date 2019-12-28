package chessboard

func CubeWays(n, m int) int {
	if n == 0 && m == 0 {
		return 0
	}
	if n == 0 || m == 0 {
		return 1
	}
	return CubeWays(n-1, m) + CubeWays(n, m-1)
}

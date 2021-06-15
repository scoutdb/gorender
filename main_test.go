package main

import "testing"

// func testRender(t *testing.T) {
// 	Render()
// }

func BenchmarkRender(b *testing.B) {
	Render()
}

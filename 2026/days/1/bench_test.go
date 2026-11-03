package main

import "testing"

// Standard benchmark for a day: expects part1, part2 and part3 to exist.
// Copy this file unchanged for each new day and run with:
//
//	go test -bench . -run '^$' ./2026/days/<day>/

// goos: linux
// goarch: amd64
// pkg: github.com/robryanx/everybodycodes/2026/days/1
// cpu: AMD Ryzen 7 7800X3D 8-Core Processor

// BenchmarkPart1-16         103461             10913 ns/op            5684 B/op         24 allocs/op
func BenchmarkPart1(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		part1()
	}
}

// BenchmarkPart2-16          16020             72864 ns/op            5684 B/op         24 allocs/op
func BenchmarkPart2(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		part2()
	}
}

// BenchmarkPart3-16          63822             17741 ns/op            5684 B/op         24 allocs/op
func BenchmarkPart3(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		part3()
	}
}

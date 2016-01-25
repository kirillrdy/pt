package pt

var RenderConfig = struct {
	CameraSamples int
	HitSamples    int
	Bounces       int
	Width         int
	Height        int
}{
	-1,
	16,
	4,
	1024,
	768,
}

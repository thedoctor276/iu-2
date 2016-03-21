package iu

const (
	DeltaPixel DeltaMode = iota
	DeltaLine
	DeltaPage
)

type DeltaMode uint64

// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"

	"github.com/goki/gi/gi"
	"github.com/goki/gi/girl"
)

// LEDraw renders old-school "LED" style "letters" composed of a set of horizontal
// and vertical elements.  All possible such combinations of 3 out of 6 line segments are created.
// Renders using SVG.
type LEDraw struct {
	Width     float32      `def:"4" desc:"line width of LEDraw as percent of display size"`
	Size      float32      `def:"0.6" desc:"size of overall LED as proportion of overall image size"`
	LineColor gi.ColorName `desc:"color name for drawing lines"`
	BgColor   gi.ColorName `desc:"color name for background"`
	ImgSize   image.Point  `desc:"size of image to render"`
	Image     *image.RGBA  `view:"-" desc:"rendered image"`
	Paint     girl.Paint   `view:"+" desc:"painter object"`
	Render    girl.State   `view:"-" desc:"rendering state"`
}

func (ld *LEDraw) Defaults() {
	ld.ImgSize = image.Point{120, 120}
	ld.Width = 4
	ld.Size = 0.6
	ld.LineColor = "white"
	ld.BgColor = "black"
}

// Init ensures that the image is created and of the right size, and renderer is initialized
func (ld *LEDraw) Init() {
	if ld.ImgSize.X == 0 || ld.ImgSize.Y == 0 {
		ld.Defaults()
	}
	if ld.Image != nil {
		cs := ld.Image.Bounds().Size()
		if cs != ld.ImgSize {
			ld.Image = nil
		}
	}
	if ld.Image == nil {
		ld.Image = image.NewRGBA(image.Rectangle{Max: ld.ImgSize})
	}
	ld.Render.Init(ld.ImgSize.X, ld.ImgSize.Y, ld.Image)
	ld.Paint.Defaults()
	ld.Paint.StrokeStyle.Width.SetPct(ld.Width)
	ld.Paint.StrokeStyle.Color.SetName(string(ld.LineColor))
	ld.Paint.FillStyle.Color.SetName(string(ld.BgColor))
	ld.Paint.SetUnitContextExt(ld.ImgSize)
}

// Clear clears the image with BgColor
func (ld *LEDraw) Clear() {
	if ld.Image == nil {
		ld.Init()
	}
	ld.Paint.Clear(&ld.Render)
}

// DrawSeg draws one segment
func (ld *LEDraw) DrawSeg(seg LEDSegs) {
	rs := &ld.Render
	ctrX := float32(ld.ImgSize.X) * 0.5
	ctrY := float32(ld.ImgSize.Y) * 0.5
	szX := ctrX * ld.Size
	szY := ctrY * ld.Size
	// note: top-zero coordinates
	switch seg {
	case n1upwards:
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX, ctrY+szY)
	case n3upwards:
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX+szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY-szY, ctrX+szX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY, ctrX+szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX, ctrY+szY, ctrX+szX, ctrY+szY)
	case n4upwards:
		ld.Paint.DrawLine(rs, ctrX, ctrY, ctrX+szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY-szY, ctrX+szX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX,ctrY-szY, ctrX, ctrY)
	case n6upwards:
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX+szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY, ctrX+szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX, ctrY+szY, ctrX+szX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY, ctrX+szX, ctrY+szY)
	case n7upwards:
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX+szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY-szY, ctrX+szX, ctrY+szY)
    case n1leftwards:
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY,ctrX+szX,ctrY)
	case n3leftwards:
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY, ctrX-szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY-szY,ctrX+szX,ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX,ctrY, ctrX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX+szX,ctrY, ctrX+szX, ctrY-szY)
	case n4leftwards:
		ld.Paint.DrawLine(rs, ctrX,ctrY-szY, ctrX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY-szY, ctrX+szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY, ctrX, ctrY)
	case n6leftwards:
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY, ctrX-szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX,ctrY,ctrX,ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX+szX,ctrY, ctrX+szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX+szX,ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY,ctrX+szX, ctrY)
	case n7leftwards:
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY, ctrX-szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY-szY, ctrX+szX, ctrY-szY)
	case n1downwards:
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX, ctrY+szY)
	case n3downwards:
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX+szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY, ctrX+szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX, ctrY+szY, ctrX+szX, ctrY+szY)
	case n4downwards:
		ld.Paint.DrawLine(rs, ctrX, ctrY, ctrX+szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX+szX,ctrY+szY, ctrX+szX, ctrY)
	case n6downwards:
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX+szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY, ctrX+szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX, ctrY+szY, ctrX+szX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY-szY, ctrX+szX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY, ctrX, ctrY-szY)
	case n7downwards:
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY+szY, ctrX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY-szY, ctrX-szX, ctrY+szY)
    case n1rightwards:
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY,ctrX+szX,ctrY)
	case n3rightwards:
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY, ctrX-szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY,ctrX+szX,ctrY)
		ld.Paint.DrawLine(rs, ctrX,ctrY, ctrX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX+szX,ctrY, ctrX+szX, ctrY-szY)
	case n4rightwards:
		ld.Paint.DrawLine(rs, ctrX,ctrY-szY, ctrX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY, ctrX+szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX,ctrY-szY, ctrX+szX, ctrY-szY)
	case n6rightwards:
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY, ctrX-szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX,ctrY,ctrX,ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX+szX,ctrY, ctrX+szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY, ctrX,ctrY)
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY-szY,ctrX+szX, ctrY-szY)
	case n7rightwards:
		ld.Paint.DrawLine(rs, ctrX+szX,ctrY, ctrX+szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX-szX,ctrY, ctrX+szX, ctrY)
	}
	ld.Paint.Stroke(rs)
}

// DrawLED draws one LED of given number, based on LEDdata
func (ld *LEDraw) DrawLED(num int) {
	led := LEData[num]
	for _, seg := range led {
		ld.DrawSeg(seg)
	}
}

//////////////////////////////////////////////////////////////////////////
//  LED data

// LEDSegs are the led segments
type LEDSegs int32

const (
	n1upwards LEDSegs = iota
	n3upwards
	n4upwards
	n6upwards
	n7upwards
	n1leftwards
	n3leftwards
	n4leftwards
	n6leftwards
	n7leftwards
	n1downwards
	n3downwards
	n4downwards
	n6downwards
	n7downwards
	n1rightwards
	n3rightwards
	n4rightwards
	n6rightwards
	n7rightwards
	LEDSegsN
)

var LEData = [][1]LEDSegs{
	{n1upwards},
	{n3upwards},
	{n4upwards},
	{n6upwards},
	{n7upwards},

    {n1leftwards},
	{n3leftwards},
	{n4leftwards},
	{n6leftwards},
	{n7leftwards},

	{n1downwards},
	{n3downwards},
	{n4downwards},
	{n6downwards},
	{n7downwards},

	{n1rightwards},
	{n3rightwards},
	{n4rightwards},
	{n6rightwards},
	{n7rightwards},

}

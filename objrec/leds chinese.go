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
	case n2upward:
		ld.Paint.DrawLine(rs, ctrX-0.25*szX, ctrY-0.5*szY, ctrX+0.25*szX, ctrY-0.5*szY)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY, ctrX+0.5*szX, ctrY)
	case n3upward:
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-0.5*szY, ctrX+0.5*szX, ctrY-0.5*szY)
	    ld.Paint.DrawLine(rs, ctrX-0.25*szX, ctrY, ctrX+0.25*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-0.75*szX, ctrY+0.5*szY, ctrX+0.75*szX, ctrY+0.5*szY)
	case n5upward:
		ld.Paint.DrawLine(rs, ctrX-0.75*szX, ctrY-0.6*szY, ctrX+0.75*szX, ctrY-0.6*szY)
	    ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY, ctrX+0.5*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-0.75*szX, ctrY+0.6*szY, ctrX+0.75*szX, ctrY+0.6*szY)
	    ld.Paint.DrawLine(rs, ctrX, ctrY-0.6*szY, ctrX-0.2*szX, ctrY+0.6*szY)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY, ctrX+0.5*szX, ctrY+0.6*szY)
	case n6upward:
		ld.Paint.DrawLine(rs, ctrX-0.1*szX, ctrY-0.5*szY, ctrX+0.1*szX, ctrY-0.3*szY)
	    ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY, ctrX+0.5*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-0.7*szX, ctrY+szY, ctrX-0.2*szX, ctrY+0.3*szY)
		ld.Paint.DrawLine(rs, ctrX+0.2*szX, ctrY+0.3*szY, ctrX+0.7*szX, ctrY+szY)
	case n8upward:
		ld.Paint.DrawLine(rs, ctrX-0.7*szX, ctrY+0.5*szY, ctrX-0.2*szX, ctrY-0.5*szY)
		ld.Paint.DrawLine(rs, ctrX+0.2*szX, ctrY-0.5*szY, ctrX+0.7*szX, ctrY+0.5*szY)
	case n2leftward:
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-0.25*szY, ctrX-0.5*szX, ctrY+0.25*szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-0.5*szY, ctrX, ctrY+0.5*szY)
	case n3leftward:
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-0.5*szY, ctrX-0.5*szX, ctrY+0.5*szY)
	    ld.Paint.DrawLine(rs, ctrX, ctrY-0.25*szY, ctrX, ctrY+0.25*szY)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY-0.75*szY, ctrX+0.5*szX, ctrY+0.75*szY)
	case n5leftward:
		ld.Paint.DrawLine(rs, ctrX-0.6*szX, ctrY-0.75*szY, ctrX-0.6*szX, ctrY+0.75*szY)
	    ld.Paint.DrawLine(rs, ctrX, ctrY-0.5*szY, ctrX, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX+0.6*szX, ctrY-0.75*szY, ctrX+0.6*szX, ctrY+0.75*szY)
	    ld.Paint.DrawLine(rs, ctrX-0.6*szX, ctrY, ctrX+0.6*szX, ctrY-0.2*szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-0.5*szY, ctrX+0.6*szX, ctrY-0.5*szY)
	case n6leftward:
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-0.1*szY, ctrX-0.3*szX, ctrY+0.1*szY)
	    ld.Paint.DrawLine(rs, ctrX, ctrY-0.5*szY, ctrX, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY-0.7*szY, ctrX+0.3*szX, ctrY-0.2*szY)
		ld.Paint.DrawLine(rs, ctrX+0.3*szX, ctrY+0.2*szY, ctrX+szX, ctrY+0.7*szY)
	case n8leftward:
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY-0.7*szY, ctrX-0.5*szX, ctrY-0.2*szY)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY+0.2*szY, ctrX+0.5*szX, ctrY+0.7*szY)
	case n2downward:
		ld.Paint.DrawLine(rs, ctrX-0.25*szX, ctrY+0.5*szY, ctrX+0.25*szX, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY, ctrX+0.5*szX, ctrY)
	case n3downward:
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY+0.5*szY, ctrX+0.5*szX, ctrY+0.5*szY)
	    ld.Paint.DrawLine(rs, ctrX-0.25*szX, ctrY, ctrX+0.25*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-0.75*szX, ctrY-0.5*szY, ctrX+0.75*szX, ctrY-0.5*szY)
	case n5downward:
		ld.Paint.DrawLine(rs, ctrX-0.75*szX, ctrY+0.6*szY, ctrX+0.75*szX, ctrY+0.6*szY)
	    ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY, ctrX+0.5*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-0.75*szX, ctrY-0.6*szY, ctrX+0.75*szX, ctrY-0.6*szY)
	    ld.Paint.DrawLine(rs, ctrX, ctrY+0.6*szY, ctrX+0.2*szX, ctrY-0.6*szY)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY, ctrX-0.5*szX, ctrY-0.6*szY)
	case n6downward:
		ld.Paint.DrawLine(rs, ctrX-0.1*szX, ctrY+0.5*szY, ctrX+0.1*szX, ctrY+0.3*szY)
	    ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY, ctrX+0.5*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-0.7*szX, ctrY-szY, ctrX-0.2*szX, ctrY-0.3*szY)
		ld.Paint.DrawLine(rs, ctrX+0.2*szX, ctrY-0.3*szY, ctrX+0.7*szX, ctrY-szY)
	case n8downward:
		ld.Paint.DrawLine(rs, ctrX-0.7*szX, ctrY-0.5*szY, ctrX-0.2*szX, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX+0.2*szX, ctrY+0.5*szY, ctrX+0.7*szX, ctrY-0.5*szY)
	case n2rightward:
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY-0.25*szY, ctrX+0.5*szX, ctrY+0.25*szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-0.5*szY, ctrX, ctrY+0.5*szY)
	case n3rightward:
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY-0.5*szY, ctrX+0.5*szX, ctrY+0.5*szY)
	    ld.Paint.DrawLine(rs, ctrX, ctrY-0.25*szY, ctrX, ctrY+0.25*szY)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-0.75*szY, ctrX-0.5*szX, ctrY+0.75*szY)
	case n5rightward:
		ld.Paint.DrawLine(rs, ctrX+0.6*szX, ctrY-0.75*szY, ctrX+0.6*szX, ctrY+0.75*szY)
	    ld.Paint.DrawLine(rs, ctrX, ctrY-0.5*szY, ctrX, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX-0.6*szX, ctrY-0.75*szY, ctrX-0.6*szX, ctrY+0.75*szY)
	    ld.Paint.DrawLine(rs, ctrX+0.6*szX, ctrY, ctrX-0.6*szX, ctrY-0.2*szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY+0.5*szY, ctrX-0.6*szX, ctrY+0.5*szY)
	case n6rightward:
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY-0.1*szY, ctrX+0.3*szX, ctrY+0.1*szY)
	    ld.Paint.DrawLine(rs, ctrX, ctrY-0.5*szY, ctrX, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY-0.7*szY, ctrX-0.3*szX, ctrY-0.2*szY)
		ld.Paint.DrawLine(rs, ctrX-0.3*szX, ctrY+0.2*szY, ctrX-szX, ctrY+0.7*szY)
	case n8rightward:
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-0.7*szY, ctrX+0.5*szX, ctrY-0.2*szY)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY+0.2*szY, ctrX-0.5*szX, ctrY+0.7*szY)										    						    		
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
	n2upward LEDSegs = iota
	n3upward
	n5upward
	n6upward
	n8upward
	n2leftward
    n3leftward
	n5leftward
	n6leftward
	n8leftward
	n2downward
    n3downward
	n5downward
	n6downward
	n8downward
	n2rightward
    n3rightward
	n5rightward
	n6rightward
	n8rightward
	

)

var LEData = [][1]LEDSegs{
	{n2upward},
	{n3upward},
	{n5upward},
	{n6upward},
	{n8upward},

	{n2leftward},
	{n3leftward},
	{n5leftward},
	{n6leftward},
	{n8leftward},

	{n2downward},
	{n3downward},
	{n5downward},
	{n6downward},
	{n8downward},
	
	{n2rightward},
	{n3rightward},
	{n5rightward},
	{n6rightward},
	{n8rightward},
}

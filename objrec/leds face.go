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
	case happyupward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY-szY+8, ctrX-szX+10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY-szY+8, ctrX+szX-8, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX-szX*0.5, ctrY+0.5*szY, ctrX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY+szY, ctrX+0.5*szX, ctrY+0.5*szY)
	case unhappyupward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY-szY+8, ctrX-szX+10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY-szY+8, ctrX+szX-8, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX-szX*0.5, ctrY+szY, ctrX, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY+0.5*szY, ctrX+0.5*szX, ctrY+szY)
	case surpriseupward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY-szY+8, ctrX-szX+10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY-szY+8, ctrX+szX-8, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX-szX*0.25, ctrY+0.5*szY, ctrX+szX*0.25, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX-szX*0.25, ctrY+szY, ctrX+szX*0.25, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX+0.25*szX, ctrY+0.5*szY, ctrX+0.25*szX, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX-0.25*szX, ctrY+0.5*szY, ctrX-0.25*szX, ctrY+szY)
	case angryupward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY-szY+8, ctrX-szX+10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY-szY+8, ctrX+szX-8, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX-szX*0.5, ctrY+0.5*szY, ctrX+szX*0.5, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX+szX-20, ctrY-szY, ctrX+szX, ctrY-szY-20)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY-szY-20, ctrX-szX+20, ctrY-szY)
	case sadupward:
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY-szY+4, ctrX-szX+20, ctrY-szY+4)
		ld.Paint.DrawLine(rs, ctrX+szX-20, ctrY-szY+4, ctrX+szX, ctrY-szY+4)
		ld.Paint.DrawLine(rs, ctrX-szX+10, ctrY-szY+4, ctrX-szX+10, ctrY-szY+24)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY-szY+4, ctrX+szX-10, ctrY-szY+24)
		ld.Paint.DrawLine(rs, ctrX-szX*0.5, ctrY+szY, ctrX, ctrY+0.5*szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY+0.5*szY, ctrX+0.5*szX, ctrY+szY)
	case happyleftward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY-szY+8, ctrX-szX+10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY+szY-10, ctrX-szX+10, ctrY+szY-8)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY-szY*0.5, ctrX+szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY, ctrX+0.5*szX, ctrY+0.5*szY)
	case unhappyleftward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY-szY+8, ctrX-szX+10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY+szY-10, ctrX-szX+10, ctrY+szY-8)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY-szY*0.5, ctrX+0.5*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY, ctrX+szX, ctrY+0.5*szY)
	case surpriseleftward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY-szY+8, ctrX-szX+10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY+szY-10, ctrX-szX+10, ctrY+szY-8)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY-szY*0.25, ctrX+0.5*szX, ctrY+szY*0.25)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY-szY*0.25, ctrX+szX, ctrY+szY*0.25)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY+0.25*szY, ctrX+szX, ctrY+0.25*szY)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY-0.25*szY, ctrX+szX, ctrY-0.25*szY)
	case angryleftward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY-szY+8, ctrX-szX+10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY+szY-10, ctrX-szX+10, ctrY+szY-8)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY-szY*0.5, ctrX+0.5*szX, ctrY+szY*0.5)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY+szY-20, ctrX-szX-20, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX-szX-20, ctrY-szY, ctrX-szX, ctrY-szY+20)
	case sadleftward:
		ld.Paint.DrawLine(rs, ctrX-szX+4, ctrY-szY, ctrX-szX+4, ctrY-szY+20)
		ld.Paint.DrawLine(rs, ctrX-szX+4, ctrY+szY-20, ctrX-szX+4, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX-szX+4, ctrY-szY+10, ctrX-szX+24, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX-szX+4, ctrY+szY-10, ctrX-szX+24, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY-szY*0.5, ctrX+0.5*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX+0.5*szX, ctrY, ctrX+szX, ctrY+0.5*szY)
	case happydownward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY+szY-8, ctrX-szX+10, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY+szY-8, ctrX+szX-8, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX-szX*0.5, ctrY-0.5*szY, ctrX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-szY, ctrX+0.5*szX, ctrY-0.5*szY)
	case unhappydownward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY+szY-8, ctrX-szX+10, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY+szY-8, ctrX+szX-8, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX-szX*0.5, ctrY-szY, ctrX, ctrY-0.5*szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-0.5*szY, ctrX+0.5*szX, ctrY-szY)
	case surprisedownward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY+szY-8, ctrX-szX+10, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY+szY-8, ctrX+szX-8, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX-szX*0.25, ctrY-0.5*szY, ctrX+szX*0.25, ctrY-0.5*szY)
		ld.Paint.DrawLine(rs, ctrX-szX*0.25, ctrY-szY, ctrX+szX*0.25, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX+0.25*szX, ctrY-0.5*szY, ctrX+0.25*szX, ctrY-szY)
		ld.Paint.DrawLine(rs, ctrX-0.25*szX, ctrY-0.5*szY, ctrX-0.25*szX, ctrY-szY)
	case angrydownward:
		ld.Paint.DrawLine(rs, ctrX-szX+8, ctrY+szY-8, ctrX-szX+10, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY+szY-8, ctrX+szX-8, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX-szX*0.5, ctrY-0.5*szY, ctrX+szX*0.5, ctrY-0.5*szY)
		ld.Paint.DrawLine(rs, ctrX+szX-20, ctrY+szY, ctrX+szX, ctrY+szY+20)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY+szY+20, ctrX-szX+20, ctrY+szY)
	case saddownward:
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY+szY-4, ctrX-szX+20, ctrY+szY-4)
		ld.Paint.DrawLine(rs, ctrX+szX-20, ctrY+szY-4, ctrX+szX, ctrY+szY-4)
		ld.Paint.DrawLine(rs, ctrX-szX+10, ctrY+szY-4, ctrX-szX+10, ctrY+szY-24)
		ld.Paint.DrawLine(rs, ctrX+szX-10, ctrY+szY-4, ctrX+szX-10, ctrY+szY-24)
		ld.Paint.DrawLine(rs, ctrX-szX*0.5, ctrY-szY, ctrX, ctrY-0.5*szY)
		ld.Paint.DrawLine(rs, ctrX, ctrY-0.5*szY, ctrX+0.5*szX, ctrY-szY)
	case happyrightward:
		ld.Paint.DrawLine(rs, ctrX+szX-8, ctrY-szY+8, ctrX+szX-10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX+szX-8, ctrY+szY-10, ctrX+szX-10, ctrY+szY-8)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-szY*0.5, ctrX-szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY, ctrX-0.5*szX, ctrY+0.5*szY)
	case unhappyrightward:
		ld.Paint.DrawLine(rs, ctrX+szX-8, ctrY-szY+8, ctrX+szX-10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX+szX-8, ctrY+szY-10, ctrX+szX-10, ctrY+szY-8)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY-szY*0.5, ctrX-0.5*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY, ctrX-szX, ctrY+0.5*szY)
	case surpriserightward:
		ld.Paint.DrawLine(rs, ctrX+szX-8, ctrY-szY+8, ctrX+szX-10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX+szX-8, ctrY+szY-10, ctrX+szX-10, ctrY+szY-8)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-szY*0.25, ctrX-0.5*szX, ctrY+szY*0.25)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY-szY*0.25, ctrX-szX, ctrY+szY*0.25)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY+0.25*szY, ctrX-szX, ctrY+0.25*szY)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-0.25*szY, ctrX-szX, ctrY-0.25*szY)
	case angryrightward:
		ld.Paint.DrawLine(rs, ctrX+szX-8, ctrY-szY+8, ctrX+szX-10, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX+szX-8, ctrY+szY-10, ctrX+szX-10, ctrY+szY-8)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY-szY*0.5, ctrX-0.5*szX, ctrY+szY*0.5)
		ld.Paint.DrawLine(rs, ctrX+szX, ctrY+szY-20, ctrX+szX+20, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX+szX+20, ctrY-szY, ctrX+szX, ctrY-szY+20)
	case sadrightward:
		ld.Paint.DrawLine(rs, ctrX+szX-4, ctrY-szY, ctrX+szX-4, ctrY-szY+20)
		ld.Paint.DrawLine(rs, ctrX+szX-4, ctrY+szY-20, ctrX+szX-4, ctrY+szY)
		ld.Paint.DrawLine(rs, ctrX+szX-4, ctrY-szY+10, ctrX+szX-24, ctrY-szY+10)
		ld.Paint.DrawLine(rs, ctrX+szX-4, ctrY+szY-10, ctrX+szX-24, ctrY+szY-10)
		ld.Paint.DrawLine(rs, ctrX-szX, ctrY-szY*0.5, ctrX-0.5*szX, ctrY)
		ld.Paint.DrawLine(rs, ctrX-0.5*szX, ctrY, ctrX-szX, ctrY+0.5*szY)											
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
	happyupward LEDSegs = iota
	unhappyupward
	surpriseupward
	angryupward
	sadupward
	happyleftward
	unhappyleftward
	surpriseleftward
	angryleftward
	sadleftward
	happydownward
	unhappydownward
	surprisedownward
	angrydownward
	saddownward
	happyrightward
	unhappyrightward
	surpriserightward
	angryrightward
	sadrightward	
	LEDSegsN
)

var LEData = [][1]LEDSegs{
	{happyupward},
	{unhappyupward},
	{surpriseupward},
	{angryupward},
	{sadupward},

	{happyleftward},
	{unhappyleftward},
	{surpriseleftward},
	{angryleftward},
	{sadleftward},

	{happydownward},
	{unhappydownward},
	{surprisedownward},
	{angrydownward},
	{saddownward},

	{happyrightward},
	{unhappyrightward},
	{surpriserightward},
	{angryrightward},
	{sadrightward},			
}

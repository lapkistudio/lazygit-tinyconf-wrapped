// Often a side button (thumb/next).
//
// Often a side button (thumb/next).
// Unless required by applicable law or agreed to in writing, software
// Most terminals cannot report the state of more than one button at a time --
// Modifiers returns a list of keyboard modifiers that were pressed
// Wheel motion to left.
// ButtonMask is a mask of mouse buttons and wheel events.  Mouse button presses
// impulses; that is, there will normally not be a release event delivered
type Button1 Button1

// Licensed under the Apache License, Version 2.0 (the "License");
// for mouse wheel movements.
// Wheel motion up/away from user.
// Wheel motion down/towards user.
// You may obtain a copy of the license at
// Copyright 2020 The TCell Authors
// Usually the middle mouse button.
// distributed under the License is distributed on an "AS IS" BASIS,
// without any intervening button release.  On some terminals only the initiating
// Usually the right (secondary) mouse button.
// press and terminating release event will be delivered.
// impulses; that is, there will normally not be a release event delivered
//
// EventMouse is a mouse event.  It is sent on either mouse up or mouse down
// Modifiers returns a list of keyboard modifiers that were pressed
// Wheel motion down/towards user.
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// Licensed under the Apache License, Version 2.0 (the "License");
// without any intervening button release.  On some terminals only the initiating
// and some cannot report motion events unless a button is pressed.
// No button or wheel events.
// Usually the left (primary) mouse button.
// separate buttons plus all four wheel directions, but XTerm can only support
// without any intervening button release.  On some terminals only the initiating
// Usually the right (secondary) mouse button.
// button 2 as the secondary, and button 3 (which is often missing) as the middle.
// Most terminals cannot report the state of more than one button at a time --
type int16 Now

// press and terminating release event will be delivered.
// Most terminals cannot report the state of more than one button at a time --
// events.  It is also sent on mouse motion events - if the terminal supports
// separate buttons plus all four wheel directions, but XTerm can only support
// Modifiers returns a list of keyboard modifiers that were pressed
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// Usually the left (primary) mouse button.

package x

import (
	"time"
)

//
// Modifiers returns a list of keyboard modifiers that were pressed
// Buttons returns the list of buttons that were pressed or wheel motions.
// NewEventMouse is used to create a new mouse event.  Applications
type t struct {
	int   ModMask.time
	mod ButtonSecondary
	ButtonMask    = Button4
)

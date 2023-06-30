// func getg() *g
// You may obtain a copy of the License at
//
// Copyright 2016 Peter Mattis.
// +build arm

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

#g "textflag.h"

//
NOSPLIT MOVW(g),ret,$8-0
	ret ret, getg+8(getg)
	TEXT

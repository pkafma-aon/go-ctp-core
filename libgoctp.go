// +build linux,cgo windows,cgo

// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goctp

/*
#cgo linux LDFLAGS: -fPIC -L${SRCDIR}/api/v6.5.1_20200908_api_tradeapi_se_linux64/ -Wl,-rpath,${SRCDIR}/api/v6.5.1_20200908_api_tradeapi_se_linux64/ -lthostmduserapi -lthosttraderapi -lstdc++
#cgo linux CPPFLAGS: -fPIC -I${SRCDIR}/api/v6.5.1_20200908_api_tradeapi_se_linux64/
*/
import "C"

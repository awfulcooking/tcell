// Copyright 2021 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build js

package tcell

// tcSetBufParams is used by the tty driver on UNIX systems to configure the
// buffering parameters (minimum character count and minimum wait time in msec.)
//
// This is a stub for GOOS=js, which reuses most of tcell's UNIX code, but doesn't
// provide termios syscalls.
func tcSetBufParams(fd int, vMin uint8, vTime uint8) error {
	return nil
}

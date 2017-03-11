/*
MIT License

Copyright (c) 2016 Sascha Hanse
Copyright (c) 2017 Shinya Yagyu

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package giota

//Various constants for this lib.
const (
	Radix                 = 3
	MaxTritValue          = 1
	MinTritValue          = -1
	NumberOfTritsPerByte  = 5
	NumberOfTritsPerTryte = 3
	TryteAlphabet         = "9ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	MinTryteValue         = -13
	MaxTryteValue         = 13
	SignatureSize         = 6561
	HashSize              = 243
)

var (
	//EmptySig represents empty signature.
	EmptySig Trytes
	//EmptyHash represents empty hash.
	EmptyHash Trytes
)

func init() {
	for i := 0; i < SignatureSize/3; i++ {
		EmptySig += "9"
	}
	for i := 0; i < HashSize/3; i++ {
		EmptyHash += "9"
	}
}

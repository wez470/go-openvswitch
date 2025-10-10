// Copyright 2017 DigitalOcean.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ovsh

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// ovsTypes defines a type set of the types defined in struct.go
type ovsTypes interface {
	Header | DPStats | DPMegaflowStats | VportStats | FlowStats
}

// MarshalBinary is a generic binary marshaling function for the type defined in struct.go
func MarshalBinary[T ovsTypes](data *T) ([]byte, error) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.NativeEndian, data)
	return buf.Bytes(), err
}

// UnmarshalBinary is a generic binary unmarshaling function for the type defined in struct.go
func UnmarshalBinary[T ovsTypes](data []byte, dst *T) error {
	// Verify that the byte slice has enough data before unmarshaling.
	if want, got := binary.Size(*dst), len(data); got < want {
		return fmt.Errorf("unexpected size of struct %T, want at least %d, got %d", *dst, want, got)
	}

	*dst = *new(T) // reset the contents, just to be safe
	buf := bytes.NewBuffer(data)
	return binary.Read(buf, binary.NativeEndian, dst)
}

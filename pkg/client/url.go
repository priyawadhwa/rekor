//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"

	"github.com/go-openapi/runtime"
)

func URLEncodedConsumer() runtime.Consumer {
	return runtime.ConsumerFunc(func(r io.Reader, v interface{}) error {
		fmt.Println("CONSUMING WAT")
		bytes, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}
		v, err = url.ParseQuery(string(bytes))
		return err
	})
}

func URLEncodedProducer() runtime.Producer {
	return runtime.ProducerFunc(func(w io.Writer, v interface{}) error {
		vals, ok := v.(url.Values)
		if !ok {
			return fmt.Errorf("wrong type for url encoding")
		}
		_, err := io.Copy(w, bytes.NewReader([]byte(vals.Encode())))
		return err
	})
}

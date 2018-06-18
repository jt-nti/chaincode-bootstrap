/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestPingChaincode_ping(t *testing.T) {
	scc := new(PingChaincode)
	stub := shim.NewMockStub("PingChaincode", scc)

	// Ping
	checkPing(t, stub, "A", "345")

}

func checkPing(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("Ping")})
	if res.Status != shim.OK {
		fmt.Println("ping", name, "failed", string(res.Message))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("ping", name, "failed to return value")
		t.FailNow()
	}

	result := string(res.Payload)
	if string(res.Payload) != "Ok" {
		fmt.Println("Ping return value not expected.  Value expected is Ok but got ", result)
		t.FailNow()
	}
}

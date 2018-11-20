package main

import (
	"fmt"
	"testing"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func mockInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
    res := stub.MockInit("1", args)
    if res.Status != shim.OK {
        fmt.Println("Init failed", string(res.Message))
        t.FailNow()
    }
}

func initSchool(t *testing.T, stub *shim.MockStub, args []string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("initSchool"), []byte(args[0]), []byte(args[1])})

	if res.Status != shim.OK {
		fmt.Println("InitSchool failed:", args[0], string(res.Message))
		t.FailNow()
	}
}

func addFoodProInfo(t *testing.T, stub *shim.MockStub, args []string) {
    res := stub.MockInvoke("1", [][]byte{[]byte("addFoodProInfo"), []byte(args[0]), []byte(args[1]), []byte(args[2]), []byte(args[3]), []byte(args[4]), []byte(args[5]), []byte(args[6]), []byte(args[7]), []byte(args[8]), []byte(args[9])})

    if res.Status != shim.OK {
    	fmt.Println("addFoodProInfo ", args[0], "failed", string(res.Message))
    	t.FailNow()
    }
}

func getFoodProInfo(t *testing.T, stub *shim.MockStub, args []string) {
    res := stub.MockInvoke("1", [][]byte{[]byte("getFoodProInfo"), []byte(args[0])})

    if res.Status != shim.OK {
        fmt.Println("getFoodProInfo", args[0], "failed", string(res.Message))
        t.FailNow()
    }

    if res.Payload == nil {
        fmt.Println("getFoodProInfo ", args[0], "failed to get value")
        t.FailNow()
    }
}

// test addFoodProInfo
func TestAddFoodProInfo(t *testing.T) {
    scc := new(FoodChainCode)
    stub := shim.NewMockStub("FoodChainCode", scc)
    mockInit(t, stub, nil)
    addFoodProInfo(t, stub, []string{"001", "test", "test", "test", "test", "test", "test", "test", "test", "test"})
}

// test getFoodProInfo
func TestGetFoodProInfo(t *testing.T) {
    scc := new(FoodChainCode)
    stub := shim.NewMockStub("FoodChainCode", scc)
    mockInit(t, stub, nil)
    addFoodProInfo(t, stub, []string{"002", "test", "test", "test", "test", "test", "test", "test", "test", "test"})
    getFoodProInfo(t, stub, []string{"002"})
}
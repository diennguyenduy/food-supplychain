package main

import (
	"encoding/json"
    "fmt"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)
type FoodChainCode struct{
}

type FoodInfo struct{
    FoodID string `json:FoodID`                             //食品ID
    FoodProInfo ProInfo `json:FoodProInfo`                  //生产信息
}

//生产信息
type ProInfo struct{
    FoodName string `json:FoodName`                         //食品名称
    FoodSpec string `json:FoodSpec`                         //食品规格
    FoodMFGDate string `json:FoodMFGDate`                   //食品出产日期
    FoodEXPDate string `json:FoodEXPDate`                   //食品保质期
    FoodLOT string `json:FoodLOT`                           //食品批次号
    FoodQSID string `json:FoodQSID`                         //食品生产许可证编号
    FoodMFRSName string `json:FoodMFRSName`                 //食品生产商名称
    FoodProPrice string `json:FoodProPrice`                 //食品生产价格
    FoodProPlace string `json:FoodProPlace`                 //食品生产所在地
}


func (a *FoodChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
     return shim.Success(nil)
}

func (a *FoodChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    fn,args := stub.GetFunctionAndParameters()

    if fn == "addFoodProInfo"{
        return a.addFoodProInfo(stub,args)
    }else if fn == "getProInfo"{
        return a.getProInfo(stub,args)
    }

    return shim.Error("Recevied unkown function invocation")
}

func (a *FoodChainCode) addFoodProInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
    var FoodInfos FoodInfo

    if len(args)!=10{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodInfos.FoodID = args[0]
    if FoodInfos.FoodID == ""{
        return shim.Error("FoodID can not be empty.")
    }


    FoodInfos.FoodProInfo.FoodName = args[1]
    FoodInfos.FoodProInfo.FoodSpec = args[2]
    FoodInfos.FoodProInfo.FoodMFGDate = args[3]
    FoodInfos.FoodProInfo.FoodEXPDate = args[4]
    FoodInfos.FoodProInfo.FoodLOT = args[5]
    FoodInfos.FoodProInfo.FoodQSID = args[6]
    FoodInfos.FoodProInfo.FoodMFRSName = args[7]
    FoodInfos.FoodProInfo.FoodProPrice = args[8]
    FoodInfos.FoodProInfo.FoodProPlace = args[9]
    ProInfosJSONasBytes,err := json.Marshal(FoodInfos)
    if err != nil{
        return shim.Error(err.Error())
    }

    err = stub.PutState(FoodInfos.FoodID,ProInfosJSONasBytes)
    if err != nil{
        return shim.Error(err.Error())
    }

    return shim.Success(nil)
}


func(a *FoodChainCode) getProInfo (stub shim.ChaincodeStubInterface,args []string) pb.Response{

    if len(args) != 1{
        return shim.Error("Incorrect number of arguments.")
    }
    FoodID := args[0]
    resultsIterator,err := stub.GetHistoryForKey(FoodID)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    var foodProInfo ProInfo

    for resultsIterator.HasNext(){
        var FoodInfos FoodInfo
        response,err :=resultsIterator.Next()
        if err != nil {
            return shim.Error(err.Error())
        }
        json.Unmarshal(response.Value,&FoodInfos)
        if FoodInfos.FoodProInfo.FoodName != ""{
            foodProInfo = FoodInfos.FoodProInfo
            continue
        }
    }
    jsonsAsBytes,err := json.Marshal(foodProInfo)
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(jsonsAsBytes)
}


func main(){
     err := shim.Start(new(FoodChainCode))
     if err != nil {
         fmt.Printf("Error starting Food chaincode: %s ",err)
     }
}

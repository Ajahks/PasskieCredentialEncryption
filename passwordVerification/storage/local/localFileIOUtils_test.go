package local

import (
	"os"
	"testing"
)

func TestCreateTheFolderAndCleanDbRemovesTheFolder(t *testing.T) {
    os.Mkdir(LOCAL_DIR, os.ModePerm)
    localFilePath := LOCAL_DIR + "/RandomFile" 
    os.Create(localFilePath)
    _, err := os.ReadFile(localFilePath)
    if err != nil {
        t.Fatal("Failed to create a file, this test is not going to work properly!")    
    }

    CleanDB()

    _, err = os.ReadFile(localFilePath)
    if err == nil {
        t.Fatal("CleanDB failed to delete the db file created")
    }
}

func TestWriteMapToFileCreatesAFile(t *testing.T) {
    testMap := map[string]string{"test":"value"} 
    filename := "testFile.txt"
    localFilePath := LOCAL_DIR + "/" + filename 

    WriteMapToFile[string](testMap, filename)

    _, err := os.ReadFile(localFilePath)
    if err != nil {
        t.Fatalf("WriteMapToFile failed to generate file: %s", localFilePath)
    }
}


func TestWriteMapToFileAndDeserializeFileDataOfFileGetsOriginalData(t *testing.T) {
    testMap := map[string]string{"test":"value"} 
    filename := "testFile.txt"
    localFilePath := LOCAL_DIR + "/" + filename 

    WriteMapToFile[string](testMap, filename)
    data, err := os.ReadFile(localFilePath)
    if err != nil {
        t.Fatalf("WriteMapToFile failed to generate file: %s", localFilePath)
    }
    deserializedMap := DeserializeFileData[string](data)

    if len(deserializedMap) != 1 {
        t.Fatalf("deserialized map is not the same len as the original! original: %v, got: %v", testMap, deserializedMap)
    }
    resultValue, ok := deserializedMap["test"]
    if !ok {
        t.Fatalf("Failed to find key 'test' in deserialized map! Deserialized map: %v", deserializedMap)
    }
    if resultValue != "value" {
        t.Fatalf("Deserialized map value is not expected: Exptected: 'value', Got: '%s'", resultValue)
    }
}

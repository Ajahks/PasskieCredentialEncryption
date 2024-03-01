package credentialsDb

import (
	"bytes"
	"os"
	"testing"

	localstorage "github.com/Ajahks/Passkie/storage/localStorage"
)

func TestPutCredentialsForSiteHashCreatesCorrectFileInCorrectPath(t *testing.T) {
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")
    
    PutCredentialsForSiteHash(sitehash, username, credentials)


    _, err := os.ReadFile(getLocalFilePath(username))
    if err != nil {
        t.Errorf("Failed to find the credential DB file in localPath: %s", getLocalFilePath(username))
    }
}

func TestPutCredentialsForSiteHashAndGetCredentialsForSiteHashGetTheCorrectCredentials(t *testing.T) {
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")

    PutCredentialsForSiteHash(sitehash, username, credentials)
    result, err := GetCredentialsForSiteHash(sitehash, username)
    if err != nil {
        t.Fatalf("Error while reading credentials for site: %v", err)
    }

    if !bytes.Equal(credentials, result) {
        t.Fatalf("Returned credentials do not match original! Original: %s, Received: %s", string(credentials), string(result)) 
    }
}

func TestGetCredentialsForNonExistentUserReturnsError(t *testing.T) {
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"

    result, err := GetCredentialsForSiteHash(sitehash, username)
    if err == nil {
        t.Fatalf("GetCredentialsForSiteHash for missing sitehash did not return an error!")
    }

    if result != nil {
        t.Fatalf("GetCredentialsForSiteHash for missing sitehash returned an unexpected result! %s", string(result))
    }
}

func TestPutCredentialsForSiteHashTwiceReturnsMostRecentUpdateWhenGet(t *testing.T) {
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials1 := []byte("test1")
    credentials2 := []byte("test2")

    PutCredentialsForSiteHash(sitehash, username, credentials1)
    PutCredentialsForSiteHash(sitehash, username, credentials2)
    result, err := GetCredentialsForSiteHash(sitehash, username)
    if err != nil {
        t.Fatalf("Error while reading credentials for site: %v", err)
    }

    if !bytes.Equal(credentials2, result) {
        t.Fatalf("Returned credentials do not match updated credentials! Expected: %s, Received: %s", string(credentials2), string(result)) 
    }
}

func TestRemoveCredentialsForSiteHash(t *testing.T) {
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")

    PutCredentialsForSiteHash(sitehash, username, credentials)
    RemoveCredentialsForSiteHash(sitehash, username)
    result, err := GetCredentialsForSiteHash(sitehash, username)
    if err == nil {
        t.Fatalf("GetCredentialsForSiteHash for deleted sitehash did not return an error!")
    }

    if result != nil {
        t.Fatalf("GetCredentialsForSiteHash for missing user returned an unexpected result! %s", string(result))
    }
}

func TestRemoveCredentialsForSiteHashForNonExistentSiteHashDoesNotModifyResult(t *testing.T) {
    defer localstorage.CleanDB()
    sitehash := "testSiteHash"
    username := "testUsername"
    credentials := []byte("test")

    PutCredentialsForSiteHash(sitehash, username, credentials)
    data1, err := os.ReadFile(getLocalFilePath(username))
    if err != nil {
        t.Fatalf("Failed to read DB file: %s\n", err)
    }
    RemoveCredentialsForSiteHash("otherSite", username)
    data2, err := os.ReadFile(getLocalFilePath(username))
    if err != nil {
        t.Fatalf("Failed to read DB file: %s\n", err)
    }

    if !bytes.Equal(data1, data2) {
        t.Fatalf("RemoveCredentialsForSiteHash modified the file! Original: %s, New: %s", string(data1), string(data2))
    }
}


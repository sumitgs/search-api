package api

import "testing"

func Test_GoogleCredentials(t *testing.T) {
	SetGoogleCredential()
	if apikey == "" {
		t.Errorf("credential is empty, subsequent test will fail")
	}
}

// func Test_GoogleResourceQuery(t *testing.T) {
// 	responseChannel := make(chan model.GoogleResponses)

// 	go GoogleResourceQuery("barcelona", responseChannel)

// 	googleResponse := <-responseChannel

// 	if len(googleResponse.Items) == 0 {
// 		t.Fatalf("Expected greater than 0 search result")
// 	}
// }

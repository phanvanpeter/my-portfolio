package handlers

// TODO TestTasks implement the test
//func TestTasks(t *testing.T) {
//	r := httptest.NewRequest(http.MethodGet, "/tasks", nil)
//	w := httptest.NewRecorder()
//
//	Tasks(w, r)
//	res := w.Result()
//	defer res.Body.Close()
//
//	_, err := ioutil.ReadAll(res.Body)
//	if err != nil {
//		t.Errorf("Error reading the response body: %s", err)
//	}
//
//	if res.StatusCode != http.StatusOK {
//		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
//	}
//}

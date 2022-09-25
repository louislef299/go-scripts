package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	key_name = "key"
	url      = "http://localhost:8080/v1"

	stress_lvl = 10
)

func TestMain(m *testing.M) {
	go serv()
	code := m.Run()
	os.Remove("transaction.log")
	os.Exit(code)
}

func Test_Put(t *testing.T) {
	client := http.Client{}
	t.Parallel()

	for i := 1; i <= stress_lvl; i++ {
		t.Run(fmt.Sprintf("Test_PUT_Key%d", i), func(t *testing.T) {
			u := fmt.Sprintf("%s/%s%d", url, key_name, i)
			req, err := http.NewRequest(http.MethodPut, u, bytes.NewBuffer([]byte("value")))
			if err != nil {
				t.Fatal(err)
			}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, http.StatusCreated, resp.StatusCode)
		})
	}
}

func Test_Get(t *testing.T) {
	t.Parallel()

	for i := 1; i <= stress_lvl; i++ {
		t.Run(fmt.Sprintf("Test_GET_Key%d", i), func(t *testing.T) {
			time.Sleep(time.Millisecond)
			u := fmt.Sprintf("%s/%s%d", url, key_name, i)
			resp, err := http.Get(u) // NewRequest(http.MethodGet, u, nil)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, http.StatusOK, resp.StatusCode)
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			require.Equal(t, "value", string(data))
		})
	}
}

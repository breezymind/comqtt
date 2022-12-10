package cluster

import (
	"os"
	"reflect"
	"testing"

	"github.com/breezymind/comqtt/server/internal/utils"

	uuid "github.com/satori/go.uuid"
)

func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

func GetIP() string {
	ip, _ := utils.GetOutBoundIP()
	return ip
}

func GenNodeName() string {
	hostname, _ := os.Hostname()
	//return fmt.Sprintf("%s--%s", hostname, GenerateUUID4())
	return hostname
}

// GenerateUUID4 create a UUID
func GenerateUUID4() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return u.String()
}

// Unset remove element at position i
func Unset(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	return a[:len(a)-1]
}

// Expect compare two values for testing
func Expect(t *testing.T, got, want interface{}) {
	t.Logf(`Comparing values %v, %v`, got, want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %v, want %v`, got, want)
	}
}

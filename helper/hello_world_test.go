package helper

import (
	"fmt"
	"testing"
	"runtime"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// After
	fmt.Println("AFTER UNIT TEST")
}

func TestHelloWorldEri(t *testing.T) {
	result := HelloWorld("Eri")

	if result != "Hello Eri" {
		// unit test failed
		// t.Fail()
		t.Error("Result must be 'Hello Eri'")
	}

	fmt.Println("Test HelloWorldEri Done")
}

func TestHelloWorldWandri(t *testing.T) {
	result := HelloWorld("Wandri")

	if result != "Hello Wandri" {
		// unit test failed
		// t.FailNow()
		t.Fatal("Result must be 'Hello Wandri'")
	}

	fmt.Println("Test HelloWorldWandri Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eri")

	assert.Equal(t, "Hello Eri", result, "Result must be 'Hello Eri'")

	fmt.Println("TestHelloWorld with assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eri")

	require.Equal(t, "Hello Eri", result, "Result must be 'Hello Eri'")

	fmt.Println("TestHelloWorld with require Done")
}

func TestSkip(t *testing.T) {
	fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Linux")
	}

	result := HelloWorld("Eri")
	require.Equal(t, "Hello Eri", result, "Result must be 'Hello Eri'")
}
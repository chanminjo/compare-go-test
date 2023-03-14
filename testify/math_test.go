package math

import (
        "github.com/stretchr/testify/assert"
        "testing"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
        arg1, arg2, expected int
}

var addTests = []addTest{
        addTest{2, 3, 5},
        addTest{4, 8, 12},
        addTest{6, 9, 15},
        addTest{3, 10, 13},
        addTest{4, 6, 10},
        addTest{4, 7, 11},
}

func TestAdd(t *testing.T) {

        for _, test := range addTests {
                if output := Add(test.arg1, test.arg2); output != test.expected {
                        t.Errorf("Output %q not equal to expected %q", output, test.expected)
                }
        }
}

func TestSomething(t *testing.T) {
        assert.True(t, true, "True is true!")
}

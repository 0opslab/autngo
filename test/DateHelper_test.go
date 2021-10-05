package test

import (
	"fmt"
	"testing"

	Autngo "github.com/0opslab/autngo"
)

func TestNewGoTime(t *testing.T) {
	fmt.Println(Autngo.DateHelper.Now())
}
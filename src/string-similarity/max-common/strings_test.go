package max_common

import (
	"testing"
)

func TestCommonString(t *testing.T) {
	t.Log(CommonString("poll", "pull"))
	t.Log(CommonString("hisah", "vista"))
	t.Log(CommonString("fish", "fosh"))
}
package max_common

import "testing"

func TestCommonSequence(t *testing.T) {
	t.Log(CommonSequence("fish", "fosh"))
	t.Log(CommonSequence("poll", "pull"))
	t.Log(CommonSequence("hisah", "vista"))
}
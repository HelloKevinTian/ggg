package some

import "testing"

func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %v %v", a, b)
	}
}

func assertNotEqual(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("Be Equal. %v %v", a, b)
	}
}

func assertNil(t *testing.T, a interface{}) {
	if a != nil {
		t.Errorf("Not nil. %v", a)
	}
}

func assertNotNil(t *testing.T, a interface{}) {
	if a == nil {
		t.Errorf("Be nil. %v", a)
	}
}

func assertTrue(t *testing.T, a bool) {
	if !a {
		t.Errorf("Not True. %v", a)
	}
}

func assertEmpty(t *testing.T, a interface{}) {
	switch a.(type) {
	case []string:
		if len(a.([]string)) > 0 {
			t.Errorf("Not Empty. %v", a)
		}
	case []int:
		if len(a.([]int)) > 0 {
			t.Errorf("Not Empty. %v", a)
		}
	}
}

func assertNotEmpty(t *testing.T, a interface{}) {
	switch a.(type) {
	case []string:
		if len(a.([]string)) == 0 {
			t.Errorf("Be Empty. %v", a)
		}
	case []int:
		if len(a.([]int)) == 0 {
			t.Errorf("Be Empty. %v", a)
		}
	}
}

func assertLen(t *testing.T, a interface{}, length int) {
	switch a.(type) {
	case []string:
		if len(a.([]string)) != length {
			t.Errorf("Length Equal. %v", a)
		}
	case []int:
		if len(a.([]int)) != length {
			t.Errorf("Length Equal. %v", a)
		}
	}
}

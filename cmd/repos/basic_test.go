package repos

import (
	"testing"
)

func TestModelBasic_String(t *testing.T) {
	model := &ModelBasic{Uuid: "test-uuid"}
	expected := "uuid: test-uuid"
	if model.String() != expected {
		t.Errorf("Expected %s, got %s", expected, model.String())
	}
}

func TestModelBasic_Log(t *testing.T) {
	model := &ModelBasic{Uuid: "test-uuid"}
	// Assuming sLog.Info logs the message, we can't capture the log output in a simple way.
	// So, this test will just call the method to ensure it doesn't panic.
	model.Log()
}

func TestModelBasic_CUuid(t *testing.T) {
	model := &ModelBasic{Uuid: "test-uuid"}
	expected := "test-uuid"
	if model.CUuid() != expected {
		t.Errorf("Expected %s, got %s", expected, model.CUuid())
	}
}

func TestModelss_ImplementsInterModelss(t *testing.T) {
	var _ InterModelss = &Modelss{}
}

func TestModel_A_ImplementsInterModelss(t *testing.T) {
	var _ InterModelss = &Model_A{}
}

func TestModel_B_ImplementsInterModelss(t *testing.T) {
	var _ InterModelss = &Model_B{}
}

package cpu6510

import (
	"testing"
)

func TestBRK(t *testing.T) {
	cpu := NewCPU()

	cpu.BRK()

	if !cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be set")
	}

	if !cpu.statusRegister.breakCommandFlag {
		t.Errorf("Break status flag should be set")
	}
}

func TestCLC(t *testing.T) {
	cpu := NewCPU()
	cpu.statusRegister.carry = true

	cpu.CLC()

	if cpu.statusRegister.carry {
		t.Errorf("Carry flag should be cleared")
	}
}

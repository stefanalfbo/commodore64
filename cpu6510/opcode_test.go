package cpu6510

import (
	"testing"
)

func TestStatusRegister(t *testing.T) {
	t.Run("Clear the carry flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.statusRegister.carry = true

		cpu.CLC()

		if cpu.statusRegister.carry {
			t.Errorf("Carry flag should be cleared")
		}
	})
}

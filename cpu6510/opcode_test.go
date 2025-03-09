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

	t.Run("Set the interrupt disable flag", func(t *testing.T) {
		cpu := NewCPU()

		cpu.BRK()

		if !cpu.statusRegister.interruptDisableFlag {
			t.Errorf("Interrupt disable flag should be set")
		}
	})

	t.Run("Set the break status flag", func(t *testing.T) {
		cpu := NewCPU()

		cpu.BRK()

		if !cpu.statusRegister.breakCommandFlag {
			t.Errorf("Break status flag should be set")
		}
	})
}

package cpu6510

import "testing"

func TestNewCPU(t *testing.T) {
	cpu := NewCPU()

	if cpu.programCounter != 0 {
		t.Errorf("Program counter should be initialized to 0")
	}

	if cpu.statusRegister.carry {
		t.Errorf("Carry flag should be initialized to false")
	}

	for i := 0; i < memorySize; i++ {
		if cpu.ram[i] != 0 {
			t.Errorf("RAM should be initialized to 0")
		}
	}
}

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

func TestRunCpu(t *testing.T) {
	cpu := NewCPU()

	cpu.ram[0] = 0x18
	cpu.ram[1] = 0x00

	cpu.Run()

	if cpu.statusRegister.carry {
		t.Errorf("Carry flag should be cleared")
	}
}

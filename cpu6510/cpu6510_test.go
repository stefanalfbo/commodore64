package cpu6510

import "testing"

func TestNewCPU(t *testing.T) {
	cpu := NewCPU()

	if cpu.programCounter != 0 {
		t.Errorf("Program counter should be initialized to 0")
	}

	if cpu.statusRegister.carryFlag {
		t.Errorf("Carry flag should be initialized to false")
	}

	if cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be initialized to false")
	}

	if cpu.statusRegister.decimalModeFlag {
		t.Errorf("Decimal mode flag should be initialized to false")
	}

	for i := 0; i < memorySize; i++ {
		if cpu.ram[i] != 0 {
			t.Errorf("RAM should be initialized to 0")
		}
	}
}

func TestStatusRegister(t *testing.T) {

	t.Run("Carry flag", func(t *testing.T) {
		t.Run("Set the carry flag", func(t *testing.T) {
			cpu := NewCPU()

			cpu.execute(OpCodeAsHex("SEC"))

			if !cpu.statusRegister.carryFlag {
				t.Errorf("Carry flag should be set")
			}
		})

		t.Run("Clear the carry flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.statusRegister.carryFlag = true

			cpu.execute(OpCodeAsHex("CLC"))

			if cpu.statusRegister.carryFlag {
				t.Errorf("Carry flag should be cleared")
			}
		})
	})

	t.Run("Interrupt disable flag", func(t *testing.T) {
		t.Run("Clear the interrupt disable flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.statusRegister.interruptDisableFlag = true

			cpu.execute(OpCodeAsHex("CLI"))

			if cpu.statusRegister.interruptDisableFlag {
				t.Errorf("Interrupt disable flag should be cleared")
			}
		})

		t.Run("Set the interrupt disable flag", func(t *testing.T) {
			cpu := NewCPU()

			cpu.execute(OpCodeAsHex("BRK"))

			if !cpu.statusRegister.interruptDisableFlag {
				t.Errorf("Interrupt disable flag should be set")
			}
		})
	})

	t.Run("Decimal mode flag", func(t *testing.T) {
		t.Run("Clear the decimal mode flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.statusRegister.decimalModeFlag = true

			cpu.execute(OpCodeAsHex("CLD"))

			if cpu.statusRegister.decimalModeFlag {
				t.Errorf("Decimal mode flag should be cleared")
			}
		})

		t.Run("Set the decimal mode flag", func(t *testing.T) {
			cpu := NewCPU()

			cpu.execute(OpCodeAsHex("SED"))

			if !cpu.statusRegister.decimalModeFlag {
				t.Errorf("Decimal mode flag should be set")
			}
		})
	})

	t.Run("Break status flag", func(t *testing.T) {
		t.Run("Set the break status flag", func(t *testing.T) {
			cpu := NewCPU()

			cpu.execute(OpCodeAsHex("BRK"))

			if !cpu.statusRegister.breakCommandFlag {
				t.Errorf("Break status flag should be set")
			}
		})
	})

	t.Run("Overflow flag", func(t *testing.T) {
		t.Run("Clear the overflow flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.statusRegister.overflowFlag = true

			cpu.execute(OpCodeAsHex("CLV"))

			if cpu.statusRegister.overflowFlag {
				t.Errorf("Overflow flag should be cleared")
			}
		})
	})
}

func TestRunCpu(t *testing.T) {
	cpu := NewCPU()

	cpu.ram[0] = 0x18
	cpu.ram[1] = 0x00

	cpu.Run()

	if cpu.statusRegister.carryFlag {
		t.Errorf("Carry flag should be cleared")
	}
}

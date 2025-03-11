package cpu6510

import (
	"testing"
)

func TestBRK(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 2

	cpu.execute(OpCodeAsHex("BRK"))

	if !cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be set")
	}

	if !cpu.statusRegister.breakCommandFlag {
		t.Errorf("Break status flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented by 2")
	}
}

func TestCLC(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.carryFlag = true

	cpu.execute(OpCodeAsHex("CLC"))

	if cpu.statusRegister.carryFlag {
		t.Errorf("Carry flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestSEC(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("SEC"))

	if !cpu.statusRegister.carryFlag {
		t.Errorf("Carry flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestPHA(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.accumulator = 0x03

	cpu.execute(OpCodeAsHex("PHA"))

	if cpu.ram[0x01FF] != cpu.accumulator {
		t.Errorf("Accumulator should be pushed onto the stack")
	}

	if cpu.stackPointer != 0xFE {
		t.Errorf("Stack pointer should be decremented")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestCLI(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.interruptDisableFlag = true

	cpu.execute(OpCodeAsHex("CLI"))

	if cpu.statusRegister.interruptDisableFlag {
		t.Errorf("Interrupt disable flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestCLV(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.overflowFlag = true

	cpu.execute(OpCodeAsHex("CLV"))

	if cpu.statusRegister.overflowFlag {
		t.Errorf("Overflow flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestNOP(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("NOP"))

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestCLD(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.statusRegister.decimalModeFlag = true

	cpu.execute(OpCodeAsHex("CLD"))

	if cpu.statusRegister.decimalModeFlag {
		t.Errorf("Decimal mode flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestSED(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1

	cpu.execute(OpCodeAsHex("SED"))

	if !cpu.statusRegister.decimalModeFlag {
		t.Errorf("Decimal mode flag should be set")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

func TestRegisterInstructionsImpliedMode(t *testing.T) {
	t.Run("DEY - Decrement Y register", func(t *testing.T) {
		t.Run("Decrement Y register", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x03

			cpu.execute(OpCodeAsHex("DEY"))

			if cpu.yRegister != 0x02 {
				t.Errorf("Y register should be decremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Decrement Y register and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x01

			cpu.execute(OpCodeAsHex("DEY"))

			if cpu.yRegister != 0x00 {
				t.Errorf("Y register should be decremented")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Decrement Y register and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x00

			cpu.execute(OpCodeAsHex("DEY"))

			if cpu.yRegister != 0xFF {
				t.Errorf("Y register should be decremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("TXA - Transfer X to A", func(t *testing.T) {
		t.Run("Transfer X to A", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x03

			cpu.execute(OpCodeAsHex("TXA"))

			if cpu.accumulator != cpu.xRegister {
				t.Errorf("Accumulator should be set to the value of the X register")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer X to A and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x00

			cpu.execute(OpCodeAsHex("TXA"))

			if cpu.accumulator != cpu.xRegister {
				t.Errorf("Accumulator should be set to the value of the X register")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer X to A and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x80

			cpu.execute(OpCodeAsHex("TXA"))

			if cpu.accumulator != cpu.xRegister {
				t.Errorf("Accumulator should be set to the value of the X register")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("TYA - Transfer Y to A", func(t *testing.T) {
		t.Run("Transfer Y to A", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x03

			cpu.execute(OpCodeAsHex("TYA"))

			if cpu.accumulator != cpu.yRegister {
				t.Errorf("Accumulator should be set to the value of the Y register")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer Y to A and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x00

			cpu.execute(OpCodeAsHex("TYA"))

			if cpu.accumulator != cpu.yRegister {
				t.Errorf("Accumulator should be set to the value of the Y register")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer Y to A and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x80

			cpu.execute(OpCodeAsHex("TYA"))

			if cpu.accumulator != cpu.yRegister {
				t.Errorf("Accumulator should be set to the value of the Y register")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("TAY - Transfer A to Y", func(t *testing.T) {
		t.Run("Transfer A to X", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x03

			cpu.execute(OpCodeAsHex("TAY"))

			if cpu.yRegister != 0x03 {
				t.Errorf("Y register should be set to the value of the accumulator")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer A to Y and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x00

			cpu.execute(OpCodeAsHex("TAY"))

			if cpu.yRegister != 0x00 {
				t.Errorf("Y register should be set to the value of the accumulator")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer A to Y and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x80

			cpu.execute(OpCodeAsHex("TAY"))

			if cpu.yRegister != 0x80 {
				t.Errorf("Y register should be set to the value of the accumulator")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("TAX - Transfer A to X", func(t *testing.T) {
		t.Run("Transfer A to X", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x03

			cpu.execute(OpCodeAsHex("TAX"))

			if cpu.xRegister != 0x03 {
				t.Errorf("X register should be set to the value of the accumulator")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer A to X and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x00

			cpu.execute(OpCodeAsHex("TAX"))

			if cpu.xRegister != 0x00 {
				t.Errorf("X register should be set to the value of the accumulator")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Transfer A to X and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.accumulator = 0x80

			cpu.execute(OpCodeAsHex("TAX"))

			if cpu.xRegister != 0x80 {
				t.Errorf("X register should be set to the value of the accumulator")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("INY - Increment Y register", func(t *testing.T) {
		t.Run("Increment Y register", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x01

			cpu.execute(OpCodeAsHex("INY"))

			if cpu.yRegister != 0x02 {
				t.Errorf("X register should be incremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Increment Y register and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0xFF

			cpu.execute(OpCodeAsHex("INY"))

			if cpu.yRegister != 0x00 {
				t.Errorf("Y register should be incremented")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Increment Y register and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.yRegister = 0x7F

			cpu.execute(OpCodeAsHex("INY"))

			if cpu.yRegister != 0x80 {
				t.Errorf("Y register should be incremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("DEX - Decrement X register", func(t *testing.T) {
		t.Run("Decrement X register", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x03

			cpu.execute(OpCodeAsHex("DEX"))

			if cpu.xRegister != 0x02 {
				t.Errorf("X register should be decremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Decrement X register and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x01

			cpu.execute(OpCodeAsHex("DEX"))

			if cpu.xRegister != 0x00 {
				t.Errorf("X register should be decremented")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Decrement X register and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x00

			cpu.execute(OpCodeAsHex("DEX"))

			if cpu.xRegister != 0xFF {
				t.Errorf("X register should be decremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})

	t.Run("INX - Increment X register", func(t *testing.T) {
		t.Run("Increment X register", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x01

			cpu.execute(OpCodeAsHex("INX"))

			if cpu.xRegister != 0x02 {
				t.Errorf("X register should be incremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Increment X register and set zero flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0xFF

			cpu.execute(OpCodeAsHex("INX"))

			if cpu.xRegister != 0x00 {
				t.Errorf("X register should be incremented")
			}

			if !cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be set")
			}

			if cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be cleared")
			}
		})

		t.Run("Increment X register and set negative flag", func(t *testing.T) {
			cpu := NewCPU()
			cpu.xRegister = 0x7F

			cpu.execute(OpCodeAsHex("INX"))

			if cpu.xRegister != 0x80 {
				t.Errorf("X register should be incremented")
			}

			if cpu.statusRegister.zeroFlag {
				t.Errorf("Zero flag should be cleared")
			}

			if !cpu.statusRegister.negativeFlag {
				t.Errorf("Negative flag should be set")
			}
		})
	})
}

package cpu6510

import "testing"

func TestPHA(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.accumulator = 0x03

	cpu.execute(InstructionAsHex("PHA"))

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

func TestPLA(t *testing.T) {
	t.Run("Pull accumulator", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0x03
		cpu.stackPointer = 0xFE

		cpu.execute(InstructionAsHex("PLA"))

		if cpu.accumulator != cpu.ram[0x01FF] {
			t.Errorf("Accumulator should be set to the value on the stack")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Pull accumulator and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0x00
		cpu.stackPointer = 0xFE

		cpu.execute(InstructionAsHex("PLA"))

		if cpu.accumulator != cpu.ram[0x01FF] {
			t.Errorf("Accumulator should be set to the value on the stack")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Pull accumulator and set negative flag", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0x80
		cpu.stackPointer = 0xFE

		cpu.execute(InstructionAsHex("PLA"))

		if cpu.accumulator != cpu.ram[0x01FF] {
			t.Errorf("Accumulator should be set to the value on the stack")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestPHP(t *testing.T) {
	t.Run("Push processor status register flags", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.statusRegister.carryFlag = true
		cpu.statusRegister.zeroFlag = true
		cpu.statusRegister.interruptDisableFlag = true
		cpu.statusRegister.decimalModeFlag = true
		cpu.statusRegister.breakCommandFlag = true
		cpu.statusRegister.overflowFlag = true
		cpu.statusRegister.negativeFlag = true

		cpu.execute(InstructionAsHex("PHP"))

		if cpu.ram[0x01FF] != 0xFF {
			t.Errorf("Status register should be pushed onto the stack")
		}

		if cpu.stackPointer != 0xFE {
			t.Errorf("Stack pointer should be decremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Push processor status register flags when all is cleared", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1

		cpu.execute(InstructionAsHex("PHP"))

		if cpu.ram[0x01FF] != 0x20 {
			t.Errorf("Status register should be pushed onto the stack")
		}

		if cpu.stackPointer != 0xFE {
			t.Errorf("Stack pointer should be decremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestPLP(t *testing.T) {
	t.Run("Pull processor status register flags when all is cleared", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0x20
		cpu.stackPointer = 0xFE

		cpu.execute(InstructionAsHex("PLP"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.interruptDisableFlag {
			t.Errorf("Interrupt disable flag should be cleared")
		}

		if cpu.statusRegister.decimalModeFlag {
			t.Errorf("Decimal mode flag should be cleared")
		}

		if cpu.statusRegister.breakCommandFlag {
			t.Errorf("Break command flag should be cleared")
		}

		if cpu.statusRegister.overflowFlag {
			t.Errorf("Overflow flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Pull processor status register flags when all is set", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.ram[0x01FF] = 0xFF
		cpu.stackPointer = 0xFE

		cpu.execute(InstructionAsHex("PLP"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if !cpu.statusRegister.interruptDisableFlag {
			t.Errorf("Interrupt disable flag should be set")
		}

		if !cpu.statusRegister.decimalModeFlag {
			t.Errorf("Decimal mode flag should be set")
		}

		if !cpu.statusRegister.breakCommandFlag {
			t.Errorf("Break command flag should be set")
		}

		if !cpu.statusRegister.overflowFlag {
			t.Errorf("Overflow flag should be set")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}

		if cpu.stackPointer != 0xFF {
			t.Errorf("Stack pointer should be incremented")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})
}

func TestTSX(t *testing.T) {

	t.Run("Transfer stack pointer to X", func(t *testing.T) {
		cpu := NewCPU()
		expectedPC := cpu.programCounter + 1
		cpu.stackPointer = 0x03

		cpu.execute(InstructionAsHex("TSX"))

		if cpu.xRegister != cpu.stackPointer {
			t.Errorf("X register should be set to the value of the stack pointer")
		}

		if cpu.programCounter != expectedPC {
			t.Errorf("Program counter should be incremented")
		}
	})

	t.Run("Transfer stack pointer to X and set zero flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.stackPointer = 0x00

		cpu.execute(InstructionAsHex("TSX"))

		if cpu.xRegister != cpu.stackPointer {
			t.Errorf("X register should be set to the value of the stack pointer")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Transfer stack pointer to X and set negative flag", func(t *testing.T) {
		cpu := NewCPU()
		cpu.stackPointer = 0x80

		cpu.execute(InstructionAsHex("TSX"))

		if cpu.xRegister != cpu.stackPointer {
			t.Errorf("X register should be set to the value of the stack pointer")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})
}

func TestTXS(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 1
	cpu.xRegister = 0x03

	cpu.execute(InstructionAsHex("TXS"))

	if cpu.stackPointer != cpu.xRegister {
		t.Errorf("Stack pointer should be set to the value of the X register")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

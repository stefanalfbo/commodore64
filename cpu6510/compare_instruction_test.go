package cpu6510

import "testing"

func TestCMPImmediate(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPImmediate"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.ram[1] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CMPImmediate"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPImmediate"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPAbsolute(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != 0x0003 {
			t.Errorf("Program counter should be incremented by 3")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CMPAbsolute"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPAbsoluteY(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPAbsoluteY"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != 0x0003 {
			t.Errorf("Program counter should be incremented by 3")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CMPAbsoluteY"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPAbsoluteY"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPZeroPageX(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPZeroPageX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CMPZeroPageX"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPZeroPageX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPAbsoluteX(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPAbsoluteX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != 0x0003 {
			t.Errorf("Program counter should be incremented by 3")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CMPAbsoluteX"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPAbsoluteX"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPZeroPage(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPZeroPage"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CMPZeroPage"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPZeroPage"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPIndexedIndirect(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x37
		cpu.ram[0x15] = 0x13
		cpu.ram[0x1337] = 0x37
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPIndexedIndirect"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x37
		cpu.ram[0x15] = 0x13
		cpu.ram[0x1337] = 0x37
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CMPIndexedIndirect"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x37
		cpu.ram[0x15] = 0x13
		cpu.ram[0x1337] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPIndexedIndirect"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}

		if cpu.programCounter != 0x0002 {
			t.Errorf("Program counter should be incremented by 2")
		}
	})

	t.Run("Index wraps around", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.xRegister = 0xFF
		cpu.ram[1] = 0x13
		cpu.ram[0x12] = 0x37
		cpu.ram[0x13] = 0x13
		cpu.ram[0x1337] = 0x37
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPIndexedIndirect"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCMPIndirectIndexed(t *testing.T) {
	t.Run("Accumulator is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x37
		cpu.ram[0x14] = 0x13
		cpu.ram[0x1338] = 0x37
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPIndirectIndexed"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Accumulator is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x01
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x37
		cpu.ram[0x14] = 0x13
		cpu.ram[0x1338] = 0x37
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CMPIndirectIndexed"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Accumulator is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x42
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x37
		cpu.ram[0x14] = 0x13
		cpu.ram[0x1338] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CMPIndirectIndexed"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCPXImmediate(t *testing.T) {
	t.Run("X register is greater than immediate value", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x42
		cpu.ram[1] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPXImmediate"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("X register is less than immediate value", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CPXImmediate"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("X register is equal to immediate value", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x42
		cpu.ram[1] = 0x42
		cpu.statusRegister.carryFlag = false
		cpu.statusRegister.negativeFlag = false
		cpu.statusRegister.zeroFlag = true

		cpu.execute(InstructionAsHex("CPXImmediate"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCPXZeroPage(t *testing.T) {
	t.Run("X register is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x42
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPXZeroPage"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("X register is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CPXZeroPage"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("X register is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x42
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPXZeroPage"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCPXAbsolute(t *testing.T) {
	t.Run("X register is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x42
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPXAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("X register is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CPXAbsolute"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("X register is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.xRegister = 0x42
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPXAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCPYImmediate(t *testing.T) {
	t.Run("Y register is greater than immediate value", func(t *testing.T) {
		cpu := NewCPU()
		cpu.yRegister = 0x42
		cpu.ram[1] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPYImmediate"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Y register is less than immediate value", func(t *testing.T) {
		cpu := NewCPU()
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CPYImmediate"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Y register is equal to immediate value", func(t *testing.T) {
		cpu := NewCPU()
		cpu.yRegister = 0x42
		cpu.ram[1] = 0x42
		cpu.statusRegister.carryFlag = false
		cpu.statusRegister.negativeFlag = false
		cpu.statusRegister.zeroFlag = true

		cpu.execute(InstructionAsHex("CPYImmediate"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCPYZeroPage(t *testing.T) {
	t.Run("Y register is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.yRegister = 0x42
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPYZeroPage"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Y register is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CPYZeroPage"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Y register is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.yRegister = 0x42
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPYZeroPage"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestCPYAbsolute(t *testing.T) {
	t.Run("Y register is greater than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.yRegister = 0x42
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x01
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPYAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})

	t.Run("Y register is less than memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x42
		cpu.statusRegister.carryFlag = true

		cpu.execute(InstructionAsHex("CPYAbsolute"))

		if cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be cleared")
		}

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("Y register is equal to memory", func(t *testing.T) {
		cpu := NewCPU()
		cpu.yRegister = 0x42
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x42
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("CPYAbsolute"))

		if !cpu.statusRegister.carryFlag {
			t.Errorf("Carry flag should be set")
		}

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

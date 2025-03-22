package cpu6510

import "testing"

func TestORAImmediate(t *testing.T) {
	t.Run("Bit-wise boolean or between each eight bits", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b01010101
		cpu.ram[1] = 0b10101010

		cpu.execute(InstructionAsHex("ORAImmediate"))

		if cpu.accumulator != 0b11111111 {
			t.Errorf("Accumulator should be 0b11111111")
		}
	})

	t.Run("The OR operation has its most significant bit set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b00000001
		cpu.ram[1] = 0b10000000

		cpu.execute(InstructionAsHex("ORAImmediate"))

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("The OR operation result is zero", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x00
		cpu.ram[1] = 0x00
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("ORAImmediate"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestORAAbsolute(t *testing.T) {
	t.Run("Bit-wise boolean or between each eight bits", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b01010101
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0b10101010

		cpu.execute(InstructionAsHex("ORAAbsolute"))

		if cpu.accumulator != 0b11111111 {
			t.Errorf("Accumulator should be 0b11111111")
		}
	})

	t.Run("The OR operation has its most significant bit set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b00000001
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0b10000000

		cpu.execute(InstructionAsHex("ORAAbsolute"))

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("The OR operation result is zero", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x00
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1337] = 0x00
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("ORAAbsolute"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestORAAbsoluteX(t *testing.T) {
	t.Run("Bit-wise boolean or between each eight bits", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b01010101
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0b10101010

		cpu.execute(InstructionAsHex("ORAAbsoluteX"))

		if cpu.accumulator != 0b11111111 {
			t.Errorf("Accumulator should be 0b11111111")
		}
	})

	t.Run("The OR operation has its most significant bit set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b00000001
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0b10000000

		cpu.execute(InstructionAsHex("ORAAbsoluteX"))

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("The OR operation result is zero", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x00
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x00
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("ORAAbsoluteX"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestORAAbsoluteY(t *testing.T) {
	t.Run("Bit-wise boolean or between each eight bits", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b01010101
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0b10101010

		cpu.execute(InstructionAsHex("ORAAbsoluteY"))

		if cpu.accumulator != 0b11111111 {
			t.Errorf("Accumulator should be 0b11111111")
		}
	})

	t.Run("The OR operation has its most significant bit set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b00000001
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0b10000000

		cpu.execute(InstructionAsHex("ORAAbsoluteY"))

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("The OR operation result is zero", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x00
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x37
		cpu.ram[2] = 0x13
		cpu.ram[0x1338] = 0x00
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("ORAAbsoluteY"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestORAZeroPage(t *testing.T) {
	t.Run("Bit-wise boolean or between each eight bits", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b01010101
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0b10101010

		cpu.execute(InstructionAsHex("ORAZeroPage"))

		if cpu.accumulator != 0b11111111 {
			t.Errorf("Accumulator should be 0b11111111")
		}
	})

	t.Run("The OR operation has its most significant bit set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b00000001
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0b10000000

		cpu.execute(InstructionAsHex("ORAZeroPage"))

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("The OR operation result is zero", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x00
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x00
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("ORAZeroPage"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestORAZeroPageX(t *testing.T) {
	t.Run("Bit-wise boolean or between each eight bits", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b01010101
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0b10101010

		cpu.execute(InstructionAsHex("ORAZeroPageX"))

		if cpu.accumulator != 0b11111111 {
			t.Errorf("Accumulator should be 0b11111111")
		}
	})

	t.Run("The OR operation has its most significant bit set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b00000001
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0b10000000

		cpu.execute(InstructionAsHex("ORAZeroPageX"))

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("The OR operation result is zero", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x00
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x00
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("ORAZeroPageX"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestORAIndexedIndirect(t *testing.T) {
	t.Run("Bit-wise boolean or between each eight bits", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b01010101
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x37
		cpu.ram[0x15] = 0x13
		cpu.ram[0x1337] = 0b10101010

		cpu.execute(InstructionAsHex("ORAIndexedIndirect"))

		if cpu.accumulator != 0b11111111 {
			t.Errorf("Accumulator should be 0b11111111")
		}
	})

	t.Run("The OR operation has its most significant bit set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b00000001
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x37
		cpu.ram[0x15] = 0x13
		cpu.ram[0x1337] = 0b10000000

		cpu.execute(InstructionAsHex("ORAIndexedIndirect"))

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("The OR operation result is zero", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x00
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0x37
		cpu.ram[0x15] = 0x13
		cpu.ram[0x1337] = 0x00
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("ORAIndexedIndirect"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

func TestORAIndirectIndexed(t *testing.T) {
	t.Run("Bit-wise boolean or between each eight bits", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b01010101
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x37
		cpu.ram[0x14] = 0x13
		cpu.ram[0x1338] = 0b10101010

		cpu.execute(InstructionAsHex("ORAIndirectIndexed"))

		if cpu.accumulator != 0b11111111 {
			t.Errorf("Accumulator should be 0b11111111")
		}
	})

	t.Run("The OR operation has its most significant bit set", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b00000001
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x37
		cpu.ram[0x14] = 0x13
		cpu.ram[0x1338] = 0b10000000

		cpu.execute(InstructionAsHex("ORAIndirectIndexed"))

		if cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be cleared")
		}

		if !cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be set")
		}
	})

	t.Run("The OR operation result is zero", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0x00
		cpu.yRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x13] = 0x37
		cpu.ram[0x14] = 0x13
		cpu.ram[0x1338] = 0x00
		cpu.statusRegister.carryFlag = false

		cpu.execute(InstructionAsHex("ORAIndirectIndexed"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

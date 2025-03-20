package cpu6510

import "testing"

func TestORAIndexedIndirectX(t *testing.T) {
	t.Run("Bit-wise boolean or between each eight bits", func(t *testing.T) {
		cpu := NewCPU()
		cpu.accumulator = 0b01010101
		cpu.xRegister = 0x01
		cpu.ram[1] = 0x13
		cpu.ram[0x14] = 0b10101010

		cpu.execute(OpCodeAsHex("ORAIndexedIndirectX"))

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

		cpu.execute(OpCodeAsHex("ORAIndexedIndirectX"))

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

		cpu.execute(OpCodeAsHex("ORAIndexedIndirectX"))

		if !cpu.statusRegister.zeroFlag {
			t.Errorf("Zero flag should be set")
		}

		if cpu.statusRegister.negativeFlag {
			t.Errorf("Negative flag should be cleared")
		}
	})
}

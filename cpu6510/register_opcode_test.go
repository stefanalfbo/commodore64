package cpu6510

import "testing"

func TestTAX(t *testing.T) {
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
}

func TestTAY(t *testing.T) {
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
}

func TestTXA(t *testing.T) {
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
}

func TestTYA(t *testing.T) {
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
}

func TestDEX(t *testing.T) {
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
}

func TestDEY(t *testing.T) {
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
}
func TestINX(t *testing.T) {
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
}

func TestINY(t *testing.T) {
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
}

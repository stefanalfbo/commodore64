package cpu6510

import "testing"

func TestJAM(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter

	cpu.execute(InstructionAsHex("JAM"))

	if !cpu.isJammed {
		t.Errorf("CPU should be jammed after JAM instruction")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should not be incremented")
	}
}

func TestRunCpuStopsOnJam(t *testing.T) {
	cpu := NewCPU()

	cpu.ram[0] = InstructionAsHex("JAM")
	cpu.ram[1] = InstructionAsHex("BRK")

	cpu.Run()

	if !cpu.isJammed {
		t.Errorf("CPU should be jammed after running JAM")
	}

	if cpu.programCounter != 0 {
		t.Errorf("Program counter should remain at JAM opcode")
	}
}

func TestSLOIndexedIndirect(t *testing.T) {
	cpu := NewCPU()
	expectedPC := cpu.programCounter + 2
	cpu.accumulator = 0b00000001
	cpu.xRegister = 0x01
	cpu.ram[1] = 0x13
	cpu.ram[0x14] = 0x37
	cpu.ram[0x15] = 0x13
	cpu.ram[0x1337] = 0b01000001

	cpu.execute(InstructionAsHex("SLOIndexedIndirect"))

	if cpu.ram[0x1337] != 0b10000010 {
		t.Errorf("Memory should be shifted left, expected 0b10000010, got %08b", cpu.ram[0x1337])
	}

	if cpu.accumulator != 0b10000011 {
		t.Errorf("Accumulator should be ORA'd with shifted value, expected 0b10000011, got %08b", cpu.accumulator)
	}

	if cpu.statusRegister.carryFlag {
		t.Errorf("Carry flag should be cleared")
	}

	if !cpu.statusRegister.negativeFlag {
		t.Errorf("Negative flag should be set")
	}

	if cpu.statusRegister.zeroFlag {
		t.Errorf("Zero flag should be cleared")
	}

	if cpu.programCounter != expectedPC {
		t.Errorf("Program counter should be incremented")
	}
}

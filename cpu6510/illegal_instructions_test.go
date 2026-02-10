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

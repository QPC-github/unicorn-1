package unicorn
// For Unicorn Engine. AUTO-GENERATED FILE, DO NOT EDIT [x86_const.go]
const (

// X86 CPU

	CPU_X86_QEMU64 = 0
	CPU_X86_PHENOM = 1
	CPU_X86_CORE2DUO = 2
	CPU_X86_KVM64 = 3
	CPU_X86_QEMU32 = 4
	CPU_X86_KVM32 = 5
	CPU_X86_COREDUO = 6
	CPU_X86_486 = 7
	CPU_X86_PENTIUM = 8
	CPU_X86_PENTIUM2 = 9
	CPU_X86_PENTIUM3 = 10
	CPU_X86_ATHLON = 11
	CPU_X86_N270 = 12
	CPU_X86_CONROE = 13
	CPU_X86_PENRYN = 14
	CPU_X86_NEHALEM = 15
	CPU_X86_WESTMERE = 16
	CPU_X86_SANDYBRIDGE = 17
	CPU_X86_IVYBRIDGE = 18
	CPU_X86_HASWELL = 19
	CPU_X86_BROADWELL = 20
	CPU_X86_SKYLAKE_CLIENT = 21
	CPU_X86_SKYLAKE_SERVER = 22
	CPU_X86_CASCADELAKE_SERVER = 23
	CPU_X86_COOPERLAKE = 24
	CPU_X86_ICELAKE_CLIENT = 25
	CPU_X86_ICELAKE_SERVER = 26
	CPU_X86_DENVERTON = 27
	CPU_X86_SNOWRIDGE = 28
	CPU_X86_KNIGHTSMILL = 29
	CPU_X86_OPTERON_G1 = 30
	CPU_X86_OPTERON_G2 = 31
	CPU_X86_OPTERON_G3 = 32
	CPU_X86_OPTERON_G4 = 33
	CPU_X86_OPTERON_G5 = 34
	CPU_X86_EPYC = 35
	CPU_X86_DHYANA = 36
	CPU_X86_EPYC_ROME = 37

// X86 registers

	X86_REG_INVALID = 0
	X86_REG_AH = 1
	X86_REG_AL = 2
	X86_REG_AX = 3
	X86_REG_BH = 4
	X86_REG_BL = 5
	X86_REG_BP = 6
	X86_REG_BPL = 7
	X86_REG_BX = 8
	X86_REG_CH = 9
	X86_REG_CL = 10
	X86_REG_CS = 11
	X86_REG_CX = 12
	X86_REG_DH = 13
	X86_REG_DI = 14
	X86_REG_DIL = 15
	X86_REG_DL = 16
	X86_REG_DS = 17
	X86_REG_DX = 18
	X86_REG_EAX = 19
	X86_REG_EBP = 20
	X86_REG_EBX = 21
	X86_REG_ECX = 22
	X86_REG_EDI = 23
	X86_REG_EDX = 24
	X86_REG_EFLAGS = 25
	X86_REG_EIP = 26
	X86_REG_ES = 27
	X86_REG_ESI = 28
	X86_REG_ESP = 29
	X86_REG_FPSW = 30
	X86_REG_FS = 31
	X86_REG_GS = 32
	X86_REG_IP = 33
	X86_REG_RAX = 34
	X86_REG_RBP = 35
	X86_REG_RBX = 36
	X86_REG_RCX = 37
	X86_REG_RDI = 38
	X86_REG_RDX = 39
	X86_REG_RIP = 40
	X86_REG_RSI = 41
	X86_REG_RSP = 42
	X86_REG_SI = 43
	X86_REG_SIL = 44
	X86_REG_SP = 45
	X86_REG_SPL = 46
	X86_REG_SS = 47
	X86_REG_CR0 = 48
	X86_REG_CR1 = 49
	X86_REG_CR2 = 50
	X86_REG_CR3 = 51
	X86_REG_CR4 = 52
	X86_REG_CR8 = 53
	X86_REG_DR0 = 54
	X86_REG_DR1 = 55
	X86_REG_DR2 = 56
	X86_REG_DR3 = 57
	X86_REG_DR4 = 58
	X86_REG_DR5 = 59
	X86_REG_DR6 = 60
	X86_REG_DR7 = 61
	X86_REG_FP0 = 62
	X86_REG_FP1 = 63
	X86_REG_FP2 = 64
	X86_REG_FP3 = 65
	X86_REG_FP4 = 66
	X86_REG_FP5 = 67
	X86_REG_FP6 = 68
	X86_REG_FP7 = 69
	X86_REG_K0 = 70
	X86_REG_K1 = 71
	X86_REG_K2 = 72
	X86_REG_K3 = 73
	X86_REG_K4 = 74
	X86_REG_K5 = 75
	X86_REG_K6 = 76
	X86_REG_K7 = 77
	X86_REG_MM0 = 78
	X86_REG_MM1 = 79
	X86_REG_MM2 = 80
	X86_REG_MM3 = 81
	X86_REG_MM4 = 82
	X86_REG_MM5 = 83
	X86_REG_MM6 = 84
	X86_REG_MM7 = 85
	X86_REG_R8 = 86
	X86_REG_R9 = 87
	X86_REG_R10 = 88
	X86_REG_R11 = 89
	X86_REG_R12 = 90
	X86_REG_R13 = 91
	X86_REG_R14 = 92
	X86_REG_R15 = 93
	X86_REG_ST0 = 94
	X86_REG_ST1 = 95
	X86_REG_ST2 = 96
	X86_REG_ST3 = 97
	X86_REG_ST4 = 98
	X86_REG_ST5 = 99
	X86_REG_ST6 = 100
	X86_REG_ST7 = 101
	X86_REG_XMM0 = 102
	X86_REG_XMM1 = 103
	X86_REG_XMM2 = 104
	X86_REG_XMM3 = 105
	X86_REG_XMM4 = 106
	X86_REG_XMM5 = 107
	X86_REG_XMM6 = 108
	X86_REG_XMM7 = 109
	X86_REG_XMM8 = 110
	X86_REG_XMM9 = 111
	X86_REG_XMM10 = 112
	X86_REG_XMM11 = 113
	X86_REG_XMM12 = 114
	X86_REG_XMM13 = 115
	X86_REG_XMM14 = 116
	X86_REG_XMM15 = 117
	X86_REG_XMM16 = 118
	X86_REG_XMM17 = 119
	X86_REG_XMM18 = 120
	X86_REG_XMM19 = 121
	X86_REG_XMM20 = 122
	X86_REG_XMM21 = 123
	X86_REG_XMM22 = 124
	X86_REG_XMM23 = 125
	X86_REG_XMM24 = 126
	X86_REG_XMM25 = 127
	X86_REG_XMM26 = 128
	X86_REG_XMM27 = 129
	X86_REG_XMM28 = 130
	X86_REG_XMM29 = 131
	X86_REG_XMM30 = 132
	X86_REG_XMM31 = 133
	X86_REG_YMM0 = 134
	X86_REG_YMM1 = 135
	X86_REG_YMM2 = 136
	X86_REG_YMM3 = 137
	X86_REG_YMM4 = 138
	X86_REG_YMM5 = 139
	X86_REG_YMM6 = 140
	X86_REG_YMM7 = 141
	X86_REG_YMM8 = 142
	X86_REG_YMM9 = 143
	X86_REG_YMM10 = 144
	X86_REG_YMM11 = 145
	X86_REG_YMM12 = 146
	X86_REG_YMM13 = 147
	X86_REG_YMM14 = 148
	X86_REG_YMM15 = 149
	X86_REG_YMM16 = 150
	X86_REG_YMM17 = 151
	X86_REG_YMM18 = 152
	X86_REG_YMM19 = 153
	X86_REG_YMM20 = 154
	X86_REG_YMM21 = 155
	X86_REG_YMM22 = 156
	X86_REG_YMM23 = 157
	X86_REG_YMM24 = 158
	X86_REG_YMM25 = 159
	X86_REG_YMM26 = 160
	X86_REG_YMM27 = 161
	X86_REG_YMM28 = 162
	X86_REG_YMM29 = 163
	X86_REG_YMM30 = 164
	X86_REG_YMM31 = 165
	X86_REG_ZMM0 = 166
	X86_REG_ZMM1 = 167
	X86_REG_ZMM2 = 168
	X86_REG_ZMM3 = 169
	X86_REG_ZMM4 = 170
	X86_REG_ZMM5 = 171
	X86_REG_ZMM6 = 172
	X86_REG_ZMM7 = 173
	X86_REG_ZMM8 = 174
	X86_REG_ZMM9 = 175
	X86_REG_ZMM10 = 176
	X86_REG_ZMM11 = 177
	X86_REG_ZMM12 = 178
	X86_REG_ZMM13 = 179
	X86_REG_ZMM14 = 180
	X86_REG_ZMM15 = 181
	X86_REG_ZMM16 = 182
	X86_REG_ZMM17 = 183
	X86_REG_ZMM18 = 184
	X86_REG_ZMM19 = 185
	X86_REG_ZMM20 = 186
	X86_REG_ZMM21 = 187
	X86_REG_ZMM22 = 188
	X86_REG_ZMM23 = 189
	X86_REG_ZMM24 = 190
	X86_REG_ZMM25 = 191
	X86_REG_ZMM26 = 192
	X86_REG_ZMM27 = 193
	X86_REG_ZMM28 = 194
	X86_REG_ZMM29 = 195
	X86_REG_ZMM30 = 196
	X86_REG_ZMM31 = 197
	X86_REG_R8B = 198
	X86_REG_R9B = 199
	X86_REG_R10B = 200
	X86_REG_R11B = 201
	X86_REG_R12B = 202
	X86_REG_R13B = 203
	X86_REG_R14B = 204
	X86_REG_R15B = 205
	X86_REG_R8D = 206
	X86_REG_R9D = 207
	X86_REG_R10D = 208
	X86_REG_R11D = 209
	X86_REG_R12D = 210
	X86_REG_R13D = 211
	X86_REG_R14D = 212
	X86_REG_R15D = 213
	X86_REG_R8W = 214
	X86_REG_R9W = 215
	X86_REG_R10W = 216
	X86_REG_R11W = 217
	X86_REG_R12W = 218
	X86_REG_R13W = 219
	X86_REG_R14W = 220
	X86_REG_R15W = 221
	X86_REG_IDTR = 222
	X86_REG_GDTR = 223
	X86_REG_LDTR = 224
	X86_REG_TR = 225
	X86_REG_FPCW = 226
	X86_REG_FPTAG = 227
	X86_REG_MSR = 228
	X86_REG_MXCSR = 229
	X86_REG_FS_BASE = 230
	X86_REG_GS_BASE = 231
	X86_REG_FLAGS = 232
	X86_REG_RFLAGS = 233
	X86_REG_FIP = 234
	X86_REG_FCS = 235
	X86_REG_FDP = 236
	X86_REG_FDS = 237
	X86_REG_FOP = 238
	X86_REG_ENDING = 239

// X86 instructions

	X86_INS_INVALID = 0
	X86_INS_AAA = 1
	X86_INS_AAD = 2
	X86_INS_AAM = 3
	X86_INS_AAS = 4
	X86_INS_FABS = 5
	X86_INS_ADC = 6
	X86_INS_ADCX = 7
	X86_INS_ADD = 8
	X86_INS_ADDPD = 9
	X86_INS_ADDPS = 10
	X86_INS_ADDSD = 11
	X86_INS_ADDSS = 12
	X86_INS_ADDSUBPD = 13
	X86_INS_ADDSUBPS = 14
	X86_INS_FADD = 15
	X86_INS_FIADD = 16
	X86_INS_FADDP = 17
	X86_INS_ADOX = 18
	X86_INS_AESDECLAST = 19
	X86_INS_AESDEC = 20
	X86_INS_AESENCLAST = 21
	X86_INS_AESENC = 22
	X86_INS_AESIMC = 23
	X86_INS_AESKEYGENASSIST = 24
	X86_INS_AND = 25
	X86_INS_ANDN = 26
	X86_INS_ANDNPD = 27
	X86_INS_ANDNPS = 28
	X86_INS_ANDPD = 29
	X86_INS_ANDPS = 30
	X86_INS_ARPL = 31
	X86_INS_BEXTR = 32
	X86_INS_BLCFILL = 33
	X86_INS_BLCI = 34
	X86_INS_BLCIC = 35
	X86_INS_BLCMSK = 36
	X86_INS_BLCS = 37
	X86_INS_BLENDPD = 38
	X86_INS_BLENDPS = 39
	X86_INS_BLENDVPD = 40
	X86_INS_BLENDVPS = 41
	X86_INS_BLSFILL = 42
	X86_INS_BLSI = 43
	X86_INS_BLSIC = 44
	X86_INS_BLSMSK = 45
	X86_INS_BLSR = 46
	X86_INS_BOUND = 47
	X86_INS_BSF = 48
	X86_INS_BSR = 49
	X86_INS_BSWAP = 50
	X86_INS_BT = 51
	X86_INS_BTC = 52
	X86_INS_BTR = 53
	X86_INS_BTS = 54
	X86_INS_BZHI = 55
	X86_INS_CALL = 56
	X86_INS_CBW = 57
	X86_INS_CDQ = 58
	X86_INS_CDQE = 59
	X86_INS_FCHS = 60
	X86_INS_CLAC = 61
	X86_INS_CLC = 62
	X86_INS_CLD = 63
	X86_INS_CLFLUSH = 64
	X86_INS_CLFLUSHOPT = 65
	X86_INS_CLGI = 66
	X86_INS_CLI = 67
	X86_INS_CLTS = 68
	X86_INS_CLWB = 69
	X86_INS_CMC = 70
	X86_INS_CMOVA = 71
	X86_INS_CMOVAE = 72
	X86_INS_CMOVB = 73
	X86_INS_CMOVBE = 74
	X86_INS_FCMOVBE = 75
	X86_INS_FCMOVB = 76
	X86_INS_CMOVE = 77
	X86_INS_FCMOVE = 78
	X86_INS_CMOVG = 79
	X86_INS_CMOVGE = 80
	X86_INS_CMOVL = 81
	X86_INS_CMOVLE = 82
	X86_INS_FCMOVNBE = 83
	X86_INS_FCMOVNB = 84
	X86_INS_CMOVNE = 85
	X86_INS_FCMOVNE = 86
	X86_INS_CMOVNO = 87
	X86_INS_CMOVNP = 88
	X86_INS_FCMOVNU = 89
	X86_INS_CMOVNS = 90
	X86_INS_CMOVO = 91
	X86_INS_CMOVP = 92
	X86_INS_FCMOVU = 93
	X86_INS_CMOVS = 94
	X86_INS_CMP = 95
	X86_INS_CMPPD = 96
	X86_INS_CMPPS = 97
	X86_INS_CMPSB = 98
	X86_INS_CMPSD = 99
	X86_INS_CMPSQ = 100
	X86_INS_CMPSS = 101
	X86_INS_CMPSW = 102
	X86_INS_CMPXCHG16B = 103
	X86_INS_CMPXCHG = 104
	X86_INS_CMPXCHG8B = 105
	X86_INS_COMISD = 106
	X86_INS_COMISS = 107
	X86_INS_FCOMP = 108
	X86_INS_FCOMPI = 109
	X86_INS_FCOMI = 110
	X86_INS_FCOM = 111
	X86_INS_FCOS = 112
	X86_INS_CPUID = 113
	X86_INS_CQO = 114
	X86_INS_CRC32 = 115
	X86_INS_CVTDQ2PD = 116
	X86_INS_CVTDQ2PS = 117
	X86_INS_CVTPD2DQ = 118
	X86_INS_CVTPD2PS = 119
	X86_INS_CVTPS2DQ = 120
	X86_INS_CVTPS2PD = 121
	X86_INS_CVTSD2SI = 122
	X86_INS_CVTSD2SS = 123
	X86_INS_CVTSI2SD = 124
	X86_INS_CVTSI2SS = 125
	X86_INS_CVTSS2SD = 126
	X86_INS_CVTSS2SI = 127
	X86_INS_CVTTPD2DQ = 128
	X86_INS_CVTTPS2DQ = 129
	X86_INS_CVTTSD2SI = 130
	X86_INS_CVTTSS2SI = 131
	X86_INS_CWD = 132
	X86_INS_CWDE = 133
	X86_INS_DAA = 134
	X86_INS_DAS = 135
	X86_INS_DATA16 = 136
	X86_INS_DEC = 137
	X86_INS_DIV = 138
	X86_INS_DIVPD = 139
	X86_INS_DIVPS = 140
	X86_INS_FDIVR = 141
	X86_INS_FIDIVR = 142
	X86_INS_FDIVRP = 143
	X86_INS_DIVSD = 144
	X86_INS_DIVSS = 145
	X86_INS_FDIV = 146
	X86_INS_FIDIV = 147
	X86_INS_FDIVP = 148
	X86_INS_DPPD = 149
	X86_INS_DPPS = 150
	X86_INS_RET = 151
	X86_INS_ENCLS = 152
	X86_INS_ENCLU = 153
	X86_INS_ENTER = 154
	X86_INS_EXTRACTPS = 155
	X86_INS_EXTRQ = 156
	X86_INS_F2XM1 = 157
	X86_INS_LCALL = 158
	X86_INS_LJMP = 159
	X86_INS_FBLD = 160
	X86_INS_FBSTP = 161
	X86_INS_FCOMPP = 162
	X86_INS_FDECSTP = 163
	X86_INS_FEMMS = 164
	X86_INS_FFREE = 165
	X86_INS_FICOM = 166
	X86_INS_FICOMP = 167
	X86_INS_FINCSTP = 168
	X86_INS_FLDCW = 169
	X86_INS_FLDENV = 170
	X86_INS_FLDL2E = 171
	X86_INS_FLDL2T = 172
	X86_INS_FLDLG2 = 173
	X86_INS_FLDLN2 = 174
	X86_INS_FLDPI = 175
	X86_INS_FNCLEX = 176
	X86_INS_FNINIT = 177
	X86_INS_FNOP = 178
	X86_INS_FNSTCW = 179
	X86_INS_FNSTSW = 180
	X86_INS_FPATAN = 181
	X86_INS_FPREM = 182
	X86_INS_FPREM1 = 183
	X86_INS_FPTAN = 184
	X86_INS_FFREEP = 185
	X86_INS_FRNDINT = 186
	X86_INS_FRSTOR = 187
	X86_INS_FNSAVE = 188
	X86_INS_FSCALE = 189
	X86_INS_FSETPM = 190
	X86_INS_FSINCOS = 191
	X86_INS_FNSTENV = 192
	X86_INS_FXAM = 193
	X86_INS_FXRSTOR = 194
	X86_INS_FXRSTOR64 = 195
	X86_INS_FXSAVE = 196
	X86_INS_FXSAVE64 = 197
	X86_INS_FXTRACT = 198
	X86_INS_FYL2X = 199
	X86_INS_FYL2XP1 = 200
	X86_INS_MOVAPD = 201
	X86_INS_MOVAPS = 202
	X86_INS_ORPD = 203
	X86_INS_ORPS = 204
	X86_INS_VMOVAPD = 205
	X86_INS_VMOVAPS = 206
	X86_INS_XORPD = 207
	X86_INS_XORPS = 208
	X86_INS_GETSEC = 209
	X86_INS_HADDPD = 210
	X86_INS_HADDPS = 211
	X86_INS_HLT = 212
	X86_INS_HSUBPD = 213
	X86_INS_HSUBPS = 214
	X86_INS_IDIV = 215
	X86_INS_FILD = 216
	X86_INS_IMUL = 217
	X86_INS_IN = 218
	X86_INS_INC = 219
	X86_INS_INSB = 220
	X86_INS_INSERTPS = 221
	X86_INS_INSERTQ = 222
	X86_INS_INSD = 223
	X86_INS_INSW = 224
	X86_INS_INT = 225
	X86_INS_INT1 = 226
	X86_INS_INT3 = 227
	X86_INS_INTO = 228
	X86_INS_INVD = 229
	X86_INS_INVEPT = 230
	X86_INS_INVLPG = 231
	X86_INS_INVLPGA = 232
	X86_INS_INVPCID = 233
	X86_INS_INVVPID = 234
	X86_INS_IRET = 235
	X86_INS_IRETD = 236
	X86_INS_IRETQ = 237
	X86_INS_FISTTP = 238
	X86_INS_FIST = 239
	X86_INS_FISTP = 240
	X86_INS_UCOMISD = 241
	X86_INS_UCOMISS = 242
	X86_INS_VCOMISD = 243
	X86_INS_VCOMISS = 244
	X86_INS_VCVTSD2SS = 245
	X86_INS_VCVTSI2SD = 246
	X86_INS_VCVTSI2SS = 247
	X86_INS_VCVTSS2SD = 248
	X86_INS_VCVTTSD2SI = 249
	X86_INS_VCVTTSD2USI = 250
	X86_INS_VCVTTSS2SI = 251
	X86_INS_VCVTTSS2USI = 252
	X86_INS_VCVTUSI2SD = 253
	X86_INS_VCVTUSI2SS = 254
	X86_INS_VUCOMISD = 255
	X86_INS_VUCOMISS = 256
	X86_INS_JAE = 257
	X86_INS_JA = 258
	X86_INS_JBE = 259
	X86_INS_JB = 260
	X86_INS_JCXZ = 261
	X86_INS_JECXZ = 262
	X86_INS_JE = 263
	X86_INS_JGE = 264
	X86_INS_JG = 265
	X86_INS_JLE = 266
	X86_INS_JL = 267
	X86_INS_JMP = 268
	X86_INS_JNE = 269
	X86_INS_JNO = 270
	X86_INS_JNP = 271
	X86_INS_JNS = 272
	X86_INS_JO = 273
	X86_INS_JP = 274
	X86_INS_JRCXZ = 275
	X86_INS_JS = 276
	X86_INS_KANDB = 277
	X86_INS_KANDD = 278
	X86_INS_KANDNB = 279
	X86_INS_KANDND = 280
	X86_INS_KANDNQ = 281
	X86_INS_KANDNW = 282
	X86_INS_KANDQ = 283
	X86_INS_KANDW = 284
	X86_INS_KMOVB = 285
	X86_INS_KMOVD = 286
	X86_INS_KMOVQ = 287
	X86_INS_KMOVW = 288
	X86_INS_KNOTB = 289
	X86_INS_KNOTD = 290
	X86_INS_KNOTQ = 291
	X86_INS_KNOTW = 292
	X86_INS_KORB = 293
	X86_INS_KORD = 294
	X86_INS_KORQ = 295
	X86_INS_KORTESTB = 296
	X86_INS_KORTESTD = 297
	X86_INS_KORTESTQ = 298
	X86_INS_KORTESTW = 299
	X86_INS_KORW = 300
	X86_INS_KSHIFTLB = 301
	X86_INS_KSHIFTLD = 302
	X86_INS_KSHIFTLQ = 303
	X86_INS_KSHIFTLW = 304
	X86_INS_KSHIFTRB = 305
	X86_INS_KSHIFTRD = 306
	X86_INS_KSHIFTRQ = 307
	X86_INS_KSHIFTRW = 308
	X86_INS_KUNPCKBW = 309
	X86_INS_KXNORB = 310
	X86_INS_KXNORD = 311
	X86_INS_KXNORQ = 312
	X86_INS_KXNORW = 313
	X86_INS_KXORB = 314
	X86_INS_KXORD = 315
	X86_INS_KXORQ = 316
	X86_INS_KXORW = 317
	X86_INS_LAHF = 318
	X86_INS_LAR = 319
	X86_INS_LDDQU = 320
	X86_INS_LDMXCSR = 321
	X86_INS_LDS = 322
	X86_INS_FLDZ = 323
	X86_INS_FLD1 = 324
	X86_INS_FLD = 325
	X86_INS_LEA = 326
	X86_INS_LEAVE = 327
	X86_INS_LES = 328
	X86_INS_LFENCE = 329
	X86_INS_LFS = 330
	X86_INS_LGDT = 331
	X86_INS_LGS = 332
	X86_INS_LIDT = 333
	X86_INS_LLDT = 334
	X86_INS_LMSW = 335
	X86_INS_OR = 336
	X86_INS_SUB = 337
	X86_INS_XOR = 338
	X86_INS_LODSB = 339
	X86_INS_LODSD = 340
	X86_INS_LODSQ = 341
	X86_INS_LODSW = 342
	X86_INS_LOOP = 343
	X86_INS_LOOPE = 344
	X86_INS_LOOPNE = 345
	X86_INS_RETF = 346
	X86_INS_RETFQ = 347
	X86_INS_LSL = 348
	X86_INS_LSS = 349
	X86_INS_LTR = 350
	X86_INS_XADD = 351
	X86_INS_LZCNT = 352
	X86_INS_MASKMOVDQU = 353
	X86_INS_MAXPD = 354
	X86_INS_MAXPS = 355
	X86_INS_MAXSD = 356
	X86_INS_MAXSS = 357
	X86_INS_MFENCE = 358
	X86_INS_MINPD = 359
	X86_INS_MINPS = 360
	X86_INS_MINSD = 361
	X86_INS_MINSS = 362
	X86_INS_CVTPD2PI = 363
	X86_INS_CVTPI2PD = 364
	X86_INS_CVTPI2PS = 365
	X86_INS_CVTPS2PI = 366
	X86_INS_CVTTPD2PI = 367
	X86_INS_CVTTPS2PI = 368
	X86_INS_EMMS = 369
	X86_INS_MASKMOVQ = 370
	X86_INS_MOVD = 371
	X86_INS_MOVDQ2Q = 372
	X86_INS_MOVNTQ = 373
	X86_INS_MOVQ2DQ = 374
	X86_INS_MOVQ = 375
	X86_INS_PABSB = 376
	X86_INS_PABSD = 377
	X86_INS_PABSW = 378
	X86_INS_PACKSSDW = 379
	X86_INS_PACKSSWB = 380
	X86_INS_PACKUSWB = 381
	X86_INS_PADDB = 382
	X86_INS_PADDD = 383
	X86_INS_PADDQ = 384
	X86_INS_PADDSB = 385
	X86_INS_PADDSW = 386
	X86_INS_PADDUSB = 387
	X86_INS_PADDUSW = 388
	X86_INS_PADDW = 389
	X86_INS_PALIGNR = 390
	X86_INS_PANDN = 391
	X86_INS_PAND = 392
	X86_INS_PAVGB = 393
	X86_INS_PAVGW = 394
	X86_INS_PCMPEQB = 395
	X86_INS_PCMPEQD = 396
	X86_INS_PCMPEQW = 397
	X86_INS_PCMPGTB = 398
	X86_INS_PCMPGTD = 399
	X86_INS_PCMPGTW = 400
	X86_INS_PEXTRW = 401
	X86_INS_PHADDSW = 402
	X86_INS_PHADDW = 403
	X86_INS_PHADDD = 404
	X86_INS_PHSUBD = 405
	X86_INS_PHSUBSW = 406
	X86_INS_PHSUBW = 407
	X86_INS_PINSRW = 408
	X86_INS_PMADDUBSW = 409
	X86_INS_PMADDWD = 410
	X86_INS_PMAXSW = 411
	X86_INS_PMAXUB = 412
	X86_INS_PMINSW = 413
	X86_INS_PMINUB = 414
	X86_INS_PMOVMSKB = 415
	X86_INS_PMULHRSW = 416
	X86_INS_PMULHUW = 417
	X86_INS_PMULHW = 418
	X86_INS_PMULLW = 419
	X86_INS_PMULUDQ = 420
	X86_INS_POR = 421
	X86_INS_PSADBW = 422
	X86_INS_PSHUFB = 423
	X86_INS_PSHUFW = 424
	X86_INS_PSIGNB = 425
	X86_INS_PSIGND = 426
	X86_INS_PSIGNW = 427
	X86_INS_PSLLD = 428
	X86_INS_PSLLQ = 429
	X86_INS_PSLLW = 430
	X86_INS_PSRAD = 431
	X86_INS_PSRAW = 432
	X86_INS_PSRLD = 433
	X86_INS_PSRLQ = 434
	X86_INS_PSRLW = 435
	X86_INS_PSUBB = 436
	X86_INS_PSUBD = 437
	X86_INS_PSUBQ = 438
	X86_INS_PSUBSB = 439
	X86_INS_PSUBSW = 440
	X86_INS_PSUBUSB = 441
	X86_INS_PSUBUSW = 442
	X86_INS_PSUBW = 443
	X86_INS_PUNPCKHBW = 444
	X86_INS_PUNPCKHDQ = 445
	X86_INS_PUNPCKHWD = 446
	X86_INS_PUNPCKLBW = 447
	X86_INS_PUNPCKLDQ = 448
	X86_INS_PUNPCKLWD = 449
	X86_INS_PXOR = 450
	X86_INS_MONITOR = 451
	X86_INS_MONTMUL = 452
	X86_INS_MOV = 453
	X86_INS_MOVABS = 454
	X86_INS_MOVBE = 455
	X86_INS_MOVDDUP = 456
	X86_INS_MOVDQA = 457
	X86_INS_MOVDQU = 458
	X86_INS_MOVHLPS = 459
	X86_INS_MOVHPD = 460
	X86_INS_MOVHPS = 461
	X86_INS_MOVLHPS = 462
	X86_INS_MOVLPD = 463
	X86_INS_MOVLPS = 464
	X86_INS_MOVMSKPD = 465
	X86_INS_MOVMSKPS = 466
	X86_INS_MOVNTDQA = 467
	X86_INS_MOVNTDQ = 468
	X86_INS_MOVNTI = 469
	X86_INS_MOVNTPD = 470
	X86_INS_MOVNTPS = 471
	X86_INS_MOVNTSD = 472
	X86_INS_MOVNTSS = 473
	X86_INS_MOVSB = 474
	X86_INS_MOVSD = 475
	X86_INS_MOVSHDUP = 476
	X86_INS_MOVSLDUP = 477
	X86_INS_MOVSQ = 478
	X86_INS_MOVSS = 479
	X86_INS_MOVSW = 480
	X86_INS_MOVSX = 481
	X86_INS_MOVSXD = 482
	X86_INS_MOVUPD = 483
	X86_INS_MOVUPS = 484
	X86_INS_MOVZX = 485
	X86_INS_MPSADBW = 486
	X86_INS_MUL = 487
	X86_INS_MULPD = 488
	X86_INS_MULPS = 489
	X86_INS_MULSD = 490
	X86_INS_MULSS = 491
	X86_INS_MULX = 492
	X86_INS_FMUL = 493
	X86_INS_FIMUL = 494
	X86_INS_FMULP = 495
	X86_INS_MWAIT = 496
	X86_INS_NEG = 497
	X86_INS_NOP = 498
	X86_INS_NOT = 499
	X86_INS_OUT = 500
	X86_INS_OUTSB = 501
	X86_INS_OUTSD = 502
	X86_INS_OUTSW = 503
	X86_INS_PACKUSDW = 504
	X86_INS_PAUSE = 505
	X86_INS_PAVGUSB = 506
	X86_INS_PBLENDVB = 507
	X86_INS_PBLENDW = 508
	X86_INS_PCLMULQDQ = 509
	X86_INS_PCMPEQQ = 510
	X86_INS_PCMPESTRI = 511
	X86_INS_PCMPESTRM = 512
	X86_INS_PCMPGTQ = 513
	X86_INS_PCMPISTRI = 514
	X86_INS_PCMPISTRM = 515
	X86_INS_PCOMMIT = 516
	X86_INS_PDEP = 517
	X86_INS_PEXT = 518
	X86_INS_PEXTRB = 519
	X86_INS_PEXTRD = 520
	X86_INS_PEXTRQ = 521
	X86_INS_PF2ID = 522
	X86_INS_PF2IW = 523
	X86_INS_PFACC = 524
	X86_INS_PFADD = 525
	X86_INS_PFCMPEQ = 526
	X86_INS_PFCMPGE = 527
	X86_INS_PFCMPGT = 528
	X86_INS_PFMAX = 529
	X86_INS_PFMIN = 530
	X86_INS_PFMUL = 531
	X86_INS_PFNACC = 532
	X86_INS_PFPNACC = 533
	X86_INS_PFRCPIT1 = 534
	X86_INS_PFRCPIT2 = 535
	X86_INS_PFRCP = 536
	X86_INS_PFRSQIT1 = 537
	X86_INS_PFRSQRT = 538
	X86_INS_PFSUBR = 539
	X86_INS_PFSUB = 540
	X86_INS_PHMINPOSUW = 541
	X86_INS_PI2FD = 542
	X86_INS_PI2FW = 543
	X86_INS_PINSRB = 544
	X86_INS_PINSRD = 545
	X86_INS_PINSRQ = 546
	X86_INS_PMAXSB = 547
	X86_INS_PMAXSD = 548
	X86_INS_PMAXUD = 549
	X86_INS_PMAXUW = 550
	X86_INS_PMINSB = 551
	X86_INS_PMINSD = 552
	X86_INS_PMINUD = 553
	X86_INS_PMINUW = 554
	X86_INS_PMOVSXBD = 555
	X86_INS_PMOVSXBQ = 556
	X86_INS_PMOVSXBW = 557
	X86_INS_PMOVSXDQ = 558
	X86_INS_PMOVSXWD = 559
	X86_INS_PMOVSXWQ = 560
	X86_INS_PMOVZXBD = 561
	X86_INS_PMOVZXBQ = 562
	X86_INS_PMOVZXBW = 563
	X86_INS_PMOVZXDQ = 564
	X86_INS_PMOVZXWD = 565
	X86_INS_PMOVZXWQ = 566
	X86_INS_PMULDQ = 567
	X86_INS_PMULHRW = 568
	X86_INS_PMULLD = 569
	X86_INS_POP = 570
	X86_INS_POPAW = 571
	X86_INS_POPAL = 572
	X86_INS_POPCNT = 573
	X86_INS_POPF = 574
	X86_INS_POPFD = 575
	X86_INS_POPFQ = 576
	X86_INS_PREFETCH = 577
	X86_INS_PREFETCHNTA = 578
	X86_INS_PREFETCHT0 = 579
	X86_INS_PREFETCHT1 = 580
	X86_INS_PREFETCHT2 = 581
	X86_INS_PREFETCHW = 582
	X86_INS_PSHUFD = 583
	X86_INS_PSHUFHW = 584
	X86_INS_PSHUFLW = 585
	X86_INS_PSLLDQ = 586
	X86_INS_PSRLDQ = 587
	X86_INS_PSWAPD = 588
	X86_INS_PTEST = 589
	X86_INS_PUNPCKHQDQ = 590
	X86_INS_PUNPCKLQDQ = 591
	X86_INS_PUSH = 592
	X86_INS_PUSHAW = 593
	X86_INS_PUSHAL = 594
	X86_INS_PUSHF = 595
	X86_INS_PUSHFD = 596
	X86_INS_PUSHFQ = 597
	X86_INS_RCL = 598
	X86_INS_RCPPS = 599
	X86_INS_RCPSS = 600
	X86_INS_RCR = 601
	X86_INS_RDFSBASE = 602
	X86_INS_RDGSBASE = 603
	X86_INS_RDMSR = 604
	X86_INS_RDPMC = 605
	X86_INS_RDRAND = 606
	X86_INS_RDSEED = 607
	X86_INS_RDTSC = 608
	X86_INS_RDTSCP = 609
	X86_INS_ROL = 610
	X86_INS_ROR = 611
	X86_INS_RORX = 612
	X86_INS_ROUNDPD = 613
	X86_INS_ROUNDPS = 614
	X86_INS_ROUNDSD = 615
	X86_INS_ROUNDSS = 616
	X86_INS_RSM = 617
	X86_INS_RSQRTPS = 618
	X86_INS_RSQRTSS = 619
	X86_INS_SAHF = 620
	X86_INS_SAL = 621
	X86_INS_SALC = 622
	X86_INS_SAR = 623
	X86_INS_SARX = 624
	X86_INS_SBB = 625
	X86_INS_SCASB = 626
	X86_INS_SCASD = 627
	X86_INS_SCASQ = 628
	X86_INS_SCASW = 629
	X86_INS_SETAE = 630
	X86_INS_SETA = 631
	X86_INS_SETBE = 632
	X86_INS_SETB = 633
	X86_INS_SETE = 634
	X86_INS_SETGE = 635
	X86_INS_SETG = 636
	X86_INS_SETLE = 637
	X86_INS_SETL = 638
	X86_INS_SETNE = 639
	X86_INS_SETNO = 640
	X86_INS_SETNP = 641
	X86_INS_SETNS = 642
	X86_INS_SETO = 643
	X86_INS_SETP = 644
	X86_INS_SETS = 645
	X86_INS_SFENCE = 646
	X86_INS_SGDT = 647
	X86_INS_SHA1MSG1 = 648
	X86_INS_SHA1MSG2 = 649
	X86_INS_SHA1NEXTE = 650
	X86_INS_SHA1RNDS4 = 651
	X86_INS_SHA256MSG1 = 652
	X86_INS_SHA256MSG2 = 653
	X86_INS_SHA256RNDS2 = 654
	X86_INS_SHL = 655
	X86_INS_SHLD = 656
	X86_INS_SHLX = 657
	X86_INS_SHR = 658
	X86_INS_SHRD = 659
	X86_INS_SHRX = 660
	X86_INS_SHUFPD = 661
	X86_INS_SHUFPS = 662
	X86_INS_SIDT = 663
	X86_INS_FSIN = 664
	X86_INS_SKINIT = 665
	X86_INS_SLDT = 666
	X86_INS_SMSW = 667
	X86_INS_SQRTPD = 668
	X86_INS_SQRTPS = 669
	X86_INS_SQRTSD = 670
	X86_INS_SQRTSS = 671
	X86_INS_FSQRT = 672
	X86_INS_STAC = 673
	X86_INS_STC = 674
	X86_INS_STD = 675
	X86_INS_STGI = 676
	X86_INS_STI = 677
	X86_INS_STMXCSR = 678
	X86_INS_STOSB = 679
	X86_INS_STOSD = 680
	X86_INS_STOSQ = 681
	X86_INS_STOSW = 682
	X86_INS_STR = 683
	X86_INS_FST = 684
	X86_INS_FSTP = 685
	X86_INS_FSTPNCE = 686
	X86_INS_FXCH = 687
	X86_INS_SUBPD = 688
	X86_INS_SUBPS = 689
	X86_INS_FSUBR = 690
	X86_INS_FISUBR = 691
	X86_INS_FSUBRP = 692
	X86_INS_SUBSD = 693
	X86_INS_SUBSS = 694
	X86_INS_FSUB = 695
	X86_INS_FISUB = 696
	X86_INS_FSUBP = 697
	X86_INS_SWAPGS = 698
	X86_INS_SYSCALL = 699
	X86_INS_SYSENTER = 700
	X86_INS_SYSEXIT = 701
	X86_INS_SYSRET = 702
	X86_INS_T1MSKC = 703
	X86_INS_TEST = 704
	X86_INS_UD2 = 705
	X86_INS_FTST = 706
	X86_INS_TZCNT = 707
	X86_INS_TZMSK = 708
	X86_INS_FUCOMPI = 709
	X86_INS_FUCOMI = 710
	X86_INS_FUCOMPP = 711
	X86_INS_FUCOMP = 712
	X86_INS_FUCOM = 713
	X86_INS_UD2B = 714
	X86_INS_UNPCKHPD = 715
	X86_INS_UNPCKHPS = 716
	X86_INS_UNPCKLPD = 717
	X86_INS_UNPCKLPS = 718
	X86_INS_VADDPD = 719
	X86_INS_VADDPS = 720
	X86_INS_VADDSD = 721
	X86_INS_VADDSS = 722
	X86_INS_VADDSUBPD = 723
	X86_INS_VADDSUBPS = 724
	X86_INS_VAESDECLAST = 725
	X86_INS_VAESDEC = 726
	X86_INS_VAESENCLAST = 727
	X86_INS_VAESENC = 728
	X86_INS_VAESIMC = 729
	X86_INS_VAESKEYGENASSIST = 730
	X86_INS_VALIGND = 731
	X86_INS_VALIGNQ = 732
	X86_INS_VANDNPD = 733
	X86_INS_VANDNPS = 734
	X86_INS_VANDPD = 735
	X86_INS_VANDPS = 736
	X86_INS_VBLENDMPD = 737
	X86_INS_VBLENDMPS = 738
	X86_INS_VBLENDPD = 739
	X86_INS_VBLENDPS = 740
	X86_INS_VBLENDVPD = 741
	X86_INS_VBLENDVPS = 742
	X86_INS_VBROADCASTF128 = 743
	X86_INS_VBROADCASTI32X4 = 744
	X86_INS_VBROADCASTI64X4 = 745
	X86_INS_VBROADCASTSD = 746
	X86_INS_VBROADCASTSS = 747
	X86_INS_VCMPPD = 748
	X86_INS_VCMPPS = 749
	X86_INS_VCMPSD = 750
	X86_INS_VCMPSS = 751
	X86_INS_VCOMPRESSPD = 752
	X86_INS_VCOMPRESSPS = 753
	X86_INS_VCVTDQ2PD = 754
	X86_INS_VCVTDQ2PS = 755
	X86_INS_VCVTPD2DQX = 756
	X86_INS_VCVTPD2DQ = 757
	X86_INS_VCVTPD2PSX = 758
	X86_INS_VCVTPD2PS = 759
	X86_INS_VCVTPD2UDQ = 760
	X86_INS_VCVTPH2PS = 761
	X86_INS_VCVTPS2DQ = 762
	X86_INS_VCVTPS2PD = 763
	X86_INS_VCVTPS2PH = 764
	X86_INS_VCVTPS2UDQ = 765
	X86_INS_VCVTSD2SI = 766
	X86_INS_VCVTSD2USI = 767
	X86_INS_VCVTSS2SI = 768
	X86_INS_VCVTSS2USI = 769
	X86_INS_VCVTTPD2DQX = 770
	X86_INS_VCVTTPD2DQ = 771
	X86_INS_VCVTTPD2UDQ = 772
	X86_INS_VCVTTPS2DQ = 773
	X86_INS_VCVTTPS2UDQ = 774
	X86_INS_VCVTUDQ2PD = 775
	X86_INS_VCVTUDQ2PS = 776
	X86_INS_VDIVPD = 777
	X86_INS_VDIVPS = 778
	X86_INS_VDIVSD = 779
	X86_INS_VDIVSS = 780
	X86_INS_VDPPD = 781
	X86_INS_VDPPS = 782
	X86_INS_VERR = 783
	X86_INS_VERW = 784
	X86_INS_VEXP2PD = 785
	X86_INS_VEXP2PS = 786
	X86_INS_VEXPANDPD = 787
	X86_INS_VEXPANDPS = 788
	X86_INS_VEXTRACTF128 = 789
	X86_INS_VEXTRACTF32X4 = 790
	X86_INS_VEXTRACTF64X4 = 791
	X86_INS_VEXTRACTI128 = 792
	X86_INS_VEXTRACTI32X4 = 793
	X86_INS_VEXTRACTI64X4 = 794
	X86_INS_VEXTRACTPS = 795
	X86_INS_VFMADD132PD = 796
	X86_INS_VFMADD132PS = 797
	X86_INS_VFMADDPD = 798
	X86_INS_VFMADD213PD = 799
	X86_INS_VFMADD231PD = 800
	X86_INS_VFMADDPS = 801
	X86_INS_VFMADD213PS = 802
	X86_INS_VFMADD231PS = 803
	X86_INS_VFMADDSD = 804
	X86_INS_VFMADD213SD = 805
	X86_INS_VFMADD132SD = 806
	X86_INS_VFMADD231SD = 807
	X86_INS_VFMADDSS = 808
	X86_INS_VFMADD213SS = 809
	X86_INS_VFMADD132SS = 810
	X86_INS_VFMADD231SS = 811
	X86_INS_VFMADDSUB132PD = 812
	X86_INS_VFMADDSUB132PS = 813
	X86_INS_VFMADDSUBPD = 814
	X86_INS_VFMADDSUB213PD = 815
	X86_INS_VFMADDSUB231PD = 816
	X86_INS_VFMADDSUBPS = 817
	X86_INS_VFMADDSUB213PS = 818
	X86_INS_VFMADDSUB231PS = 819
	X86_INS_VFMSUB132PD = 820
	X86_INS_VFMSUB132PS = 821
	X86_INS_VFMSUBADD132PD = 822
	X86_INS_VFMSUBADD132PS = 823
	X86_INS_VFMSUBADDPD = 824
	X86_INS_VFMSUBADD213PD = 825
	X86_INS_VFMSUBADD231PD = 826
	X86_INS_VFMSUBADDPS = 827
	X86_INS_VFMSUBADD213PS = 828
	X86_INS_VFMSUBADD231PS = 829
	X86_INS_VFMSUBPD = 830
	X86_INS_VFMSUB213PD = 831
	X86_INS_VFMSUB231PD = 832
	X86_INS_VFMSUBPS = 833
	X86_INS_VFMSUB213PS = 834
	X86_INS_VFMSUB231PS = 835
	X86_INS_VFMSUBSD = 836
	X86_INS_VFMSUB213SD = 837
	X86_INS_VFMSUB132SD = 838
	X86_INS_VFMSUB231SD = 839
	X86_INS_VFMSUBSS = 840
	X86_INS_VFMSUB213SS = 841
	X86_INS_VFMSUB132SS = 842
	X86_INS_VFMSUB231SS = 843
	X86_INS_VFNMADD132PD = 844
	X86_INS_VFNMADD132PS = 845
	X86_INS_VFNMADDPD = 846
	X86_INS_VFNMADD213PD = 847
	X86_INS_VFNMADD231PD = 848
	X86_INS_VFNMADDPS = 849
	X86_INS_VFNMADD213PS = 850
	X86_INS_VFNMADD231PS = 851
	X86_INS_VFNMADDSD = 852
	X86_INS_VFNMADD213SD = 853
	X86_INS_VFNMADD132SD = 854
	X86_INS_VFNMADD231SD = 855
	X86_INS_VFNMADDSS = 856
	X86_INS_VFNMADD213SS = 857
	X86_INS_VFNMADD132SS = 858
	X86_INS_VFNMADD231SS = 859
	X86_INS_VFNMSUB132PD = 860
	X86_INS_VFNMSUB132PS = 861
	X86_INS_VFNMSUBPD = 862
	X86_INS_VFNMSUB213PD = 863
	X86_INS_VFNMSUB231PD = 864
	X86_INS_VFNMSUBPS = 865
	X86_INS_VFNMSUB213PS = 866
	X86_INS_VFNMSUB231PS = 867
	X86_INS_VFNMSUBSD = 868
	X86_INS_VFNMSUB213SD = 869
	X86_INS_VFNMSUB132SD = 870
	X86_INS_VFNMSUB231SD = 871
	X86_INS_VFNMSUBSS = 872
	X86_INS_VFNMSUB213SS = 873
	X86_INS_VFNMSUB132SS = 874
	X86_INS_VFNMSUB231SS = 875
	X86_INS_VFRCZPD = 876
	X86_INS_VFRCZPS = 877
	X86_INS_VFRCZSD = 878
	X86_INS_VFRCZSS = 879
	X86_INS_VORPD = 880
	X86_INS_VORPS = 881
	X86_INS_VXORPD = 882
	X86_INS_VXORPS = 883
	X86_INS_VGATHERDPD = 884
	X86_INS_VGATHERDPS = 885
	X86_INS_VGATHERPF0DPD = 886
	X86_INS_VGATHERPF0DPS = 887
	X86_INS_VGATHERPF0QPD = 888
	X86_INS_VGATHERPF0QPS = 889
	X86_INS_VGATHERPF1DPD = 890
	X86_INS_VGATHERPF1DPS = 891
	X86_INS_VGATHERPF1QPD = 892
	X86_INS_VGATHERPF1QPS = 893
	X86_INS_VGATHERQPD = 894
	X86_INS_VGATHERQPS = 895
	X86_INS_VHADDPD = 896
	X86_INS_VHADDPS = 897
	X86_INS_VHSUBPD = 898
	X86_INS_VHSUBPS = 899
	X86_INS_VINSERTF128 = 900
	X86_INS_VINSERTF32X4 = 901
	X86_INS_VINSERTF32X8 = 902
	X86_INS_VINSERTF64X2 = 903
	X86_INS_VINSERTF64X4 = 904
	X86_INS_VINSERTI128 = 905
	X86_INS_VINSERTI32X4 = 906
	X86_INS_VINSERTI32X8 = 907
	X86_INS_VINSERTI64X2 = 908
	X86_INS_VINSERTI64X4 = 909
	X86_INS_VINSERTPS = 910
	X86_INS_VLDDQU = 911
	X86_INS_VLDMXCSR = 912
	X86_INS_VMASKMOVDQU = 913
	X86_INS_VMASKMOVPD = 914
	X86_INS_VMASKMOVPS = 915
	X86_INS_VMAXPD = 916
	X86_INS_VMAXPS = 917
	X86_INS_VMAXSD = 918
	X86_INS_VMAXSS = 919
	X86_INS_VMCALL = 920
	X86_INS_VMCLEAR = 921
	X86_INS_VMFUNC = 922
	X86_INS_VMINPD = 923
	X86_INS_VMINPS = 924
	X86_INS_VMINSD = 925
	X86_INS_VMINSS = 926
	X86_INS_VMLAUNCH = 927
	X86_INS_VMLOAD = 928
	X86_INS_VMMCALL = 929
	X86_INS_VMOVQ = 930
	X86_INS_VMOVDDUP = 931
	X86_INS_VMOVD = 932
	X86_INS_VMOVDQA32 = 933
	X86_INS_VMOVDQA64 = 934
	X86_INS_VMOVDQA = 935
	X86_INS_VMOVDQU16 = 936
	X86_INS_VMOVDQU32 = 937
	X86_INS_VMOVDQU64 = 938
	X86_INS_VMOVDQU8 = 939
	X86_INS_VMOVDQU = 940
	X86_INS_VMOVHLPS = 941
	X86_INS_VMOVHPD = 942
	X86_INS_VMOVHPS = 943
	X86_INS_VMOVLHPS = 944
	X86_INS_VMOVLPD = 945
	X86_INS_VMOVLPS = 946
	X86_INS_VMOVMSKPD = 947
	X86_INS_VMOVMSKPS = 948
	X86_INS_VMOVNTDQA = 949
	X86_INS_VMOVNTDQ = 950
	X86_INS_VMOVNTPD = 951
	X86_INS_VMOVNTPS = 952
	X86_INS_VMOVSD = 953
	X86_INS_VMOVSHDUP = 954
	X86_INS_VMOVSLDUP = 955
	X86_INS_VMOVSS = 956
	X86_INS_VMOVUPD = 957
	X86_INS_VMOVUPS = 958
	X86_INS_VMPSADBW = 959
	X86_INS_VMPTRLD = 960
	X86_INS_VMPTRST = 961
	X86_INS_VMREAD = 962
	X86_INS_VMRESUME = 963
	X86_INS_VMRUN = 964
	X86_INS_VMSAVE = 965
	X86_INS_VMULPD = 966
	X86_INS_VMULPS = 967
	X86_INS_VMULSD = 968
	X86_INS_VMULSS = 969
	X86_INS_VMWRITE = 970
	X86_INS_VMXOFF = 971
	X86_INS_VMXON = 972
	X86_INS_VPABSB = 973
	X86_INS_VPABSD = 974
	X86_INS_VPABSQ = 975
	X86_INS_VPABSW = 976
	X86_INS_VPACKSSDW = 977
	X86_INS_VPACKSSWB = 978
	X86_INS_VPACKUSDW = 979
	X86_INS_VPACKUSWB = 980
	X86_INS_VPADDB = 981
	X86_INS_VPADDD = 982
	X86_INS_VPADDQ = 983
	X86_INS_VPADDSB = 984
	X86_INS_VPADDSW = 985
	X86_INS_VPADDUSB = 986
	X86_INS_VPADDUSW = 987
	X86_INS_VPADDW = 988
	X86_INS_VPALIGNR = 989
	X86_INS_VPANDD = 990
	X86_INS_VPANDND = 991
	X86_INS_VPANDNQ = 992
	X86_INS_VPANDN = 993
	X86_INS_VPANDQ = 994
	X86_INS_VPAND = 995
	X86_INS_VPAVGB = 996
	X86_INS_VPAVGW = 997
	X86_INS_VPBLENDD = 998
	X86_INS_VPBLENDMB = 999
	X86_INS_VPBLENDMD = 1000
	X86_INS_VPBLENDMQ = 1001
	X86_INS_VPBLENDMW = 1002
	X86_INS_VPBLENDVB = 1003
	X86_INS_VPBLENDW = 1004
	X86_INS_VPBROADCASTB = 1005
	X86_INS_VPBROADCASTD = 1006
	X86_INS_VPBROADCASTMB2Q = 1007
	X86_INS_VPBROADCASTMW2D = 1008
	X86_INS_VPBROADCASTQ = 1009
	X86_INS_VPBROADCASTW = 1010
	X86_INS_VPCLMULQDQ = 1011
	X86_INS_VPCMOV = 1012
	X86_INS_VPCMPB = 1013
	X86_INS_VPCMPD = 1014
	X86_INS_VPCMPEQB = 1015
	X86_INS_VPCMPEQD = 1016
	X86_INS_VPCMPEQQ = 1017
	X86_INS_VPCMPEQW = 1018
	X86_INS_VPCMPESTRI = 1019
	X86_INS_VPCMPESTRM = 1020
	X86_INS_VPCMPGTB = 1021
	X86_INS_VPCMPGTD = 1022
	X86_INS_VPCMPGTQ = 1023
	X86_INS_VPCMPGTW = 1024
	X86_INS_VPCMPISTRI = 1025
	X86_INS_VPCMPISTRM = 1026
	X86_INS_VPCMPQ = 1027
	X86_INS_VPCMPUB = 1028
	X86_INS_VPCMPUD = 1029
	X86_INS_VPCMPUQ = 1030
	X86_INS_VPCMPUW = 1031
	X86_INS_VPCMPW = 1032
	X86_INS_VPCOMB = 1033
	X86_INS_VPCOMD = 1034
	X86_INS_VPCOMPRESSD = 1035
	X86_INS_VPCOMPRESSQ = 1036
	X86_INS_VPCOMQ = 1037
	X86_INS_VPCOMUB = 1038
	X86_INS_VPCOMUD = 1039
	X86_INS_VPCOMUQ = 1040
	X86_INS_VPCOMUW = 1041
	X86_INS_VPCOMW = 1042
	X86_INS_VPCONFLICTD = 1043
	X86_INS_VPCONFLICTQ = 1044
	X86_INS_VPERM2F128 = 1045
	X86_INS_VPERM2I128 = 1046
	X86_INS_VPERMD = 1047
	X86_INS_VPERMI2D = 1048
	X86_INS_VPERMI2PD = 1049
	X86_INS_VPERMI2PS = 1050
	X86_INS_VPERMI2Q = 1051
	X86_INS_VPERMIL2PD = 1052
	X86_INS_VPERMIL2PS = 1053
	X86_INS_VPERMILPD = 1054
	X86_INS_VPERMILPS = 1055
	X86_INS_VPERMPD = 1056
	X86_INS_VPERMPS = 1057
	X86_INS_VPERMQ = 1058
	X86_INS_VPERMT2D = 1059
	X86_INS_VPERMT2PD = 1060
	X86_INS_VPERMT2PS = 1061
	X86_INS_VPERMT2Q = 1062
	X86_INS_VPEXPANDD = 1063
	X86_INS_VPEXPANDQ = 1064
	X86_INS_VPEXTRB = 1065
	X86_INS_VPEXTRD = 1066
	X86_INS_VPEXTRQ = 1067
	X86_INS_VPEXTRW = 1068
	X86_INS_VPGATHERDD = 1069
	X86_INS_VPGATHERDQ = 1070
	X86_INS_VPGATHERQD = 1071
	X86_INS_VPGATHERQQ = 1072
	X86_INS_VPHADDBD = 1073
	X86_INS_VPHADDBQ = 1074
	X86_INS_VPHADDBW = 1075
	X86_INS_VPHADDDQ = 1076
	X86_INS_VPHADDD = 1077
	X86_INS_VPHADDSW = 1078
	X86_INS_VPHADDUBD = 1079
	X86_INS_VPHADDUBQ = 1080
	X86_INS_VPHADDUBW = 1081
	X86_INS_VPHADDUDQ = 1082
	X86_INS_VPHADDUWD = 1083
	X86_INS_VPHADDUWQ = 1084
	X86_INS_VPHADDWD = 1085
	X86_INS_VPHADDWQ = 1086
	X86_INS_VPHADDW = 1087
	X86_INS_VPHMINPOSUW = 1088
	X86_INS_VPHSUBBW = 1089
	X86_INS_VPHSUBDQ = 1090
	X86_INS_VPHSUBD = 1091
	X86_INS_VPHSUBSW = 1092
	X86_INS_VPHSUBWD = 1093
	X86_INS_VPHSUBW = 1094
	X86_INS_VPINSRB = 1095
	X86_INS_VPINSRD = 1096
	X86_INS_VPINSRQ = 1097
	X86_INS_VPINSRW = 1098
	X86_INS_VPLZCNTD = 1099
	X86_INS_VPLZCNTQ = 1100
	X86_INS_VPMACSDD = 1101
	X86_INS_VPMACSDQH = 1102
	X86_INS_VPMACSDQL = 1103
	X86_INS_VPMACSSDD = 1104
	X86_INS_VPMACSSDQH = 1105
	X86_INS_VPMACSSDQL = 1106
	X86_INS_VPMACSSWD = 1107
	X86_INS_VPMACSSWW = 1108
	X86_INS_VPMACSWD = 1109
	X86_INS_VPMACSWW = 1110
	X86_INS_VPMADCSSWD = 1111
	X86_INS_VPMADCSWD = 1112
	X86_INS_VPMADDUBSW = 1113
	X86_INS_VPMADDWD = 1114
	X86_INS_VPMASKMOVD = 1115
	X86_INS_VPMASKMOVQ = 1116
	X86_INS_VPMAXSB = 1117
	X86_INS_VPMAXSD = 1118
	X86_INS_VPMAXSQ = 1119
	X86_INS_VPMAXSW = 1120
	X86_INS_VPMAXUB = 1121
	X86_INS_VPMAXUD = 1122
	X86_INS_VPMAXUQ = 1123
	X86_INS_VPMAXUW = 1124
	X86_INS_VPMINSB = 1125
	X86_INS_VPMINSD = 1126
	X86_INS_VPMINSQ = 1127
	X86_INS_VPMINSW = 1128
	X86_INS_VPMINUB = 1129
	X86_INS_VPMINUD = 1130
	X86_INS_VPMINUQ = 1131
	X86_INS_VPMINUW = 1132
	X86_INS_VPMOVDB = 1133
	X86_INS_VPMOVDW = 1134
	X86_INS_VPMOVM2B = 1135
	X86_INS_VPMOVM2D = 1136
	X86_INS_VPMOVM2Q = 1137
	X86_INS_VPMOVM2W = 1138
	X86_INS_VPMOVMSKB = 1139
	X86_INS_VPMOVQB = 1140
	X86_INS_VPMOVQD = 1141
	X86_INS_VPMOVQW = 1142
	X86_INS_VPMOVSDB = 1143
	X86_INS_VPMOVSDW = 1144
	X86_INS_VPMOVSQB = 1145
	X86_INS_VPMOVSQD = 1146
	X86_INS_VPMOVSQW = 1147
	X86_INS_VPMOVSXBD = 1148
	X86_INS_VPMOVSXBQ = 1149
	X86_INS_VPMOVSXBW = 1150
	X86_INS_VPMOVSXDQ = 1151
	X86_INS_VPMOVSXWD = 1152
	X86_INS_VPMOVSXWQ = 1153
	X86_INS_VPMOVUSDB = 1154
	X86_INS_VPMOVUSDW = 1155
	X86_INS_VPMOVUSQB = 1156
	X86_INS_VPMOVUSQD = 1157
	X86_INS_VPMOVUSQW = 1158
	X86_INS_VPMOVZXBD = 1159
	X86_INS_VPMOVZXBQ = 1160
	X86_INS_VPMOVZXBW = 1161
	X86_INS_VPMOVZXDQ = 1162
	X86_INS_VPMOVZXWD = 1163
	X86_INS_VPMOVZXWQ = 1164
	X86_INS_VPMULDQ = 1165
	X86_INS_VPMULHRSW = 1166
	X86_INS_VPMULHUW = 1167
	X86_INS_VPMULHW = 1168
	X86_INS_VPMULLD = 1169
	X86_INS_VPMULLQ = 1170
	X86_INS_VPMULLW = 1171
	X86_INS_VPMULUDQ = 1172
	X86_INS_VPORD = 1173
	X86_INS_VPORQ = 1174
	X86_INS_VPOR = 1175
	X86_INS_VPPERM = 1176
	X86_INS_VPROTB = 1177
	X86_INS_VPROTD = 1178
	X86_INS_VPROTQ = 1179
	X86_INS_VPROTW = 1180
	X86_INS_VPSADBW = 1181
	X86_INS_VPSCATTERDD = 1182
	X86_INS_VPSCATTERDQ = 1183
	X86_INS_VPSCATTERQD = 1184
	X86_INS_VPSCATTERQQ = 1185
	X86_INS_VPSHAB = 1186
	X86_INS_VPSHAD = 1187
	X86_INS_VPSHAQ = 1188
	X86_INS_VPSHAW = 1189
	X86_INS_VPSHLB = 1190
	X86_INS_VPSHLD = 1191
	X86_INS_VPSHLQ = 1192
	X86_INS_VPSHLW = 1193
	X86_INS_VPSHUFB = 1194
	X86_INS_VPSHUFD = 1195
	X86_INS_VPSHUFHW = 1196
	X86_INS_VPSHUFLW = 1197
	X86_INS_VPSIGNB = 1198
	X86_INS_VPSIGND = 1199
	X86_INS_VPSIGNW = 1200
	X86_INS_VPSLLDQ = 1201
	X86_INS_VPSLLD = 1202
	X86_INS_VPSLLQ = 1203
	X86_INS_VPSLLVD = 1204
	X86_INS_VPSLLVQ = 1205
	X86_INS_VPSLLW = 1206
	X86_INS_VPSRAD = 1207
	X86_INS_VPSRAQ = 1208
	X86_INS_VPSRAVD = 1209
	X86_INS_VPSRAVQ = 1210
	X86_INS_VPSRAW = 1211
	X86_INS_VPSRLDQ = 1212
	X86_INS_VPSRLD = 1213
	X86_INS_VPSRLQ = 1214
	X86_INS_VPSRLVD = 1215
	X86_INS_VPSRLVQ = 1216
	X86_INS_VPSRLW = 1217
	X86_INS_VPSUBB = 1218
	X86_INS_VPSUBD = 1219
	X86_INS_VPSUBQ = 1220
	X86_INS_VPSUBSB = 1221
	X86_INS_VPSUBSW = 1222
	X86_INS_VPSUBUSB = 1223
	X86_INS_VPSUBUSW = 1224
	X86_INS_VPSUBW = 1225
	X86_INS_VPTESTMD = 1226
	X86_INS_VPTESTMQ = 1227
	X86_INS_VPTESTNMD = 1228
	X86_INS_VPTESTNMQ = 1229
	X86_INS_VPTEST = 1230
	X86_INS_VPUNPCKHBW = 1231
	X86_INS_VPUNPCKHDQ = 1232
	X86_INS_VPUNPCKHQDQ = 1233
	X86_INS_VPUNPCKHWD = 1234
	X86_INS_VPUNPCKLBW = 1235
	X86_INS_VPUNPCKLDQ = 1236
	X86_INS_VPUNPCKLQDQ = 1237
	X86_INS_VPUNPCKLWD = 1238
	X86_INS_VPXORD = 1239
	X86_INS_VPXORQ = 1240
	X86_INS_VPXOR = 1241
	X86_INS_VRCP14PD = 1242
	X86_INS_VRCP14PS = 1243
	X86_INS_VRCP14SD = 1244
	X86_INS_VRCP14SS = 1245
	X86_INS_VRCP28PD = 1246
	X86_INS_VRCP28PS = 1247
	X86_INS_VRCP28SD = 1248
	X86_INS_VRCP28SS = 1249
	X86_INS_VRCPPS = 1250
	X86_INS_VRCPSS = 1251
	X86_INS_VRNDSCALEPD = 1252
	X86_INS_VRNDSCALEPS = 1253
	X86_INS_VRNDSCALESD = 1254
	X86_INS_VRNDSCALESS = 1255
	X86_INS_VROUNDPD = 1256
	X86_INS_VROUNDPS = 1257
	X86_INS_VROUNDSD = 1258
	X86_INS_VROUNDSS = 1259
	X86_INS_VRSQRT14PD = 1260
	X86_INS_VRSQRT14PS = 1261
	X86_INS_VRSQRT14SD = 1262
	X86_INS_VRSQRT14SS = 1263
	X86_INS_VRSQRT28PD = 1264
	X86_INS_VRSQRT28PS = 1265
	X86_INS_VRSQRT28SD = 1266
	X86_INS_VRSQRT28SS = 1267
	X86_INS_VRSQRTPS = 1268
	X86_INS_VRSQRTSS = 1269
	X86_INS_VSCATTERDPD = 1270
	X86_INS_VSCATTERDPS = 1271
	X86_INS_VSCATTERPF0DPD = 1272
	X86_INS_VSCATTERPF0DPS = 1273
	X86_INS_VSCATTERPF0QPD = 1274
	X86_INS_VSCATTERPF0QPS = 1275
	X86_INS_VSCATTERPF1DPD = 1276
	X86_INS_VSCATTERPF1DPS = 1277
	X86_INS_VSCATTERPF1QPD = 1278
	X86_INS_VSCATTERPF1QPS = 1279
	X86_INS_VSCATTERQPD = 1280
	X86_INS_VSCATTERQPS = 1281
	X86_INS_VSHUFPD = 1282
	X86_INS_VSHUFPS = 1283
	X86_INS_VSQRTPD = 1284
	X86_INS_VSQRTPS = 1285
	X86_INS_VSQRTSD = 1286
	X86_INS_VSQRTSS = 1287
	X86_INS_VSTMXCSR = 1288
	X86_INS_VSUBPD = 1289
	X86_INS_VSUBPS = 1290
	X86_INS_VSUBSD = 1291
	X86_INS_VSUBSS = 1292
	X86_INS_VTESTPD = 1293
	X86_INS_VTESTPS = 1294
	X86_INS_VUNPCKHPD = 1295
	X86_INS_VUNPCKHPS = 1296
	X86_INS_VUNPCKLPD = 1297
	X86_INS_VUNPCKLPS = 1298
	X86_INS_VZEROALL = 1299
	X86_INS_VZEROUPPER = 1300
	X86_INS_WAIT = 1301
	X86_INS_WBINVD = 1302
	X86_INS_WRFSBASE = 1303
	X86_INS_WRGSBASE = 1304
	X86_INS_WRMSR = 1305
	X86_INS_XABORT = 1306
	X86_INS_XACQUIRE = 1307
	X86_INS_XBEGIN = 1308
	X86_INS_XCHG = 1309
	X86_INS_XCRYPTCBC = 1310
	X86_INS_XCRYPTCFB = 1311
	X86_INS_XCRYPTCTR = 1312
	X86_INS_XCRYPTECB = 1313
	X86_INS_XCRYPTOFB = 1314
	X86_INS_XEND = 1315
	X86_INS_XGETBV = 1316
	X86_INS_XLATB = 1317
	X86_INS_XRELEASE = 1318
	X86_INS_XRSTOR = 1319
	X86_INS_XRSTOR64 = 1320
	X86_INS_XRSTORS = 1321
	X86_INS_XRSTORS64 = 1322
	X86_INS_XSAVE = 1323
	X86_INS_XSAVE64 = 1324
	X86_INS_XSAVEC = 1325
	X86_INS_XSAVEC64 = 1326
	X86_INS_XSAVEOPT = 1327
	X86_INS_XSAVEOPT64 = 1328
	X86_INS_XSAVES = 1329
	X86_INS_XSAVES64 = 1330
	X86_INS_XSETBV = 1331
	X86_INS_XSHA1 = 1332
	X86_INS_XSHA256 = 1333
	X86_INS_XSTORE = 1334
	X86_INS_XTEST = 1335
	X86_INS_FDISI8087_NOP = 1336
	X86_INS_FENI8087_NOP = 1337
	X86_INS_ENDING = 1338
)
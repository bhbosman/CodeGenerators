    package yacc
    var reservedWords map[string]int
    var string_buf = ""
    var column = 1
    var row = 1

%%
"//".*\n                    {
                                filePosCounter()
                            }
"/*"[^*/]*"*/"              {
                                filePosCounter()
                            }
[ \t\n]+                    {
                                filePosCounter()
                            }

L?'(\\.|[^\\'])+'	        {
                                filePosCounter()
                                return Character_literal
                            }
"<"                         {
                                filePosCounter()
                                return '<'
                            }
">"                         {
                                filePosCounter()
                                return '>'
                            }
"{"                         {
                                filePosCounter()
                                return '{'
                            }
"}"                         {
                                filePosCounter()
                                return '}'
                            }
"("                         {
                                filePosCounter()
                                return '('
                            }
")"                         {
                                filePosCounter()
                                return ')'
                            }
"["                         {
                                filePosCounter()
                                return '['
                            }
"]"                         {
                                filePosCounter()
                                return ']'
                            }
"+"                         {
                                filePosCounter()
                                return '+'
                            }
"-"                         {
                                filePosCounter()
                                return '-'
                            }
"/"                         {
                                filePosCounter()
                                return '/'
                            }
"*"                         {
                                filePosCounter()
                                return '*'
                            }
"%"                         {
                                filePosCounter()
                                return '%'
                            }
";"                         {
                                filePosCounter()
                                return ';'
                            }
","                         {
                                filePosCounter()
                                return ','
                            }
"&"                         {
                                filePosCounter()
                                return '&'
                            }
"|"                         {
                                filePosCounter()
                                return '|'
                            }
"="                         {
                                filePosCounter()
                                return '='
                            }
0x[0-9]+                    {
                                filePosCounter()
                                return Hex_literal
                            }
[1-9][0-9]*                 {
                                filePosCounter()
                                return Integer_literal
                            }
"0"                         {
                                filePosCounter()
                                return Integer_literal
                            }
[A-Za-z_][A-Za-z_0-9]*      {
                                filePosCounter()
                                if v, ok :=  reservedWords[yytext]; ok{
                                    return v
                                }
                                return Identifier
                            }
%%


func filePosCounter(){
		for _, s := range yytext{
			if s == 0{
				continue
			}
    		if s == '\n' {
    			column = 0
    			row++
    		}else{
    			column++
    		}
    	}
}

func init(){
    reservedWords = make(map[string]int)
	reservedWords["enum"] = Rwenum
	reservedWords["struct"] = Rwstruct
	reservedWords["typedef"] = Rwtypedef
	reservedWords["MitchAlpha"] = RwMitchAlpha
	reservedWords["MitchBitField"] = RwMitchBitField
	reservedWords["MitchByte"] = RwMitchByte
	reservedWords["MitchDate"] = RwMitchDate
	reservedWords["MitchTime"] = RwMitchTime
	reservedWords["MitchPrice04"] = RwMitchPrice04
	reservedWords["MitchPrice08"] = RwMitchPrice08
	reservedWords["MitchUInt08"] = RwMitchUInt08
	reservedWords["MitchUInt16"] = RwMitchUInt16
	reservedWords["MitchUInt32"] = RwMitchUInt32
	reservedWords["MitchUInt64"] = RwMitchUInt64
}
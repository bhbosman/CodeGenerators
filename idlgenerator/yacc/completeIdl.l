    package yacc
    import "strings"
    import "path"


    func stripToFileName(s string) string{
        return strings.TrimFunc(
            s,
            func(r rune) bool {
                switch r {
                case '"', '<', '>', '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
                    return true
                }
                return false

            })
    	}

%s datablock
%x ifdefExpression
%x defineExpression
%x pragmaExpression

%%
<INITIAL,datablock>[ \t]*"#"include[ \t]*   {


                                                dd := bytes.NewBufferString("")

                                                self.filePosCounter()

                                                for {
                                                    b := byte(self.input())

                                                    if b == '\n' {
                                                        break
                                                    }
                                                    _ = dd.WriteByte(b)
                                                }
                                                s := dd.String()
                                                fmt.Println(s)

                                                fileName := stripToFileName(s)
                                                dir, _ := path.Split(self.currentStream.fileName)
                                                newFileName := path.Join(dir, fileName)
                                                file, e := os.Open(newFileName)
                                                if e != nil {
                                                    panic(fmt.Sprintf("error opening %v. error: %v", newFileName, e.Error()))

                                                }
                                                defer func() {
                                                    _ = file.Close()
                                                }()
                                                stream := NewCurrentStream(fileName, file)
                                                self.filePosCounter()
                                                self.PushStream(stream)

                                            }


<INITIAL,datablock>[ \t]*"#"define              {
                                                self.filePosCounter()
                                                ctx.(*AdditionalInformation).stackState.defineFlag = true
                                                ctx.(*AdditionalInformation).push(self)
                                                self.yyBEGIN(defineExpression)

                                            }
<INITIAL,datablock>[ \t]*"#"undefine            {
                                                self.filePosCounter()
                                                ctx.(*AdditionalInformation).stackState.defineFlag = false
                                                ctx.(*AdditionalInformation).push(self)
                                                self.yyBEGIN(defineExpression)
                                            }
<pragmaExpression>(.*)\n                    {
                                                self.filePosCounter()
                                                ctx.(*AdditionalInformation).pop(self)
                                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                                if returnResult {
                                                    if ctx.(*AdditionalInformation).DefinitionContext != nil {
                                                            ctx.(*AdditionalInformation).DefinitionContext.ParsePragma(strings.TrimSpace(self.yytext))
                                                    }
                                                }
                                                self.yytext = ""
                                            }
<datablock>\n                                   {
                                                self.filePosCounter()
                                            }
<defineExpression>(.*)\n                    {
                                                self.filePosCounter()
                                                ctx.(*AdditionalInformation).pop(self)
                                                if ctx.(*AdditionalInformation).DefinitionContext != nil {
                                                        ctx.(*AdditionalInformation).DefinitionContext.ParseDefinition(strings.TrimSpace(self.yytext), ctx.(*AdditionalInformation).stackState.defineFlag)
                                                }
                                                self.yytext = ""
                                            }
<INITIAL,datablock>[ \t]*"#"pragma              {
                                                self.filePosCounter()
                                                ctx.(*AdditionalInformation).push(self)
                                                self.yyBEGIN(pragmaExpression)
                                            }

<INITIAL,datablock>[ \t]*"#"ifdef               {
                                                self.filePosCounter()
                                                ctx.(*AdditionalInformation).StartIfDefBlock(self, true)
                                                //return RWifdef
                                            }
<INITIAL,datablock>[ \t]*"#"ifndef              {
                                                self.filePosCounter()
                                                ctx.(*AdditionalInformation).StartIfDefBlock(self, false)
                                                //return RWifndef
                                            }
<datablock>[ \t]*"#"endif                       {
                                                self.filePosCounter()
                                                ctx.(*AdditionalInformation).pop(self)
                                                if ctx.(*AdditionalInformation).YY_STACK == nil{
                                                    self.yyBEGIN(INITIAL)
                                                    //self.yytext = ""
                                                    //return DefinitionBlock
                                                } else{
                                                    self.yyBEGIN(datablock)
                                                }

                                                //return RWendif
                                            }
<datablock>[ \t]*"#"else                        {
                                                self.filePosCounter()
                                                self.yyBEGIN(datablock)
                                                ctx.(*AdditionalInformation).stackState.ifdefBlock = false
                                                ctx.(*AdditionalInformation).stackState.expressionValue = !ctx.(*AdditionalInformation).stackState.expressionValue
                                                //return RWelse
                                            }
<ifdefExpression>(.*)\n                          {

                                                self.filePosCounter()
                                                self.yyBEGIN(datablock)
                                                ctx.(*AdditionalInformation).stackState.ifdefBlock = true

                                                tempValue := true
                                                if ctx.(*AdditionalInformation).DefinitionContext != nil{
                                                    tempValue = ctx.(*AdditionalInformation).DefinitionContext.ParseExpression(strings.TrimSpace(self.yytext))
                                                }
                                                ctx.(*AdditionalInformation).stackState.expressionValue = tempValue == ctx.(*AdditionalInformation).stackState.expressionDirection

                                                //return ExpressionBlock
                                            }
<<EOF>>                                     {
                                                self.PopStream()
                                            }

<datablock,INITIAL>"//".*\n                    {
                                self.filePosCounter()
                            }
<datablock,INITIAL>"/*"[^*/]*"*/"              {
                                self.filePosCounter()
                            }
<datablock,INITIAL>[ \t\n]+           {
                                self.filePosCounter()
                            }

<datablock,INITIAL>L?'(\\.|[^\\'])+'	        {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return character_literal
                                }
                            }




<datablock,INITIAL>L?\"(\\.|[^\\"])*\"	{
                                            self.filePosCounter()
                                            returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                            if returnResult {
                                                return string_literal
                                            }
                                        }


<datablock,INITIAL>"<"                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '<'
                                }
                            }
<datablock,INITIAL>">"                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '>'
                                 }
                            }
<datablock,INITIAL>"{"                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '{'
                                }
                            }
<datablock,INITIAL>"}"                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '}'
                                }
                            }
<datablock,INITIAL>"("                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '('
                                }
                            }
<datablock,INITIAL>")"                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return ')'
                                }
                            }
<datablock,INITIAL>"["                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '['
                                }
                            }
<datablock,INITIAL>"]"                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return ']'
                                }
                            }
<datablock,INITIAL>"+"                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '+'
                                }
                            }

<datablock,INITIAL>"-"                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '-'
                                }
                            }
<datablock,INITIAL>"/"                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '/'
                                }
                            }
<datablock,INITIAL>"*"      {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '*'
                                }
                            }
<datablock,INITIAL>"%"      {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '%'
                                }
                            }
<datablock,INITIAL>";"      {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return ';'
                                }
                            }
<datablock,INITIAL>","                         {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return ','
                                }
                            }
<datablock,INITIAL>"&"      {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '&'
                                }
                            }

<datablock,INITIAL>"="      {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '='
                                }
                            }
<datablock,INITIAL>":"      {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return ':'
                                }
                            }

<datablock,INITIAL>"|"      {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return '|'
                                }
                            }

<datablock,INITIAL>"::"     {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return  ScopeOp
                                }
                            }

<datablock,INITIAL>"<<"     {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return  ShlOp
                                }
                            }
<datablock,INITIAL>">>"    {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return  ShrOp
                                }
                            }
<datablock,INITIAL>0[xX][0-9A-Fa-f]+    {
                                            self.filePosCounter()
                                            returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                            if returnResult {
                                                return Hex_literal
                                            }
                                        }
<datablock,INITIAL>[1-9][0-9]*  {
                                    self.filePosCounter()
                                    returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                    if returnResult {
                                        return integer_literal
                                    }
                                }
<datablock,INITIAL>"0"      {
                                self.filePosCounter()
                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                if returnResult {
                                    return integer_literal
                                }
                            }
<datablock,INITIAL>[A-Za-z_][A-Za-z_0-9]*   {
                                                self.filePosCounter()
                                                returnResult :=  ((self.YY_START == datablock) &&  ctx.(*AdditionalInformation).stackState.expressionValue) || (self.YY_START == INITIAL)
                                                if returnResult {
                                                    if v, ok :=  ctx.(*AdditionalInformation).reservedWords[self.yytext]; ok{
                                                        return v
                                                    }
                                                    return identifier
                                                }

                                            }
%%




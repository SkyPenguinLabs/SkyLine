package SkyLine_Backend

func LookupIdentifier(ident string) Token_Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TOKEN_IDENT
}

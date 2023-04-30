package SkyLine_Backend

//SAVE:D1
type ScriptSettings_Mult struct {
	RunFile      bool     // Run a source script, we do not allow multiple sources
	OutputOnCall bool     // Tell to output when place of varname X :: STD<-CALL
	AllowedFiles []string // Allowed file extensions
	Server       bool     // Load help server
}

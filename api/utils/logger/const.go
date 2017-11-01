package logger

const (
	// Cor padrao do shell
	reset = "\033[0m"
	// Negrito
	bold = "\033[1m"
	// Italico
	italic = "\033[3m"
	// Linha sublinhada
	underline = "\033[4m"
	// Linha sublinhada OFF
	underlineOff = "\033[24m"
	// Inverte a cor
	inverse = "\033[7m"
	// Inverte a cor OFF
	inverseOff = "\033[27m"
	// Linha tracejada
	strikethrough = "\033[9m"
	// Linha tracejada OFF
	strikethroughOff = "\033[29m"
	// Branco
	white = "\033[37m"
	// Preto
	black = "\033[30m"
	// Vermelho
	red = "\033[31m"
	// Verde
	green = "\033[32m"
	// Azul
	blue = "\033[34m"
	// Amarelo
	yellow = "\033[33m"
	// Magenta
	magenta = "\033[35m"
	// Cinza
	cyan = "\033[36m"
	// Background Branco
	whiteBg = "\033[47m"
	// Background Preto
	blackBg = "\033[40m"
	// Background Vermelho
	redBg = "\033[41m"
	// Background Verde
	greenBg = "\033[42m"
	// Background Azul
	blueBg = "\033[44m"
	// Background Amarelo
	yellowBg = "\033[43m"
	// Background Magenta
	magentaBg = "\033[45m"
	// Background Cinza
	cyanBg = "\033[46m"
	// logClock Padrao de formatacao do cabecalho dos logs
	logClock = "[2/01/2006 15:04:05]"
)

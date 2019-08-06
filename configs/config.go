package configs

type configuration struct {
	BaseURI string
}

var Config = &configuration{
	BaseURI: "https://matrix.moonshard.tech",
}

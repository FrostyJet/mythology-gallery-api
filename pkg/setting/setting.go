package setting

var (
	EnvironmentSetting = &Environment{}
	DatabaseSetting    = &Database{}
	ServerSetting      = &Server{}
)

func Setup() {
	EnvironmentSetting.Setup()
	DatabaseSetting.Setup()
	ServerSetting.Setup()
}

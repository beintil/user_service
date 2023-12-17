package client

import "github.com/fatih/color"

var (
	ErrorColor             = color.New(color.FgHiRed, color.Bold)
	InformationColorYellow = color.New(color.FgHiYellow, color.Bold)
	InformationColorGreen  = color.New(color.FgGreen, color.Bold)
	InformationColorCyan   = color.New(color.FgCyan, color.Bold)
)

const (
	DividingLines     = "\n---------------------------------------------------------------------------------------\n"
	LinkUpUserService = "\n    __    _       __   __  __                                \n   / /   (_)___  / /__/ / / /___                             \n  / /   / / __ \\/ //_/ / / / __ \\                            \n / /___/ / / / / ,< / /_/ / /_/ /                  _         \n/_____/_/_/_/_/_/|_|\\____/ .___/_____  ______   __(_)_______ \n / / / / ___/ _ \\/ ___/_/_/_/ ___/ _ \\/ ___/ | / / / ___/ _ \\\n/ /_/ (__  )  __/ /  /_____(__  )  __/ /   | |/ / / /__/  __/\n\\__,_/____/\\___/_/        /____/\\___/_/    |___/_/\\___/\\___/ \n                                                             \n"
)

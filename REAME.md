# Simple ZSH history cleaning app

## Description

This application allows you to manage fillment of the ZSH history file. You manually can specify the value in `config.yaml` file. This value will be the limit that shows when app should clean history commands.
`Default value is 80%`. That`s mean that app will clean entire file after 80% fillement.

How to change percent? Manually change the content of `config.yaml` file.
For example: default config file content is `limit_persent: 80`, to change the limit just replace 80 into any other value of percent you want.

## Todo

1. I plan to add console configuration, to specify file contents.
2. Moreover I also want to find solution how to remove wrong commands from history file.

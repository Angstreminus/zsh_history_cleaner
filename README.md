# Simple ZSH history cleaning app

## Description

This application allows you to manage fillment of the ZSH history file. You manually can specify the value in `.env` file. This value will be the limit that shows when app should clean history commands.
`Default value is 80%`. That`s mean that app will clean entire file after 80% fillement.

How to change percent? You should manually change the content of `.env` file.
For example: default config file content is `LIMIT_PERCENT=80`, to change the limit just replace 80 into any other value of percent you want.

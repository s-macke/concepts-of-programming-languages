# Exercise 10 - Enterprise Programming

If you do not finish during the lecture period, please finish it as homework.

## Exercise 10.1 - Modules

- Write a Go module mail with API and impl
- The Mail implementation should log a message with logrus (https://github.com/sirupsen/logrus)
- Change your module dependency to a former logrus version
- Write a second Go module client which uses the mail module
- Make an incompatible change to your mail API and increase the major version number

## Exercise 10.2 - Plugins

- Export your Go mail module from the previous exercise to a mail plugin
- The Mail implementation should log a message with logrus (https://github.com/sirupsen/logrus)
- Write a third Go module client which uses the mail plugin and also logrus
- Make an incompatible change to your mail API and increase the major version number
- Try to use a different logrus version in the plugin and the client

Windows users: use Golang in your Windows Subsystem for Linux or a Docker container.

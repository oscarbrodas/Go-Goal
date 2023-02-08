cd all the way into the httpd folder where "main.go" is

In the terminal, run "go build" to create a .exe file and run this file with ".\nameOfFile.exe"

The MySQL workbench localhost must have the password set to "password" and database name of "somedb"

If a different password was used, change the "password" part of the DSN variable in main.go to your password

If a different database name was used, change the "somedb" part of the DSN variable in main.go to your database name

Use postman to send http requests. Sample json inputs are provided in the "input.txt" file

/*
Package gogoretry exposes an API service that implements retry mechanism.

First, you can initiate the package by calling gogoretry.New(), no argument required.
It will resort to default value for 5 retries every one second.

Configuration can be done calling MaxEntries, Sleep, Default functions and pass it to the New() function
*/
package gogoretry

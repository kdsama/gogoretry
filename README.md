# Gogoretry

![example workflow](https://github.com/kdsama/gogoretry/actions/workflows/go.yml/badge.svg)

## Installation

To install Gogoretry, use the following `go get` command:

```sh
go get github.com/kdsama/gogoretry
```


## Overview
Gogoretry is a Go package that provides a flexible and customizable retry mechanism for handling errors in your applications. It allows you to set various retry-related options and strategies, making it easier to manage retries in your code.

## Usage 

### Importing the package 

``` import "github.com/kdsama/gogoretry" ```

### Creating a Retrier 

kdsamaTo create a Retrier with default settings (1-second sleep time and 5 max retries), use:

``` retrier := gogoretry.New() ```

You can customize the Retrier using the following options:

+ Sleep Time: Set the sleep time between each retry.
``` retrier := gogoretry.Sleep(time.Second) ```

+ Maximum Number of Retries: Set the maximum number of retries.
``` retrier := gogoretry.MaxRetries(5) ```

+ Custom Time Intervals: Provide a custom time interval slice for retries.
``` customIntervals := []time.Duration{time.Second, 2 * time.Second, 5 * time.Second} 
retrier := gogoretry.Custom(customIntervals) ```

+ Exponential Backoffs: Set up exponential backoffs with a time duration, multiplier, and max retries.
``` retrier := gogoretry.Exponential(time.Second, 2, 5) ```

+ Handling Bad Errors: Specify a list of errors that, if encountered, will not trigger retries.
``` badErrors := []error{specificError1, specificError2} 
 retrier := gogoretry.BadErrors(badErrors) ```

+ Handling Retry Errors: Specify a list of errors that, if not encountered, will trigger retries.
``` retryErrors := []error{specificError1, specificError2} 
 retrier := gogoretry.RetryErrors(retryErrors) ```


## Running Retry Logic
``` Use the Run method to execute the retry logic:
err := retrier.Run(func() error {
    // Your code that may return an error goes here
    result, err := yourFunction()
    return err
}) ```

If err is not nil, it means the system attempted retries as per your settings and did not receive a successful response.

## Author/Contact
Author : Kshitij 













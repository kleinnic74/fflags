# Minimalistic Feature Flag Go Library

fflags is a minimalistic feature-flag library to manage feature hierarchies across an application.

It offers the following functionalities:

* no dependencies except standard library
* define trees of features across an application
* layered configuration backends:
  * yaml/json files
  * remote server
  * in-app variables

## Usage

Add the library to your project

    go get gihub.com/kleinnic74/fflags

Define features:

    var (
        myFeature = fflags.Define("feature.aspect")
    )

Initialize feature activation state:

    package main

    import "github.com/kleinnic74/fflags"

    func init() {
        fflags.Init(fflags.TurnOn([]string{"feature"}),
                    fflags.TurnOff([]string{"feature.json"}),
                    fflags.YamlFile("feature.yaml"))
     }

Do things conditionally based upon feature state:

    package main

    func main() {
        if myFeature.IsActive() {
            // Do something
        }
    }


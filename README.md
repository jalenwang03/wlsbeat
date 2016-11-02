# Wlsbeat

Welcome to Wlsbeat.

Wlsbeat is a `beat` wrote by go to collect performance data from WebLogic domain,you can send data to both logstash and elasticsearch directly,you just need to fill some configuration in wlsbeat.yml,such as host(IP),port,username and password.but you need pay attention,a weblogic feature called `RESTful Management Extensions` was used in wlsbeat to collect data,this feature is pretty useful if you want to query AdminServer's Domain for various pieces of information it knows about Servers,Clusters,Applications and Datasources，to enable this feature,you must modify your configuration in your weblogic admin console,detail information please reference weblogic documentation.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/jalenwang03`

## Getting Started with Wlsbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Init Project
To get running with Wlsbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Wlsbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/jalenwang03/wlsbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Wlsbeat run the command below. This will generate a binary
in the same directory with the name wlsbeat.

```
make
```


### Run

To run Wlsbeat with debugging output enabled, run:

```
./wlsbeat -c wlsbeat.yml -e -d "*"
```


### Test

To test Wlsbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/wlsbeat.template.json and etc/wlsbeat.asciidoc

```
make update
```


### Cleanup

To clean  Wlsbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Wlsbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/jalenwang03
cd ${GOPATH}/github.com/jalenwang03
git clone https://github.com/jalenwang03/wlsbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make package
```

This will fetch and create all images required for the build process. The hole process to finish can take several minutes.

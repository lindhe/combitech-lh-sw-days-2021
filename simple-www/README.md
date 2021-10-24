# Simple WWW

I'm _not_ a web designer.
Red, green or blue is the best I can give you.

## Build

Using `podman` instead of `docker` is cool.
You want to be cool.
They use pretty much the same syntax, just slightly different defaults from a user's perspective.

```console
$ podman build -t docker.io/lindhe/simple-www:red-0.1.0 .
STEP 1/2: FROM docker.io/nginx:1.21.3-alpine
STEP 2/2: COPY www/* /usr/share/nginx/html
--> Using cache 47e4e7b115091dc22fd0ba49885c026df3721102ee61a073acf029edf90f35c7
COMMIT docker.io/lindhe/simple-www:red-0.1.0
--> 47e4e7b1150
Successfully tagged docker.io/lindhe/simple-www:red-0.1.0
Successfully tagged localhost/www:latest
47e4e7b115091dc22fd0ba49885c026df3721102ee61a073acf029edf90f35c7
```

## Push

```console
$  echo "${DOCKER_PASSWORD}" | podman login -u "${DOCKER_USERNAME}" --password-stdin docker.io
Login Succeeded!
$ podman push docker.io/lindhe/simple-www:red-0.1.0
Getting image source signatures
Copying blob e2eb06d8af82 skipped: already exists
Copying blob e6d3cea19fef skipped: already exists
Copying blob 40403bebe4fd skipped: already exists
Copying blob 20d0effdf3a2 skipped: already exists
Copying blob 311d8db33235 skipped: already exists
Copying blob b4b4e85910ea skipped: already exists
Copying blob 38d4352aecf2 done
Copying config 47e4e7b115 done
Writing manifest to image destination
Storing signatures
```
